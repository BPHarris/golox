package errors

import "testing"

func Test_HasHadError(t *testing.T) {
	// assert HasHadError is false initially
	if HasHadError() {
		t.Fatalf("expect HasHadError() to be false")
	}

	// set has_had_error
	has_had_error = true

	if !HasHadError() {
		t.Fatalf("expect HasHadError() to be true")
	}

	// reset
	has_had_error = false
}

func Test_Error(t *testing.T) {
	// assert HasHadError is false initially
	if HasHadError() {
		t.Fatalf("expect HasHadError() to be false")
	}

	// have an error
	Error(SyntaxError, 999, "oh no!")

	if !HasHadError() {
		t.Fatalf("expect HasHadError() to be true")
	}

	// reset
	has_had_error = false
}

func ExampleError_lox_error() {
	Error(LoxError, 66, "my error")
	// Output: LoxError: line 66: my error
}

func ExampleError_syntax_error() {
	Error(SyntaxError, 999, "oh no!")
	// Output: SyntaxError: line 999: oh no!
}

func Test_Report(t *testing.T) {
	// assert HasHadError is false initially
	if HasHadError() {
		t.Fatalf("expect HasHadError() to be false")
	}

	// have an error
	Report(SyntaxError, 999, "location", "oh no!")

	if !HasHadError() {
		t.Fatalf("expect HasHadError() to be true")
	}

	// reset
	has_had_error = false
}

func ExampleReport_lox_error() {
	Report(LoxError, 66, "somewhere", "my error")
	// Output: LoxError (somewhere): line 66: my error
}

func ExampleReport_syntax_error() {
	Report(SyntaxError, 999, "somewhen", "oh no!")
	// Output: SyntaxError (somewhen): line 999: oh no!
}
