package strings

import (
	"testing"
)

func TestIsPalindromePermutation(t *testing.T) {
	type input struct {
		S string
	}

	type output struct {
		isPalPerm bool
	}

	type test struct {
		Input  input
		Output output
	}

	tests := []test{
		{input{"aabbccddeeff"}, output{true}},
		{input{"aabbccddeeffg"}, output{true}},
		{input{"a"}, output{true}},
		{input{""}, output{true}},
		{input{"aabc"}, output{false}},
		{input{"abcdc"}, output{false}},
	}

	for i, tst := range tests {
		actual := IsPalindromePermutationTable(tst.Input.S)
		expected := tst.Output.isPalPerm
		if actual != expected {
			t.Errorf("Test #%d: got=%v, want=%v", i+1, actual, expected)
		}
	}
}
