package day4

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/seb-emmot/advent-2023/utils"
)

type Card struct {
	Id         int
	OwnNumbers []int
	Winners    []int
}

func NewCard(line string) Card {
	line = strings.ReplaceAll(line, "   ", " ")
	line = strings.ReplaceAll(line, "  ", " ")
	s := strings.Split(line, ": ")
	id, e := strconv.Atoi(strings.Split(s[0], " ")[1])
	utils.Check(e)

	n := strings.Split(s[1], " | ")
	ownStr := n[0]
	winStr := n[1]

	ownNumbers := []int{}
	for _, num := range strings.Split(ownStr, " ") {
		n, e := strconv.Atoi(num)
		utils.Check(e)
		ownNumbers = append(ownNumbers, n)
	}

	winNumbers := []int{}
	for _, num := range strings.Split(winStr, " ") {
		n, e := strconv.Atoi(num)
		utils.Check(e)
		winNumbers = append(winNumbers, n)
	}

	return Card{Id: id, OwnNumbers: ownNumbers, Winners: winNumbers}
}

func (c Card) CalcPoints() int {
	pts := 0
	for _, num := range c.OwnNumbers {
		if slices.Contains(c.Winners, num) {
			if pts == 0 {
				pts = 1
			} else {
				pts *= 2
			}
		}
	}
	return pts
}

func (c Card) Print() {
	fmt.Println(c.Id, c.OwnNumbers, c.Winners)
}

func Day4(input []string) {
	points := 0
	for _, line := range input {
		c := NewCard(line)

		pts := c.CalcPoints()
		points += pts

		fmt.Println(pts, points)
	}
}
