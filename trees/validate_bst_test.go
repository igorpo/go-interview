package trees

import "testing"

func TestIsBST(t *testing.T) {
	type input struct {
		T *BTIntNode
	}

	type output struct {
		IsBST bool
	}

	type test struct {
		Input  input
		Output output
	}

	tests := []test{
		{input{testNotBST1()}, output{false}},
		{input{testBigBST()}, output{true}},
		{input{testBigNotBST()}, output{false}},
		{input{nil}, output{true}},
		{input{&BTIntNode{Data: 1}}, output{true}},
	}

	for i, tst := range tests {
		actual := IsBST(tst.Input.T)
		expected := tst.Output.IsBST
		if actual != expected {
			t.Errorf("Test #%d: got=%v, want=%v", i+1, actual, expected)
		}
	}
}

func testBST1() *BTIntNode {
	one := &BTIntNode{Data: 1}
	onePrime := &BTIntNode{Data: 1}

	one.Left = onePrime
	return one
}

func testNotBST1() *BTIntNode {
	one := &BTIntNode{Data: 1}
	onePrime := &BTIntNode{Data: 1}

	one.Right = onePrime
	return one
}

func testBigBST() *BTIntNode {
	three := &BTIntNode{Data: 3}
	five := &BTIntNode{Data: 5}
	seven := &BTIntNode{Data: 7}
	ten := &BTIntNode{Data: 10}
	fifteen := &BTIntNode{Data: 15}
	seventeen := &BTIntNode{Data: 17}
	twenty := &BTIntNode{Data: 20}
	thirty := &BTIntNode{Data: 30}
	thirtyP := &BTIntNode{Data: 30}

	twenty.Left = ten
	twenty.Right = thirty
	ten.Left = five
	ten.Right = fifteen
	fifteen.Right = seventeen
	five.Left = three
	five.Right = seven
	thirty.Left = thirtyP

	return twenty
}

func testBigNotBST() *BTIntNode {
	three := &BTIntNode{Data: 3}
	five := &BTIntNode{Data: 5}
	seven := &BTIntNode{Data: 7}
	ten := &BTIntNode{Data: 10}
	fifteen := &BTIntNode{Data: 15}
	seventeen := &BTIntNode{Data: 17}
	twenty := &BTIntNode{Data: 20}
	thirty := &BTIntNode{Data: 30}
	thirtyP := &BTIntNode{Data: 30}

	twenty.Left = ten
	twenty.Right = thirty
	ten.Left = five
	ten.Right = fifteen
	fifteen.Right = seventeen
	five.Left = three
	five.Right = seven
	thirty.Right = thirtyP

	return twenty
}
