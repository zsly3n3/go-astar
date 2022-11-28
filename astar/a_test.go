package astar

import (
	"fmt"
	"testing"
)

func TestAstar_H(t *testing.T) {
	startNode := &Node{X: 1, Y: 2}
	endNode := &Node{X: 10, Y: 2}
	startNode.CreateTag()
	endNode.CreateTag()

	obstacleNodes := []*Node{
		{X: 9, Y: 3, Tag: `9-3`},
		{X: 10, Y: 3, Tag: `10-3`},
		{X: 11, Y: 3, Tag: `11-3`},
		{X: 9, Y: 2, Tag: `9-2`},
		{X: 11, Y: 2, Tag: `11-2`},
		//{X: 9, Y: 1, Tag: `9-1`},
		//{X: 10, Y: 1, Tag: `10-1`},
		//{X: 11, Y: 1, Tag: `11-1`},
	}
	//w := 1
	waterNodes := []*Node{
		//{X: 4, Y: 1, Weighting: w, Tag: `4-1`},
		//{X: 4, Y: 2, Weighting: w, Tag: `4-2`},
		//{X: 5, Y: 1, Weighting: w, Tag: `5-1`},
		//{X: 5, Y: 2, Weighting: w, Tag: `5-2`},
		//{X: 5, Y: 3, Weighting: w, Tag: `5-3`},
		//{X: 5, Y: 4, Weighting: w, Tag: `5-4`},
		//{X: 6, Y: 1, Weighting: w, Tag: `6-1`},
		//{X: 6, Y: 2, Weighting: w, Tag: `6-2`},
		//{X: 6, Y: 3, Weighting: w, Tag: `6-3`},
		//{X: 7, Y: 2, Weighting: w, Tag: `7-2`},
	}

	// set nodes to the config
	aConfig := &Config{
		GridWidth:     12,
		GridHeight:    5,
		InvalidNodes:  obstacleNodes,
		WeightedNodes: waterNodes,
	}

	// create the algo with defined config
	algo, err := New(aConfig)
	if err != nil {
		fmt.Println("invalid astar config", err)
		return
	}

	// run it
	foundPath, err := algo.FindPath(startNode, endNode)
	if err != nil {
		fmt.Println(`err:`, err.Error())
		return
	}
	if len(foundPath) == 0 {
		fmt.Println("No path found ...")
		return
	}

	// the foundPath has now the way to the target

	// IMPORTANT:
	// the path is in the opposite way so the endpoint node is on index 0
	// you can avoid it by switching the startNode<>endNode parameter
	for _, node := range foundPath {
		fmt.Println(node)
	}
}
