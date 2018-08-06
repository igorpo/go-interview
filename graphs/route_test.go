package graphs

import (
	"testing"
)

func TestRoute(t *testing.T) {
	g := make(Graph, 13, 13)
	for i := range g {
		g[i] = &Node{Data: i + 1}
	}

	g[0].Adjacent = []*Node{g[1], g[2], g[3]}
	g[1].Adjacent = []*Node{g[2], g[4], g[5]}
	g[2].Adjacent = []*Node{g[1], g[5]}
	g[3].Adjacent = []*Node{g[7]}
	g[4].Adjacent = []*Node{g[1], g[5]}
	g[5].Adjacent = []*Node{g[6]}
	g[6].Adjacent = []*Node{g[4], g[5], g[7]}
	g[7].Adjacent = []*Node{g[2]}
	g[8].Adjacent = []*Node{g[3], g[9], g[11]}
	g[9].Adjacent = []*Node{g[10]}
	g[10].Adjacent = []*Node{g[9]}
	g[11].Adjacent = []*Node{g[7]}

	type input struct {
		G Graph
		S *Node
		T *Node
	}

	type output struct {
		Path []int
	}

	type test struct {
		Input  input
		Output output
	}

	tests := []test{
		{input{g, g[0], g[2]}, output{[]int{1, 3}}},
		{input{g, g[0], g[3]}, output{[]int{1, 4}}},
		{input{g, g[1], g[5]}, output{[]int{2, 6}}},
		{input{g, g[1], g[7]}, output{[]int{2, 6, 7, 8}}},
		{input{g, g[1], g[5]}, output{[]int{2, 6}}},
		{input{g, g[3], g[10]}, output{[]int{}}},
		{input{g, g[3], g[4]}, output{[]int{4, 8, 3, 2, 5}}},
	}

	for i, tst := range tests {
		actual := convertToInt(Route(tst.Input.S, tst.Input.T, tst.Input.G))
		expected := tst.Output.Path
		if !equalsPath(actual, expected) {
			t.Errorf("Test #%d failed, got=%v, want=%v", i+1, actual, expected)
		}
	}
}

func convertToInt(path []*Node) []int {
	pathI := []int{}
	for _, v := range path {
		pathI = append(pathI, v.Data)
	}
	return pathI
}

func equalsPath(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
