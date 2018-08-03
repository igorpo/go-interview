package trees

// Given a binary tree, validate whether or not it is a BST

// We can recursively check whether a tree is a BST by keeping track of the min, max values
func IsBST(root *BTIntNode) bool {
	return validate(root, nil, nil)
}

// Pass the current node's data as max when moving left, min when moving right
func validate(t *BTIntNode, min *int, max *int) bool {
	if t == nil {
		return true
	}

	if (min != nil && t.Data <= *min) || (max != nil && t.Data > *max) {
		return false
	}

	data := t.Data
	return validate(t.Left, min, &data) && validate(t.Right, &data, max)
}
