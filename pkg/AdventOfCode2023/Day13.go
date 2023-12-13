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
		summary += getMirrorSummary(grid, false)
	}

	return fmt.Sprintf("%d", summary)
}


func (s Day13) part02() string {

	summary := 0

	for _, grid := range s.grids {
		summary += getMirrorSummary(grid, true)
	}

	return fmt.Sprintf("%d", summary)
}


func getMirrorSummary(grid []string, cleanUp bool) int {
	orientation, index := findMirrorIndex(grid, cleanUp)

	fmt.Printf("> FOUND %s MIRROR AT INDEX %d\n", orientation, index)
	switch orientation {
	case "HORIZONTAL": return 100 * index
	case "VERTICAL": return index
	}

	return 0
}


func findMirrorIndex(grid []string, cleanUp bool) (orientation string, index int) {
	for i, line := range grid {
		fmt.Printf("[%d]: %v\n", i+1, line)
	}

	orientation = "HORIZONTAL"
	index = findMirror(grid, cleanUp)

	if index < 0 {
		fmt.Printf(">>> TRANSPOSING <<<\n")

		for i, line := range transpose(grid) {
			fmt.Printf("[%d]: %v\n", i+1, line)
		}
		orientation = "VERTICAL"
		index = findMirror(transpose(grid), cleanUp)
	}

	return orientation, index
}


func findMirror(grid []string, cleanUp bool) int {
	
	index := -1
	var isMirror func(string, string) (bool, int)
	if cleanUp {
		isMirror = cleanSmudge
	} else {
		isMirror = simpleMirror
	}

	// Searching horizontally, can we find any reflections?
	step: for pivot := 0; pivot < len(grid)-1; pivot++ {
		smudges := 0

		for offset := 0; offset <= pivot && pivot+offset+1 < len(grid); offset++ {
			subject := pivot-offset // Where our subject is standing
			image := pivot+offset+1 // Where our reflection should be
			// fmt.Printf("[pivot: %d, offset: %d]: Comparing %d and %d...\n", pivot, offset, subject, image)

			reflects, smudge := isMirror(grid[subject], grid[image])
			fmt.Printf("[pivot: %d, offset: %d]: reflects? %t, smudgecount: %d\n", pivot+1, offset, reflects, smudge)

			if !reflects {
				fmt.Printf("> Sorry! No reflection at pivot %d. Found %d smudges.\n", pivot, smudge)
				continue step
			}

			smudges += smudge
		}

		fmt.Printf("> What's going on here??? %d pivot, %d smudges\n", pivot, smudges)

		if cleanUp && smudges != 1 {
			fmt.Printf("> %d is too many/not enough smudges!\n", smudges)
			continue step
		}

		// fmt.Printf("> NOW? %d", pivot)
		index = pivot+1
		break
	}

	return index
}


func simpleMirror(subject string, image string) (bool, int) {
	return subject == image, 0
}


func cleanSmudge(subject string, image string) (bool, int) {
	// If the mirror contains only one smudge, should we need to clean it up...?
	// Yes, we need to make sure there is only one character different
	smudges := 0
	for i := 0; i < len(subject); i++ {
		if subject[i] != image[i] {
			smudges++
		}
	}

	return smudges <= 1, smudges
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

