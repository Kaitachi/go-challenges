package AdventOfCode2022

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/kaitachi/go-challenges/internal/lib"
)

type Day14 struct {
	data map[int]map[int]piece
}


type piece string

const (
	ROCK	piece = "に"
	AIR		piece = "・"
	SPOUT	piece = "十"
	SAND	piece = "の"
	MOVE	piece = "〜"
)

const (
	SPOUT_ROW int = 0
	SPOUT_COL int = 500
)


// 1. Assemble - How should we transform the data from our input files?
func (s *Day14) Assemble(tc *lib.TestCase) {

	s.data = make(map[int]map[int]piece, 0)

	re_rocks := regexp.MustCompile(`(\d+),(\d+)`)

	// Insert rocks in our grid
	for _, line := range strings.Split(tc.Input, "\n") {

		rocks := re_rocks.FindAllStringSubmatch(line, -1)
		s.draw(rocks)
	}

	// Insert spout in our grid
	if _, ok := s.data[SPOUT_ROW]; !ok {
		s.data[SPOUT_ROW] = make(map[int]piece, 0)
	}

	s.data[SPOUT_ROW][SPOUT_COL] = SPOUT
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
	return s.simulate()
}


func (s Day14) part02() string {

	return fmt.Sprintf("%d", -1)
}


func (s Day14) simulate() string {

	max_row, min_col, max_col := s.bounds()

	s.print(max_row, min_col, max_col)
	grains := 0

	end: for {
		grains++

		row := SPOUT_ROW
		col := SPOUT_COL

		settle: for step := 0; step < s.lowestRow() + 3; step++ {
			//fmt.Printf("start for loop (%d, %d)...\n", row, col)
			// Whenever layer right below us is found,
			if line, ok := s.data[row+1]; ok {
				// we need to ensure the three spots below us are not occupied

				// Once we realize all three spots are occupied, we break out
				_, filled_l := line[col-1]
				_, filled_c := line[col+0]
				_, filled_r := line[col+1]
	
				switch {
				case !filled_c: // Move through gap in center
					//fmt.Println("Moving through gap in centre")

				case !filled_l: // Move through gap in left
					//fmt.Println("Moving through gap in left")
					col--

				case !filled_r: // Move through gap in right
					//fmt.Println("Moving through gap in right")
					col++

				default:
					//fmt.Println("Settling on current row...")
					break settle
				}
			}

			//fmt.Println("Moving to next row")
			row++
		}

		if row == s.lowestRow() + 3 {
			break end
		}

		// Only need to create rows when settling grain of sand on it
		if _, ok := s.data[row]; !ok {
			s.data[row] = make(map[int]piece, 0)
		}

		//fmt.Printf("Settling grain at (%d, %d)\n", row, col)
		s.data[row][col] = SAND
	}

	s.print(max_row, min_col, max_col)

	return fmt.Sprintf("%d", grains-1)
}


func (s Day14) lowestRow() int {
	max_row := 0

	for row, _ := range s.data {
		max_row = int(math.Max(float64(max_row), float64(row)))
	}

	return max_row
}


func (s *Day14) draw(rocks [][]string) {
	var prev_row *int = nil
	var prev_col *int = nil

	for _, segment := range rocks {
		row, _ := strconv.Atoi(segment[2])
		col, _ := strconv.Atoi(segment[1])

		if prev_row != nil && prev_col != nil {
			dy0 := int(math.Min(float64(row), float64(*prev_row)))
			dy1 := int(math.Max(float64(row), float64(*prev_row)))

			for _ = dy0; dy0 <= dy1; dy0++ {
				// Do we need to define this row?
				if _, ok := s.data[dy0]; !ok {
					s.data[dy0] = make(map[int]piece, 0)
				}

				s.data[dy0][col] = ROCK
			}

			dx0 := int(math.Min(float64(col), float64(*prev_col)))
			dx1 := int(math.Max(float64(col), float64(*prev_col)))

			for _ = dx0; dx0 <= dx1; dx0++ {
				s.data[row][dx0] = ROCK
			}
		}

		prev_row = &row
		prev_col = &col
	}
}


func (s Day14) print(max_row int, min_col int, max_col int) {

	fmt.Printf("--------------------------\n")

	// Print grid, bottom down
	for row := 0; row <= max_row; row++ {
		fmt.Printf("[%3d] ", row)

		items, ok := s.data[row]
		if !ok {
			items = make(map[int]piece, 0)
		}

		for col := min_col; col <= max_col; col++ {
			if item, ok := items[col]; ok {
				fmt.Printf("%s", item)
			} else {
				fmt.Printf("%s", AIR)
			}
			
		}

		fmt.Println()
	}

	fmt.Printf("Dimensions: %d max_row, %d min_col, %d max_col\n", max_row, min_col, max_col)
}


func (s Day14) bounds() (max_row int, min_col int, max_col int) {
	max_row = 0
	min_col = math.MaxInt
	max_col = math.MinInt

	// Grab max and min values for our grid
	for row, items := range s.data {
		max_row = int(math.Max(float64(max_row), float64(row)))

		for col, _ := range items {
			min_col = int(math.Min(float64(min_col), float64(col)))
			max_col = int(math.Max(float64(max_col), float64(col)))
		}
	}

	return max_row, min_col, max_col
}

