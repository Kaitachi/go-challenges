package AdventOfCode2023

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/kaitachi/go-challenges/internal/lib"
)

type Day10 struct {
	start *cell
	grid []map[int]*cell
}


type cell struct {
	name		string
	row			int
	col			int
	directions	map[string]*cell
}


// 1. Assemble - How should we transform the data from our input files?
func (s *Day10) Assemble(tc *lib.TestCase) {

	s.start = nil
	s.grid = make([]map[int]*cell, 0)

	re_pipe := regexp.MustCompile(`[|\-LJ7FS]`)

	for row, line := range strings.Split(tc.Input, "\n") {
		s.grid = append(s.grid, make(map[int]*cell, 0))
		matches := re_pipe.FindAllString(line, -1)
		count := re_pipe.FindAllStringIndex(line, -1)

		for i := 0; i < len(count); i++ {
			s.grid[row][count[i][0]] = &cell{
				name: matches[i],
				row: row,
				col: count[i][0],
				directions: make(map[string]*cell, 0),
			}

			if matches[i] == "S" {
				s.start = s.grid[row][count[i][0]]
			}
		}
	}

	s.linkPipes()
}


// 2. Activate - Take our transformed input data and make the core logic needed to resolve this Problem
func (s *Day10) Activate(tc *lib.TestCase) {
	// Assign final value to TestCase.Actual field
	switch tc.Algorithm {
	case "part01":
		tc.Actual = s.part01()
		
	case "part02":
		tc.Actual = s.part02()
	}
}


func (s Day10) part01() string {

	// fmt.Printf("%+v\n", s.start)
	// for i := 0; i < len(s.grid); i++ {
	// 	fmt.Printf("%v\n", s.grid[i])
	// 
	// 	for j, pipe := range s.grid[i] {
	// 		fmt.Printf("[%d]: %+v\n", j, pipe)
	// 	}
	//
	// 	fmt.Println()
	// }

	// Let's make two pointers, one travels left and the other travels right
	var left, right *cell
	var prev_left, prev_right *cell
	steps := 0

	for _, path := range s.start.directions {
		switch {
		case left == nil: left = path
		default: right = path
		}
	}

	prev_left = right
	prev_right = left

	for steps = 1; left != right; steps++ {
		prev_left, left = left.next(prev_left)
		prev_right, right = right.next(prev_right)
	}

	fmt.Printf("> steps: %d\n", steps)
	fmt.Printf("> left: %+v\n", left)
	fmt.Printf("> right: %+v\n", right)

	return fmt.Sprintf("%d", steps)
}


func (s Day10) part02() string {

	return fmt.Sprintf("%d", -1)
}


func (s *Day10) linkPipes() {
	rows := 0
	cols := 0

	for _, line := range s.grid {
		rows++

		for col, _ := range line {
			cols = int(math.Max(float64(cols), float64(col)))
		}
	}

	for row, line := range s.grid {
		for col, pipe := range line {
			isStart := pipe.name == "S"

			// Looking North, can we connect to our neighbour?
			if row-1 >= 0 && (pipe.hasPipeFacing("N") || isStart) {
				north := s.grid[row-1][col]

				if north != nil && north.hasPipeFacing("S") {
					pipe.directions["N"] = north
				}
			}

			// Looking East, can we connect to our neighbour?
			if col+1 <= cols && (pipe.hasPipeFacing("E") || isStart) {
				east := s.grid[row][col+1]

				if east != nil && east.hasPipeFacing("W") {
					pipe.directions["E"] = east
				}
			}

			// Looking South, can we connect to our neighbour?
			if row+1 < rows && (pipe.hasPipeFacing("S") || isStart) {
				south := s.grid[row+1][col]

				if south != nil && south.hasPipeFacing("N") {
					pipe.directions["S"] = south
				}
			}

			// Looking West, can we connect to our neighbour?
			if col-1 >= 0 && (pipe.hasPipeFacing("W") || isStart) {
				west := s.grid[row][col-1]

				if west != nil && west.hasPipeFacing("E") {
					pipe.directions["W"] = west
				}
			}
		}
	}
}


func (c cell) hasPipeFacing(d string) bool {

	// Looking North, can we connect to our neighbour?
	switch c.name {
	case "|", "L", "J": if d == "N" { return true }
	}

	// Looking East, can we connect to our neighbour?
	switch c.name {
	case "-", "L", "F": if d == "E" { return true }
	}

	// Looking South, can we connect to our neighbour?
	switch c.name {
	case "|", "7", "F": if d == "S" { return true }
	}

	// Looking West, can we connect to our neighbour?
	switch c.name {
	case "-", "J", "7": if d == "W" { return true }
	}

	return false
}


func (c *cell) next(prev *cell) (*cell, *cell) {
	for _, neighbor := range c.directions {
		if neighbor != prev {
			return c, neighbor
		}
	}

	return nil, nil
}

