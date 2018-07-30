package strings

import "testing"

func TestIsOneEditAway(t *testing.T) {
	type input struct {
		S string
		T string
	}

	type output struct {
		IsOneEdit bool
	}

	type test struct {
		Input  input
		Output output
	}

	tests := []test{
		// replace
		{input{"abc", "abb"}, output{true}},
		{input{"axe", "exe"}, output{true}},
		{input{"I like to eat!", "U like to eat!"}, output{true}},
		{input{"ssss", "sssa"}, output{true}},
		{input{"a", "b"}, output{true}},

		{input{"abc", "acb"}, output{false}},
		{input{"I like to eat!", "U lik3 to eat!"}, output{false}},
		{input{"aabbcc", "bbccaa"}, output{false}},

		// insert
		{input{"axe", "axes"}, output{true}},
		{input{"axe", "laxe"}, output{true}},
		{input{"axe", "axle"}, output{true}},
		{input{"", "a"}, output{true}},
		{input{"aaaaa", "abaaaa"}, output{true}},
		{input{"aaaaa", "aaabaa"}, output{true}},
		{input{"aaaaa", "aaaaav"}, output{true}},
		{input{"aaaaa", "faaaaa"}, output{true}},
		{input{"baby", "bodys"}, output{false}},
		{input{"aaa", "aabb"}, output{false}},

		// remove
		{input{"aaa", "aa"}, output{true}},
		{input{"abcd", "acd"}, output{true}},
		{input{"123123123", "12123123"}, output{true}},
		{input{"one", "ne"}, output{true}},
		{input{"one", "on"}, output{true}},
		{input{"o", ""}, output{true}},
		{input{"baby", "bod"}, output{false}},
		{input{"aaaabbbb", "bbbbaaa"}, output{false}},
		{input{"aaann", "aanm"}, output{false}},

		// misc
		{input{"", "123"}, output{false}},
		{input{"", ""}, output{true}},
		{input{"avc", "a"}, output{false}},
	}

	for i, tst := range tests {
		actual := IsOneEditAway(tst.Input.S, tst.Input.T)
		expected := tst.Output.IsOneEdit
		if actual != expected {
			t.Errorf("Test #%d: got=%v, want=%v", i+1, actual, expected)
		}
	}
}
