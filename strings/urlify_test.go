package strings

import (
	"testing"
)

func TestUrlifyBuilder(t *testing.T) {
	type input struct {
		S string
	}

	type output struct {
		Url string
	}

	type test struct {
		Input  input
		Output output
	}

	tests := []test{
		{input{"I am happy!"}, output{"I%20am%20happy!"}},
		{input{"Hey  , good morning  sir! "}, output{"Hey%20%20,%20good%20morning%20%20sir!%20"}},
		{input{" "}, output{"%20"}},
		{input{"nospaceshere"}, output{"nospaceshere"}},
	}

	for i, tst := range tests {
		actual := UrlifyBuilder(tst.Input.S)
		expected := tst.Output.Url
		if actual != expected {
			t.Errorf("Test #%d: got=%v, want=%v", i+1, actual, expected)
		}
	}
}
