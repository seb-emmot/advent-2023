package day08

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

func (n *Node) Print(prefix string) {
	fmt.Println(prefix + n.Value)
	if n.Value == "ZZZ" {
		return
	}
	if n.Left != n {
		n.Left.Print(prefix + " ")
	}
	if n.Right != n {
		n.Right.Print(prefix + " ")
	}
}

func Day8(input []string) {
	moveSeq := input[0]

	nodeMap := map[string]*Node{}

	// fmt.Println(moveSeq)

	for _, row := range input[2:] {
		parts := strings.Split(row, " = ")

		value := parts[0]
		nPairs := strings.Split(strings.Trim(parts[1], "()"), ", ")
		n, ok := nodeMap[value]

		if !ok {
			n = &Node{Value: value}
		}

		lValue := nPairs[0]
		rValue := nPairs[1]

		l, ok := nodeMap[lValue]

		if !ok {
			l = &Node{Value: lValue}
			nodeMap[lValue] = l
		}

		r, ok := nodeMap[rValue]

		if !ok {
			r = &Node{Value: rValue}
			nodeMap[rValue] = r
		}

		n.SetPaths(l, r)
		nodeMap[value] = n
	}

	node := nodeMap["AAA"]

	stepsP1 := 0
	doneP1 := false

	for !doneP1 {
		for _, r := range moveSeq {
			if !doneP1 {
				if node.Value == "ZZZ" {
					fmt.Println("DONE", "steps", stepsP1)
					doneP1 = true
				} else {
					stepsP1++
				}
				if r == 'L' {
					node = node.Left
				} else if r == 'R' {
					node = node.Right
				} else {
					panic("invalid move")
				}
			}
		}
	}
	fmt.Println("P1", stepsP1)

	p2nodes := []*Node{}

	for k, v := range nodeMap {
		if strings.HasSuffix(k, "A") {
			p2nodes = append(p2nodes, v)
		}
	}

	as := []int{}

	for _, n := range p2nodes[:] {
		cnt := 0
		node := n

		a := 0
		b := 0
		c := 0
		for a == 0 || b == 0 || c == 0 {
			for _, r := range moveSeq {
				if r == 'L' {
					node = node.Left
				} else if r == 'R' {
					node = node.Right
				} else {
					panic("invalid move")
				}
				cnt++
				if strings.HasSuffix(node.Value, "Z") {
					if a == 0 {
						a = cnt
					} else if b == 0 {
						b = cnt
					} else {
						c = cnt
						break
					}
					// fmt.Print(cnt, " ")
				}
			}

			if node.Value == n.Value {
				fmt.Println("break")
				break
			}
		}
		as = append(as, a)
		fmt.Println(a, b-(a*2), c-(a*3))
	}

	result := as[0] * as[1] / GCD(as[0], as[1])

	for i := 2; i < len(as); i++ {
		result = result * as[i] / GCD(result, as[i])
	}

	fmt.Println("P2 res", result)
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
