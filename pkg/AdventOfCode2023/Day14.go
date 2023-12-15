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


var (
	// {dy, dx}
	NORTH = [2]int{-1, 0}
	EAST = [2]int{0, 1}
	SOUTH = [2]int{1, 0}
	WEST = [2]int{0, -1}
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

	tilt(&s.data, NORTH)

	totalLoad := 0

	for row, line := range s.data {
		for _, cell := range line {
			if cell == SLIDE {
				totalLoad += len(s.data)-row
			}
		}
	}

	return fmt.Sprintf("%d", totalLoad)
}


func (s Day14) part02() string {

	cycles := 1000000000
	directions := [][2]int{
		NORTH,
		WEST,
		SOUTH,
		EAST,
	}
	print(&s.data)

	for steps := 0; steps < cycles * len(directions); steps++ {
		tilt(&s.data, directions[steps%len(directions)])
		// print(&s.data)
	}

	totalLoad := 0

	for row, line := range s.data {
		for _, cell := range line {
			if cell == SLIDE {
				totalLoad += len(s.data)-row
			}
		}
	}

	return fmt.Sprintf("%d", totalLoad)
}


func tilt(grid *[][]string, kernel [2]int) {
	// Find direction using direction kernel
	dy, dx := kernel[0], kernel[1]

	// fmt.Printf("> dy: %d, dx: %d\n", dy, dx)
	r0, c0 := 0, 0
	rn, cn := len(*grid), len((*grid)[0])
	delta := 1

	// We should reverse iteration order when tilting to the South or East
	if dy > 0 || dx > 0 {
		r0, c0 = rn, cn
		rn, cn = 0, 0
		delta = -1
	}

	for row := r0; row != rn + delta; row += delta {
		for col := c0; col != cn + delta; col += delta {
			if row < 0 || len(*grid) <= row { continue }
			if col < 0 || len((*grid)[0]) <= col { continue }
			if (*grid)[row][col] != SLIDE { continue }

			source_x, source_y := col, row
			target_x, target_y := col, row

			// Once we know in which direction to move, let's move forward until no more moves are left
			for {
				target_x += dx
				target_y += dy

				if target_y < 0 || len(*grid) <= target_y { break }
				if target_x < 0 || len((*grid)[0]) <= target_x { break }
				if (*grid)[target_y][target_x] != SPACE { break }

				// fmt.Printf("> Swapping (%d, %d) with (%d, %d)\n", source_y, source_x, target_y, target_x)
				
				// Swap source/target items
				temp := (*grid)[target_y][target_x]
				(*grid)[target_y][target_x] = (*grid)[source_y][source_x]
				(*grid)[source_y][source_x] = temp

				// Move source pointer
				source_x += dx
				source_y += dy
			}
		}
	}
}


func print(grid *[][]string) {
	fmt.Printf("-----\n")
	for row, line := range *grid {
		fmt.Printf("[%d]: %+v\n", row, line)
	}
}

