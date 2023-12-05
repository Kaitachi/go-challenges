package AdventOfCode2023

import (
	"fmt"
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/kaitachi/go-challenges/internal/lib"
)

type Day04 struct {
	data []card
}


type card struct {
	id		int
	numbers	[]int
	winning	[]int
}


// 1. Assemble - How should we transform the data from our input files?
func (s *Day04) Assemble(tc *lib.TestCase) {
	
	input := strings.Split(tc.Input, "\n")
	s.data = make([]card, len(input))

	for row, line := range input {
		card := card{
			id: row,
		}

		re_parts := regexp.MustCompile(`^(?P<Card>Card\s+\d+:)\s+(?P<Numbers>.*)\s+\|\s+(?P<Calls>.*)$`)
		re_nums := regexp.MustCompile(`\d+`)

		found_parts := re_parts.FindStringSubmatch(line)
		found_board := re_nums.FindAllString(found_parts[2], -1)
		found_calls := re_nums.FindAllString(found_parts[3], -1)

		board := make([]int, len(found_board))
		for i, v := range found_board {
			board[i], _ = strconv.Atoi(v)
		}

		calls := make([]int, len(found_calls))
		for i, v := range found_calls {
			calls[i], _ = strconv.Atoi(v)
		}

		card.numbers = board
		card.winning = calls

		s.data[row] = card
	}
}


// 2. Activate - Take our transformed input data and make the core logic needed to resolve this Problem
func (s *Day04) Activate(tc *lib.TestCase) {
	// Assign final value to TestCase.Actual field
	switch tc.Algorithm {
	case "part01":
		tc.Actual = s.part01()
		
	case "part02":
		tc.Actual = s.part02()
	}
}


func (s Day04) part01() string {

	total := 0

	for _, card := range s.data {
		matches := float64(0)

		for _, call := range card.winning {
			if slices.Contains(card.numbers, call) {
				matches += 1
			}
		}

		if matches > 0 {
			total += int(math.Pow(2, matches - 1))
		}
	}

	return fmt.Sprintf("%d", total)
}


func (s Day04) part02() string {

	return fmt.Sprintf("%d", -1)
}

