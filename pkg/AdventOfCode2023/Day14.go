package AdventOfCode2023

import (
	"fmt"
	"strings"

	"github.com/kaitachi/go-challenges/internal/lib"
)

type Day14 struct {
	data [][]string
}


const (
	SPACE = "."
	ROCK = "#"
	SLIDE = "O"
)


// 1. Assemble - How should we transform the data from our input files?
func (s *Day14) Assemble(tc *lib.TestCase) {

	s.data = make([][]string, 0)

	for row, line := range strings.Split(tc.Input, "\n") {
		s.data = append(s.data, make([]string, len(line)))
		for col, cell := range line {
			s.data[row][col] = string(cell)
		}
	}
}


// 2. Activate - Take our transformed input data and make the core logic needed to resolve this Problem
func (s *Day14) Activate(tc *lib.TestCase) {
	// Assign final value to TestCase.Actual field
	switch tc.Algorithm {
	case "part01":
		tc.Actual = s.part01()
		
	case "part02":
		tc.Actual = s.part02()
	}
}


func (s Day14) part01() string {

	for row, line := range s.data {
		fmt.Printf("[%d]: %+v\n", row, line)
	}

	s.data = tiltNorth(s.data)

	totalLoad := 0

	for row, line := range s.data {
		fmt.Printf("[%d]: %+v\n", row, line)

		for _, cell := range line {
			if cell == SLIDE {
				totalLoad += len(s.data)-row
			}
		}
	}

	return fmt.Sprintf("%d", totalLoad)
}


func (s Day14) part02() string {

	return fmt.Sprintf("%d", -1)
}


func tiltNorth(grid [][]string) [][]string {
	for row, line := range grid {
		for col, cell := range line {
			if row <= 0 {
				continue
			}

			if cell != SLIDE {
				continue
			}

			for i := 1; i <= row; i++ {
				fmt.Printf(">> (%d, %d) checking i=%d\n", row, col, i)
				if grid[row-i][col] != SPACE {
					// Destination is already occupied; let's stop here
					break
				}

				// Swap these two positions
				temp := grid[row-i+1][col]
				grid[row-i+1][col] = grid[row-i][col]
				grid[row-i][col] = temp
			}

			fmt.Printf("> %d, %d\n", row, col)
		}
	}

	return grid
}

