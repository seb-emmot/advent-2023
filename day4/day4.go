package day4

import (
	"fmt"
	"math"
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
	wins := c.GetWinnerCount()

	var pts int
	if wins > 0 {
		pts = int(math.Pow(2.0, float64(wins-1)))

	}
	return pts
}

func (c Card) GetWinnerCount() int {
	wins := 0
	for _, num := range c.OwnNumbers {
		if slices.Contains(c.Winners, num) {
			wins++
		}
	}
	return wins
}

func (c Card) Print() {
	fmt.Println(c.Id, c.OwnNumbers, c.Winners)
}

func Day4(input []string) {
	points := 0

	duplicates := map[int]int{}

	for _, line := range input {
		c := NewCard(line)
		cardCount, ok := duplicates[c.Id]
		if !ok {
			cardCount = 1
		} else {
			// dupes + original
			cardCount++
		}

		pts := c.CalcPoints()
		points += pts

		wins := c.GetWinnerCount()
		for i := 1; i <= wins; i++ {
			dupes, ok := duplicates[c.Id+i]
			if !ok {
				dupes = 0
			}
			dupes += cardCount
			duplicates[c.Id+i] = dupes
		}
	}

	fmt.Println("p1 points", points)

	sumDupes := 0
	for _, val := range duplicates {
		sumDupes += val
	}

	total := len(input) + sumDupes

	fmt.Println("p2 total cards", total)
}
