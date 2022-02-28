package lexer

import "fmt"
import e "golox/errors"

// Store the lexer state.
type Lexer struct {
	source  string
	lexemes []Lexeme

	start   int
	current int
	line    int
}

// Lex returns the list of tokens in the given lox source code.
func Lex(source string) []Lexeme {
	lexer := Lexer{
		source:  source,
		lexemes: make([]Lexeme, 0),
		start:   0,
		current: 0,
		line:    1,
	}

	for !lexer.IsAtEnd() {
		lexer.ConsumeLexeme()
	}

	return lexer.lexemes
}

// Return true if the lexer has reached the end of the file, false otherwise.
func (l Lexer) IsAtEnd() bool {
	return l.current >= len(l.source)
}

// Consume the next lexeme and update the lexer state accordingly.
func (l Lexer) ConsumeLexeme() {
	c := l.Advance()

	switch c {
	case '(':
		l.AddLexeme(LeftParenthesis)
	case ')':
		l.AddLexeme(RightParenthesis)
	case '{':
		l.AddLexeme(LeftBrace)
	case '}':
		l.AddLexeme(RightBrace)
	case ',':
		l.AddLexeme(Comma)
	case '.':
		l.AddLexeme(Dot)
	case '-':
		l.AddLexeme(Minus)
	case '+':
		l.AddLexeme(Plus)
	case ';':
		l.AddLexeme(Semicolon)
	case '*':
		l.AddLexeme(Star)
	case '!':
		l.AddLexemeWithLookAhead('=', BangEqual, Bang)
	case '=':
		l.AddLexemeWithLookAhead('=', EqualEqual, Equal)
	case '<':
		l.AddLexemeWithLookAhead('=', LessEqual, Less)
	case '>':
		l.AddLexemeWithLookAhead('=', GreaterEqual, Greater)
	case '/':
		if l.Match('/') {
			l.ConsumeComment()
		} else if l.Match('*') {
			l.ConsumeMultiLineComment()
		} else {
			l.AddLexeme(Slash)
		}
	case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
		'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '_':
		l.AddIdentifier()
	case '"':
		l.AddLiteralString()
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		l.AddLiteralNumber()
	case ' ', '\r', '\t':
		break
	case '\n':
		l.line++
	default:
		e.Error(e.SyntaxError, l.line, fmt.Sprintf("Unexpected character '%c'", c))
	}

	l.start = l.current
}

// Consume characters until the end of the started line comment.
func (l Lexer) ConsumeComment() {
	for l.LookAhead() != '\n' && !l.IsAtEnd() {
		l.Advance()
	}
}

// Consume multi-line comment.
func (l Lexer) ConsumeMultiLineComment() {
	for l.LookAhead() != '*' && l.LookAheadNext() != '/' && !l.IsAtEnd() {
		if l.LookAhead() == '\n' {
			l.line++
		}
		l.Advance()
	}

	if l.IsAtEnd() {
		e.Error(e.SyntaxError, l.line, "Unterminated multi-line comment.")
		return
	}

	// Consume "*/"
	l.Advance()
	l.Advance()
}

// Return the current character and advance the lexer by one.
func (l Lexer) Advance() byte {
	char := l.source[l.current]

	l.current++

	return char
}

// Add a new Lexeme of the given type to the lexer.
func (l Lexer) AddLexeme(lexeme_type LexemeType) {
	l.lexemes = append(l.lexemes, Lexeme{lexeme_type, l.source[l.start:l.current], l.line})
}

// If the current character matches the expected then add the first given
// lexeme, otherwise add the second given lexeme.
// Consume the expected character if seen, see Lexer.Match().
func (l Lexer) AddLexemeWithLookAhead(expected byte, on_match LexemeType, otherwise LexemeType) {
	if l.Match(expected) {
		l.AddLexeme(on_match)
	} else {
		l.AddLexeme(otherwise)
	}
}

// Add a new Identifier Lexeme to the lexer.
func (l Lexer) AddIdentifier() {
	for IsAlphanumeric(l.LookAhead()) {
		l.Advance()
	}

	matched := l.source[l.start:l.current]
	lexeme_type, is_keyword := keywords[matched]

	if !is_keyword {
		lexeme_type = Identifier
	}

	l.AddLexeme(lexeme_type)
}

// Add a new LiteralString Lexeme to the lexer.
func (l Lexer) AddLiteralString() {
	for l.LookAhead() != '"' && !l.IsAtEnd() {
		if l.LookAhead() == '\n' {
			l.line++
		}
		l.Advance()
	}

	if l.IsAtEnd() {
		e.Error(e.SyntaxError, l.line, "Unterminated string.")
		return
	}

	// Consume closing quote
	l.Advance()

	l.AddLexeme(LiteralString)
}

// Add a new LiteralNumber Lexeme to the lexer.
func (l Lexer) AddLiteralNumber() {
	for IsDigit(l.LookAhead()) {
		l.Advance()
	}

	if l.LookAhead() == '.' && IsDigit(l.LookAheadNext()) {
		// Consume dot
		l.Advance()

		for IsDigit(l.LookAhead()) {
			l.Advance()
		}
	}

	l.AddLexeme(LiteralNumber)
}

// Return the next character without consuming it.
func (l Lexer) LookAhead() byte {
	if l.IsAtEnd() {
		return 0
	} else {
		return l.source[l.current]
	}
}

// Return the character after the next without consuming either.
func (l Lexer) LookAheadNext() byte {
	if l.current+1 >= len(l.source) {
		return 0
	} else {
		return l.source[l.current+1]
	}
}

// Return true if the current character matches the expected character, false
// otherwise.
// If the character matches then it is consumed.
func (l Lexer) Match(expected byte) bool {
	if l.IsAtEnd() {
		return false
	}

	if l.source[l.current] == expected {
		l.current++
		return true
	} else {
		return false
	}
}

// Return true if the given character is a digit.
func IsDigit(b byte) bool {
	return '0' <= b && b <= '9'
}

// Return true if the given character is a letter or an underscore.
func IsAlpha(b byte) bool {
	return ('a' <= b && b <= 'z') || ('A' <= b && b <= 'Z') || (b == '_')
}

// Return true if the given character is a letter, an underscore, or a digit.
func IsAlphanumeric(b byte) bool {
	return IsAlpha(b) || IsDigit(b)
}
