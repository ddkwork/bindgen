package libdemo

import (
	"testing"

	"github.com/ddkwork/bindgen/dll/clang"
	"github.com/ddkwork/bindgen/dll/gengo"
	"github.com/ddkwork/golibrary/mylog"
)

func TestDemoDll(t *testing.T) {
	pkg := gengo.NewPackage("bug")
	path := "1.h"
	mylog.Check(pkg.Transform("bug", &clang.Options{
		Sources:          []string{path},
		AdditionalParams: []string{},
	}),
	)
	mylog.Check(pkg.WriteToDir("./tmp"))
}
