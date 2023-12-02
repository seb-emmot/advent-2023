package day2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/seb-emmot/advent-2023/utils"
)

type Game struct {
	Id   int
	Sets []map[string]int
}

func NewGame(s string) Game {
	sSplit := strings.Split(s, ": ")
	gameInfo := sSplit[0]
	gameNr, e := strconv.Atoi(strings.Split(gameInfo, " ")[1])
	utils.Check(e)

	sets := []map[string]int{}

	i := sSplit[1]
	subs := strings.Split(i, "; ")

	for _, set := range subs {
		setElement := strings.Split(set, ", ")
		m := map[string]int{}
		for _, q := range setElement {
			parts := strings.Split(q, " ")
			value, e := strconv.Atoi(parts[0])
			utils.Check(e)
			m[parts[1]] = value
		}
		sets = append(sets, m)
	}

	g := Game{Id: gameNr, Sets: sets}

	return g
}

func (g Game) IsConstructable(comparable map[string]int) bool {
	for _, set := range g.Sets {
		for k, v := range set {
			// fmt.Println("checking ", k, v)
			sValue, ok := comparable[k]
			if !ok {
				fmt.Println("set does not contain any ", k, "cubes")
				panic("invalid input comparable set")
			}
			if sValue < v {
				return false
			}
		}
	}
	return true
}

func (g Game) GetMinimumViableSet() map[string]int {
	mvs := map[string]int{}
	for _, set := range g.Sets {
		for setColor, setCubes := range set {
			minCubes, ok := mvs[setColor]
			if !ok || minCubes < setCubes {
				mvs[setColor] = setCubes
			}
		}
	}

	return mvs
}

func CalculateSetPower(set map[string]int) int {
	setPower := 1
	for _, count := range set {
		setPower *= count
	}

	return setPower
}

func (g Game) Print() {
	fmt.Println(g.Id, g.Sets)
}

func Day2(input []string) {

	buildableGames := []int{}
	setPowers := []int{}
	compareSet := map[string]int{"red": 12, "green": 13, "blue": 14}

	for _, e := range input {

		g := NewGame(e)
		isBuildable := g.IsConstructable(compareSet)

		if isBuildable {
			buildableGames = append(buildableGames, g.Id)
		}

		mvs := g.GetMinimumViableSet()
		setPower := CalculateSetPower(mvs)
		setPowers = append(setPowers, setPower)

		fmt.Println("mvs for", g.Id, "is", mvs, "setpower is", setPower)
	}
	sumId := utils.SumArray(buildableGames)
	sumSetPower := utils.SumArray(setPowers)

	fmt.Println("sum of buildable games ids", sumId, "with compare set", compareSet)
	fmt.Println("sum of setpowers for minimum viable sets for all games:", sumSetPower)
}
