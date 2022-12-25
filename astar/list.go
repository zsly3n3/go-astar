package astar

import (
	"container/heap"
	"errors"
)

// List represents a list of nodes
type List struct {
	nodes map[string]Node //map[tag]Node,如果保存地址数据再外可能被修改
	data  []*heapObj      //最小堆
}

// NewList creates a new list
func NewList() *List {
	l := new(List)
	l.nodes = make(map[string]Node)
	l.data = make([]*heapObj, 0, 4)
	heap.Init(l)
	return l
}

// Add one or more nodes to the list
func (l *List) Add(nodes ...Node) {
	for _, v := range nodes {
		l.nodes[v.Tag] = v
		obj := new(heapObj)
		obj.Key = v.Tag
		obj.F = v.f
		heap.Push(l, obj)
	}
}

// Remove a node from the list
// if the node is not found we do nothing
func (l *List) remove(tag string) {
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
	l.nodes = nil
	l.data = nil
}

// GetMinFNode returns the node with the smallest node.F value
func (l *List) GetMinFNode() (Node, error) {
	if len(l.nodes) <= 0 || len(l.data) <= 0 {
		return Node{}, errors.New(`list is empty`)
	}
	obj := heap.Pop(l).(*heapObj)
	minNode, tf := l.nodes[obj.Key]
	if obj != nil && tf {
		l.remove(obj.Key)
		return minNode, nil
	}
	return Node{}, errors.New(`get min node err`)
}

type heapObj struct {
	F   int
	Key string
}

func (l *List) Len() int           { return len(l.data) }
func (l *List) Less(i, j int) bool { return l.data[i].F < l.data[j].F }
func (l *List) Swap(i, j int)      { l.data[i], l.data[j] = l.data[j], l.data[i] }

func (l *List) Push(x any) {
	l.data = append(l.data, x.(*heapObj))
}

func (l *List) Pop() any {
	old := l.data
	n := len(old)
	x := old[n-1]
	l.data = old[0 : n-1]
	return x
}
