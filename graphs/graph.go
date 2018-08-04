package graphs

type Node struct {
	Data     int
	Adjacent []*Node
}

type Graph []*Node
