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

// Returns a slice of data from the tree via preorder traversal: (root, L, R)
func Preorder(root *BTNode) []string {
	return preorderHelper([]string{}, root)
}

func preorderHelper(list []string, n *BTNode) []string {
	if n == nil {
		return list
	}

	ret := []string{n.Data}
	lList := preorderHelper(list, n.Left)
	rList := preorderHelper(list, n.Right)
	ret = append(ret, lList...)
	ret = append(ret, rList...)
	return ret
}

// Returns a slice of data from the tree via postorder traversal: (L, R, root)
func Postorder(root *BTNode) []string {
	return postorderHelper([]string{}, root)
}

func postorderHelper(list []string, n *BTNode) []string {
	if n == nil {
		return list
	}

	lList := postorderHelper(list, n.Left)
	rList := postorderHelper(list, n.Right)
	lList = append(lList, rList...)
	lList = append(lList, n.Data)
	return lList
}
