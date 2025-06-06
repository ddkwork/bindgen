package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/ddkwork/golibrary/std/mylog"
)

type pattern struct{}

// Find    c api name: Script::Pattern::Find
// ┌────┬──────────────────────┬──────────────────────────────────────────────────────────────┬─────────┐
// │ id │         name         │                            c type                            │ go type │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 0  │ data                 │ unsigned char *                                              │ *byte   │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 1  │ datasize             │ duint                                                        │ uint    │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 2  │ pattern              │ const char *                                                 │ *int8   │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ return               │ duint                                                        │ uint    │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\_scriptapi_pattern.h                                 │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.name         │ Find                                                         │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.mangledName  │ ?Find@Pattern@Script@@YA_KPEAE_KPEBD@Z                       │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.file         │                                                              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.includedFrom │                                                              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.expansionLoc │ D:\workspace\workspace\bindgen\project\x64dbg\include\_scri- │         │
// │    │                      │ ptapi_pattern.h                                              │         │
// └────┴──────────────────────┴──────────────────────────────────────────────────────────────┴─────────┘
func (p *pattern) Find(data *byte, datasize uint, pattern *int8) uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_pattern.h/Find").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "data",
				Type:  "*byte ",
				Value: fmt.Sprintf("%v", data),
			},
			{
				Name:  "datasize",
				Type:  "uint",
				Value: fmt.Sprintf("%v", datasize),
			},
			{
				Name:  "pattern",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", pattern),
			},
		},
	))).Request()
	panic("not support return type: duint")
}

// FindMem    c api name: Script::Pattern::FindMem
// ┌────┬──────────────────────┬──────────────────────────────────────────────────────────────┬─────────┐
// │ id │         name         │                            c type                            │ go type │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 0  │ start                │ duint                                                        │ uint    │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 1  │ size                 │ duint                                                        │ uint    │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 2  │ pattern              │ const char *                                                 │ *int8   │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ return               │ duint                                                        │ uint    │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\_scriptapi_pattern.h                                 │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.name         │ FindMem                                                      │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.mangledName  │ ?FindMem@Pattern@Script@@YA_K_K0PEBD@Z                       │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.file         │                                                              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.includedFrom │                                                              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.expansionLoc │ D:\workspace\workspace\bindgen\project\x64dbg\include\_scri- │         │
// │    │                      │ ptapi_pattern.h                                              │         │
// └────┴──────────────────────┴──────────────────────────────────────────────────────────────┴─────────┘
func (p *pattern) FindMem(start uint, size uint, pattern *int8) uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_pattern.h/FindMem").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "start",
				Type:  "uint",
				Value: fmt.Sprintf("%v", start),
			},
			{
				Name:  "size",
				Type:  "uint",
				Value: fmt.Sprintf("%v", size),
			},
			{
				Name:  "pattern",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", pattern),
			},
		},
	))).Request()
	panic("not support return type: duint")
}

// Write    c api name: Script::Pattern::Write
// ┌────┬──────────────────────┬──────────────────────────────────────────────────────────────┬─────────┐
// │ id │         name         │                            c type                            │ go type │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 0  │ data                 │ unsigned char *                                              │ *byte   │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 1  │ datasize             │ duint                                                        │ uint    │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 2  │ pattern              │ const char *                                                 │ *int8   │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ return               │ void                                                         │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\_scriptapi_pattern.h                                 │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.name         │ Write                                                        │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.mangledName  │ ?Write@Pattern@Script@@YAXPEAE_KPEBD@Z                       │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.file         │                                                              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.includedFrom │                                                              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.expansionLoc │ D:\workspace\workspace\bindgen\project\x64dbg\include\_scri- │         │
// │    │                      │ ptapi_pattern.h                                              │         │
// └────┴──────────────────────┴──────────────────────────────────────────────────────────────┴─────────┘
func (p *pattern) Write(data *byte, datasize uint, pattern *int8) {
	Client.Post().Url("http://localhost:8888/_scriptapi_pattern.h/Write").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "data",
				Type:  "*byte ",
				Value: fmt.Sprintf("%v", data),
			},
			{
				Name:  "datasize",
				Type:  "uint",
				Value: fmt.Sprintf("%v", datasize),
			},
			{
				Name:  "pattern",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", pattern),
			},
		},
	))).Request()
}

// WriteMem    c api name: Script::Pattern::WriteMem
// ┌────┬──────────────────────┬──────────────────────────────────────────────────────────────┬─────────┐
// │ id │         name         │                            c type                            │ go type │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 0  │ start                │ duint                                                        │ uint    │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 1  │ size                 │ duint                                                        │ uint    │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 2  │ pattern              │ const char *                                                 │ *int8   │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ return               │ void                                                         │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\_scriptapi_pattern.h                                 │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.name         │ WriteMem                                                     │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.mangledName  │ ?WriteMem@Pattern@Script@@YAX_K0PEBD@Z                       │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.file         │                                                              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.includedFrom │                                                              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.expansionLoc │ D:\workspace\workspace\bindgen\project\x64dbg\include\_scri- │         │
// │    │                      │ ptapi_pattern.h                                              │         │
// └────┴──────────────────────┴──────────────────────────────────────────────────────────────┴─────────┘
func (p *pattern) WriteMem(start uint, size uint, pattern *int8) {
	Client.Post().Url("http://localhost:8888/_scriptapi_pattern.h/WriteMem").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "start",
				Type:  "uint",
				Value: fmt.Sprintf("%v", start),
			},
			{
				Name:  "size",
				Type:  "uint",
				Value: fmt.Sprintf("%v", size),
			},
			{
				Name:  "pattern",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", pattern),
			},
		},
	))).Request()
}

// SearchAndReplace    c api name: Script::Pattern::SearchAndReplace
// ┌────┬──────────────────────┬──────────────────────────────────────────────────────────────┬─────────┐
// │ id │         name         │                            c type                            │ go type │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 0  │ data                 │ unsigned char *                                              │ *byte   │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 1  │ datasize             │ duint                                                        │ uint    │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 2  │ searchpattern        │ const char *                                                 │ *int8   │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 3  │ replacepattern       │ const char *                                                 │ *int8   │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ return               │ bool                                                         │ bool    │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\_scriptapi_pattern.h                                 │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.name         │ SearchAndReplace                                             │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.mangledName  │ ?SearchAndReplace@Pattern@Script@@YA_NPEAE_KPEBD2@Z          │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.file         │                                                              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.includedFrom │                                                              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.expansionLoc │ D:\workspace\workspace\bindgen\project\x64dbg\include\_scri- │         │
// │    │                      │ ptapi_pattern.h                                              │         │
// └────┴──────────────────────┴──────────────────────────────────────────────────────────────┴─────────┘
func (p *pattern) SearchAndReplace(data *byte, datasize uint, searchpattern *int8, replacepattern *int8) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_pattern.h/SearchAndReplace").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "data",
				Type:  "*byte ",
				Value: fmt.Sprintf("%v", data),
			},
			{
				Name:  "datasize",
				Type:  "uint",
				Value: fmt.Sprintf("%v", datasize),
			},
			{
				Name:  "searchpattern",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", searchpattern),
			},
			{
				Name:  "replacepattern",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", replacepattern),
			},
		},
	))).Request()
	return true
}

// SearchAndReplaceMem    c api name: Script::Pattern::SearchAndReplaceMem
// ┌────┬──────────────────────┬──────────────────────────────────────────────────────────────┬─────────┐
// │ id │         name         │                            c type                            │ go type │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 0  │ start                │ duint                                                        │ uint    │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 1  │ size                 │ duint                                                        │ uint    │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 2  │ searchpattern        │ const char *                                                 │ *int8   │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 3  │ replacepattern       │ const char *                                                 │ *int8   │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ return               │ bool                                                         │ bool    │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\_scriptapi_pattern.h                                 │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.name         │ SearchAndReplaceMem                                          │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.mangledName  │ ?SearchAndReplaceMem@Pattern@Script@@YA_N_K0PEBD1@Z          │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.file         │                                                              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.includedFrom │                                                              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.expansionLoc │ D:\workspace\workspace\bindgen\project\x64dbg\include\_scri- │         │
// │    │                      │ ptapi_pattern.h                                              │         │
// └────┴──────────────────────┴──────────────────────────────────────────────────────────────┴─────────┘
func (p *pattern) SearchAndReplaceMem(start uint, size uint, searchpattern *int8, replacepattern *int8) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_pattern.h/SearchAndReplaceMem").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "start",
				Type:  "uint",
				Value: fmt.Sprintf("%v", start),
			},
			{
				Name:  "size",
				Type:  "uint",
				Value: fmt.Sprintf("%v", size),
			},
			{
				Name:  "searchpattern",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", searchpattern),
			},
			{
				Name:  "replacepattern",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", replacepattern),
			},
		},
	))).Request()
	return true
}
