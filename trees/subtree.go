package trees

// Given a tree T1 and another tree T2, determine if T2 is a subtree of T1.
// A tree T2 is a subtre of T1 if there exists a node n such that the subtree of n is identical to T2.
// That is, if you cut off the tree at node n, the two trees would be identical.

func IsSubtree(T1 *BTNode, T2 *BTNode) bool {
	if T2 == nil {
		return true
	}
	return checkTree(T1, T2)
}

// Walks down T1 to find nodes that match T2's root, and then try to validate that
// the subtrees are the same.
func checkTree(T1 *BTNode, T2 *BTNode) bool {
	if T1 == nil {
		return false
	}

	if T1.Data == T2.Data && isIdenticalTree(T1, T2) {
		return true
	}

	return checkTree(T1.Left, T2) || checkTree(T1.Right, T2)
}

// T1: nil, T2: non-nil     -> non-equal
// T1: nil, T2: nil         -> equal
// T1: non-nil, T2: nil     -> non-equal
// T1: non-nil, T2: non-nil -> check roots and L, R subtrees
func isIdenticalTree(T1 *BTNode, T2 *BTNode) bool {
	if T1 == nil && T2 == nil {
		return true
	} else if (T1 == nil && T2 != nil) || (T2 == nil && T1 != nil) {
		return false
	}

	return T1.Data == T2.Data && isIdenticalTree(T1.Left, T2.Left) && isIdenticalTree(T1.Right, T2.Right)
}
