package strings

import (
	"testing"
)

func TestPermutations(t *testing.T) {
	type input struct {
		S string
	}

	type output struct {
		Permutations []string
	}

	type test struct {
		Input  input
		Output output
	}

	tests := []test{
		{input{"a"}, output{[]string{"a"}}},
		{input{"ab"}, output{[]string{"ab", "ba"}}},
		{input{"abc"}, output{[]string{"cab", "acb", "abc", "cba", "bca", "bac"}}},
		{input{"abcd"}, output{[]string{
			"dcab", "cdab", "cadb", "cabd",
			"dacb", "adcb", "acdb", "acbd",
			"dabc", "adbc", "abdc", "abcd",
			"dcba", "cdba", "cbda", "cbad",
			"dbca", "bdca", "bcda", "bcad",
			"dbac", "bdac", "badc", "bacd",
		}}},
	}

	for i, tst := range tests {
		actual := Permutations(tst.Input.S)
		expected := tst.Output.Permutations
		if !stringSliceEqual(actual, expected) {
			t.Errorf("Test #%d: got=%v, want=%v", i+1, actual, expected)
		}
	}
}

func stringSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	m := make(map[string]int)
	for _, v := range a {
		m[v] += 1
	}

	for _, v := range b {
		m[v] -= 1
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}
