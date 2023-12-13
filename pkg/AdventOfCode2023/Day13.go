package AdventOfCode2023

import (
	"fmt"
	"strings"

	"github.com/kaitachi/go-challenges/internal/lib"
)

type Day13 struct {
	grids [][]string
}


// 1. Assemble - How should we transform the data from our input files?
func (s *Day13) Assemble(tc *lib.TestCase) {

	s.grids = make([][]string, 0)

	for _, grid := range strings.Split(tc.Input, "\n\n") {
		s.grids = append(s.grids, strings.Split(grid, "\n"))
	}
}


// 2. Activate - Take our transformed input data and make the core logic needed to resolve this Problem
func (s *Day13) Activate(tc *lib.TestCase) {
	// Assign final value to TestCase.Actual field
	switch tc.Algorithm {
	case "part01":
		tc.Actual = s.part01()
		
	case "part02":
		tc.Actual = s.part02()
	}
}


func (s Day13) part01() string {

	summary := 0

	for _, grid := range s.grids {
		orientation, index := findMirrorIndex(grid)

		fmt.Printf("> FOUND %s MIRROR AT INDEX %d\n", orientation, index)
		switch orientation {
		case "HORIZONTAL": summary += 100 * index
		case "VERTICAL": summary += index
		}
	}

	return fmt.Sprintf("%d", summary)
}


func (s Day13) part02() string {

	return fmt.Sprintf("%d", -1)
}


func findMirrorIndex(grid []string) (orientation string, index int) {
	for i, line := range grid {
		fmt.Printf("[%d]: %v\n", i+1, line)
	}

	orientation = "HORIZONTAL"
	index = findMirror(grid)

	fmt.Printf(">>> TRANSPOSING <<<\n")

	for i, line := range transpose(grid) {
		fmt.Printf("[%d]: %v\n", i+1, line)
	}

	if index < 0 {
		orientation = "VERTICAL"
		index = findMirror(transpose(grid))
	}

	return orientation, index
}


func findMirror(grid []string) int {
	
	index := -1

	// Searching horizontally, can we find any reflections?
	step: for pivot := 0; pivot < len(grid)-1; pivot++ {
		for offset := 0; offset <= pivot && pivot+offset+1 < len(grid); offset++ {
			subject := pivot-offset // Where our subject is standing
			image := pivot+offset+1 // Where our reflection should be
			// fmt.Printf("[pivot: %d, offset: %d]: Comparing %d and %d...\n", pivot, offset, subject, image)

			if grid[subject] != grid[image] {
				continue step
			}
		}

		// fmt.Printf("> NOW? %d", pivot)
		index = pivot+1
		break
	}

	return index
}


func transpose(grid []string) []string {
	t := make([]string, len(grid[0]))

	for _, line := range grid {
		for col, item := range line {
			t[col] += string(item)
		}
	}

	return t
}

