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

	rows int
	cols int
}


type cell struct {
	name		string
	row			int
	col			int
	isMainLoop	bool
	isFilled	bool
	directions	map[string]*cell
}


// 1. Assemble - How should we transform the data from our input files?
func (s *Day10) Assemble(tc *lib.TestCase) {

	s.rows = 0
	s.cols = 0
	s.start = nil
	s.grid = make([]map[int]*cell, 0)

	re_pipe := regexp.MustCompile(`[|\-LJ7FS.]`)

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

	steps := s.drawLoop()

	return fmt.Sprintf("%d", steps)
}


func (s Day10) part02() string {

	s.drawLoop()


	// Let's replace our starting position
	_, okNorth := s.start.directions["N"]
	_, okEast := s.start.directions["E"]
	_, okSouth := s.start.directions["S"]
	_, okWest := s.start.directions["W"]
	links := fmt.Sprintf("%5t-%5t-%5t-%5t", okNorth, okEast, okSouth, okWest)

	switch links {
	case "false- true- true-false": s.start.name = "F"
	case " true- true-false-false": s.start.name = "L"
	case "false-false- true- true": s.start.name = "7"
	case " true-false-false- true": s.start.name = "J"
	case " true-false- true-false": s.start.name = "|"
	case "false- true-false- true": s.start.name = "-"
	}

	inner := make([]*cell, 0)
	inLoopCol := make([]bool, s.cols) // Scanning from top to bottom, could we say this cell could be within our loop?
	loopColCorner := make([]string, s.cols)

	for row, _ := range s.grid {
		inLoopRow := false

		for col := 0; col < s.cols; col++ {
			cell := s.grid[row][col]

			if cell.isMainLoop {
				// Are we entering/exiting the loop?
				switch cell.name {
				case "|":
					inLoopRow = !inLoopRow

				case "-", "7", "F":
					inLoopCol[col] = !inLoopCol[col]
					switch cell.name {
					case "7", "F":
						loopColCorner[col] = cell.name
					}

				case "L":
					inLoopRow = !inLoopRow

					if loopColCorner[col] == "F" {
						inLoopCol[col] = !inLoopCol[col]
						loopColCorner[col] = ""
					}

				case "J":
					inLoopRow = !inLoopRow

					if loopColCorner[col] == "7" {
						inLoopCol[col] = !inLoopCol[col]
						loopColCorner[col] = ""
					}

				}
			} else if inLoopCol[col] && inLoopRow {
				// We're not touching the loop. This cell would then be inside of our loop.
				cell.isFilled = true
				inner = append(inner, cell)
			}

			if row == 4 && col == 7 {
				fmt.Printf(">>>> (%d, %d): inLoopRow: %t; inLoopCol: %t\n", row, col, inLoopRow, inLoopCol[col])
			}
		}
	}

	fmt.Printf("%+v\n", s.start)
	// for i := 0; i < len(s.grid); i++ {
	// 	fmt.Printf("[%d]: %v\n", i, s.grid[i])
	//
	// 	// for j, pipe := range s.grid[i] {
	// 	// 	fmt.Printf("[%d]: %+v\n", j, pipe)
	// 	// }
	// }

	fmt.Printf("%v\n", inner)

	s.printMainLoop()

	return fmt.Sprintf("%d", len(inner))
}


func (s *Day10) drawLoop() int {
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

	// Mark first and last nodes that are visited as part of the loop
	s.start.isMainLoop = true
	left.isMainLoop = true

	return steps
}


func (s *Day10) linkPipes() {

	for _, line := range s.grid {
		s.rows++

		for col, _ := range line {
			s.cols = int(math.Max(float64(s.cols), float64(col)))
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
			if col+1 <= s.cols && (pipe.hasPipeFacing("E") || isStart) {
				east := s.grid[row][col+1]

				if east != nil && east.hasPipeFacing("W") {
					pipe.directions["E"] = east
				}
			}

			// Looking South, can we connect to our neighbour?
			if row+1 < s.rows && (pipe.hasPipeFacing("S") || isStart) {
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


func (c cell) isPipe() bool {
	return c.name != "."
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
	c.isMainLoop = true

	for _, neighbor := range c.directions {
		if neighbor != prev {
			return c, neighbor
		}
	}

	return nil, nil
}


func (s *Day10) printMainLoop() {

	fmt.Printf("%d rows, %d cols\n", s.rows, s.cols)
	for row := 0; row < s.rows; row++ {
		fmt.Printf("[%d]: ", row)
	
		for col := 0; col < s.cols; col++ {
			cell := s.grid[row][col]

			switch {
			case cell.isMainLoop:	fmt.Printf("%s", cell.name)
			case cell.isFilled:		fmt.Printf("I")
			default:				fmt.Printf(".")
			}
		}

		fmt.Printf("\n")
	}
}

