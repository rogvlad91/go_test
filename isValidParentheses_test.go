package main

import "testing"

func TestIsValidParenthesesEmpty(t *testing.T) {
	flag := isValidParentheses("")

	if !flag {
		t.Error("result must be true")
	}
}

func TestIsValidParenthesesEmptyMultipleSpaces(t *testing.T) {
	flag := isValidParentheses("     ")

	if !flag {
		t.Error("result must be true")
	}
}

func TestIsValidParenthesesOnlyOtherChars(t *testing.T) {
	flag := isValidParentheses("foobar!")

	if !flag {
		t.Error("result must be true")
	}
}

func TestIsValidParenthesesWithOtherCharsTrue(t *testing.T) {
	flag := isValidParentheses("(!)u[?]")
	if !flag {
		t.Error("result must be true")
	}
}

func TestIsValidParenthesesWithOtherCharsFalse(t *testing.T) {
	flag := isValidParentheses("(!)}u[[?]")
	if flag {
		t.Error("result must be false")
	}
}

func TestIsValidParenthesesTrue1(t *testing.T) {
	flag := isValidParentheses("[]")
	if !flag {
		t.Error("result must be true")
	}
}

func TestIsValidParenthesesTrue2(t *testing.T) {
	flag := isValidParentheses("(({{[{[]}]}}))")
	if !flag {
		t.Error("result must be true")
	}
}

func TestIsValidParenthesesTrue3(t *testing.T) {
	flag := isValidParentheses("{[]}")
	if !flag {
		t.Error("result must be true")
	}
}

func TestIsValidParenthesesSingleBracket(t *testing.T) {
	flag := isValidParentheses("(")
	if flag {
		t.Error("result must be false")
	}
}

func TestIsValidParenthesesSingleReversedBracket(t *testing.T) {
	flag := isValidParentheses("}")
	if flag {
		t.Error("result must be false")
	}
}

func TestIsValidParenthesesFalse1(t *testing.T) {
	flag := isValidParentheses("{[()")
	if flag {
		t.Error("result must be false")
	}
}

func TestIsValidParenthesesFalse2(t *testing.T) {
	flag := isValidParentheses("{{([(])}}")
	if flag {
		t.Error("result must be false")
	}
}
