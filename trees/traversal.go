package trees

// Returns a slice of data from the tree via inorder traversal: (L, root, R)
func Inorder(root *BTNode) []string {
	return inorderHelper([]string{}, root)
}

func inorderHelper(list []string, n *BTNode) []string {
	if n == nil {
		return list
	}

	lList := inorderHelper(list, n.Left)
	rList := inorderHelper(list, n.Right)

	lList = append(lList, n.Data)
	lList = append(lList, rList...)
	return lList
}
