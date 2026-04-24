// Example for https://gitlab.com/cznic/cc/-/issues/163#note_2258256336

package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"modernc.org/cc/v4"
)

const src = `
struct inner {
	char flags;
	int  dlen;
};

struct test {
	const char * str;
	short stats;
	struct inner meta;
	short inlen;
	int len;
};
`

const predefined = `
int __predefined_declarator;

#if defined(__i386__) || defined(__arm__)
typedef unsigned __predefined_size_t;
#else
typedef unsigned long long __predefined_size_t;
#endif
`

func fail(rc int, s string, args ...any) {
	fmt.Fprintf(os.Stderr, s, args...)
	os.Exit(rc)
}

func main() {
	goarch := flag.String("arch", runtime.GOARCH, "target GOARCH")
	goos := flag.String("os", runtime.GOOS, "target GOOS")
	flag.Parse()
	abi, err := cc.NewABI(*goos, *goarch)
	if err != nil {
		fail(2, "%v\n", err)
	}

	ast, err := cc.Translate(
		&cc.Config{ABI: abi},
		[]cc.Source{
			{Name: "<predefined>", Value: predefined},
			{Name: "struct.c", Value: src},
		},
	)
	if err != nil {
		fail(1, "%v\n", err)
	}

	for l := ast.TranslationUnit; l != nil; l = l.TranslationUnit {
		ed := l.ExternalDeclaration
		switch ed.Case {
		case cc.ExternalDeclarationDecl:
			switch t := ed.Declaration.DeclarationSpecifiers.Type().(type) {
			case *cc.StructType:
				fmt.Printf("==== %s align=%v size=%v padding=%v\n", t, t.Align(), t.Size(), t.Padding())
				for i := 0; i < t.NumFields(); i++ {
					f := t.FieldByIndex(i)
					ft := f.Type()
					fmt.Printf("#%d: %q ofs=%v align=%v size=%v type=%s\n", i, f.Name(), f.Offset(), ft.Align(), ft.Size(), ft)
				}
			}
		}
	}
}
