package repl

import "bufio"
import "fmt"
import "os"
import "golox/interpreter"

func RunPrompt() {
	reader := bufio.NewReader(os.Stdin)

	for true {
		fmt.Print("> ")

		line, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
			os.Exit(65)
		}

		if line == "\n" {
			break
		}

		interpreter.Run(line)
	}
}
