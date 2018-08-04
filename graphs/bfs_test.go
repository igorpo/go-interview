package graphs

import (
	"testing"
)

func TestBfs(t *testing.T) {
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

	traversal := Bfs(g)
	for _, n := range g {
		if !contains(n.Data, traversal) {
			t.Errorf("Node %d is not in dfs traversal!", n.Data)
		}
	}
}
