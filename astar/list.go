package astar

import (
	"errors"
)

// List represents a list of nodes
type List struct {
	nodes map[string]Node //map[tag]Node,如果保存地址数据再外可能被修改
}

// NewList creates a new list
func NewList() *List {
	l := new(List)
	l.nodes = make(map[string]Node)
	return l
}

// Add one or more nodes to the list
func (l *List) Add(nodes ...Node) {
	for _, v := range nodes {
		l.nodes[v.Tag] = v
	}
}

// Remove a node from the list
// if the node is not found we do nothing
func (l *List) Remove(tag string) {
	delete(l.nodes, tag)
}

// Contains check if a node is in the list
func (l *List) Contains(tag string) bool {
	_, tf := l.nodes[tag]
	if tf {
		return true
	}
	return false
}

// IsEmpty returns if the nodes list has nodes or not
func (l *List) IsEmpty() bool {
	return len(l.nodes) == 0
}

// Clear removes all nodes from the list
func (l *List) Clear() {
	l.nodes = make(map[string]Node)
}

// GetMinFNode returns the node with the smallest node.F value
func (l *List) GetMinFNode() (Node, error) {
	if len(l.nodes) <= 0 {
		return Node{}, errors.New(`list is empty`)
	}
	var minNode Node
	var minF int
	isFirst := true
	for _, v := range l.nodes {
		if isFirst {
			minF = v.f
			minNode = v
			isFirst = false
		} else {
			if v.f < minF {
				minF = v.f
				minNode = v
			}
		}
	}
	return minNode, nil
}
