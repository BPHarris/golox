package lexer

import "errors"
import "fmt"
import "strconv"

// Lexeme stores the information for a lexeme.
// Lexeme.lexeme:
//   Identifier    => the name as is
//   LiteralString => the string literal enclosed in quotes
//   LiteralNumber => the number as is, float64 as a string
//   Otherwise     => as expected
type Lexeme struct {
	lexeme_type LexemeType
	lexeme      string
	line        int
}

// Return a lexeme's literal value as a string.
// If the given lexeme is not is not a Lox literal an error is returned.
func (l Lexeme) Literal() (string, error) {
	switch l.lexeme_type {
	case Identifier, LiteralString, LiteralNumber:
		return l.lexeme, nil
	default:
		return "", errors.New(fmt.Sprintf("lexeme %s is not a literal", l))
	}
}

// Return the given lexeme's literal value without the enclosing "".
// If the given lexeme is not a LiteralString an error is returned.
// If the given LiteralString is malformed an error is returned.
func (l Lexeme) ParseString() (string, error) {
	if l.lexeme_type != LiteralString {
		msg := fmt.Sprintf(
			"lexeme of type '%s' has no string literal",
			l.lexeme_type.String(),
		)
		return "", errors.New(msg)
	}

	literal, _ := l.Literal()

	if len(literal) < 2 || literal[0] != '"' || literal[len(literal)-1] != '"' {
		return "", errors.New("lexeme is malformed, no enclosing \"")
	}

	return literal[1 : len(literal)-1], nil
}

// Return the given lexeme's literal value as a float64.
// If the given lexeme is not a LiteralNumber an error is returned.
// If the given lexeme fails strconv.ParseFloat the error is propogated.
func (l Lexeme) ParseFloat() (float64, error) {
	if l.lexeme_type != LiteralNumber {
		msg := fmt.Sprintf(
			"lexeme of type '%s' has no numeric literal",
			l.lexeme_type.String(),
		)
		return 0.0, errors.New(msg)
	}

	literal, _ := l.Literal()

	return strconv.ParseFloat(literal, 64)
}

// Cast the given lexeme to a string.
func (l Lexeme) String() string {
	const base = "Lexeme(type=%s, lexeme=\"%s\")"
	const literal = "Lexeme(type=%s, lexeme=\"%s\", literal=%s)"

	switch l.lexeme_type {
	case Identifier, LiteralString, LiteralNumber:
		val, err := l.Literal()

		if err != nil {
			val = "!error"
		}

		// If the lexeme is a LiteralString escape the enclosing ""
		var lexeme_string string
		if l.lexeme_type == LiteralString {
			lexeme_string = "\\\"" + l.lexeme[1:len(l.lexeme)-1] + "\\\""
		} else {
			lexeme_string = l.lexeme
		}

		return fmt.Sprintf(literal, l.lexeme_type, lexeme_string, val)
	default:
		return fmt.Sprintf(base, l.lexeme_type, l.lexeme)
	}
}
