// Example for https://gitlab.com/cznic/cc/-/issues/163#note_2258256336

package main

import (
	"fmt"
	"os"

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
`

func main() {
	ast, err := cc.Parse(
		&cc.Config{},
		[]cc.Source{
			{Name: "<predefined>", Value: predefined},
			{Name: "struct.c", Value: src},
		},
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(ast.TranslationUnit)
}
