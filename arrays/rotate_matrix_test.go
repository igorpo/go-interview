package arrays

import "testing"

func TestRotateMatrix90CW(t *testing.T) {
	type input struct {
		Matrix [][]int
	}

	type output struct {
		RotatedMatrix [][]int
	}

	type test struct {
		Input  input
		Output output
	}

	tests := []test{
		{
			input{
				[][]int{
					{1},
				},
			},
			output{
				[][]int{
					{1},
				},
			},
		},
		{
			input{
				[][]int{
					{1, 2},
					{3, 4},
				},
			},
			output{
				[][]int{
					{3, 1},
					{4, 2},
				},
			},
		},
		{
			input{
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			output{
				[][]int{
					{7, 4, 1},
					{8, 5, 2},
					{9, 6, 3},
				},
			},
		},
	}

	for i, tst := range tests {
		actual := tst.Input.Matrix
		RotateMatrix90CW(actual)
		expected := tst.Output.RotatedMatrix
		if !equalMatrix(actual, expected) {
			t.Errorf("Test #%d: got=%v, want=%v", i+1, actual, expected)
		}
	}
}

func equalMatrix(a, b [][]int) bool {
	for i := range a {
		for j := range a[0] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
