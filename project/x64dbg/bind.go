package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/ddkwork/golibrary/mylog"
)

func main() {

	fmt.Println("Is Debugging:", IsDebugging("1"))
	fmt.Println("Execute Console Command:", ExeConsoleCmd("restartadmin")) // restartadmin FileRestartAdmin
	fmt.Printf("Memory Base: %v\n", MemoryBase("0x7FF12345"))
	fmt.Printf("Module List: %v\n", GetModuleList())
}

const (
	DefaultX64dbgServer = "http://127.0.0.1:8888"
)

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
		s := strings.ReplaceAll(string(bodyBytes), `\`, `\\`)
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

func IsDebugging(Dbgcheck string) string {
	return safeGet("IsDebugActive", map[string]string{"dbgcheck": Dbgcheck})[0]
}

func ExeConsoleCmd(Command string) []string {
	return safeGet("ExeConsoleCmd", map[string]string{"Command": Command})
}

func MemoryBase(addr string) map[string]any {
	fmt.Printf("Calling MemoryBase with address: %s\n", addr)

	response := safeGet("FindMemBase", map[string]string{"addr": addr})
	fmt.Printf("Raw response: '%s'\n", response)

	if len(response) == 0 {
		fmt.Println("No response received")
		return map[string]any{"error": "No response received"}
	}

	if len(response) > 0 && strings.Contains(response[0], ",") {
		parts := strings.Split(response[0], ",")
		fmt.Printf("Split parts: %v\n", parts)
		baseAddress := parts[0]
		size, err := strconv.Atoi(parts[1])
		if err != nil {
			return map[string]any{"error": err.Error()}
		}
		return map[string]any{"base_address": baseAddress, "size": size}
	} else {
		fmt.Println("Response format not as expected")
		return map[string]any{"raw_response": response}
	}
}

func GetModuleList() any {
	response := safeGet("GetModuleList", map[string]string{})
	if len(response) == 0 {
		return nil
	}

	var moduleList []map[string]any
	err := json.Unmarshal([]byte(response[0]), &moduleList)
	if err != nil {
		return map[string]any{"error": fmt.Sprintf("Failed to parse response: %s", err.Error()), "raw_response": response}
	}

	return moduleList
}
