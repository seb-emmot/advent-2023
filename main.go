package main

import (
	"github.com/seb-emmot/advent-2023/day10"
	"github.com/seb-emmot/advent-2023/utils"
)

func main() {
	input := utils.ReadFile("inputs/day10.txt")
	// input := utils.ReadFile("inputs/day10_sample.txt")
	// input := utils.ReadFile("inputs/day10_sample2.txt")
	// input := utils.ReadFile("inputs/day10_sample3.txt")
	// input := utils.ReadFile("inputs/day10_sample4.txt")

	day10.Day10(input)
}
