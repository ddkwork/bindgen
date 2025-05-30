package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
)

type stack struct{}

func (s *stack) Pop() {
	Client.Post().Url("http://localhost:8888/_scriptapi_stack.h/Pop").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (s *stack) Push(value uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_stack.h/Push").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (s *stack) Peek(offset int) {
	Client.Post().Url("http://localhost:8888/_scriptapi_stack.h/Peek").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "offset",
				Type:  "int",
				Value: fmt.Sprintf("%v", offset),
			},
		},
	))).Request()
	// todo handle response into result
}
