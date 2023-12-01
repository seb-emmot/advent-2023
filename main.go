package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/seb-emmot/advent-2023/utils"
)

func main() {
	input := utils.ReadFile("inputs/day1a.txt")

	newStrings := make([]string, 0)

	for _, element := range input {

		firstIndex := -1
		firstFound := ""

		toCheck := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

		for _, digit := range toCheck {
			idx := strings.Index(element, digit)
			if idx > -1 && idx < firstIndex || firstIndex == -1 {
				firstIndex = idx
				firstFound = digit
			}
		}

		lastIndex := -1
		lastFound := ""

		for _, digit := range toCheck {
			idx := strings.LastIndex(element, digit)
			if idx > -1 && idx > lastIndex || lastIndex == -1 {
				lastIndex = idx
				lastFound = digit
			}
		}

		finalstring := firstFound + lastFound

		finalstring = strings.Replace(finalstring, "one", "1", -1)
		finalstring = strings.Replace(finalstring, "two", "2", -1)
		finalstring = strings.Replace(finalstring, "three", "3", -1)
		finalstring = strings.Replace(finalstring, "four", "4", -1)
		finalstring = strings.Replace(finalstring, "five", "5", -1)
		finalstring = strings.Replace(finalstring, "six", "6", -1)
		finalstring = strings.Replace(finalstring, "seven", "7", -1)
		finalstring = strings.Replace(finalstring, "eight", "8", -1)
		finalstring = strings.Replace(finalstring, "nine", "9", -1)

		fmt.Println(finalstring)

		newStrings = append(newStrings, finalstring)
	}

	total := 0
	for _, news := range newStrings {
		i, e := strconv.Atoi(news)
		utils.Check(e)
		total += i
	}

	fmt.Println(total)
}
