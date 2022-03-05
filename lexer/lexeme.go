package lexer

import "fmt"

// Lexeme stores the information for a lexeme.
// Lexeme.lexeme:
//   Identifier    => the name as is
//   LiteralString => the string literal enclosed in quotes
//   LiteralNumber => the number as is, float64
//   Otherwise     => as expected
type Lexeme struct {
	lexeme_type LexemeType
	lexeme      string
	line        int
}

// Return a lexeme's literal value as a string, or the empty string if the
// lexeme is non-literal.
func (l Lexeme) Literal() string {
	switch l.lexeme_type {
	case Identifier, LiteralString, LiteralNumber:
		return l.lexeme
	default:
		return ""
	}
}

// Cast a lexeme to a string.
func (l Lexeme) String() string {
	const base = "Lexeme(type=%s, lexeme=\"%s\")"
	const literal = "Lexeme(type=%s, lexeme=\"%s\", literal=%s)"

	switch l.lexeme_type {
	case Identifier, LiteralString, LiteralNumber:
		return fmt.Sprintf(literal, l.lexeme_type, l.lexeme, l.Literal())
	default:
		return fmt.Sprintf(base, l.lexeme_type, l.lexeme)
	}
}
