package main

import (
	"github.com/seb-emmot/advent-2023/day8"
	"github.com/seb-emmot/advent-2023/utils"
)

func main() {
	input := utils.ReadFile("inputs/day8.txt")
	// input := utils.ReadFile("inputs/day8_sample.txt")
	// input := utils.ReadFile("inputs/day8_sample2.txt")
	// input := utils.ReadFile("inputs/day8_sampleb.txt")

	day8.Day8(input)
}
