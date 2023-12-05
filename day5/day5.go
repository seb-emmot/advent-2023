package day5

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/seb-emmot/advent-2023/utils"
)

type Interval struct {
	min    int
	max    int
	offset int
}

func (i Interval) GetDestination(src int) (int, bool) {
	if src >= i.min && src < i.max {
		return src + i.offset, true
	}
	return 0, false
}

type ReagentMap struct {
	Kind      string
	Intervals []Interval
}

func ParseReagentMap(input []string) ReagentMap {
	rm := ReagentMap{Kind: input[0], Intervals: make([]Interval, 0)}

	for _, line := range input[1:] {
		nums := strings.Split(line, " ")
		dst, e := strconv.Atoi(nums[0])
		utils.Check(e)
		src, e := strconv.Atoi(nums[1])
		utils.Check(e)
		r, e := strconv.Atoi(nums[2])
		utils.Check(e)

		i := Interval{min: src, max: src + r, offset: dst - src}
		rm.Intervals = append(rm.Intervals, i)
	}

	return rm
}

func (rm ReagentMap) GetDestination(src int) int {
	for _, interval := range rm.Intervals {
		// fmt.Println("checking interval", interval)
		val, ok := interval.GetDestination(src)
		if !ok {
			continue
		}
		return val
	}
	// if no mapping exist, destination = src
	return src
}

func part1Seeds(line string) []int {
	seeds := []int{}
	nums := strings.Split(line, " ")

	for _, n := range nums[1:] {
		nr, e := strconv.Atoi(n)
		utils.Check(e)
		seeds = append(seeds, nr)
	}
	return seeds
}

func calculateP2(line string, rms []ReagentMap) int {
	nums := strings.Split(line, " ")

	seed := 0
	curMin := -1
	for i, n := range nums[1:] {
		fmt.Println("processing", i, "/", len(nums[1:]))
		isSeed := (i % 2) == 0
		if isSeed {
			tmp, e := strconv.Atoi(n)
			utils.Check(e)
			seed = tmp
		} else {
			r, e := strconv.Atoi(n)
			utils.Check(e)
			fmt.Println("adding", r, "numbers to seed")
			for i := 0; i < r; i++ {
				cur := seed + i
				for _, rm := range rms {
					cur = rm.GetDestination(cur)
				}
				if curMin == -1 || curMin > cur {
					curMin = cur
				}
			}
		}
		fmt.Println("cur min", curMin)
	}
	return curMin
}

func MyTask(input []string) {
	seeds := []int{}
	maps := []ReagentMap{}
	buf := []string{}
	for _, line := range input[2:] {
		if line == "" {
			if len(buf) > 0 {
				rm := ParseReagentMap(buf)
				maps = append(maps, rm)
				buf = buf[:0]
			}
		} else {
			buf = append(buf, line)
		}
	}
	if len(buf) != 0 {
		rm := ParseReagentMap(buf)
		maps = append(maps, rm)
	}

	seedLine := input[0]
	seeds = part1Seeds(seedLine)

	locations := []int{}
	for _, seed := range seeds {
		cur := seed
		for _, rm := range maps {
			cur = rm.GetDestination(cur)
		}
		locations = append(locations, cur)
	}

	fmt.Println("p1 -- min location", slices.Min(locations))

	// p2
	p2Min := calculateP2(seedLine, maps)
	fmt.Println("current p2 min location", p2Min)
}
