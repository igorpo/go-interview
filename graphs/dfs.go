package graphs

func Dfs(g Graph) []int {
	visited := make(map[*Node]bool)
	traversed := []int{}

	// pick a node to search
	for _, n := range g {
		// this completes a dfs of a disconnected graph, we loop once the
		// search is done on a connected graph to see if there are any
		// more unseen vertices
		if !visited[n] {
			dfsHelper(n, visited, &traversed)
		}

	}
	return traversed
}

func dfsHelper(n *Node, visited map[*Node]bool, list *[]int) {
	if n == nil {
		return
	}

	*list = append(*list, n.Data)
	visited[n] = true

	for _, v := range n.Adjacent {
		if !visited[v] {
			dfsHelper(v, visited, list)
		}
	}
}
