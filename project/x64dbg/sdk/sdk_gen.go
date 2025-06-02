package sdk

import (
	"encoding/json"
	"github.com/ddkwork/golibrary/stream/net/httpClient"
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

type Param struct {
	Name  string
	Type  string
	Value any
}

type ApiResponse struct {
	Success bool
	Type    string
	Result  json.RawMessage
}

var Client = httpClient.New()
