package day10

import (
	"fmt"
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (d Direction) IsInverse(d2 Direction) bool {
	if Abs(int(d-d2)) == 2 {
		return true
	}
	return false
}

func (d Direction) GetInverse() Direction {
	switch d {
	case North:
		return South
	case South:
		return North
	case West:
		return East
	case East:
		return West
	default:
		panic("erros")
	}
}

func CalcType(d1, d2 Direction) PipeType {
	fmt.Println(d1, d2)
	switch d1 {
	case North:
		switch d2 {
		case South:
			return NorthSouth
		case East:
			return NorthEast
		case West:
			return NorthWest
		}
	case East:
		switch d2 {
		case West:
			return EastWest
		}
	case South:
		switch d2 {
		case West:
			return SouthWest
		case East:
			return SouthEast
		}
	}
	panic("did not find case")
}

func GetDirections(s PipeType) (Direction, Direction, bool) {
	switch s {
	case NorthSouth:
		return North, South, true
	case EastWest:
		return East, West, true
	case NorthEast:
		return North, East, true
	case NorthWest:
		return North, West, true
	case SouthWest:
		return South, West, true
	case SouthEast:
		return South, East, true
	default:
		return -1, -1, false
	}
}

type PipeType string

const (
	NorthSouth PipeType = "|"
	EastWest   PipeType = "-"
	NorthEast  PipeType = "L"
	NorthWest  PipeType = "J"
	SouthWest  PipeType = "7"
	SouthEast  PipeType = "F"
)

func ToPipeType(s string) (PipeType, bool) {
	switch s {
	case string(NorthSouth):
		return NorthSouth, true
	case string(EastWest):
		return EastWest, true
	case string(NorthEast):
		return NorthEast, true
	case string(NorthWest):
		return NorthWest, true
	case string(SouthWest):
		return SouthWest, true
	case string(SouthEast):
		return SouthEast, true
	default:
		return "", false
	}
}

type Pipe struct {
	pType   PipeType
	pos     Coord
	prev    *Pipe
	prevDir Direction
	next    *Pipe
}

func (p Pipe) GetNextPipePos() (Coord, Direction) {
	if p.prev == nil {
		panic("Has no previous pipe")
	}
	if p.pType == "" {
		panic("Has no pipe type set")
	}

	d1, d2, ok := GetDirections(p.pType)
	if !ok {
		panic("once again")
	}

	var c Coord
	var usedDir Direction
	if d1 != p.prevDir {
		usedDir = d1
	} else if d2 != p.prevDir {
		usedDir = d2
	} else {
		panic("invalid")
	}

	c = GetCoord(p.pos, usedDir)

	return c, usedDir
}

func (p Pipe) GetLoopLength() int {
	cur := p.next
	cnt := 1
	for cur.pos != p.pos {
		cur = cur.next
		cnt++
	}
	return cnt
}

func (p Pipe) Print(start *Pipe) {
	fmt.Println(p.pos, p.pType, p.next.pos)
	if p.next != start {
		p.next.Print(start)
	}
}

func GetCoord(pos Coord, dir Direction) Coord {
	var coord Coord
	switch dir {
	case North:
		coord = Coord{pos.X, pos.Y - 1}
	case East:
		coord = Coord{pos.X + 1, pos.Y}
	case South:
		coord = Coord{pos.X, pos.Y + 1}
	case West:
		coord = Coord{pos.X - 1, pos.Y}
	}
	return coord
}

type Coord struct {
	X, Y int
}

func Day10(input []string) {
	pipes := map[Coord]*Pipe{}

	var startPipe Pipe

	for y, row := range input {
		for x, r := range row {
			if r == 'S' {
				// start pos found.
				fmt.Println("Found start pos")

				dirs := []Direction{North, East, South, West}

				startCoord := Coord{X: x, Y: y}
				startPipe = Pipe{pos: startCoord}
				pipes[startCoord] = &startPipe

				for _, fromStartdir := range dirs {
					potential := GetCoord(startPipe.pos, fromStartdir)
					if potential.X < 0 || potential.Y < 0 || potential.X >= len(row) || potential.Y >= len(input) {
						// invalid position
						continue
					}

					pipeType, ok := ToPipeType(string(rune(input[potential.Y][potential.X])))
					if !ok {
						continue
					}

					dir1, dir2, ok := GetDirections(pipeType)
					if !ok {
						panic("invalid input in sequence")
					}

					if fromStartdir.IsInverse(dir1) || fromStartdir.IsInverse(dir2) {
						newPipe := Pipe{pos: potential, pType: pipeType, prevDir: fromStartdir.GetInverse()}
						startPipe.next = &newPipe
						newPipe.prev = &startPipe
						pipes[potential] = &newPipe
						break
					}
				}
				done := false
				curPipe := startPipe.next
				fmt.Println(pipes)
				for !done {
					potential, fromCurToNew := curPipe.GetNextPipePos()

					start, ok := pipes[potential]

					if ok {
						done = true
						fmt.Println("We're done")
						curPipe.next = start
						start.prev = curPipe

						fmt.Println(fromCurToNew, start.next.prevDir)
						startPType := CalcType(start.next.prevDir.GetInverse(), fromCurToNew.GetInverse())
						start.pType = startPType

						continue
					}

					if potential.X < 0 || potential.Y < 0 || potential.X >= len(row) || potential.Y >= len(input) {
						// invalid position
						panic("pipe ran off edge")
					}

					pipeType, ok := ToPipeType(string(rune(input[potential.Y][potential.X])))
					if !ok {
						panic("pipe ran out")
					}

					newPipe := Pipe{
						pos:     potential,
						pType:   pipeType,
						prev:    curPipe,
						prevDir: fromCurToNew.GetInverse(),
					}

					pipes[potential] = &newPipe

					curPipe.next = &newPipe
					curPipe = &newPipe
				}
			}
		}
	}
	startPipe.Print(&startPipe)
	fmt.Println("len of loop", startPipe.GetLoopLength())

	RunP2(input, pipes)
}

func RunP2(input []string, pipes map[Coord]*Pipe) {
	inside := false
	areaCnt := 0
	for y, row := range input {
		cameFrom := West
		for x, _ := range row {
			isPipe := false
			coord := Coord{x, y}
			pipe, ok := pipes[coord]

			if !ok {
				fmt.Println(coord, "not in path")
			} else {
				fmt.Println(coord, pipe.pType, "in path")
				d1, _, ok := GetDirections(pipe.pType)
				if !ok {
					// not a pipe
					continue
				} else {
					isPipe = true
					if pipe.pType == NorthSouth {
						// passing NS border
						inside = !inside
						fmt.Println("Passing border, now ", inside)
					} else if pipe.pType == SouthEast || pipe.pType == NorthEast {
						inside = !inside
						cameFrom = d1
					} else if pipe.pType == SouthWest || pipe.pType == NorthWest {
						if d1 == cameFrom.GetInverse() {
							// part of same border, do nothing
						} else {
							inside = !inside
						}
					}
				}
			}
			if !isPipe && inside {
				areaCnt++
				fmt.Println("-- inside area", areaCnt)
			}
		}
	}

	fmt.Println("areacount", areaCnt)
}
