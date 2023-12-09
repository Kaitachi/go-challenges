package AdventOfCode2023

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/kaitachi/go-challenges/internal/lib"
)

type Day09 struct {
	data [][]int
}


// 1. Assemble - How should we transform the data from our input files?
func (s *Day09) Assemble(tc *lib.TestCase) {

	s.data = make([][]int, 0)

	re_numbers := regexp.MustCompile(`-?\d+`)

	for _, line := range strings.Split(tc.Input, "\n") {
		
		matches := re_numbers.FindAllString(line, -1)
		numbers := make([]int, len(matches))

		for i, match := range matches {
			number, _ := strconv.Atoi(match)
			numbers[i] = number
		}

		s.data = append(s.data, numbers)
	}
}


// 2. Activate - Take our transformed input data and make the core logic needed to resolve this Problem
func (s *Day09) Activate(tc *lib.TestCase) {
	// Assign final value to TestCase.Actual field
	switch tc.Algorithm {
	case "part01":
		tc.Actual = s.part01()
		
	case "part02":
		tc.Actual = s.part02()
	}
}


func (s Day09) part01() string {

	sum := 0

	for _, seq := range s.data {
		sum += nextNumberByDerivative(seq)
	}

	return fmt.Sprintf("%d", sum)
}


func (s Day09) part02() string {

	return fmt.Sprintf("%d", -1)
}


func nextNumberByDerivative(seq []int) int {

	d := make([][]int, 1)
	d[0] = seq
	m := 0

	// Calculate nth derivative until slope = 0
	for dx := 1; dx == 1 || m != 0; dx++ {
		d = append(d, make([]int, len(d[dx-1])-1))
		m = 0

		for i := 0; i < len(d[dx-1])-1; i++ {
			x0, x1 := d[dx-1][i], d[dx-1][i+1]
			slope := x1 - x0

			m = int(math.Max(math.Abs(float64(m)), math.Abs(float64(slope))))
			d[dx][i] = slope
		}
	}

	fmt.Println("> m = 0")
	fmt.Printf("%v\n", d)

	for dx := len(d)-1; dx > 0; dx-- {
		prev, curr := d[dx-1][len(d[dx-1])-1], d[dx][len(d[dx])-1]

		d[dx-1] = append(d[dx-1], curr + prev)
	}

	fmt.Printf("> Partial = %d\n", d[0][len(d[0])-1])

	return d[0][len(d[0])-1]
}

