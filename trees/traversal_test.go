package trees

import "testing"

func TestInorderTraversal(t *testing.T) {
	type input struct {
		Tree *BTNode
	}

	type output struct {
		Inorder []string
	}

	type test struct {
		Input  input
		Output output
	}

	tests := []test{
		{input{testTree1()}, output{[]string{"1", "2", "3", "4", "5", "6", "7"}}},
		{input{testTree2()}, output{[]string{"1", "2", "3", "4"}}},
		{input{testTree3()}, output{[]string{"2", "1", "3", "7", "4", "6", "5"}}},
		{input{&BTNode{Data: "1"}}, output{[]string{"1"}}},
		{input{nil}, output{[]string{}}},
	}

	for i, tst := range tests {
		actual := Inorder(tst.Input.Tree)
		expected := tst.Output.Inorder
		if !stringSliceEqual(actual, expected) {
			t.Errorf("Test #%d: got=%v, want=%v", i+1, actual, expected)
		}
	}
}

func TestPreorderTraversal(t *testing.T) {
	type input struct {
		Tree *BTNode
	}

	type output struct {
		Preorder []string
	}

	type test struct {
		Input  input
		Output output
	}

	tests := []test{
		{input{testTree1()}, output{[]string{"4", "2", "1", "3", "6", "5", "7"}}},
		{input{testTree2()}, output{[]string{"3", "2", "1", "4"}}},
		{input{testTree3()}, output{[]string{"2", "3", "1", "4", "7", "5", "6"}}},
		{input{&BTNode{Data: "1"}}, output{[]string{"1"}}},
		{input{nil}, output{[]string{}}},
	}

	for i, tst := range tests {
		actual := Preorder(tst.Input.Tree)
		expected := tst.Output.Preorder
		if !stringSliceEqual(actual, expected) {
			t.Errorf("Test #%d: got=%v, want=%v", i+1, actual, expected)
		}
	}
}

// (('1') '2' ('3')) '4' (('5') '6' ('7'))
func testTree1() *BTNode {
	one := &BTNode{Data: "1"}
	two := &BTNode{Data: "2"}
	three := &BTNode{Data: "3"}
	four := &BTNode{Data: "4"}
	five := &BTNode{Data: "5"}
	six := &BTNode{Data: "6"}
	seven := &BTNode{Data: "7"}

	four.Left = two
	four.Right = six
	two.Left = one
	two.Right = three
	six.Left = five
	six.Right = seven

	return four
}

// ('1' ('2')) '3' (('4'))
func testTree2() *BTNode {
	one := &BTNode{Data: "1"}
	two := &BTNode{Data: "2"}
	three := &BTNode{Data: "3"}
	four := &BTNode{Data: "4"}

	three.Left = two
	three.Right = four
	two.Left = one

	return three
}

// () '2' (('1') '3' (('7') '4' (('6') '5')))
func testTree3() *BTNode {
	one := &BTNode{Data: "1"}
	two := &BTNode{Data: "2"}
	three := &BTNode{Data: "3"}
	four := &BTNode{Data: "4"}
	five := &BTNode{Data: "5"}
	six := &BTNode{Data: "6"}
	seven := &BTNode{Data: "7"}

	two.Right = three
	three.Left = one
	three.Right = four
	four.Left = seven
	four.Right = five
	five.Left = six

	return two
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
