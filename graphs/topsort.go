package graphs

// Will not detect cycles unfortunately
func Topsort(g Graph) []int {
	visited := make(map[*Node]bool)
	topsorted := []int{}

	for _, n := range g {
		if !visited[n] {
			topsortHelper(n, visited, &topsorted)
		}
	}

	for i, j := 0, len(topsorted)-1; i < len(topsorted)/2; i, j = i+1, j-1 {
		topsorted[i], topsorted[j] = topsorted[j], topsorted[i]
	}
	return topsorted
}

func topsortHelper(n *Node, visited map[*Node]bool, list *[]int) {
	if n == nil {
		return
	}

	visited[n] = true

	for _, v := range n.Adjacent {
		topsortHelper(v, visited, list)
	}

	*list = append(*list, n.Data)

}
