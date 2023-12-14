package day03

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/seb-emmot/advent-2023/utils"
)

type Part struct {
	Num       int
	C         Coord
	Adjacents []Symbol
}

type Symbol struct {
	C Coord
	S string
}

type Schematic struct {
	Numbers map[Coord]int
	Symbols map[Coord]string
}

type Coord struct {
	X int
	Y int
}

func NewSchematic(grid []string) Schematic {
	s := Schematic{
		Numbers: map[Coord]int{},
		Symbols: map[Coord]string{},
	}

	var buf string
	var coord Coord

	for y, row := range grid {
		for x, r := range row {
			if r == '.' && len(buf) == 0 {
				continue

			} else if r == '.' {
				n, e := strconv.Atoi(string(buf))
				utils.Check(e)
				s.Numbers[coord] = n
				buf = ""

			} else if unicode.IsDigit(r) {
				if len(buf) == 0 {
					coord = Coord{X: x, Y: y}
				}
				buf = buf + string(r)

			} else {
				if len(buf) > 0 {
					n, e := strconv.Atoi(string(buf))
					utils.Check(e)
					s.Numbers[coord] = n
					buf = ""
				}
				c := Coord{X: x, Y: y}
				s.Symbols[c] = string(r)
			}
		}
		if len(buf) > 0 {
			n, e := strconv.Atoi(string(buf))
			utils.Check(e)
			s.Numbers[coord] = n
			buf = ""
		}
	}

	return s
}

func (s Schematic) GetParts() []Part {
	parts := []Part{}
	for pos, candidatePart := range s.Numbers {
		symbols := []Symbol{}
		adjCandidates := s.GetAdajacentPositions(pos)
		isValidPart := false
		for _, c := range adjCandidates {
			symbol, isReal := s.GetSymbol(c)
			if isReal {
				s := Symbol{C: c, S: symbol}
				symbols = append(symbols, s)
				isValidPart = true
			}
		}

		if isValidPart {
			p := Part{Num: candidatePart, C: pos, Adjacents: symbols}
			parts = append(parts, p)
		}
	}
	return parts
}

func (s Schematic) GetSymbol(c Coord) (string, bool) {
	sym, ok := s.Symbols[c]
	if !ok {
		return "", false
	}

	return sym, true
}

func (s Schematic) GetAdajacentPositions(c Coord) []Coord {
	num, ok := s.Numbers[c]
	if !ok {
		panic("num does not exist")
	}
	l := len(strconv.Itoa(num))

	coords := []Coord{}

	for i := 0; i < l; i++ {
		x := c.X + i
		y := c.Y
		if i == 0 {
			// first pos of number
			coords = append(coords,
				Coord{X: x - 1, Y: y - 1},
				Coord{X: x - 1, Y: y},
				Coord{X: x - 1, Y: y + 1},
				Coord{X: x, Y: y - 1},
				Coord{X: x, Y: y + 1},
			)
		}
		if i == l-1 {
			// last pos of number
			coords = append(coords,
				Coord{X: x, Y: y - 1},
				Coord{X: x, Y: y + 1},
				Coord{X: x + 1, Y: y - 1},
				Coord{X: x + 1, Y: y},
				Coord{X: x + 1, Y: y + 1},
			)
		}
		if i > 0 && i < l-1 {
			// middle pos of number
			coords = append(coords,
				Coord{X: x, Y: y - 1},
				Coord{X: x, Y: y + 1},
			)
		}
	}

	return coords
}

func GetGears(parts []Part) map[Symbol][]int {
	m := map[Symbol][]int{}
	for _, p := range parts {
		for _, adj := range p.Adjacents {
			if adj.S != "*" {
				continue
			}
			v, ok := m[adj]
			if !ok {
				v = make([]int, 0)

			}
			v = append(v, p.Num)
			m[adj] = v
		}
	}

	return m
}

func (s Schematic) Print() {
	fmt.Println(s.Numbers)
	fmt.Println(s.Symbols)
}

func Day3(input []string) {
	s := NewSchematic(input)

	// s.Print()

	parts := s.GetParts()

	sum := 0
	for _, p := range parts {
		sum += p.Num
	}

	fmt.Println("Part Id Sum", sum)

	gearMap := GetGears(parts)
	gearValue := 0
	for _, v := range gearMap {
		if len(v) != 2 {
			continue
		}
		gearValue += v[0] * v[1]
	}

	fmt.Println("Gear value", gearValue)

}
