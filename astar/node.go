package astar

import (
	"fmt"
)

type Node struct {
	f, g, h   int
	X, Y      int
	Weighting int
	Tag       string
	parent    *Node
}

// String returns formatted values of the node
func (n *Node) String() string {
	return fmt.Sprintf("Node [X:%d Y:%d F:%d G:%d H:%d Tag:%s]", n.X, n.Y, n.f, n.g, n.h, n.Tag)
}

func (n *Node) CreateTag() {
	n.Tag = fmt.Sprintf(`%d-%d`, n.X, n.Y)
}
