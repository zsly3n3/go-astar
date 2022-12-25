package astar

import (
	"errors"
	"fmt"
	"math"
)

// Config holds important settings
// to perform the calculation
//
// GridWidth and GridHeight are required and represents
// the size of the grid
//
// InvalidNodes can be used to add not accessible nodes like obstacles etc.
// WeightedNodes can be used to add nodes to be avoided like mud or mountains
type Config struct {
	GridWidth, GridHeight int
	InvalidNodes          []*Node          //常量
	WeightedNodes         map[string]*Node //常量
}

type astar struct {
	closedList         map[string]struct{} //递增
	startNode, endNode *Node               //常量
	openList           *List               //递增
	config             *Config             //常量
}

// New creates a new astar instance
func New(config *Config) (*astar, error) {
	if config.GridWidth < 2 || config.GridHeight < 2 {
		return nil, errors.New("GridWidth and GridHeight must be min 2")
	}
	a := &astar{config: config}
	a.openList = NewList()
	a.closedList = make(map[string]struct{})
	return a.init(), nil
}

// init initialised needed properties
// internal function
func (a *astar) init() *astar {
	// add invalidNodes directly to the closedList
	for _, v := range a.config.InvalidNodes {
		a.closedList[v.Tag] = struct{}{}
	}
	return a
}

// H caluclates the absolute distance between
// nodeA and nodeB calculates by the manhattan distance
func (a *astar) H(nodeA *Node, nodeB *Node) int {
	absX := math.Abs(float64(nodeA.X - nodeB.X))
	absY := math.Abs(float64(nodeA.Y - nodeB.Y))
	return int(absX + absY)
}

// GetNeighborNodes calculates the next neighbors of the given node
// if a neighbor node is not accessible the node will be ignored
func (a *astar) GetNeighborNodes(node Node) []Node {
	neighborNodes := make([]Node, 0, 4) //4个放向

	upNode := Node{X: node.X, Y: node.Y + 1, parent: &node, g: node.g + 1}
	upNode.CreateTag()
	if a.isAccessible(&upNode) {
		neighborNodes = append(neighborNodes, upNode)
	}

	downNode := Node{X: node.X, Y: node.Y - 1, parent: &node, g: node.g + 1}
	downNode.CreateTag()
	if a.isAccessible(&downNode) {
		neighborNodes = append(neighborNodes, downNode)
	}

	leftNode := Node{X: node.X - 1, Y: node.Y, parent: &node, g: node.g + 1}
	leftNode.CreateTag()
	if a.isAccessible(&leftNode) {
		neighborNodes = append(neighborNodes, leftNode)
	}

	rightNode := Node{X: node.X + 1, Y: node.Y, parent: &node, g: node.g + 1}
	rightNode.CreateTag()
	if a.isAccessible(&rightNode) {
		neighborNodes = append(neighborNodes, rightNode)
	}

	return neighborNodes
}

// isAccessible checks if the node is reachable in the grid
// and is not in the invalidNodes slice
func (a *astar) isAccessible(node *Node) bool {

	// if node is out of bound
	if node.X < 0 || node.Y < 0 || node.X > a.config.GridWidth-1 || node.Y > a.config.GridHeight-1 {
		return false
	}

	// check if the node is in the closedList
	// the predefined invalidNodes are also in this list
	_, tf := a.closedList[node.Tag]
	if tf {
		return false
	}
	return true
}

// IsEndNode checks if the given node has
// equal node coordinates with the end node
func (a *astar) IsEndNode(checkNode, endNode *Node) bool {
	return checkNode.X == endNode.X && checkNode.Y == endNode.Y
}

// FindPath starts the a* algorithm for the given start and end node
// The return value will be the fastest way represented as a nodes slice
//
// If no path was found it returns nil and an error
func (a *astar) FindPath(startNode, endNode *Node) ([]*Node, error) {

	a.startNode = startNode
	a.endNode = endNode

	defer func() {
		a.openList.Clear()
		a.closedList = make(map[string]struct{})
	}()

	a.openList.Add(*startNode)

	for !a.openList.IsEmpty() {

		currentNode, err := a.openList.GetMinFNode()
		if err != nil {
			return nil, fmt.Errorf("cannot get minF node %v", err)
		}

		a.closedList[currentNode.Tag] = struct{}{}

		// we found the path
		if a.IsEndNode(&currentNode, endNode) {
			return a.getNodePath(&currentNode), nil
		}

		neighbors := a.GetNeighborNodes(currentNode)
		for _, neighbor := range neighbors {
			_, tf := a.closedList[neighbor.Tag]
			if tf {
				continue
			}
			if !a.openList.Contains(neighbor.Tag) {
				a.calculateNode(&neighbor)
				a.openList.Add(neighbor)
			}
		}

	}

	return nil, errors.New(`no path found`)
}

// calculateNode calculates the F, G and H value for the given node
func (a *astar) calculateNode(node *Node) {
	// check for special node weighting
	wNode, tf := a.config.WeightedNodes[node.Tag]
	if tf {
		node.g = node.g + wNode.Weighting
	}
	node.h = a.H(node, a.endNode)
	node.f = node.g + node.h
}

// getNodePath returns the chain of parent nodes
// the given node will be still included in the nodes slice
func (a *astar) getNodePath(currentNode *Node) []*Node {
	nodePath := make([]*Node, 0, 30)
	nodePath = append(nodePath, currentNode)
	for {
		if currentNode.parent == nil {
			break
		}

		parentNode := currentNode.parent

		// if the end of node chain
		if parentNode.parent == nil {
			break
		}

		nodePath = append(nodePath, parentNode)
		currentNode = parentNode
	}
	return nodePath
}
