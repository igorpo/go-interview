package trees

import "testing"

func TestIsBalanced(t *testing.T) {
	type input struct {
		T *BTNode
	}

	type output struct {
		IsBal bool
	}

	type test struct {
		Input  input
		Output output
	}

	tests := []test{
		{input{testTree1()}, output{true}},
		{input{testTree2()}, output{true}},
		{input{testTree3()}, output{false}},
		{input{&BTNode{Data: "1"}}, output{true}},
		{input{nil}, output{true}},
	}

	for i, tst := range tests {
		actual := IsBalanced(tst.Input.T)
		expected := tst.Output.IsBal
		if actual != expected {
			t.Errorf("Test #%d: got=%v, want=%v", i+1, actual, expected)
		}
	}
}
