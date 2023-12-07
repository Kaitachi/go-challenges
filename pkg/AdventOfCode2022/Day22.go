package AdventOfCode2022

import (
	"fmt"
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

	map_line_count := strings.Count(tc.Input, "\n")-2

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

	// for i := 0; i < len(s.board); i++ {
	// 	fmt.Printf("[%d] %+v\n", i, s.board[i])
	// 	for k, v := range s.board[i] {
	// 		fmt.Printf(">> [%d](%p): %#v\n", k, v, v)
	// 	}
	// 	fmt.Println()
	// }

	direction := "E"
	cursor := s.spawn
	steps := make([]*cell, 0)

	for _, instruction := range s.instructions {

		// Where are we now?
		fmt.Printf(">>>>>>> NOW: (%d, %d)\n", cursor.row+1, cursor.col+1)
		fmt.Println(instruction, direction)

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
			fmt.Printf(">> STEP (%d, %d)\n", cursor.row+1, cursor.col+1)
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

	fmt.Printf("Final position: (%d, %d)\n", cursor.row+1, cursor.col+1)
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

