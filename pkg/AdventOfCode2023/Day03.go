package AdventOfCode2023

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/kaitachi/go-challenges/internal/lib"
)

type Day03 struct {
	numbers map[int]map[int]*int
	symbols map[int]map[int]string
}


// 1. Assemble - How should we transform the data from our input files?
func (s *Day03) Assemble(tc *lib.TestCase) {

	s.numbers = make(map[int]map[int]*int, 0)
	s.symbols = make(map[int]map[int]string, 0)

	nums_re := regexp.MustCompile(`\d+`)
	parts_re := regexp.MustCompile(`[^\d.]`)

	for row, line := range strings.Split(tc.Input, "\n") {
		
		s.numbers[row] = make(map[int]*int, 0)

		// Parse all numbers found in row
		nums := nums_re.FindAllString(line, -1)
		nums_indices := nums_re.FindAllStringIndex(line, -1)

		for j := 0; j < len(nums); j++ {
			number, _ := strconv.Atoi(nums[j])

			for col := nums_indices[j][0]; col < nums_indices[j][1]; col++ {
				s.numbers[row][col] = &number
			}
		}

		s.symbols[row] = make(map[int]string, 0)

		// Parse all symbols found in row
		sym := parts_re.FindAllString(line, -1)
		sym_indices := parts_re.FindAllStringIndex(line, -1)

		for j := 0; j < len(sym); j++ {
			col := sym_indices[j][0]
			s.symbols[row][col] = sym[j]
		}
	}

	fmt.Printf("----- GRID -----\n")
	for k, v := range s.numbers {
		fmt.Printf("%d: %v\n", k, v)
	}

	fmt.Printf("----- SYMBOLS -----\n")
	for k, v := range s.symbols {
		fmt.Printf("%d: %v\n", k, v)
	}

}


// 2. Activate - Take our transformed input data and make the core logic needed to resolve this Problem
func (s *Day03) Activate(tc *lib.TestCase) {
	// Assign final value to TestCase.Actual field
	switch tc.Algorithm {
	case "part01":
		tc.Actual = s.part01()
		
	case "part02":
		tc.Actual = s.part02()
	}
}


func (s Day03) part01() string {

	found := make(map[*int]bool, 0)

	// Iterate through symbols list
	for row, line := range s.symbols {
		for col, _ := range line {
			numbers := findPointersAround(&s, row, col)

			// If a number hasn't been found, append it to found map
			for ptr, _ := range numbers {
				found[ptr] = true
			}
		}
	}

	total := 0

	fmt.Printf("----- NUMBERS FOUND ADJACENT TO SYMBOLS -----\n")
	for k, v := range found {
		fmt.Printf("%d: %v\n", k, v)
		total += *k
	}

	fmt.Println(">> ", total)

	return fmt.Sprintf("%d", total)
}


func (s Day03) part02() string {

	total := 0

	// Iterate through symbols list
	for row, line := range s.symbols {
		for col, part := range line {
			if part == "*" {
				subtotal := 1
				numbers := findPointersAround(&s, row, col)

				// Only append multiplication if there are exactly two numbers around this part
				if len(numbers) == 2 {
					for ptr, _ := range numbers {
						subtotal *= *ptr
					}

					total += subtotal
				}
			}
		}
	}

	return fmt.Sprintf("%d", total)
}


func findPointersAround(s *Day03, row int, col int) map[*int]bool {
	found := make(map[*int]bool, 0)

	for i := row-1; i <= row+1; i++ {
		if i < 0 { continue }
		
		for j := col-1; j <= col+1; j++ {
			if j < 0 { continue }
			if s.numbers[i][j] != nil {
				found[s.numbers[i][j]] = true
			}
		}
	}

	return found
}

