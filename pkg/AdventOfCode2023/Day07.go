package AdventOfCode2023

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/kaitachi/go-challenges/internal/lib"
)

type Day07 struct {
	data []hand
}


type HandType int

const (
	FIVE_OF_A_KIND	HandType = 5
	FOUR_OF_A_KIND	HandType = 4
	FULL_HOUSE		HandType = 3
	THREE_OF_A_KIND	HandType = 2
	TWO_PAIR		HandType = 1
	ONE_PAIR		HandType = 0
	HIGH_CARD		HandType = -1
)

var (
	CARD_RANKS = map[rune]int{
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'T': 10,
		'J': 11,
		'Q': 12,
		'K': 13,
		'A': 14,
	}
)


type hand struct {
	cards	string
	bid		int
}


// 1. Assemble - How should we transform the data from our input files?
func (s *Day07) Assemble(tc *lib.TestCase) {

	s.data = make([]hand, 0)

	re_line := regexp.MustCompile(`^(\w{5}) (\d+)$`)

	for _, line := range strings.Split(tc.Input, "\n") {

		for _, match := range re_line.FindAllStringSubmatch(line, -1) {
			bid, _ := strconv.Atoi(match[2])

			s.data = append(s.data, hand{
				cards: match[1],
				bid: bid,
			})
		}
	}
}


// 2. Activate - Take our transformed input data and make the core logic needed to resolve this Problem
func (s *Day07) Activate(tc *lib.TestCase) {
	// Assign final value to TestCase.Actual field
	switch tc.Algorithm {
	case "part01":
		tc.Actual = s.part01()
		
	case "part02":
		tc.Actual = s.part02()
	}
}


func (s Day07) part01() string {

	//fmt.Println(s.data)

	sort.SliceStable(s.data, func(i int, j int) bool {
		hi := s.data[i]
		hj := s.data[j]

		switch {
		// First, sort by Hand Type
		case hi.HandType() != hj.HandType():
			return sortHandsByType(hi, hj)

		// Then, sort by Card arrangement
		default:
			return sortHandsByCards(hi, hj)
		}
	})

	//fmt.Println(s.data)

	winnings := 0

	for i, hand := range s.data {
		winnings += (i+1) * hand.bid
	}

	return fmt.Sprintf("%d", winnings)
}


func (s Day07) part02() string {

	return fmt.Sprintf("%d", -1)
}


func (h hand) HandType() HandType {
	count := make(map[rune]int, 0)

	for _, r := range h.cards {
		count[r]++
	}

	switch len(count) {
	case 1:
		return FIVE_OF_A_KIND

	case 2:
		a := 0

		// Grab card with highest count
		for _, c := range count {
			a = int(math.Max(float64(a), float64(c)))
		}

		switch a {
		case 4:
			return FOUR_OF_A_KIND

		default:
			return FULL_HOUSE
		}

	case 3:
		a := 0

		// Grab card with highest count
		for _, c := range count {
			a = int(math.Max(float64(a), float64(c)))
		}

		switch a {
		case 3:
			return THREE_OF_A_KIND

		default:
			return TWO_PAIR
		}

	case 4:
		return ONE_PAIR

	default:
		return HIGH_CARD
	}
}


func sortHandsByType(l hand, r hand) bool {
	return l.HandType() < r.HandType()
}


func sortHandsByCards(l hand, r hand) bool {

	left := []rune(l.cards)
	right := []rune(r.cards)

	for i := 0; i < 5; i++ {
		if left[i] != right[i] {
			return CARD_RANKS[left[i]] < CARD_RANKS[right[i]]
		}
	}

	// We'll assume that all hands will be distinct for now
	return true
}

