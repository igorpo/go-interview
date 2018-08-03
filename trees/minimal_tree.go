package trees

// Given an array of increasing unique integer values create a minimal height binary search tree
func MinimalHeightTree(list []int) *BTIntNode {
	if len(list) == 0 {
		return nil
	} else if len(list) == 1 {
		return &BTIntNode{Data: list[0]}
	}

	med := len(list) / 2
	medElem := list[med]
	root := &BTIntNode{Data: medElem}
	root.Left = MinimalHeightTree(list[:med])
	root.Right = MinimalHeightTree(list[med:])
	return root
}
