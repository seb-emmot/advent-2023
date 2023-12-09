package day9

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/seb-emmot/advent-2023/utils"
)

type History struct {
	Values []int
}

func (h History) IsZero() bool {
	for _, e := range h.Values {
		if e != 0 {
			return false
		}
	}
	return true
}

func NewHistory(line string) History {
	nums := strings.Split(line, " ")
	vals := []int{}
	for _, n := range nums {
		v, e := strconv.Atoi(n)
		utils.Check(e)
		vals = append(vals, v)
	}
	h := History{Values: vals}
	return h
}

func (h History) GetDiff() History {
	diffs := []int{}
	for i := 1; i < len(h.Values); i++ {
		diff := h.Values[i] - h.Values[i-1]
		diffs = append(diffs, diff)
	}
	return History{Values: diffs}
}

func (h History) PredictNext() int {
	fmt.Println(h)
	if h.IsZero() {
		return 0
	} else {
		lastElement := h.Values[len(h.Values)-1]
		pred := lastElement + h.GetDiff().PredictNext()
		return pred
	}
}

func (h History) PredictPrevious() int {
	fmt.Println(h)
	if h.IsZero() {
		return 0
	} else {
		first := h.Values[0]
		pred := first - h.GetDiff().PredictPrevious()
		return pred
	}
}

func Day9(input []string) {
	nexts := []int{}
	prevs := []int{}

	for _, row := range input {
		h := NewHistory(row)
		next := h.PredictNext()
		fmt.Println(next)
		nexts = append(nexts, next)

		previous := h.PredictPrevious()
		fmt.Println(previous)
		prevs = append(prevs, previous)
	}

	p1Sum := utils.SumArray(nexts)
	fmt.Println("p1 sum", p1Sum)

	p2Sum := utils.SumArray(prevs)
	fmt.Println("p2 sum", p2Sum)
}
