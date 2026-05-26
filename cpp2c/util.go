package cpp2c

import (
	"encoding/json"
	"strconv"
	"strings"
)

func maybeSuffix(counter int) string {
	if counter == 0 {
		return ""
	}

	return strconv.Itoa(counter + 1)
}

func titleCase(s string) string {
	return strings.ToUpper(s[0:1]) + s[1:]
}

func prettyPrint(obj interface{}) {
	jb, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		panic(err)
	}

	_ = string(jb)
}

func ifv[T any](condition bool, trueval T, falseval T) T {
	if condition {
		return trueval
	}
	return falseval
}

func Addr[T any](s T) *T {
	return &s
}

func MsvcIncludeFlags() string {
	msvcInc := `"E:/Program Files/Microsoft Visual Studio/2022/BuildTools/VC/Tools/MSVC/14.44.35207/include"`
	ucrtInc := `"E:/Program Files/Windows Kits/10/Include/10.0.28000.0/ucrt"`
	sharedInc := `"E:/Program Files/Windows Kits/10/Include/10.0.28000.0/shared"`
	umInc := `"E:/Program Files/Windows Kits/10/Include/10.0.28000.0/um"`
	return " -isystem " + msvcInc + " -isystem " + ucrtInc + " -isystem " + sharedInc + " -isystem " + umInc
}

func slice_to_set[T comparable](list []T) map[T]struct{} {
	ret := make(map[T]struct{}, len(list))
	for _, v := range list {
		ret[v] = struct{}{}
	}
	return ret
}
