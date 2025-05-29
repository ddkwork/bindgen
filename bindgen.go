package bindgen

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"github.com/ddkwork/ddk/vswhere"
	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/safemap"
	"github.com/ddkwork/golibrary/stream"
	"github.com/ddkwork/golibrary/waitgroup"
	"github.com/tidwall/gjson"
)

type (
	ClangCModelCallback func(path string) bool
	SkipFileCallback    ClangCModelCallback
)

func Walk(root, targetDir string, skipFileCallback SkipFileCallback, cModelCallback ClangCModelCallback) {
	var paths []string
	mylog.Check(filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if !skipFileCallback(path) {
			return err
		}
		switch filepath.Ext(path) {
		case ".cpp", ".c", ".h":
			paths = append(paths, path)
		}
		return err
	}))
	Bind(targetDir, cModelCallback, paths...)
}

func Bind(targetDir string, cModelCallback ClangCModelCallback, paths ...string) {
	result := Result{
		Enums:     new(safemap.M[string, EnumInfo]),
		Structs:   new(safemap.M[string, StructInfo]),
		Functions: new(safemap.M[string, FunctionInfo]),
		Typedefs:  new(safemap.M[string, string]),
	}
	w := waitgroup.New()
	for _, path := range paths {
		w.Go(func() {
			root := gjson.ParseBytes(runClangASTDump(path, cModelCallback))
			typedefsNameByID(root, &result)
		})
	}
	w.Wait()

	w2 := waitgroup.New()
	for _, path := range paths {
		mylog.Warning("bind", path)
		w2.Go(func() {
			root := gjson.ParseBytes(runClangASTDump(path, cModelCallback))
			traverseNode(root, &result, path)
		})
	}

	w2.Wait()
	os.RemoveAll("tmp")
	os.RemoveAll(targetDir)

	genGoFile(result, targetDir, paths...)
	mylog.Success("bind success")
}

func genGoFile(results Result, targetDir string, paths ...string) {
	var methods []string
	packet := new(safemap.M[string, string]) // urlPath api
	pkgName := filepath.Base(targetDir)
	for _, path := range paths {
		buffer := bytes.NewBufferString("")
		buffer.WriteString("package " + pkgName + "\n")
		for _, e := range results.Enums.Range() {
			if e.Comment.currentFile != path {
				continue
			}
			buffer.WriteString(fmt.Sprintf("type %s int // %s\nconst (\n", e.Name, e.Loc))
			for i, m := range e.Members {
				line := fmt.Sprintf("\t%s", m.Name)
				if i == 0 {
					line += " " + e.Name + " = iota"
				} else if m.ExplicitValue != "" {
					line += " = " + m.ExplicitValue
				}
				buffer.WriteString(fmt.Sprintf("%s // %d\n", line, m.ComputedValue))
			}
			buffer.WriteString(")\n")
		}

		for _, s := range results.Structs.Range() {
			if s.Comment.currentFile != path {
				continue
			}
			buffer.WriteString(fmt.Sprintf("// %s (%s)\n", s.Name, s.Loc))
			buffer.WriteString(fmt.Sprintf("type %s struct {\n", s.Name))
			for _, f := range s.Fields {
				buffer.WriteString(fmt.Sprintf("\t%s %s // C type: %s\n",
					strings.Title(f.Name),
					f.TypeDecl,
					f.Type))
			}
			buffer.WriteString("}\n")
		}

		fileName := strings.TrimPrefix(filepath.Base(path), "_")
		methodName := strings.TrimPrefix(fileName, "scriptapi_")
		methodName = stream.TrimExtension(methodName)
		methods = append(methods, methodName)
		buffer.WriteString(fmt.Sprintf("type %s struct {}\n", methodName))

		for _, m := range results.Functions.Range() {
			if m.Comment.currentFile != path {
				continue
			}
			m.Name = strings.TrimPrefix(m.Name, path)
			buffer.WriteString(fmt.Sprintf("func ("+string(methodName[0])+" *%s) %s(", methodName, m.Name))
			params := make([]string, len(m.Params))
			for i, p := range m.Params {
				params[i] = fmt.Sprintf("%s %s", p.Name, p.Type)
			}

			urlPath := fmt.Sprintf("%s/%s", filepath.Base(path), m.Name)
			//?GetList@Bookmark@Script@@YA_NPEAUListInfo@@@Z
			api := m.Comment.mangledName
			if strings.HasPrefix(api, "?") { // namespace
				split := strings.Split(api, "@")
				api = split[2] + "::" + split[1] + "::" + strings.TrimPrefix(split[0], "?")
			}
			// mylog.Trace(m.name, m.Comment.mangledName)
			packet.Set(urlPath, api)
			body := "response := safeGet(" +
				strconv.Quote(urlPath) +
				", map[string]string{})\n\tif len(response) == 0 {\n\t\treturn \n\t}"
			buffer.WriteString(fmt.Sprintf("%s) %s {\n\t"+
				body+
				"}\n\n",
				strings.Join(params, ", "),
				m.ReturnType))
		}

		mylog.Info(path, fileName+"_gen.go")
		stream.WriteGoFile(filepath.Join(targetDir, fileName+"_gen.go"), buffer.String())
	}
	buffer := bytes.NewBufferString("")
	buffer.WriteString("package " + pkgName + "\n")
	buffer.WriteString(`
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/ddkwork/golibrary/mylog"
)
`)
	buffer.WriteString("type Sdk struct {" + "\n")
	for _, m := range methods {
		buffer.WriteString(fmt.Sprintf("\t%s *%s\n", stream.ToCamelUpper(m), m))
	}
	buffer.WriteString("}\n")
	buffer.WriteString(fmt.Sprintf("func New() *Sdk {\n\treturn &Sdk{" + "\n"))
	for _, m := range methods {
		buffer.WriteString(fmt.Sprintf("\t\t%s: new(%s),\n", stream.ToCamelUpper(m), m))
	}
	buffer.WriteString("}}\n")
	// todo rename for any project
	body := `
const DefaultX64dbgServer = "http://127.0.0.1:8888"
var x64dbgServerURL string

func init() {
	if len(os.Args) > 1 {
		x64dbgServerURL = os.Args[1]
	} else {
		x64dbgServerURL = DefaultX64dbgServer
	}
}

func safeGet(endpoint string, params map[string]string) []string {
	if params == nil {
		params = make(map[string]string)
	}

	url := fmt.Sprintf("%s/%s", x64dbgServerURL, endpoint)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []string{fmt.Sprintf("Request failed: %s", err.Error())}
	}

	q := request.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	request.URL.RawQuery = q.Encode()
	// mylog.Info("Sending request to", request.URL.String())
	mylog.Request(request, true)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return []string{fmt.Sprintf("Request failed: %s", err.Error())}
	}
	defer response.Body.Close()
	mylog.Response(response, true)
	log.Println()
	log.Println()
	log.Println()

	if response.StatusCode == http.StatusOK {
		var lines []string
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			return []string{fmt.Sprintf("Error reading response: %s", err.Error())}
		}
		s := string(bodyBytes)
		buffer := new(bytes.Buffer)
		err = json.Indent(buffer, []byte(s), "", " ")
		if err != nil {
			lines = strings.Split(s, "\n")
		} else {
			// todo GetModuleList 直接解码到结构体
			lines = strings.Split(buffer.String(), "\n")
		}
		return lines
	}
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return []string{fmt.Sprintf("Error reading response: %s", err.Error())}
	}
	return []string{fmt.Sprintf("Error %d: %s", response.StatusCode, strings.TrimSpace(string(bodyBytes)))}
}

func safePost(endpoint string, data any) string {
	var body []byte
	var err error
	if strData, ok := data.(string); ok {
		body = []byte(strData)
	} else if dataMap, ok := data.(map[string]string); ok {
		body, err = json.Marshal(dataMap)
		if err != nil {
			return fmt.Sprintf("Request failed: %s", err.Error())
		}
	} else {
		return "Request failed: unsupported data type"
	}

	url := fmt.Sprintf("%s/%s", x64dbgServerURL, endpoint)
	request, err := http.NewRequest("POST", url, strings.NewReader(string(body)))
	if err != nil {
		return fmt.Sprintf("Request failed: %s", err.Error())
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return fmt.Sprintf("Request failed: %s", err.Error())
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			return fmt.Sprintf("Error reading response: %s", err.Error())
		}
		return strings.TrimSpace(string(bodyBytes))
	} else {
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			return fmt.Sprintf("Error reading response: %s", err.Error())
		}
		return fmt.Sprintf("Error %d: %s", response.StatusCode, strings.TrimSpace(string(bodyBytes)))
	}
}
`
	buffer.WriteString(body)
	stream.WriteGoFile(filepath.Join(targetDir, "sdk_gen.go"), buffer.String())

	//todo
	/*
		APIS = [
		    {
		        "path": "/_scriptapi_memory.h/ReadByte",
		        "func": "Script::Memory::ReadByte",
		        "params": ["address"]//类型，名称，value的切片
		        "return_type": "uint8"
		    },
		    {
		        "path": "/_scriptapi_memory.h/ReadDword",
		        "func": "Script::Memory::ReadDword",
		        "params": ["address"],
		        "return_type": "uint32"
		    },
		    # 添加更多API...
		]
	*/
	// mylog.Struct(packet.Map())
	stream.MarshalJsonToFile(packet.Map(), filepath.Join(targetDir, "api.json"))
	genMcpCppServerCode(packet)
}

func genMcpCppServerCode(m *safemap.M[string, string]) {
	start := `
//
// Created by Admin on 28/05/2025.
//
#ifndef MCPX64DBG_ECHO_GEN_H
#define MCPX64DBG_ECHO_GEN_H
#include "pluginsdk/bridgemain.h"
#include "pluginsdk/_plugins.h"
#include "pluginsdk/_scriptapi_argument.h"
#include "pluginsdk/_scriptapi_assembler.h"
#include "pluginsdk/_scriptapi_bookmark.h"
#include "pluginsdk/_scriptapi_comment.h"
#include "pluginsdk/_scriptapi_debug.h"
#include "pluginsdk/_scriptapi_flag.h"
#include "pluginsdk/_scriptapi_function.h"
#include "pluginsdk/_scriptapi_gui.h"
#include "pluginsdk/_scriptapi_label.h"
#include "pluginsdk/_scriptapi_memory.h"
#include "pluginsdk/_scriptapi_misc.h"
#include "pluginsdk/_scriptapi_module.h"
#include "pluginsdk/_scriptapi_pattern.h"
#include "pluginsdk/_scriptapi_register.h"
#include "pluginsdk/_scriptapi_stack.h"
#include "pluginsdk/_scriptapi_symbol.h"
#include "pluginsdk/DeviceNameResolver/DeviceNameResolver.h"
#include "pluginsdk/jansson/jansson.h"
#include "pluginsdk/lz4/lz4file.h"
#include "pluginsdk/TitanEngine/TitanEngine.h"
#include "pluginsdk/XEDParse/XEDParse.h"
#include <winsock2.h>
#include <ws2tcpip.h>
#include <map>
#include <string>
#include <functional>
#include <mutex>
#include <ios>
#include <sstream>
#include <algorithm>
//=============================================================================
// HTTP Server Implementation
//=============================================================================
// Default settings
#define DEFAULT_PORT     8888
#define MAX_REQUEST_SIZE 8192
// Global variables
int        g_pluginHandle;
HANDLE     g_httpServerThread = NULL;
bool       g_httpServerRunning = false;
int        g_httpPort = DEFAULT_PORT;
std::mutex g_httpMutex;
SOCKET     g_serverSocket = INVALID_SOCKET;
// Forward declarations
bool                                         startHttpServer();
void                                         stopHttpServer();
void                                         dispatch(SOCKET clientSocket);
DWORD WINAPI                                 HttpServerThread(LPVOID lpParam);
std::string                                  readHttpRequest(SOCKET clientSocket);
void                                         sendHttpResponse(SOCKET clientSocket, int statusCode, const std::string &contentType, const std::string &responseBody);
void                                         parseHttpRequest(const std::string &request, std::string &method, std::string &path, std::string &query, std::string &body);
std::unordered_map<std::string, std::string> parseQueryParams(const std::string &query);
// Command callback declarations
bool cbEnableHttpServer(int argc, char *argv[]);
bool cbSetHttpPort(int argc, char *argv[]);

void dispatch(SOCKET clientSocket) {
    /*
// Read the HTTP request
        std::string requestData = readHttpRequest(clientSocket);

        if (!requestData.empty()) {
            // Parse the HTTP request
            std::string method, path, query, body;
            parseHttpRequest(requestData, method, path, query, body);

            _plugin_logprintf("HTTP Request: %s %s\n", method.c_str(), path.c_str());

            // Parse query parameters
            std::unordered_map<std::string, std::string> queryParams = parseQueryParams(query);

            // Handle different endpoints
            try {
                if (path == "/IsDebugActive") {
                    bool active = DbgIsRunning();
                    sendHttpResponse(clientSocket, 200, "text/plain", active ? "true" : "false");
                } else if (path == "/ExeConsoleCmd") {
                    std::string cmd = queryParams["Command"];
                    if (cmd.empty() && !body.empty()) {
                        cmd = body;
                    }
                    bool success = DbgCmdExec(cmd.c_str());
                    sendHttpResponse(clientSocket, 200, "text/plain",
                                     success ? "Command executed successfully" : "Command execution failed");
                } else if (path == "/FindMemBase") {
                    std::string addrStr = queryParams["addr"];
                    if (addrStr.empty() && !body.empty()) {
                        addrStr = body;
                    }
                    _plugin_logprintf("FindMemBase endpoint called with addr: %s\n", addrStr.c_str());
                    // Convert string address to duint
                    duint addr = 0;
                    try {
                        addr = std::stoull(addrStr, nullptr, 16); // Parse as hex
                    }
                    catch (const std::exception &e) {
                        sendHttpResponse(clientSocket, 400, "text/plain", "Invalid address format");
                        continue;
                    }
                    _plugin_logprintf("Converted address: 0x%llx\n", addr);

                    // Get the base address and size
                    duint size = 0;
                    duint baseAddr = DbgMemFindBaseAddr(addr, &size);
                    _plugin_logprintf("Base address found: 0x%llx, size: %llu\n", baseAddr, size);
                    if (baseAddr == 0) {
                        sendHttpResponse(clientSocket, 404, "text/plain", "No module found for this address");
                    } else {
                        // Format the response with base address and size
                        std::stringstream ss;
                        ss << std::hex << "0x" << baseAddr << "," << size;
                        sendHttpResponse(clientSocket, 200, "text/plain", ss.str());
                    }
                } else if (path == "/GetModuleList") {
                    // Create a list to store the module information
                    ListInfo moduleList;

                    // Get the list of modules
                    bool success = Script::Module::GetList(&moduleList);

                    if (!success) {
                        sendHttpResponse(clientSocket, 500, "text/plain", "Failed to get module list");
                    } else {
                        // Create a JSON array to hold the module information
                        std::stringstream jsonResponse;
                        jsonResponse << "[";

                        // Iterate through each module in the list
                        size_t count = moduleList.count;
                        Script::Module::ModuleInfo *modules = (Script::Module::ModuleInfo *) moduleList.data;

                        for (size_t i = 0; i < count; i++) {
                            if (i > 0) jsonResponse << ",";

                            // Add module info as JSON object
                            jsonResponse << "{";
                            jsonResponse << "\"name\":\"" << modules[i].name << "\",";
                            jsonResponse << "\"base\":\"0x" << std::hex << modules[i].base << "\",";
                            jsonResponse << "\"size\":\"0x" << std::hex << modules[i].size << "\",";
                            jsonResponse << "\"entry\":\"0x" << std::hex << modules[i].entry << "\",";
                            jsonResponse << "\"sectionCount\":" << std::dec << modules[i].sectionCount << ",";
                            jsonResponse << "\"path\":\"" << modules[i].path << "\"";
                            jsonResponse << "}";
                        }

                        jsonResponse << "]";

                        // Free the list
                        BridgeFree(moduleList.data);

                        // Send the response
                        sendHttpResponse(clientSocket, 200, "application/json", jsonResponse.str());
                    }
                } else {
                    // Unknown URL
                    sendHttpResponse(clientSocket, 404, "text/plain", "Not Found");
                }
            }
            catch (const std::exception &e) {
                // Exception in handling request
                sendHttpResponse(clientSocket, 500, "text/plain", std::string("Internal Server Error: ") + e.what());
            }
        }
 * */
    // Read the HTTP request
    std::string requestData = readHttpRequest(clientSocket);
    if (!requestData.empty()) {
        // Parse the HTTP request
        std::string method, path, query, body;
        parseHttpRequest(requestData, method, path, query, body);
        _plugin_logprintf("HTTP Request: %s %s\n", method.c_str(), path.c_str());
        // Parse query parameters
        std::unordered_map<std::string, std::string>                        queryParams = parseQueryParams(query);
        static const std::unordered_map<std::string, std::function<void()>> handlers{
`

	end := `

        };
        if (auto it = handlers.find(path); it != handlers.end()) it->second();//todo: 这里的路径解析有问题，发包恢复客户端
    };
}

// HTTP server thread function using standard Winsock
DWORD WINAPI HttpServerThread(LPVOID lpParam) {
    WSADATA wsaData;
    int     result = WSAStartup(MAKEWORD(2, 2), &wsaData);
    if (result != 0) {
        _plugin_logprintf("WSAStartup failed with error: %d\n", result);
        return 1;
    }
    // Create a socket for the server
    g_serverSocket = socket(AF_INET, SOCK_STREAM, IPPROTO_TCP);
    if (g_serverSocket == INVALID_SOCKET) {
        _plugin_logprintf("Failed to create socket, error: %d\n", WSAGetLastError());
        WSACleanup();
        return 1;
    }
    // Setup the server address structure
    sockaddr_in serverAddr;
    serverAddr.sin_family = AF_INET;
    serverAddr.sin_addr.s_addr = htonl(INADDR_LOOPBACK);// localhost only
    serverAddr.sin_port = htons((u_short) g_httpPort);
    // Bind the socket
    if (bind(g_serverSocket, (sockaddr *) &serverAddr, sizeof(serverAddr)) == SOCKET_ERROR) {
        _plugin_logprintf("Bind failed with error: %d\n", WSAGetLastError());
        closesocket(g_serverSocket);
        WSACleanup();
        return 1;
    }
    // Listen for incoming connections
    if (listen(g_serverSocket, SOMAXCONN) == SOCKET_ERROR) {
        _plugin_logprintf("Listen failed with error: %d\n", WSAGetLastError());
        closesocket(g_serverSocket);
        WSACleanup();
        return 1;
    }
    _plugin_logprintf("HTTP server started at http://localhost:%d/\n", g_httpPort);
    // Set socket to non-blocking mode
    u_long mode = 1;
    ioctlsocket(g_serverSocket, FIONBIO, &mode);
    // Main server loop
    while (g_httpServerRunning) {
        // Accept a client connection
        sockaddr_in clientAddr;
        int         clientAddrSize = sizeof(clientAddr);
        SOCKET      clientSocket = accept(g_serverSocket, (sockaddr *) &clientAddr, &clientAddrSize);
        if (clientSocket == INVALID_SOCKET) {
            // Check if we need to exit
            if (!g_httpServerRunning) { break; }
            // Non-blocking socket may return WOULD_BLOCK when no connections are pending
            if (WSAGetLastError() != WSAEWOULDBLOCK) { _plugin_logprintf("Accept failed with error: %d\n", WSAGetLastError()); }
            Sleep(100);// Avoid tight loop
            continue;
        }
        dispatch(clientSocket);
        closesocket(clientSocket);// Close the client socket
    }
    // Clean up
    if (g_serverSocket != INVALID_SOCKET) {
        closesocket(g_serverSocket);
        g_serverSocket = INVALID_SOCKET;
    }
    WSACleanup();
    return 0;
}

// Function to read the HTTP request
std::string readHttpRequest(SOCKET clientSocket) {
    std::string request;
    char        buffer[MAX_REQUEST_SIZE];
    int         bytesReceived;
    // Set socket to blocking mode to receive full request
    u_long mode = 0;
    ioctlsocket(clientSocket, FIONBIO, &mode);
    // Receive data
    bytesReceived = recv(clientSocket, buffer, sizeof(buffer) - 1, 0);
    if (bytesReceived > 0) {
        buffer[bytesReceived] = '\0';
        request = buffer;
    }
    return request;
}

// Function to parse an HTTP request
void parseHttpRequest(const std::string &request, std::string &method, std::string &path, std::string &query, std::string &body) {
    // Parse the request line
    size_t firstLineEnd = request.find("\r\n");
    if (firstLineEnd == std::string::npos) { return; }
    std::string requestLine = request.substr(0, firstLineEnd);
    // Extract method and URL
    size_t methodEnd = requestLine.find(' ');
    if (methodEnd == std::string::npos) { return; }
    method = requestLine.substr(0, methodEnd);
    size_t urlEnd = requestLine.find(' ', methodEnd + 1);
    if (urlEnd == std::string::npos) { return; }
    std::string url = requestLine.substr(methodEnd + 1, urlEnd - methodEnd - 1);
    // Split URL into path and query
    size_t queryStart = url.find('?');
    if (queryStart != std::string::npos) {
        path = url.substr(0, queryStart);
        query = url.substr(queryStart + 1);
    } else {
        path = url;
        query = "";
    }
    // Find the end of headers and start of body
    size_t headersEnd = request.find("\r\n\r\n");
    if (headersEnd == std::string::npos) { return; }
    // Extract body
    body = request.substr(headersEnd + 4);
}

// Function to send HTTP response
void sendHttpResponse(SOCKET clientSocket, int statusCode, const std::string &contentType, const std::string &responseBody) {
    // Prepare status line
    std::string statusText;
    switch (statusCode) {
        case 200:
            statusText = "OK";
            break;
        case 404:
            statusText = "Not Found";
            break;
        case 500:
            statusText = "Internal Server Error";
            break;
        default:
            statusText = "Unknown";
    }
    // Build the response
    std::stringstream response;
    response << "HTTP/1.1 " << statusCode << " " << statusText << "\r\n";
    response << "Content-Type: " << contentType << "\r\n";
    response << "Content-Length: " << responseBody.length() << "\r\n";
    response << "Connection: close\r\n";
    response << "\r\n";
    response << responseBody;
    // Send the response
    std::string responseStr = response.str();
    send(clientSocket, responseStr.c_str(), (int) responseStr.length(), 0);
}

// Parse query parameters from URL
std::unordered_map<std::string, std::string> parseQueryParams(const std::string &query) {
    std::unordered_map<std::string, std::string> params;
    size_t                                       pos = 0;
    size_t                                       nextPos;
    while (pos < query.length()) {
        nextPos = query.find('&', pos);
        if (nextPos == std::string::npos) { nextPos = query.length(); }
        std::string pair = query.substr(pos, nextPos - pos);
        size_t      equalPos = pair.find('=');
        if (equalPos != std::string::npos) {
            std::string key = pair.substr(0, equalPos);
            std::string value = pair.substr(equalPos + 1);
            params[key] = value;
        }
        pos = nextPos + 1;
    }
    return params;
}

// Command callback for toggling HTTP server
bool cbEnableHttpServer(int argc, char *argv[]) {
    if (g_httpServerRunning) {
        _plugin_logputs("Stopping HTTP server...");
        stopHttpServer();
        _plugin_logputs("HTTP server stopped");
    } else {
        _plugin_logputs("Starting HTTP server...");
        if (startHttpServer()) {
            _plugin_logprintf("HTTP server started on port %d\n", g_httpPort);
        } else {
            _plugin_logputs("Failed to start HTTP server");
        }
    }
    return true;
}

// Command callback for changing HTTP server port
bool cbSetHttpPort(int argc, char *argv[]) {
    if (argc < 2) {
        _plugin_logputs("Usage: httpport [port_number]");
        return false;
    }
    int port;
    try {
        port = std::stoi(argv[1]);
    } catch (const std::exception &) {
        _plugin_logputs("Invalid port number");
        return false;
    }
    if (port <= 0 || port > 65535) {
        _plugin_logputs("Port number must be between 1 and 65535");
        return false;
    }
    g_httpPort = port;
    if (g_httpServerRunning) {
        _plugin_logputs("Restarting HTTP server with new port...");
        stopHttpServer();
        if (startHttpServer()) {
            _plugin_logprintf("HTTP server restarted on port %d\n", g_httpPort);
        } else {
            _plugin_logputs("Failed to restart HTTP server");
        }
    } else {
        _plugin_logprintf("HTTP port set to %d\n", g_httpPort);
    }
    return true;
}

// Start the HTTP server
bool startHttpServer() {
    std::lock_guard<std::mutex> lock(g_httpMutex);
    // Stop existing server if running
    if (g_httpServerRunning) { stopHttpServer(); }
    // Create and start the server thread
    g_httpServerThread = CreateThread(NULL, 0, HttpServerThread, NULL, 0, NULL);
    if (g_httpServerThread == NULL) {
        _plugin_logputs("Failed to create HTTP server thread");
        return false;
    }
    g_httpServerRunning = true;
    return true;
}

// Stop the HTTP server
void stopHttpServer() {
    std::lock_guard<std::mutex> lock(g_httpMutex);
    if (g_httpServerRunning) {
        g_httpServerRunning = false;
        // Close the server socket to unblock any accept calls
        if (g_serverSocket != INVALID_SOCKET) {
            closesocket(g_serverSocket);
            g_serverSocket = INVALID_SOCKET;
        }
        // Wait for the thread to exit
        if (g_httpServerThread != NULL) {
            WaitForSingleObject(g_httpServerThread, 1000);
            CloseHandle(g_httpServerThread);
            g_httpServerThread = NULL;
        }
    }
}

/*
void echo(const std::string path) {
    // 使用 std::map 来存储路径和对应的操作函数
    std::map<std::string, std::function<void()>> pathActions = {
            {"_scriptapi_argument.h/Add",   []() {
                Script::Argument::Add(nullptr);
            }},
            {"_scriptapi_argument.h/Clear", []() {
                Script::Argument::Clear();
            }},
    };

    // 查找并执行对应的操作
    auto it = pathActions.find(path);
    if (it != pathActions.end()) {
        it->second();
    } else {
        // 默认情况下什么都不做
    }
}


#include <iostream>
#include <string>

// 编译期哈希计算（DJB2算法）
constexpr size_t hash(const char *str) noexcept {
    size_t hash = 5381;
    while (*str != '\0') {
        hash = ((hash << 5) + hash) + *str; // hash = hash * 33 + c
        ++str;
    }
    return hash;
}

void dispatch2(const std::string &path) {
    const size_t h = hash(path.c_str());
    switch (h) {
        case hash("Namespace1::myFunc"):
            if (path != "Namespace1::myFunc") break; // 防哈希碰撞
            std::cout << "Call Namespace1::myFunc\n";
            break;
        case hash("Namespace2::myFunc"):
            if (path != "Namespace2::myFunc") break;
            std::cout << "Call Namespace2::myFunc\n";
            break;
        default:
            std::cerr << "Unknown path: " << path << "\n";
    }
}

int main() {
    dispatch("Namespace1::myFunc");
    dispatch("Namespace2::myFunc");
    dispatch("InvalidPath");
    return 0;
}
 */
#endif//MCPX64DBG_ECHO_GEN_H
`

	g := stream.NewGeneratedFile()
	g.P(start)
	for path, api := range m.Map() {
		g.P("\t{")
		g.P("\t\t", strconv.Quote(path), ",")
		g.P("\t\t\t", "[&queryParams, &body, &clientSocket] {")
		api = api
		//参数处理代码生成逻辑
		//取出ast得到的函数参数和返回值，queryParams解码参数名称。如果没有参数则从body中取出参数
		//处理返回值，cpp只有一个返回值应该好处理
		//判断返回值类型
		//字符串，返回空字符串发包回复客户端，否则回复字符串封包
		//整型，返回0或者-1，发包回复客户端，否则使用std::hex格式化返回值，封包回复客户端
		//bool，返回false或者true，发包回复客户端，否则使用std::string格式化返回值，封包回复客户端
		//结构体，返回空结构体，发包回复客户端，否则使用json格式化返回值，封包回复客户端
		//list，生成json结构体对象切片语法回复客户端
		//todo 逻辑：参数解析，传入参数执行api，api返回检查和回复，api返回json封包回复客户端
		//  std::string name = queryParams["name"];
		//      if (name.empty() && !body.empty()) { name = body; }
		//      _plugin_logprintf("bridgemain.h/DbgModBaseFromName endpoint called with name: %s\n", name.c_str());
		//      auto base = DbgModBaseFromName(name.c_str());
		//      _plugin_logprintf("Base address found: 0x%llx\n", base);
		//      if (base == 0) {
		//          sendHttpResponse(clientSocket, 404, "text/plain", "No module found for this address");
		//      } else {
		//          // Format the response with base address and size
		//          std::stringstream ss;
		//          ss << std::hex << "0x" << base;
		//          sendHttpResponse(clientSocket, 200, "text/plain", ss.str());
		//      }
		g.P("\t\t\t },")
		g.P("\t },")
	}
	g.P(end)
	stream.WriteTruncate("echo_gen.h", g.String())
}

func typedefsNameByID(root gjson.Result, result *Result) {
	root.Get("inner").ForEach(func(_, child gjson.Result) bool {
		if child.Get("kind").String() == "TypedefDecl" {
			id := child.Get("inner.0.ownedTagDecl.id").String()
			name := child.Get("name").String()
			if id == "" {
				return true
			}
			result.Typedefs.Update(id, name)
		}
		return true
	})
}

// -DWIN32_LEAN_AND_MEAN
//
//	-D_WINSOCKAPI_
var Flags = `
#define WIN32_LEAN_AND_MEAN
#define WINSOCK_DEPRECATED_NO_WARNINGS
//#define _WINSOCKAPI_
typedef char bool;
`

func extractFlags() []string {
	f := make([]string, 0)
	for s := range strings.Lines(Flags) {
		switch {
		case strings.HasPrefix(s, "//"):
			s = strings.TrimSpace(s)
		case strings.HasPrefix(s, "#define"):
			s = strings.TrimSpace(strings.TrimPrefix(s, "#define")) // todo -D
			f = append(f, "-D"+s)
		default:
			s = strings.TrimSpace(s)
		}
	}
	return f
}

func runClangASTDump(path string, cModelCallback ClangCModelCallback) []byte {
	path = mylog.Check2(filepath.Abs(path))
	jsonPath := filepath.Join("cache", filepath.Base(path)+".json")
	if stream.FileExists(jsonPath) {
		return mylog.Check2(os.ReadFile(jsonPath))
	}
	arg := []string{
		"clang",
		"-x", "c++", // bridgemain.h need c model，否则会找不解析很多函数，原因不明
		"-Xclang", "-ast-dump=json", "-fsyntax-only",
	}
	if cModelCallback(path) {
		arg = slices.Delete(arg, 1, 3)
	}
	arg = append(arg, extractFlags()...)
	includes := vswhere.New().VisualStudio().Includes
	for _, include := range includes {
		arg = append(arg, "-I", include)
	}
	arg = append(arg, "-I", filepath.Dir(path))
	arg = append(arg, path)
	out := stream.RunCommandArgs(arg...)
	stream.WriteTruncate(jsonPath, out.Output)
	return out.Output.Bytes()
}

func traverseNode(node gjson.Result, result *Result, path string) {
	var processNode func(gjson.Result)
	processNode = func(n gjson.Result) {
		name := n.Get("name").String()
		comment := Comment{
			currentFile:  path,
			name:         name,
			mangledName:  n.Get("mangledName").String(),
			file:         n.Get("range.begin.file").String(),
			includedFrom: n.Get("range.begin.includedFrom.file").String(),
			expansionLoc: n.Get("range.begin.expansionLoc.file").String(),
		}
		kind := n.Get("kind").String()
		if kind == "BuiltinType" {
			return
		}
		if !strings.Contains(n.Raw, "TranslationUnitDecl") && strings.Contains(n.Raw, "Program Files") {
			return
		}
		switch kind {
		case "EnumDecl":
			if name == "" {
				id := n.Get("id").String()
				value, exist := result.Typedefs.Get(id)
				if !exist {
					mylog.Struct(comment)
					mylog.Todo(id + " not found")
				}
				name = value
			}
			enum := parseEnum(n)
			if enum.Name == "" {
				if name == "" {
					return
				}
				enum.Name = name
			}
			enum.Comment = comment
			result.Enums.Update(enum.Name, enum)
		case "RecordDecl", "CXXRecordDecl":
			if name == "" {
				id := n.Get("id").String()
				value, exist := result.Typedefs.Get(id)
				if !exist {
					mylog.Struct(comment)
					mylog.Todo(id + " not found")
				}
				name = value
			}
			if n.Get("tagUsed").String() != "struct" {
				return
			}
			if slices.Contains([]string{
				"_GUID",
			}, name) {
				return
			}
			for _, s2 := range []string{"PLUG_", "json_"} {
				if strings.HasPrefix(name, s2) {
					return
				}
			}
			object := parseStruct(n)
			if object.Name == "" {
				if name == "" {
					return
				}
				object.Name = name
			}
			if object.Fields == nil {
				return
			}
			object.Comment = comment
			result.Structs.Update(object.Name, object)
		case "FunctionDecl", "CXXMethodDecl":
			if name == "" {
				id := n.Get("id").String()
				value, exist := result.Typedefs.Get(id)
				if !exist {
					mylog.Struct(comment)
					mylog.Todo(id + " not found")
				}
				name = value
			}
			skipFunc := []string{
				//"ToVector",
			}
			if slices.Contains(skipFunc, name) {
				return
			}
			skip2 := []string{ // todo add callback
				"operator",
				"json_",
				"_plugin_",
			}
			for _, s2 := range skip2 {
				if strings.HasPrefix(name, s2) {
					return
				}
			}
			function := parseFunction(n)
			if function.Name == "" {
				if name == "" {
					return
				}
				function.Name = name
			}
			function.Comment = comment
			result.Functions.Update(path+function.Name, function)
		}

		n.Get("inner").ForEach(func(_, child gjson.Result) bool {
			processNode(child)
			return true
		})
	}

	processNode(node)
}

func parseEnum(node gjson.Result) EnumInfo {
	info := EnumInfo{
		Name: node.Get("name").String(),
		Loc:  formatLoc(node.Get("loc")),
	}

	nextValue := 0
	node.Get("inner").ForEach(func(_, member gjson.Result) bool {
		if member.Get("kind").String() == "EnumConstantDecl" {
			explicit, computed := resolveEnumValue(member, nextValue)
			if explicit != "" {
				if val, e := parseNumber(explicit); e == nil {
					nextValue = val + 1
				}
			} else {
				nextValue++
			}

			info.Members = append(info.Members, EnumMember{
				Name:          member.Get("name").String(),
				ExplicitValue: explicit,
				ComputedValue: computed,
			})
		}
		return true
	})

	return info
}

func parseStruct(node gjson.Result) StructInfo {
	info := StructInfo{
		Name:   node.Get("name").String(),
		Loc:    formatLoc(node.Get("loc")),
		IsImpl: node.Get("isImplicit").Bool(),
	}

	node.Get("inner").ForEach(func(_, child gjson.Result) bool {
		if child.Get("isBitfield").Bool() {
			mylog.Json("bitfield", child.Raw)
			mylog.Todo("# 生成结构体内存布局报告（支持C/C++）\n" +
				"clang -Xclang -fdump-record-layouts -fsyntax-only input.c")
		}
		switch kind := child.Get("kind").String(); kind {
		case "RecordDecl", "CXXRecordDecl":
		case "FieldDecl":
			info.Fields = append(info.Fields, StructField{
				Name:     child.Get("name").String(),
				Type:     child.Get("type.qualType").String(),
				TypeDecl: resolveType(child.Get("type")),
			})
		case "CXXMethodDecl":
			fn := parseFunction(child)
			fn.IsMethod = true
			fn.ReceiverType = info.Name
			info.Methods = append(info.Methods, fn)
		}
		return true
	})

	return info
}

func parseFunction(node gjson.Result) FunctionInfo {
	if node.Get("name").String() == "AddNode" {
		println()
	}
	info := FunctionInfo{
		Name:       node.Get("name").String(),
		Loc:        formatLoc(node.Get("loc")),
		ReturnType: resolveType(node.Get("type.resultType")),
	}
	node.Get("inner").ForEach(func(_, param gjson.Result) bool {
		if param.Get("kind").String() == "ParmVarDecl" {
			s := param.Get("name").String()
			switch s { // todo more syntax check
			case "type":
				s = "Type"
			case "string":
				s = "s"
			}
			s = strings.NewReplacer(
				"type", "Type",
				"const", "",
				"&", "",
			).Replace(s)
			info.Params = append(info.Params, FunctionParam{
				Name: s,
				Type: resolveType(param.Get("type")),
			})
		}
		return true
	})

	return info
}

func resolveEnumValue(node gjson.Result, defaultVal int) (string, int) {
	var explicit string
	node.Get("inner").ForEach(func(_, child gjson.Result) bool {
		if child.Get("value").Exists() {
			explicit = child.Get("value").String()
			return false
		}
		return true
	})
	if explicit == "" {
		explicit = node.Get("init.value").String()
	}
	if explicit != "" {
		val, e := parseNumber(explicit)
		if e != nil {
			mylog.CheckIgnore(e)
			return "", defaultVal
		}
		return explicit, val
	}
	return "", defaultVal
}

func parseNumber(s string) (int, error) {
	base := 10
	if strings.HasPrefix(s, "0x") {
		s = s[2:]
		base = 16
		i, e := strconv.ParseInt(s, base, 32)
		return int(i), e
	} else if strings.HasPrefix(s, "0") && len(s) > 1 {
		base = 8
		i, e := strconv.ParseInt(s, base, 32)
		return int(i), e
	}
	return strconv.Atoi(s)
}

func resolveType(typeNode gjson.Result) string {
	return handleQualType(typeNode.Get("qualType").String())
}

func handleQualType(qualType string) string {
	m := safemap.NewOrdered[string, string](func(yield func(string, string) bool) {
		yield("long double", "float128")
		yield("int", "int")
		yield("HANDLE", "uintptr")
		yield("ULONG_PTR", "uintptr")
		yield("WORD", "uint16")
		yield("DWORD", "uint32")
		yield("DWORD32", "uint32")
		yield("DWORD64", "uint64")
		yield("Long", "int32")
		yield("ULong", "uint32")
		yield("ULONG64", "uint64")
		yield("ULONGLONG", "uint64")
		yield("ULong32", "uint32")
		yield("ULong64", "uint64")
		yield("UINT64", "uint64")
		yield("UShort32", "uint16")
		yield("ULongLong", "uint64")
		yield("UShort", "uint16")
		yield("UChar", "byte")
		yield("BYTE", "byte")
		yield("UByte", "byte")
		yield("CHAR", "int8")
		yield("unsigned int", "uint")
		yield("long long", "int64")
		yield("long", "int32")
		yield("unsigned long long", "uint64")
		yield("unsigned long", "uint32")
		yield("short", "int16")
		yield("unsigned short", "uint16")
		yield("char", "int8")
		yield("unsigned char", "byte")
		yield("float", "float32")
		yield("double", "float64")
		yield("bool", "bool")
		yield("BOOL", "bool")
		yield("void", "void")
		yield("PVOID", "uintptr")
		yield("PCHAR", "int8*")

		yield("int8_t", "int8")
		yield("int16_t", "int16")
		yield("int32_t", "int32")
		yield("int64_t", "int64")
		yield("uint8_t", "uint8")
		yield("uint16_t", "uint16")
		yield("uint32_t", "uint32")
		yield("uint64_t", "uint64")

		yield("size_t", "uint") // todo
		yield("uintptr_t", "uintptr")
		yield("intmax_t", "int64")
		yield("uintmax_t", "uint64")
		yield("ptrdiff_t", "int64")
		yield("ssize_t", "int64")
		yield("off_t", "int64")
		yield("time_t", "int64")
		yield("clock_t", "int64")
		yield("va_list", "uintptr")
		yield("jmp_buf", "uintptr")
		yield("sigjmp_buf", "uintptr")
		yield("sig_atomic_t", "int32")
		yield("locale_t", "uintptr")
		yield("fpos_t", "int64")

		yield("const char *", "string")
		yield("duint", "uint")                               // todo test duint
		yield("BRIDGEBP", "uintptr")                         // todo  以下大部分是不需要绑定的，这样这是为了让go语法通过格式化
		yield("FILETIME", "uint64")                          // todo "time.Duration"
		yield("XMMREGISTER", "uintptr")                      // todo
		yield("FLAGS", "uintptr")                            // todo
		yield("SYMBOLPTR", "uintptr")                        // todo
		yield("VALUE_SIZE", "uintptr")                       // todo
		yield("MEMORY_BASIC_INFORMATION", "uintptr")         // todo
		yield("MSG", "uintptr")                              // todo
		yield("BASIC_INSTRUCTION_INFO", "uintptr")           // todo
		yield("CBSYMBOLENUM", "uintptr")                     // todo
		yield("MXCSRFIELDS", "uintptr")                      // todo
		yield("X87STATUSWORDFIELDS", "uintptr")              // todo
		yield("X87CONTROLWORDFIELDS", "uintptr")             // todo
		yield("SYMBOLINFO", "uintptr")                       // todo
		yield("GUISCRIPTEXECUTE", "uintptr")                 // todo
		yield("GUISCRIPTCOMPLETER", "uintptr")               // todo
		yield("GUICALLBACKEX", "uintptr")                    // todo
		yield("GUICALLBACK", "uintptr")                      // todo
		yield("TYPEDESCRIPTOR", "uintptr")                   // todo
		yield("BridgeCFGraphList", "uintptr")                // todo
		yield("std::vector<BridgeCFInstruction>", "uintptr") // todo
		// yield("xed_operand_enum_t", "uintptr")               // todo
	})
	var arg []string
	for k, v := range m.Range() {
		arg = append(arg, k, v)
	}
	s := strings.NewReplacer(arg...).Replace(qualType)
	if strings.Contains(s, "*") {
		s = strings.NewReplacer(
			"const ", "",
			"*", "",
			"void", "uintptr", // todo bug
		).Replace(s)
		return "*" + s
	}

	if strings.Contains(s, "[") {
		index := strings.Index(s, "[")
		arrayType := s[:index]
		arrayType = strings.NewReplacer(arg...).Replace(arrayType)
		arrayLength := s[index:]
		return arrayLength + arrayType
	}
	if strings.Contains(s, "&") {
		s = strings.NewReplacer(
			"const ", "",
			"&", "",
		).Replace(s)
		return "*" + s
	}
	return s
}

func formatLoc(loc gjson.Result) string {
	if !loc.Exists() {
		return ""
	}
	return fmt.Sprintf("%s:%d",
		loc.Get("file").String(),
		loc.Get("line").Int(),
	)
}

type (
	EnumMember struct {
		Name          string
		ExplicitValue string
		ComputedValue int
	}

	EnumInfo struct {
		Name    string
		Loc     string
		Members []EnumMember
		Comment
	}

	StructField struct {
		Name     string
		Type     string
		TypeDecl string
	}

	StructInfo struct {
		Name    string
		Loc     string
		IsImpl  bool
		Fields  []StructField
		Methods []FunctionInfo
		Comment
	}

	FunctionParam struct {
		Name string
		Type string
		Comment
	}

	FunctionInfo struct {
		Name         string
		Loc          string
		ReturnType   string
		Params       []FunctionParam
		IsMethod     bool
		ReceiverType string
		Comment
	}

	Result struct {
		Enums     *safemap.M[string, EnumInfo]
		Structs   *safemap.M[string, StructInfo]
		Functions *safemap.M[string, FunctionInfo]
		Typedefs  *safemap.M[string, string] // id -> name
	}
	Comment struct {
		currentFile  string
		name         string
		mangledName  string
		file         string
		includedFrom string
		expansionLoc string
	}
)

func (c Comment) String() string {
	// return "/*\n" + pretty.Format(c) + "\n*/\n"
	return strings.Join([]string{
		"// currentFile: " + c.currentFile + ":1",
		"//        name: " + c.name,
		"//  mangleName: " + c.mangledName,
		"//        file: " + c.file,
		"//includedFrom: " + c.includedFrom,
		"//expansionLoc: " + c.expansionLoc,
	}, "\n") + "\n"
}
