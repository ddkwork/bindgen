package sdk

import (
	"encoding/json"
	"github.com/ddkwork/golibrary/mylog"
)

type pattern struct{}

func (p *pattern) Find(data *byte, datasize uint, pattern string) {
	Client.Post().Url("http://localhost:8888/_scriptapi_pattern.h/Find").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "data",
				Type:  "*byte ",
				Value: data,
			},
			Param{
				Name:  "datasize",
				Type:  "uint",
				Value: datasize,
			},
			Param{
				Name:  "pattern",
				Type:  "string",
				Value: pattern,
			},
		},
	))).Request()
	// todo handle response into result
}
func (p *pattern) FindMem(start uint, size uint, pattern string) {
	Client.Post().Url("http://localhost:8888/_scriptapi_pattern.h/FindMem").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "start",
				Type:  "uint",
				Value: start,
			},
			Param{
				Name:  "size",
				Type:  "uint",
				Value: size,
			},
			Param{
				Name:  "pattern",
				Type:  "string",
				Value: pattern,
			},
		},
	))).Request()
	// todo handle response into result
}
func (p *pattern) Write(data *byte, datasize uint, pattern string) {
	Client.Post().Url("http://localhost:8888/_scriptapi_pattern.h/Write").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "data",
				Type:  "*byte ",
				Value: data,
			},
			Param{
				Name:  "datasize",
				Type:  "uint",
				Value: datasize,
			},
			Param{
				Name:  "pattern",
				Type:  "string",
				Value: pattern,
			},
		},
	))).Request()
	// todo handle response into result
}
func (p *pattern) WriteMem(start uint, size uint, pattern string) {
	Client.Post().Url("http://localhost:8888/_scriptapi_pattern.h/WriteMem").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "start",
				Type:  "uint",
				Value: start,
			},
			Param{
				Name:  "size",
				Type:  "uint",
				Value: size,
			},
			Param{
				Name:  "pattern",
				Type:  "string",
				Value: pattern,
			},
		},
	))).Request()
	// todo handle response into result
}
func (p *pattern) SearchAndReplace(data *byte, datasize uint, searchpattern string, replacepattern string) {
	Client.Post().Url("http://localhost:8888/_scriptapi_pattern.h/SearchAndReplace").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "data",
				Type:  "*byte ",
				Value: data,
			},
			Param{
				Name:  "datasize",
				Type:  "uint",
				Value: datasize,
			},
			Param{
				Name:  "searchpattern",
				Type:  "string",
				Value: searchpattern,
			},
			Param{
				Name:  "replacepattern",
				Type:  "string",
				Value: replacepattern,
			},
		},
	))).Request()
	// todo handle response into result
}
func (p *pattern) SearchAndReplaceMem(start uint, size uint, searchpattern string, replacepattern string) {
	Client.Post().Url("http://localhost:8888/_scriptapi_pattern.h/SearchAndReplaceMem").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "start",
				Type:  "uint",
				Value: start,
			},
			Param{
				Name:  "size",
				Type:  "uint",
				Value: size,
			},
			Param{
				Name:  "searchpattern",
				Type:  "string",
				Value: searchpattern,
			},
			Param{
				Name:  "replacepattern",
				Type:  "string",
				Value: replacepattern,
			},
		},
	))).Request()
	// todo handle response into result
}
