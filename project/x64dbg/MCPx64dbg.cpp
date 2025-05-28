//+build ignore


#define _WINSOCK_DEPRECATED_NO_WARNINGS
#define WIN32_LEAN_AND_MEAN

// Include Windows headers before socket headers
#include <Windows.h>
// Include x64dbg SDK
#include "pluginsdk/bridgemain.h"
#include "pluginsdk/_plugins.h"
#include "pluginsdk/_scriptapi_module.h"

// Socket includes - after Windows.h
#include <winsock2.h>
#include <ws2tcpip.h>

// Standard library includes
#include <string>
#include <vector>
#include <unordered_map>
#include <sstream>
#include <mutex>
#include <thread>
#include <algorithm>
#include <memory>

// Link with ws2_32.lib
#pragma comment(lib, "ws2_32.lib")

// Plugin information
#define PLUGIN_NAME "x64dbg HTTP Server"
#define PLUGIN_VERSION 1

// Default settings
#define DEFAULT_PORT 8888
#define MAX_REQUEST_SIZE 8192

// Global variables
int g_pluginHandle;
HANDLE g_httpServerThread = NULL;
bool g_httpServerRunning = false;
int g_httpPort = DEFAULT_PORT;
std::mutex g_httpMutex;
SOCKET g_serverSocket = INVALID_SOCKET;

// Forward declarations
bool startHttpServer();
void stopHttpServer();
DWORD WINAPI HttpServerThread(LPVOID lpParam);
std::string readHttpRequest(SOCKET clientSocket);
void sendHttpResponse(SOCKET clientSocket, int statusCode, const std::string& contentType, const std::string& responseBody);
void parseHttpRequest(const std::string& request, std::string& method, std::string& path, std::string& query, std::string& body);
std::unordered_map<std::string, std::string> parseQueryParams(const std::string& query);

// Command callback declarations
bool cbEnableHttpServer(int argc, char* argv[]);
bool cbSetHttpPort(int argc, char* argv[]);
void registerCommands();

//=============================================================================
// Plugin Interface Implementation
//=============================================================================

// Initialize the plugin
bool pluginInit(PLUG_INITSTRUCT* initStruct) {
    initStruct->pluginVersion = PLUGIN_VERSION;
    initStruct->sdkVersion = PLUG_SDKVERSION;
    strncpy_s(initStruct->pluginName, PLUGIN_NAME, _TRUNCATE);
    g_pluginHandle = initStruct->pluginHandle;
    
    _plugin_logputs("x64dbg HTTP Server plugin loading...");
    
    // Register commands
    registerCommands();
    
    // Start the HTTP server
    if (startHttpServer()) {
        _plugin_logprintf("x64dbg HTTP Server started on port %d\n", g_httpPort);
    } else {
        _plugin_logputs("Failed to start HTTP server!");
    }
    
    _plugin_logputs("x64dbg HTTP Server plugin loaded!");
    return true;
}

// Stop the plugin
void pluginStop() {
    _plugin_logputs("Stopping x64dbg HTTP Server...");
    stopHttpServer();
    _plugin_logputs("x64dbg HTTP Server stopped.");
}

// Plugin setup
bool pluginSetup() {
    return true;
}

// Plugin exports
extern "C" __declspec(dllexport) bool pluginit(PLUG_INITSTRUCT* initStruct) {
    return pluginInit(initStruct);
}

extern "C" __declspec(dllexport) void plugstop() {
    pluginStop();
}

extern "C" __declspec(dllexport) void plugsetup(PLUG_SETUPSTRUCT* setupStruct) {
    pluginSetup();
}

//=============================================================================
// HTTP Server Implementation
//=============================================================================

// Start the HTTP server
bool startHttpServer() {
    std::lock_guard<std::mutex> lock(g_httpMutex);
    
    // Stop existing server if running
    if (g_httpServerRunning) {
        stopHttpServer();
    }
    
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

// HTTP server thread function using standard Winsock
DWORD WINAPI HttpServerThread(LPVOID lpParam) {
    WSADATA wsaData;
    int result = WSAStartup(MAKEWORD(2, 2), &wsaData);
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
    serverAddr.sin_addr.s_addr = htonl(INADDR_LOOPBACK); // localhost only
    serverAddr.sin_port = htons((u_short)g_httpPort);
    
    // Bind the socket
    if (bind(g_serverSocket, (sockaddr*)&serverAddr, sizeof(serverAddr)) == SOCKET_ERROR) {
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
        int clientAddrSize = sizeof(clientAddr);
        SOCKET clientSocket = accept(g_serverSocket, (sockaddr*)&clientAddr, &clientAddrSize);
        
        if (clientSocket == INVALID_SOCKET) {
            // Check if we need to exit
            if (!g_httpServerRunning) {
                break;
            }
            
            // Non-blocking socket may return WOULD_BLOCK when no connections are pending
            if (WSAGetLastError() != WSAEWOULDBLOCK) {
                _plugin_logprintf("Accept failed with error: %d\n", WSAGetLastError());
            }
            
            Sleep(100); // Avoid tight loop
            continue;
        }
        
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
                }
                else if (path == "/ExeConsoleCmd") {
                    std::string cmd = queryParams["Command"];
                    if (cmd.empty() && !body.empty()) {
                        cmd = body;
                    }
                    bool success = DbgCmdExec(cmd.c_str());
                    sendHttpResponse(clientSocket, 200, "text/plain", 
                                  success ? "Command executed successfully" : "Command execution failed");
                }
                else if (path == "/FindMemBase") {
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
                    catch (const std::exception& e) {
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
                    }
                    else {
                        // Format the response with base address and size
                        std::stringstream ss;
                        ss << std::hex << "0x" << baseAddr << "," << size;
                        sendHttpResponse(clientSocket, 200, "text/plain", ss.str());
                    }
                }
                else if (path == "/GetModuleList") {
                    // Create a list to store the module information
                    ListInfo moduleList;
                    
                    // Get the list of modules
                    bool success = Script::Module::GetList(&moduleList);
                    
                    if (!success) {
                        sendHttpResponse(clientSocket, 500, "text/plain", "Failed to get module list");
                    }
                    else {
                        // Create a JSON array to hold the module information
                        std::stringstream jsonResponse;
                        jsonResponse << "[";
                        
                        // Iterate through each module in the list
                        size_t count = moduleList.count;
                        Script::Module::ModuleInfo* modules = (Script::Module::ModuleInfo*)moduleList.data;
                        
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
                }
                else {
                    // Unknown URL
                    sendHttpResponse(clientSocket, 404, "text/plain", "Not Found");
                }
            }
            catch (const std::exception& e) {
                // Exception in handling request
                sendHttpResponse(clientSocket, 500, "text/plain", std::string("Internal Server Error: ") + e.what());
            }
        }
        
        // Close the client socket
        closesocket(clientSocket);
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
    char buffer[MAX_REQUEST_SIZE];
    int bytesReceived;
    
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
void parseHttpRequest(const std::string& request, std::string& method, std::string& path, std::string& query, std::string& body) {
    // Parse the request line
    size_t firstLineEnd = request.find("\r\n");
    if (firstLineEnd == std::string::npos) {
        return;
    }
    
    std::string requestLine = request.substr(0, firstLineEnd);
    
    // Extract method and URL
    size_t methodEnd = requestLine.find(' ');
    if (methodEnd == std::string::npos) {
        return;
    }
    
    method = requestLine.substr(0, methodEnd);
    
    size_t urlEnd = requestLine.find(' ', methodEnd + 1);
    if (urlEnd == std::string::npos) {
        return;
    }
    
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
    if (headersEnd == std::string::npos) {
        return;
    }
    
    // Extract body
    body = request.substr(headersEnd + 4);
}

// Function to send HTTP response
void sendHttpResponse(SOCKET clientSocket, int statusCode, const std::string& contentType, const std::string& responseBody) {
    // Prepare status line
    std::string statusText;
    switch (statusCode) {
        case 200: statusText = "OK"; break;
        case 404: statusText = "Not Found"; break;
        case 500: statusText = "Internal Server Error"; break;
        default: statusText = "Unknown";
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
    send(clientSocket, responseStr.c_str(), (int)responseStr.length(), 0);
}

// Parse query parameters from URL
std::unordered_map<std::string, std::string> parseQueryParams(const std::string& query) {
    std::unordered_map<std::string, std::string> params;
    
    size_t pos = 0;
    size_t nextPos;
    
    while (pos < query.length()) {
        nextPos = query.find('&', pos);
        if (nextPos == std::string::npos) {
            nextPos = query.length();
        }
        
        std::string pair = query.substr(pos, nextPos - pos);
        size_t equalPos = pair.find('=');
        
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
bool cbEnableHttpServer(int argc, char* argv[]) {
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
bool cbSetHttpPort(int argc, char* argv[]) {
    if (argc < 2) {
        _plugin_logputs("Usage: httpport [port_number]");
        return false;
    }
    
    int port;
    try {
        port = std::stoi(argv[1]);
    }
    catch (const std::exception&) {
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

// Register plugin commands
void registerCommands() {
    _plugin_registercommand(g_pluginHandle, "httpserver", cbEnableHttpServer, 
                           "Toggle HTTP server on/off");
    _plugin_registercommand(g_pluginHandle, "httpport", cbSetHttpPort, 
                           "Set HTTP server port");
}