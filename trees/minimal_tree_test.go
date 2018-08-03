package trees

import "testing"

func TestMinHeightTree(t *testing.T) {
	type input struct {
		List []int
	}

	type output struct {
		MinHeightTree *BTIntNode
	}

	type test struct {
		Input  input
		Output output
	}

	one := &BTIntNode{Data: 1}
	two := &BTIntNode{Data: 2}
	three := &BTIntNode{Data: 3}
	two.Left = one
	two.Right = three

	tests := []test{
		{input{[]int{1}}, output{&BTIntNode{Data: 1}}},
		{input{[]int{}}, output{nil}},
		{input{[]int{1, 2, 3}}, output{two}},
	}

	for i, tst := range tests {
		actual := MinimalHeightTree(tst.Input.List)
		expected := tst.Output.MinHeightTree
		if !checkTreeEquality(actual, expected) {
			t.Errorf("Test #%d: got=%v, want=%v", i+1, actual, expected)
		}
	}
}

func checkTreeEquality(T1 *BTIntNode, T2 *BTIntNode) bool {
	if T1 == nil || T2 == nil {
		return true
	} else if (T1 == nil && T2 != nil) || (T1 != nil && T2 == nil) {
		return false
	}

	return T1.Data == T2.Data && checkTreeEquality(T1.Left, T2.Left) && checkTreeEquality(T1.Right, T2.Right)
}
