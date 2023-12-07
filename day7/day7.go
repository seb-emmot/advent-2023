package day7

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/seb-emmot/advent-2023/utils"
)

type HandType int

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	TheeOfKind
	FullHouse
	FourOfKind
	FiveOfKind
)

type Hand struct {
	Cards   string
	P1Value HandType
	P2Value HandType
	Bet     int
}

func (h Hand) IsLess(other Hand) bool {
	if h.P1Value != other.P1Value {
		return h.P1Value < other.P1Value
	} else {
		valor := []rune{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}
		for i := 0; i < len(h.Cards); i++ {
			r1 := rune(h.Cards[i])
			r2 := rune(other.Cards[i])

			r1Value := slices.Index(valor, r1)
			r2Value := slices.Index(valor, r2)

			if r1Value != r2Value {
				return r1Value < r2Value
			}
		}
	}
	panic("could not determine greater haand")
}

func (h Hand) IsLessP2(other Hand) bool {
	if h.P2Value != other.P2Value {
		return h.P2Value < other.P2Value
	} else {
		valor := []rune{'J', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A'}
		for i := 0; i < len(h.Cards); i++ {
			r1 := rune(h.Cards[i])
			r2 := rune(other.Cards[i])

			r1Value := slices.Index(valor, r1)
			r2Value := slices.Index(valor, r2)

			if r1Value != r2Value {
				return r1Value < r2Value
			}
		}
	}
	panic("could not determine greater haand")
}

func GetType(s string) HandType {
	valor := []rune{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}

	counts := []int{}

	for _, v := range valor {
		cnt := strings.Count(s, string(v))
		// fmt.Println("found", cnt, string(v))
		counts = append(counts, cnt)
	}
	slices.Sort(counts)

	lastI := len(counts) - 1

	switch counts[lastI] {
	case 5:
		return FiveOfKind
	case 4:
		return FourOfKind
	case 3:
		if counts[lastI-1] == 2 {
			return FullHouse
		} else {
			return TheeOfKind
		}
	case 2:
		if counts[lastI-1] == 2 {
			return TwoPair
		} else {
			return OnePair
		}
	case 1:
		return HighCard
	}
	return HighCard
}

func GetP2Type(s string) HandType {
	valor := []rune{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A'}

	counts := []int{}

	jokers := strings.Count(s, string('J'))

	for _, v := range valor {
		cnt := strings.Count(s, string(v))
		// fmt.Println("found", cnt, string(v))
		counts = append(counts, cnt)
	}
	slices.Sort(counts)

	lastI := len(counts) - 1

	switch counts[lastI] + jokers {
	case 5:
		return FiveOfKind
	case 4:
		return FourOfKind
	case 3:
		if counts[lastI-1] == 2 {
			return FullHouse
		} else {
			return TheeOfKind
		}
	case 2:
		if counts[lastI-1] == 2 {
			return TwoPair
		} else {
			return OnePair
		}
	case 1:
		return HighCard
	}
	return HighCard
}

func Day7(input []string) {
	hands := []Hand{}
	for _, line := range input {
		parts := strings.Split(line, " ")
		bet, e := strconv.Atoi(parts[1])
		utils.Check(e)
		h := Hand{
			Cards:   parts[0],
			P1Value: GetType(parts[0]),
			P2Value: GetP2Type(parts[0]),
			Bet:     bet,
		}
		fmt.Println(h)
		hands = append(hands, h)
	}

	hCpy := make([]Hand, len(hands))
	copy(hCpy, hands)

	sort.Slice(hCpy, func(i, j int) bool { return hCpy[i].IsLess(hCpy[j]) })
	fmt.Println(hCpy)

	score := 0
	for i, hand := range hCpy {
		inc := (i + 1) * hand.Bet
		score += inc
		fmt.Println("score hand", inc)
	}

	fmt.Println("P1 Total Score", score)

	// p2
	hCpy = make([]Hand, len(hands))
	copy(hCpy, hands)

	sort.Slice(hCpy, func(i, j int) bool { return hCpy[i].IsLessP2(hCpy[j]) })
	fmt.Println(hCpy)

	p2score := 0
	for i, hand := range hCpy {
		inc := (i + 1) * hand.Bet
		p2score += inc
		fmt.Println("score hand", inc)
	}

	fmt.Println("P2 Total Score", p2score)
}
