package lexer

import "testing"

func Test_Literal(t *testing.T) {
	// Test lexemes with literals
	has_literal := map[Lexeme]string{
		{Identifier, "foo", 1}:                  "foo",
		{Identifier, "foobar", 1}:               "foobar",
		{LiteralString, "\"str\"", 1}:           "\"str\"",
		{LiteralString, "\"a long string\"", 1}: "\"a long string\"",
		{LiteralNumber, "1", 1}:                 "1",
		{LiteralNumber, "1.0", 1}:               "1.0",
		{LiteralNumber, "1.000", 1}:             "1.000",
		{LiteralNumber, "12.34", 1}:             "12.34",
	}

	for lexeme, expected := range has_literal {
		literal, err := lexeme.Literal()

		if err != nil {
			t.Logf("%s.Liter() returned the error: %s", lexeme, err)
			t.Fail()
		}

		if literal != expected {
			t.Logf("%s expects literal to be '%s'", lexeme, expected)
			t.Fail()
		}
	}

	// Test lexemes without literals
	no_literal := []Lexeme{
		// Sinle-character lexemes
		{LeftParenthesis, "(", 1},
		// One-or-two character lexemes
		{Bang, "!", 1},
		{BangEqual, "!=", 1},
		{EqualEqual, "==", 1},
		// Keywords
		{If, "if", 1},
		{Print, "print", 1},
		{False, "false", 1},
	}

	for _, lexeme := range no_literal {
		literal, err := lexeme.Literal()

		if err == nil || literal != "" {
			t.Logf("%s.Literal() expects error", lexeme)
			t.Fail()
		}
	}
}

func Test_ParseString(t *testing.T) {
	const malformed = "lexeme is malformed, no enclosing \""

	// Invalid: not a LiteralString
	non_literal_strings := []Lexeme{
		{If, "if", 1},
		{Class, "class", 1},
		{Bang, "!", 1},
		{BangEqual, "!=", 1},
		{Identifier, "foobar", 1},
		{LiteralNumber, "12.34", 1},
	}

	for _, lexeme := range non_literal_strings {
		_, err := lexeme.ParseString()

		if err == nil {
			t.Logf("%s.ParseString() expects error", lexeme)
			t.Fail()
		}
	}

	// Invalid: lexeme is malformed -- no enclosing ""
	malformed_strings := []Lexeme{
		{LiteralString, "", 1},
		{LiteralString, "\"", 1},
		{LiteralString, "\"non terminated string", 1},
		{LiteralString, "non enclosed string", 1},
	}

	for _, lexeme := range malformed_strings {
		_, err := lexeme.ParseString()

		if err == nil || err.Error() != malformed {
			t.Logf("%s.ParseString() expects malformed string error", lexeme)
			t.Fail()
		}
	}

	// Valid
	wellformed_strings := map[Lexeme]string{
		{LiteralString, "\"\"", 1}:                  "",
		{LiteralString, "\"\"\"", 1}:                "\"",
		{LiteralString, "\"'\"", 1}:                 "'",
		{LiteralString, "\"terminated string\"", 1}: "terminated string",
	}

	for lexeme, expected := range wellformed_strings {
		val, err := lexeme.ParseString()

		if err != nil {
			t.Logf("%s.ParseString() failed with error: %s", lexeme, err)
			t.Fail()
		}

		if val != expected {
			t.Logf(
				"%s.ParseString() expects '%s' but recieved '%s'",
				lexeme,
				expected,
				val,
			)
			t.Fail()
		}
	}
}

func Test_ParseFloat(t *testing.T) {
	// Invalid: non a LiteralNumber
	non_literal_numbers := []Lexeme{
		{If, "if", 1},
		{Class, "class", 1},
		{Bang, "!", 1},
		{BangEqual, "!=", 1},
		{Identifier, "foobar", 1},
		{LiteralString, "\"foobar\"", 1},
	}

	for _, lexeme := range non_literal_numbers {
		_, err := lexeme.ParseFloat()

		if err == nil {
			t.Logf("%s.ParseFloat() expects error", lexeme)
			t.Fail()
		}
	}

	// Invalid: malformed float
	malformed_floats := []Lexeme{
		{LiteralNumber, "12..34", 1},
	}

	for _, lexeme := range malformed_floats {
		_, err := lexeme.ParseFloat()

		if err == nil {
			t.Logf("%s.ParseFloat() expects error from strconv", lexeme)
			t.Fail()
		}
	}

	// Valid
	wellformed_floats := map[Lexeme]float64{
		{LiteralNumber, "12.34", 1}:   12.34,
		{LiteralNumber, "0.00", 1}:    0.0,
		{LiteralNumber, "0.0", 1}:     0.0,
		{LiteralNumber, "0", 1}:       0.0,
		{LiteralNumber, "000", 1}:     0.0,
		{LiteralNumber, "000.000", 1}: 0.0,
		{LiteralNumber, "001.000", 1}: 1.0,
		{LiteralNumber, "00.0001", 1}: 0.0001,
		{LiteralNumber, "3", 1}:       3.0,
		{LiteralNumber, "003", 1}:     3.0,
	}

	for lexeme, expected := range wellformed_floats {
		val, err := lexeme.ParseFloat()

		if err != nil {
			t.Logf("%s.ParseFloat() failed with error: %s", lexeme, err)
			t.Fail()
		}

		if val != expected {
			t.Logf(
				"%s.ParseFloat() expects '%f' but recieved '%f'",
				lexeme,
				expected,
				val,
			)
			t.Fail()
		}
	}
}

func Test_String(t *testing.T) {
	cases := map[Lexeme]string{
		// Single-character lexemes
		{Dot, ".", 1}:  "Lexeme(type=Dot, lexeme=\".\")",
		{Star, "*", 1}: "Lexeme(type=Star, lexeme=\"*\")",
		// One-or-two character lexemes
		{Bang, "!", 1}:       "Lexeme(type=Bang, lexeme=\"!\")",
		{BangEqual, "!=", 1}: "Lexeme(type=BangEqual, lexeme=\"!=\")",
		// Keywords
		{Fun, "fun", 1}:     "Lexeme(type=Fun, lexeme=\"fun\")",
		{Class, "class", 1}: "Lexeme(type=Class, lexeme=\"class\")",
		// Literals
		{Identifier, "foobar", 1}:       "Lexeme(type=Identifier, lexeme=\"foobar\", literal=foobar)",
		{LiteralString, "\"a str\"", 1}: "Lexeme(type=LiteralString, lexeme=\"\\\"a str\\\"\", literal=\"a str\")",
		{LiteralNumber, "12.34", 1}:     "Lexeme(type=LiteralNumber, lexeme=\"12.34\", literal=12.34)",
	}

	for lexeme, expected := range cases {
		if lexeme.String() != expected {
			t.Logf("%s expects literal to be '%s'", lexeme, expected)
			t.Fail()
		}
	}
}
