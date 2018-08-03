package trees

// Given a binary tree, check if it is balanced (i.e. height of node any node n's two subtrees differs by one at most)

// We can recursively find the heights of the subtrees. If we are in a subtree, we know right away whether it is balanced
// or not, so we can return a -1 (or error, etc.) and bubble that "error" code back to the top for a more efficient algo.
// We only take O(n) time and O(h) space, where h = height of tree,
func IsBalanced(root *BTNode) bool {
	return height(root) != -1
}

func height(t *BTNode) int {
	if t == nil {
		return 0
	}

	lH := height(t.Left)
	if lH == -1 {
		return -1
	}

	rH := height(t.Right)
	if rH == -1 {
		return -1
	}

	if abs(lH-rH) > 1 {
		return -1
	} else {
		return max(lH, rH) + 1
	}

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
