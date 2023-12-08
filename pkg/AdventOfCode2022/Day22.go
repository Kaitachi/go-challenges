package AdventOfCode2022

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/kaitachi/go-challenges/internal/lib"
)

type Day22 struct {
	board			[]map[int]*cell
	instructions	[]instruction
	spawn			*cell
}


const (
	OPEN = "."
	WALL = "#"
)


type cell struct {
	i	string
	row	int
	col	int
	N	*cell
	E	*cell
	S	*cell
	W	*cell
}


type Direction string
const (
	L Direction = "L"
	R Direction = "R"
)


type instruction struct {
	steps	int
	rot		string
}


// 1. Assemble - How should we transform the data from our input files?
func (s *Day22) Assemble(tc *lib.TestCase) {

	s.board = make([]map[int]*cell, 0)
	s.instructions = make([]instruction, 0)
	s.spawn = nil

	re_item := regexp.MustCompile(`[^\s']`)
	re_open := regexp.MustCompile(`\` + string(OPEN))
	re_rock := regexp.MustCompile(string(WALL))
	re_instructions := regexp.MustCompile(`(\d+)(R|L)?`)

	map_line_count := strings.Count(tc.Input, "\n")-1

	north_pole := make(map[int]*cell, 0) // Collection of topmost items
	south_pole := make(map[int]*cell, 0) // Collection of bottommost items

	for row, line := range strings.Split(tc.Input, "\n")[:map_line_count] {
		// fmt.Printf("Reading line %d...\n", row)
		// fmt.Println(line)

		s.board = append(s.board, make(map[int]*cell, 0))

		item_idx := re_item.FindAllStringIndex(line, -1)

		var first_west *cell // used for wrap E/W <->

		// Add all open spaces in this line
		for _, item := range re_open.FindAllStringIndex(line, -1) {
			col := item[0]
			// fmt.Printf("> item: %v\n", item)
			// fmt.Printf("OPEN at (%d, %d)\n", row, item[0])

			// So far, we only know about what's directly West of us
			current := &cell{
				i: OPEN,
				row: row,
				col: col,
			}

			// Assign ourselves to the board
			s.board[row][col] = current

			// Can we say we're the first/last cell in the array?
			switch item[0] {
			case item_idx[0][0]: // We're the first item in this row
				first_west = current

				// We might as well be the spawn point
				if s.spawn == nil {
					s.spawn = current
				}

			case item_idx[len(item_idx)-1][0]: // We're the last item in this row
				if first_west != nil {
					current.E = first_west
					first_west.W = current
				}
			}

			// [N] Is there anything North of us?
			if row != 0 {
				if N, ok := s.board[row-1][col]; ok {
					N.S = current
					current.N = N
				}
			}

			// [S] There should be nothing to the South just yet!

			// [E/W] Can we backlink whatever is West of us?
			if W, ok := s.board[row][col-1]; ok {
				W.E = current
				current.W = W
			}

			// [N] Are we near the North Pole?
			if _, ok := north_pole[col]; !ok {
				north_pole[col] = current
			}

			// [S] Are we near the South Pole?
			south_pole[col] = current
		}

		// fmt.Printf(">> PRE-ROCKS [N]: %v\n", north_pole)
		// fmt.Printf(">> PRE-ROCKS [S]: %v\n", south_pole)

		// Add gaps to the North/South Pole (nil entries imply rock location)
		for _, rock := range re_rock.FindAllStringIndex(line, -1) {
			if _, ok := north_pole[rock[0]]; !ok {
				north_pole[rock[0]] = nil
			}

			if _, ok := south_pole[rock[0]]; !ok {
				south_pole[rock[0]] = nil
			}
		}

		// fmt.Printf(">> POST-ROCKS [N]: %v\n", north_pole)
		// fmt.Printf(">> POST-ROCKS [S]: %v\n", south_pole)
		//
		// fmt.Printf("first_west: %v\n", first_west)
	}


	// Bind North/South poles
	for col, current := range north_pole {
		if south, ok := south_pole[col]; ok && current != nil {
			current.N = south_pole[col]
			south.S = current
		}
	}

	for _, match := range re_instructions.FindAllStringSubmatch(tc.Input, -1) {
		steps, _ := strconv.Atoi(match[1])
		s.instructions = append(s.instructions, instruction{
			steps: steps,
			rot: match[2],
		})
	}
}


// 2. Activate - Take our transformed input data and make the core logic needed to resolve this Problem
func (s *Day22) Activate(tc *lib.TestCase) {
	// Assign final value to TestCase.Actual field
	switch tc.Algorithm {
	case "part01":
		tc.Actual = s.part01()
		
	case "part02":
		tc.Actual = s.part02()
	}
}


func (s Day22) part01() string {

	for i := 0; i < len(s.board); i++ {
		fmt.Printf("[%d] %+v\n", i, s.board[i])
		for k, v := range s.board[i] {
			fmt.Printf(">> [%d](%p): %#v\n", k, v, v)
		}
		fmt.Println()
	}

	s.validateGrid()
	// return ""

	direction := "E"
	cursor := s.spawn
	steps := make([]*cell, 0)

	for _, instruction := range s.instructions {

		// Where are we now?
		//fmt.Printf(">>>>>>> NOW: (%d, %d)\n", cursor.row+1, cursor.col+1)
		//fmt.Println(instruction, direction)

		// Walk the steps we're meant to Walk
		for i := instruction.steps; i > 0; i-- {
			switch direction {
			case "N":
				if cursor.N != nil {
					cursor = cursor.N
				}

			case "E":
				if cursor.E != nil {
					cursor = cursor.E
				}

			case "S":
				if cursor.S != nil {
					cursor = cursor.S
				}

			case "W":
				if cursor.W != nil {
					cursor = cursor.W
				}
			}
			
			steps = append(steps, cursor)
			//fmt.Printf(">> STEP (%d, %d)\n", cursor.row+1, cursor.col+1)
		}

		// Rotate once we're done moving
		switch instruction.rot {
		case "L":
			switch direction {
			case "N":	direction = "W"
			case "E":	direction = "N"
			case "S":	direction = "E"
			case "W":	direction = "S"
			}

		case "R":
			switch direction {
			case "N":	direction = "E"
			case "E":	direction = "S"
			case "S":	direction = "W"
			case "W":	direction = "N"
			}
		}
	}

	fmt.Printf("Final position: (%d, %d), facing %s\n", cursor.row+1, cursor.col+1, direction)
	spin := -1

	switch direction {
	case "N":	spin = 3
	case "E":	spin = 0
	case "S":	spin = 1
	case "W":	spin = 2
	}

	// Formula: 1000 * row + 4 * col + facing
	password := 1000 * (cursor.row+1) + 4 * (cursor.col+1) + spin

	// fmt.Printf("PASSWORD: %d\n", password)

	return fmt.Sprintf("%d", password)
}


func (s Day22) part02() string {

	return fmt.Sprintf("%d", -1)
}


func (s Day22) validateGrid() {

	rows := 0
	cols := 0
	north_poles := make(map[int]*cell, cols) // [col]*cell
	south_poles := make(map[int]*cell, cols) // [col]*cell
	east_poles := make(map[int]*cell, rows) // [row]*cell
	west_poles := make(map[int]*cell, rows) // [row]*cell

	for row, line := range s.board {
		for col, cell := range line {
			// Gather board size
			cols = int(math.Max(float64(cols), float64(col)))

			// Gather board poles

			// Gather North Poles
			if temp, ok := north_poles[col]; ok {
				if cell.row < temp.row {
					north_poles[col] = cell
				}
			} else {
				north_poles[col] = cell
			}

			// Gather South Poles
			if temp, ok := south_poles[col]; ok {
				if temp.row < cell.row {
					south_poles[col] = cell
				}
			} else {
				south_poles[col] = cell
			}

			// Gather East Poles
			if temp, ok := east_poles[row]; ok {
				if temp.col < cell.col {
					east_poles[row] = cell
				}
			} else {
				east_poles[row] = cell
			}

			// Gather West Poles
			if temp, ok := west_poles[row]; ok {
				if cell.col < temp.col {
					west_poles[row] = cell
				}
			} else {
				west_poles[row] = cell
			}
		}
		rows++
	}

	fmt.Printf("%v\n", s.board)
	fmt.Println("--- EAST POLES ---")
	fmt.Printf("%v\n", east_poles)
	fmt.Println("--- WEST POLES ---")
	fmt.Printf("%v\n", west_poles)

	fmt.Println("Board size ", rows, cols)

	// Traverse entire board
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if cell, ok := s.board[row][col]; ok {
				fmt.Printf("Found [%p]: %+v\n", cell, cell)

				// Cell must match current index!
				if cell.row != row || cell.col != col {
					panic(fmt.Sprintf("Cell at %d %d is misplaced! Found %p instead.", row, col, cell))
				}

				// Let's traverse North, what will we find?
				north := cell.N
				expected_south_pole := south_poles[col]
				if north != nil {
					if (north.row != cell.row-1 && north.row != expected_south_pole.row) || north.col != cell.col {
						panic(fmt.Sprintf("North cell is mismatching! Should be (%d, %d) or pole (%d, %d), found (%d, %d)", cell.row-1, cell.col, expected_south_pole.row, expected_south_pole.col, north.row, north.col))
					}
				}

				// Let's traverse South, what will we find?
				south := cell.S
				expected_north_pole := north_poles[col]
				if south != nil {
					if (south.row != cell.row+1 && south.row != expected_north_pole.row) || south.col != cell.col {
						panic(fmt.Sprintf("South cell is mismatching! Should be (%d, %d) or pole (%d, %d), found (%d, %d)", cell.row+1, cell.col, expected_north_pole.row, expected_north_pole.col, south.row, south.col))
					}
				}

				// Let's traverse East, what will we find?
				east := cell.E
				expected_west_pole := west_poles[row]
				if east != nil {
					if east.row != cell.row || (east.col != cell.col+1 && east.col != expected_west_pole.col) {
						panic(fmt.Sprintf("East cell is mismatching! Should be (%d, %d) or pole (%d, %d), found (%d, %d)", cell.row, cell.col+1, expected_west_pole.row, expected_west_pole.col, east.row, east.col))
					}
				}

				// Let's traverse West, what will we find?
				west := cell.W
				expected_east_pole := east_poles[row]
				if west != nil {
					if west.row != cell.row || (west.col != cell.col-1 && west.col != expected_east_pole.col) {
						panic(fmt.Sprintf("West cell is mismatching! Should be (%d, %d) or pole (%d, %d), found (%d, %d)", cell.row, cell.col-1, expected_east_pole.row, expected_east_pole.col, west.row, west.col))
					}
				}
			}
		}
	}
}

