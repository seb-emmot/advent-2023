package day06

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/seb-emmot/advent-2023/utils"
)

type Race struct {
	Time       int
	RecordDist int
}

func (r Race) GetDistances() []int {
	results := []int{}
	for i := 0; i < r.Time; i++ {
		speed := 1 * i
		dist := (r.Time - i) * speed
		results = append(results, dist)
	}
	return results
}

func Day6(input []string) {
	timeRow := input[0]

	r, e := regexp.Compile("\\d+")
	utils.Check(e)

	matches := r.FindAllString(timeRow, -1)
	times := []int{}

	for _, m := range matches {
		n, e := strconv.Atoi(m)
		utils.Check(e)
		times = append(times, n)
	}

	distRow := input[1]
	matches = r.FindAllString(distRow, -1)
	records := []int{}
	for _, m := range matches {
		n, e := strconv.Atoi(m)
		utils.Check(e)
		records = append(records, n)
	}

	waysOfWin := 1

	for i := 0; i < len(times); i++ {
		r := Race{Time: times[i], RecordDist: records[i]}
		distances := r.GetDistances()

		candidates := []int{}

		for _, d := range distances {
			if d > r.RecordDist {
				candidates = append(candidates, d)
			}
		}

		fmt.Println("ways of winning", len(candidates))
		waysOfWin *= len(candidates)
	}
	fmt.Println("p1 total ways of winning multiplied", waysOfWin)

	//p2
	p2time, e := strconv.Atoi(strings.Join(utils.ToString(times), ""))
	utils.Check(e)

	p2record, e := strconv.Atoi(strings.Join(utils.ToString(records), ""))
	utils.Check(e)

	fmt.Println(p2time, p2record)

	race := Race{Time: p2time, RecordDist: p2record}

	distances := race.GetDistances()

	candidates := 0

	for _, d := range distances {
		if d > race.RecordDist {
			candidates++
		}
	}
	fmt.Println("p2 ways of winning", candidates)
}
