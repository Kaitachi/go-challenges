package AdventOfCode2023

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/kaitachi/go-challenges/internal/lib"
)

type Day11 struct {
	data		[]map[int]*galaxy
	galaxies	[]*galaxy
	rows		int
	cols		int
}


type galaxy struct {
	id	int
	row	int
	col	int
}


// 1. Assemble - How should we transform the data from our input files?
func (s *Day11) Assemble(tc *lib.TestCase) {

	s.rows = 0
	s.cols = 0
	s.data = make([]map[int]*galaxy, 0)
	re_galaxy := regexp.MustCompile(`#`)
	count := 0

	s.galaxies = make([]*galaxy, 0)
	galaxyCols := make(map[int][]*galaxy, 0)

	for row, line := range strings.Split(tc.Input, "\n") {
		s.data = append(s.data, make(map[int]*galaxy, 0))
		s.rows++

		for _, match := range re_galaxy.FindAllStringIndex(line, -1) {
			col := match[0]
			s.cols = int(math.Max(float64(s.cols), float64(col)))

			g := &galaxy{
				id:		count,
				row:	row,
				col:	col,
			}

			s.galaxies = append(s.galaxies, g)

			s.data[row][col] = g

			if _, ok := galaxyCols[col]; !ok {
				galaxyCols[col] = make([]*galaxy, 0)
			}

			galaxyCols[col] = append(galaxyCols[col], g)

			count++
		}
	}

	// Expand galaxy
	dRows := 0
	dCols := 0

	fmt.Printf("%+v\n", galaxyCols)

	for row := 0; row < s.rows; row++ {
		if len(s.data[row]) == 0 {
			dRows++
		} else {
			for _, galaxy := range s.data[row] {
				galaxy.row += dRows
			}
		}
	}

	for col := 0; col <= s.cols; col++ {
		if _, ok := galaxyCols[col]; !ok {
			dCols++
		} else {
			for _, galaxy := range galaxyCols[col] {
				galaxy.col += dCols
			}
		}
	}

	s.rows += dRows
	s.cols += dCols
}


// 2. Activate - Take our transformed input data and make the core logic needed to resolve this Problem
func (s *Day11) Activate(tc *lib.TestCase) {
	// Assign final value to TestCase.Actual field
	switch tc.Algorithm {
	case "part01":
		tc.Actual = s.part01()
		
	case "part02":
		tc.Actual = s.part02()
	}
}


func (s Day11) part01() string {

	// for row := 0; row < s.rows; row++ {
	// 	fmt.Printf("[%d]: %+v\n", row, s.data[row])
	// }

	for i := 0; i < len(s.galaxies); i++ {
		fmt.Printf("[%d]: %+v\n", i, s.galaxies[i])
	}

	minDistanceSum := 0

	for i := 0; i < len(s.galaxies); i++ {
		for j := 0; j < i; j++ {
			// fmt.Printf("Comparing [%d]: %+v with [%d]: %+v\n", i, s.galaxies[i], j, s.galaxies[j], )

			minDistanceSum += s.galaxies[i].manhattan(s.galaxies[j])
		}
	}

	return fmt.Sprintf("%d", minDistanceSum)
}


func (s Day11) part02() string {

	return fmt.Sprintf("%d", -1)
}


func (g0 *galaxy) manhattan(g1 *galaxy) int {
	dRow := g1.row - g0.row
	dCol := g1.col - g0.col
	return int(math.Abs(float64(dRow)) + math.Abs(float64(dCol)))
}

