package day8

import (
	"fmt"
	"strings"
)

type Node struct {
	Value string
	Left  *Node
	Right *Node
}

func (n *Node) SetPaths(left, right *Node) {
	if left != nil {
		if n.Left == nil {
			n.Left = left
		}
	}
	if right != nil {
		if n.Right == nil {
			n.Right = right
		}
	}
}

func (n *Node) Print() {
	fmt.Println(n.Value)
	fmt.Println(n.Left)
	fmt.Println(n.Right)
}

func Day8(input []string) {
	moveSeq := input[0]

	nodeMap := map[string]Node{}

	fmt.Println(moveSeq)

	for _, row := range input[2:] {
		parts := strings.Split(row, " = ")

		value := parts[0]
		nPairs := strings.Split(strings.Trim(parts[1], "()"), ", ")
		n, ok := nodeMap[value]

		if !ok {
			n = Node{Value: value}
		}

		lValue := nPairs[0]
		rValue := nPairs[1]

		l, ok := nodeMap[lValue]

		if !ok {
			l = Node{Value: lValue}
			nodeMap[lValue] = l
		}

		r, ok := nodeMap[rValue]

		if !ok {
			r = Node{Value: rValue}
			nodeMap[rValue] = r
		}

		fmt.Println("setting", n, "l and r to", l, r)
		n.SetPaths(&r, &l)

		fmt.Println(n)

		nodeMap[value] = n

		fmt.Println(row)
	}

	fmt.Println(nodeMap)
	aNode := nodeMap["AAA"]

	fmt.Println(aNode)

	aNode.Print()
}
