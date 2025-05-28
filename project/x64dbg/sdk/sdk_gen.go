package sdk

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

type Sdk struct {
	Scriptapi   *scriptapi
	Argument    *argument
	Assembler   *assembler
	Bookmark    *bookmark
	Comment     *comment
	Debug       *debug
	Flag        *flag
	Function    *function
	Gui         *gui
	Label       *label
	Memory      *memory
	Misc        *misc
	Module      *module
	Pattern     *pattern
	Register    *register
	Stack       *stack
	Symbol      *symbol
	Bridgegraph *bridgegraph
	Bridgelist  *bridgelist
	Bridgemain  *bridgemain
}

func New() *Sdk {
	return &Sdk{
		Scriptapi:   new(scriptapi),
		Argument:    new(argument),
		Assembler:   new(assembler),
		Bookmark:    new(bookmark),
		Comment:     new(comment),
		Debug:       new(debug),
		Flag:        new(flag),
		Function:    new(function),
		Gui:         new(gui),
		Label:       new(label),
		Memory:      new(memory),
		Misc:        new(misc),
		Module:      new(module),
		Pattern:     new(pattern),
		Register:    new(register),
		Stack:       new(stack),
		Symbol:      new(symbol),
		Bridgegraph: new(bridgegraph),
		Bridgelist:  new(bridgelist),
		Bridgemain:  new(bridgemain),
	}
}

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
