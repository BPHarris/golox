package errors

type ErrorType int64

const (
	LoxError ErrorType = iota

	// Lexer errors
	SyntaxError

	// Parser errors
	// ...

	// Run-time errors
	// ...
)

func (e ErrorType) String() string {
	switch e {
	case LoxError:
		return "LoxError"
	case SyntaxError:
		return "SyntaxError"
	default:
		return "UndefinedError"
	}
}
