package trees

import "testing"

func TestSubtree(t *testing.T) {
	type input struct {
		T1 *BTNode
		T2 *BTNode
	}

	type output struct {
		IsSub bool
	}

	type test struct {
		Input  input
		Output output
	}

	tests := []test{
		{input{testTree1(), subtree1()}, output{true}},
		{input{testTree1(), testTree1()}, output{true}},
		{input{testTree1(), &BTNode{Data: "1"}}, output{true}},
		{input{testTree1(), nil}, output{true}},
		{input{nil, nil}, output{true}},
		{input{&BTNode{Data: "1"}, nil}, output{true}},

		{input{nil, testTree1()}, output{false}},
		{input{testTree1(), subtree2()}, output{false}},
		{input{testTree1(), testTree2()}, output{false}},
		{input{testTree1(), subtree2()}, output{false}},
	}

	for i, tst := range tests {
		actual := IsSubtree(tst.Input.T1, tst.Input.T2)
		expected := tst.Output.IsSub
		if actual != expected {
			t.Errorf("Test #%d: got=%v, want=%v", i+1, actual, expected)
		}
	}
}

func subtree1() *BTNode {
	one := &BTNode{Data: "1"}
	two := &BTNode{Data: "2"}
	three := &BTNode{Data: "3"}

	two.Left = one
	two.Right = three
	return two
}

func subtree2() *BTNode {
	one := &BTNode{Data: "1"}
	three := &BTNode{Data: "3"}
	four := &BTNode{Data: "4"}

	three.Left = one
	three.Right = four
	return three
}
