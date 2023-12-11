package day11

import (
	"fmt"
	"slices"

	"github.com/seb-emmot/advent-2023/utils"
)

type Coord struct {
	X, Y int
}

type Galaxy struct {
	id  int
	pos Coord
}

type GalaxyPair struct {
	p1 Galaxy
	p2 Galaxy
}

func (gp GalaxyPair) GetUID() string {
	mi := min(gp.p1.id, gp.p2.id)
	ma := max(gp.p1.id, gp.p2.id)
	return fmt.Sprint(mi, ma)
}

type Universe struct {
	m        []string
	galaxies []Galaxy
}

func (u Universe) Expand(multiplier int) Universe {
	input := u.m

	emptyRows := []int{}
	emptyCols := []int{}

	for y, row := range input {
		emptyRow := true
		for _, r := range row {
			if r != '.' {
				emptyRow = false
				break
			}
		}
		if emptyRow {
			emptyRows = append(emptyRows, y)
		}
	}

	for x := 0; x < len(input[0]); x++ {
		emptyCol := true
		for y := 0; y < len(input); y++ {
			r := input[y][x]
			if r != '.' {
				emptyCol = false
				break
			}
		}
		if emptyCol {
			emptyCols = append(emptyCols, x)
		}
	}

	// fmt.Println(emptyRows, emptyCols)

	newUniverse := []string{}

	for y, row := range input {
		if multiplier > 1000 {
			fmt.Println("processing row", y, "of", len(input))
		}
		newRow := ""
		for x, r := range row {
			if slices.Contains(emptyCols, x) {
				for i := 0; i < multiplier; i++ {
					// fmt.Println(i)
					newRow += "."
				}
			}
			newRow += string(r)
		}
		if slices.Contains(emptyRows, y) {
			for i := 0; i < multiplier; i++ {
				newUniverse = append(newUniverse, newRow)
			}
		}
		newUniverse = append(newUniverse, newRow)
	}

	return Universe{m: newUniverse}
}

func (u *Universe) MapCoords() {
	galaxies := []Galaxy{}
	id := 0
	for y, row := range u.m {
		for x, r := range row {
			if r != '.' {
				galaxies = append(galaxies, Galaxy{
					pos: Coord{x, y},
					id:  id,
				})
				id++
			}
		}
	}

	u.galaxies = galaxies
}

func (u Universe) Print() {
	fmt.Println(len(u.m), len(u.m[0]))
	for _, row := range u.m {
		fmt.Println(row)
	}
}

func Day11(input []string) {

	uni := Universe{m: input}

	expUni := uni.Expand(1)

	// uni.Print()

	// expUni.Print()
	expUni.MapCoords()

	// fmt.Println(expUni.galaxies)

	pairs := map[string]GalaxyPair{}

	for _, g1 := range expUni.galaxies {
		for _, g2 := range expUni.galaxies {
			if g1.id != g2.id {
				gp := GalaxyPair{p1: g1, p2: g2}
				_, ok := pairs[gp.GetUID()]
				if !ok {
					// fmt.Println("add pair", g1, g2)
					pairs[gp.GetUID()] = gp
				}
			}
		}
	}

	distances := []int{}

	for _, v := range pairs {
		mhdist := utils.Abs(v.p1.pos.X-v.p2.pos.X) + utils.Abs(v.p1.pos.Y-v.p2.pos.Y)
		distances = append(distances, mhdist)
		// fmt.Println(k, v.p1.id, v.p2.id, mhdist)
	}

	total := utils.SumArray(distances)

	fmt.Println("Total dist", total)
}
