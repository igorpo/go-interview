package strings

import (
	"testing"
)

func TestUniqueASCII(t *testing.T) {
	type input struct {
		S string
	}

	type output struct {
		IsUnique bool
	}

	type test struct {
		Input  input
		Output output
	}

	tests := []test{
		{input{"a"}, output{true}},
		{input{"ab"}, output{true}},
		{input{"abc"}, output{true}},
		{input{"abc123"}, output{true}},
		{input{"abcabcabc"}, output{false}},
		{input{"aaa"}, output{false}},
		{input{"abc1233"}, output{false}},
	}

	for i, tst := range tests {
		actual := UniqueASCII(tst.Input.S)
		expected := tst.Output.IsUnique
		if actual != expected {
			t.Errorf("Test #%d: got=%v, want=%v", i+1, actual, expected)
		}
	}
}

func TestUnique(t *testing.T) {
	type input struct {
		S string
	}

	type output struct {
		IsUnique bool
	}

	type test struct {
		Input  input
		Output output
	}

	tests := []test{
		{input{"a"}, output{true}},
		{input{"abAB"}, output{true}},
		{input{"abcC"}, output{true}},
		{input{"abc123"}, output{true}},
		{input{"ß∂ƒ©"}, output{true}},
		{input{"abcabcabc"}, output{false}},
		{input{"aaa"}, output{false}},
		{input{"abc1233B"}, output{false}},
		{input{"abc123ååA"}, output{false}},
		{input{"®ƒs∂qwß≈∂"}, output{false}},
	}

	for i, tst := range tests {
		actual := Unique(tst.Input.S)
		expected := tst.Output.IsUnique
		if actual != expected {
			t.Errorf("Test #%d: got=%v, want=%v", i+1, actual, expected)
		}
	}
}
