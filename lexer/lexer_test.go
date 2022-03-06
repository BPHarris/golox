package lexer

import "fmt"
import "testing"

// TODO end-to-end style tests
// TODO test Lex
// TODO test consume lexeme
// TODO test consume comment
// TODO test consume multi-line comment
// TODO consume Add*

const ascii_lower = "abcdefghijklmnopqrstuvwxyz"
const ascii_upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const ascii_digits = "0123456789"
const ascii_special = "!\"£$%^&*()_+-=[]{};'#:@~,./<>?'"
const ascii_whitespace = " \n\t"

func Test_IsAtEnd(t *testing.T) {
	f := "Lexer{'%s', start=%d, current=%d}"
	no_lexemes := make([]Lexeme, 0)

	cases := map[*Lexer]bool {
		// The empty string
		{"", no_lexemes, 0, 0, 1}: true,
		{"", no_lexemes, 0, 1, 1}: true,
		{"", no_lexemes, 1, 1, 1}: true,
		// Full lexeme
		{"hello", no_lexemes, 0, 0, 1}: false,
		{"hello", no_lexemes, 0, 1, 1}: false,
		{"hello", no_lexemes, 0, 2, 1}: false,
		{"hello", no_lexemes, 0, 3, 1}: false,
		{"hello", no_lexemes, 0, 4, 1}: false,
		{"hello", no_lexemes, 0, 5, 1}: true,
	}

	for l, expected := range cases {
		if l.IsAtEnd() != expected {
			t.Logf(
				"%s.IsAtEnd() expects '%t' received '%t'",
				fmt.Sprintf(f, l.source, l.start, l.current),
				expected,
				l.IsAtEnd(),
			)
			t.Fail()
		}
	}
}

func Test_Advance(t *testing.T) {
	f := "Lexer{'%s', start=%d, current=%d}"
	no_lexemes := make([]Lexeme, 0)

	cases := map[*Lexer]byte {
		// Full lexeme
		{"hello", no_lexemes, 0, 0, 1}: 'h',
		{"hello", no_lexemes, 0, 1, 1}: 'e',
		{"hello", no_lexemes, 0, 2, 1}: 'l',
		{"hello", no_lexemes, 0, 3, 1}: 'l',
		{"hello", no_lexemes, 0, 4, 1}: 'o',
	}

	for l, expected := range cases {
		previous := l.current
		next_char := l.Advance()

		if next_char != expected {
			t.Logf(
				"%s.Advance() expects '%c' received '%c'",
				fmt.Sprintf(f, l.source, l.start, previous),
				expected,
				next_char,
			)
			t.Fail()
		}

		if l.current != previous + 1 {
			t.Logf(
				"(lexer).Advance() expected to increment current %d -> %d, got %d",
				previous,
				previous + 1,
				l.current,
			)
			t.Fail()
		}
	}
}

func Test_LookAhead(t *testing.T) {
	f := "Lexer{'%s', start=%d, current=%d}"
	no_lexemes := make([]Lexeme, 0)

	cases := map[*Lexer]byte {
		// The empty string
		{"", no_lexemes, 0, 0, 1}: 0,
		{"", no_lexemes, 0, 1, 1}: 0,
		{"", no_lexemes, 1, 1, 1}: 0,
		// Full lexeme
		{"hello", no_lexemes, 0, 0, 1}: 'h',
		{"hello", no_lexemes, 0, 1, 1}: 'e',
		{"hello", no_lexemes, 0, 2, 1}: 'l',
		{"hello", no_lexemes, 0, 3, 1}: 'l',
		{"hello", no_lexemes, 0, 4, 1}: 'o',
		{"hello", no_lexemes, 0, 5, 1}: 0,
	}

	for l, expected := range cases {
		if l.LookAhead() != expected {
			t.Logf(
				"%s.LookAhead() expects '%c' received '%c'",
				fmt.Sprintf(f, l.source, l.start, l.current),
				expected,
				l.LookAhead(),
			)
			t.Fail()
		}
	}
}

func Test_LookAheadNext(t *testing.T) {
	f := "Lexer{'%s', start=%d, current=%d}"
	no_lexemes := make([]Lexeme, 0)

	cases := map[*Lexer]byte {
		// The empty string
		{"", no_lexemes, 0, 0, 1}: 0,
		{"", no_lexemes, 0, 1, 1}: 0,
		{"", no_lexemes, 1, 1, 1}: 0,
		// Full lexeme
		{"hello", no_lexemes, 0, 0, 1}: 'e',
		{"hello", no_lexemes, 0, 1, 1}: 'l',
		{"hello", no_lexemes, 0, 2, 1}: 'l',
		{"hello", no_lexemes, 0, 3, 1}: 'o',
		{"hello", no_lexemes, 0, 4, 1}: 0,
	}

	for l, expected := range cases {
		if l.LookAheadNext() != expected {
			t.Logf(
				"%s.LookAhead() expects '%c' received '%c'",
				fmt.Sprintf(f, l.source, l.start, l.current),
				expected,
				l.LookAheadNext(),
			)
			t.Fail()
		}
	}
}

func Test_Match(t *testing.T) {
	f := "Lexer{'%s', start=%d, current=%d}"
	all_chars := ascii_lower + ascii_upper + ascii_digits + ascii_special + ascii_whitespace
	no_lexemes := make([]Lexeme, 0)

	// Neative cases -- at end
	negative_cases := []Lexer{
		{"", no_lexemes, 0, 0, 1},
		{"", no_lexemes, 0, 0, 1},
		{"if", no_lexemes, 0, 2, 1},
		{"if", no_lexemes, 1, 2, 1},
		{"if", no_lexemes, 2, 2, 1},
	}

	for _, lexer := range negative_cases {
		for _, char := range all_chars {
			if lexer.Match(byte(char)) {
				t.Logf(
					"%s.Match('%c') expects false",
					fmt.Sprintf(f, lexer.source, lexer.start, lexer.current),
					char,
				)
				t.Fail()
			}
		}
	}

	// Test matches and non-matches
	complex_source := "if foo == 4{\n\tprint \"hi\"\n}"

	positive_cases := map[*Lexer]byte{
		// Simple -- keyword
		{"if", no_lexemes, 0, 1, 1}:  'f',
		{"if", no_lexemes, 1, 1, 1}:  'f',
		{"if", no_lexemes, 99, 1, 1}: 'f',
		// Complex -- keyword
		{complex_source, no_lexemes, 0, 0, 1}: 'i',
		{complex_source, no_lexemes, 0, 1, 1}: 'f',
		// Complex -- identifier
		{complex_source, no_lexemes, 3, 3, 1}: 'f',
		{complex_source, no_lexemes, 3, 4, 1}: 'o',
		{complex_source, no_lexemes, 3, 5, 1}: 'o',
		// Complex -- match EqualEqual
		{complex_source, no_lexemes, 7, 7, 1}: '=',
		{complex_source, no_lexemes, 7, 8, 1}: '=',
		// Complex -- literal
		{complex_source, no_lexemes, 10, 10, 1}: '4',
		// Complex -- closing }
		{complex_source, no_lexemes, 25, 25, 1}: '}',
		// Complex -- whitespace
		{complex_source, no_lexemes, 2, 2, 1}:   ' ',
		{complex_source, no_lexemes, 6, 6, 1}:   ' ',
		{complex_source, no_lexemes, 12, 12, 1}: '\n',
		{complex_source, no_lexemes, 13, 13, 1}: '\t',
	}

	for l, expected := range positive_cases {
		// Negative cases
		for _, char := range all_chars {

			if byte(char) == expected {
				continue
			}

			if l.Match(byte(char)) {
				t.Logf(
					"%s.Match('%c') expects false",
					fmt.Sprintf(f, l.source, l.start, l.current),
					char,
				)
				t.Fail()
				break
			}
		}

		// Positive case
		// Positive case second as it mutates state
		if !l.Match(byte(expected)) {
			t.Logf(
				"%s.Match('%c') expects true",
				fmt.Sprintf(f, l.source, l.start, l.current),
				expected,
			)
			t.Fail()
		}
	}
}

func Test_IsDigit(t *testing.T) {
	// Invalid
	invalid := ascii_lower + ascii_upper + ascii_special

	for _, char := range invalid {
		if IsDigit(byte(char)) {
			t.Logf("IsDigit('%c') returned true", char)
			t.Fail()
		}
	}

	// Invalid and valid: number literals and ASCII numbers
	for i, digit := range ascii_digits {
		if IsDigit(byte(i)) {
			t.Logf("IsDigit(%d) returned true", i)
			t.Fail()
		}

		if !IsDigit(byte(digit)) {
			t.Logf("IsDigit('%c') returned false", digit)
			t.Fail()
		}
	}
}

func Test_IsAlpha(t *testing.T) {
	// Invalid
	invalid := ascii_digits + ascii_special

	for _, char := range invalid {
		if IsAlpha(byte(char)) && char != '_' {
			t.Logf("IsAlpha('%c') returned true", char)
			t.Fail()
		}
	}

	// Valid
	valid := ascii_lower + ascii_upper + "_"

	for _, char := range valid {
		if !IsAlpha(byte(char)) {
			t.Logf("IsAlpha('%c') returned false", char)
			t.Fail()
		}
	}
}

func Test_IsAlphanumeric(t *testing.T) {
	// Invalid
	invalid := ascii_special

	for _, char := range invalid {
		if IsAlphanumeric(byte(char)) && char != '_' {
			t.Logf("IsAlphanumeric('%c') returned true", char)
			t.Fail()
		}
	}

	// Valid
	valid := ascii_lower + ascii_upper + "_" + ascii_digits

	for _, char := range valid {
		if !IsAlphanumeric(byte(char)) {
			t.Logf("IsAlphanumeric('%c') returned false", char)
			t.Fail()
		}
	}
}
