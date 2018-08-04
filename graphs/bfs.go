package graphs

func Bfs(g Graph) []int {
	visited := make(map[*Node]bool)
	traversed := []int{}
	queue := []*Node{}

	for _, root := range g {
		if !visited[root] {
			visited[root] = true
			queue = append(queue, root)

			for len(queue) != 0 {
				var n *Node
				// dequeue
				n, queue = queue[0], queue[1:]
				traversed = append(traversed, n.Data)

				for _, v := range n.Adjacent {
					if !visited[v] {
						visited[v] = true
						queue = append(queue, v)
					}

				}
			}
		}
	}

	return traversed
}
