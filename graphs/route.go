package graphs

// Given a directed graph, find a path, if one exists, between two nodes, or return []

// BFS will find the shortest path since the graph in unweighted (otherwise, say hello to Dijkstra!)
func Route(s, t *Node, g Graph) []*Node {
	visited := make(map[*Node]bool)
	queue := []*Node{s}
	prevMap := make(map[*Node]*Node)

	visited[s] = true
	prevMap[s] = nil

	// perform bfs tracking previous nodes
	for len(queue) != 0 {
		var curr *Node
		curr, queue = queue[0], queue[1:]

		if curr == t {
			break
		}

		for _, v := range curr.Adjacent {
			if !visited[v] {
				visited[v] = true
				queue = append(queue, v)
				prevMap[v] = curr
			}
		}
	}

	// if t does not have a value in the previous table, we know that no
	// path was found, otherwise we find the path by following prev pointers
	// note that path with be in reverse order
	path := []*Node{}
	p, ok := prevMap[t]
	if ok {
		path = append(path, t)
		for p != nil {
			path = append(path, p)
			p = prevMap[p]
		}
	}

	// reverse list
	for i, j := 0, len(path)-1; i < len(path)/2; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}
