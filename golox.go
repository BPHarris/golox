package main

import "fmt"
import "golox/interpreter"
import "golox/repl"
import "os"

func main() {
	args := os.Args[1:]

	if len(args) > 1 {
		fmt.Println("Usage: golox [script]")
		os.Exit(64)
	} else if len(args) == 1 {
		interpreter.RunFile(args[0])
		os.Exit(0)
	} else {
		repl.RunPrompt()
		os.Exit(0)
	}
}
