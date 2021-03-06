// Package interpreter provides the interpreter for golox.
package interpreter

import "fmt"
import "golox/lexer"
import "golox/errors"
import "io/ioutil"
import "os"

// RunFile reads, parses, and excecutes the source from the given file.
func RunFile(path string) {
	source, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println(err)
		os.Exit(74)
	}

	Run(string(source))

	if errors.HasHadError() {
		os.Exit(65)
	}
}

// Run lexes, parses, and excecutes the given source code.
func Run(source string) {
	lexemes := lexer.Lex(source)

	for _, lexeme := range lexemes {
		fmt.Println(lexeme)
	}
}
