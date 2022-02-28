package loxerror

import (
	"fmt"
)

// Global error state
var has_had_error = false

// Return true if an error has been reached while running the lox source.
func HasHadError() bool {
	return has_had_error
}

func Error(et ErrorType, line int, message string) {
	Report(et, line, "", message)
}

func Report(et ErrorType, line int, where string, message string) {
	has_had_error = true

	if where != "" {
		where = fmt.Sprintf(" (%s)", where)
	}

	const base = "%s%s: %d: %s"

	fmt.Printf("%s%s: line %d: %s\n", et, where, line, message)
}
