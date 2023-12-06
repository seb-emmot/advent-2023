package utils

import "strconv"

func SumArray(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}

	return sum
}

func ToString(ints []int) []string {
	strs := []string{}
	for _, i := range ints {
		s := strconv.Itoa(i)
		strs = append(strs, s)
	}
	return strs
}
