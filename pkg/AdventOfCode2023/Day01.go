package AdventOfCode2023

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/kaitachi/go-challenges/internal/lib"
)

type Day01 struct {
	data []string
}


// 1. Assemble - How should we transform the data from our input files?
func (s *Day01) Assemble(tc *lib.TestCase) {

	lines := strings.Split(tc.Input, "\n")
	
	s.data = lines[:len(lines)-1]
}


// 2. Activate - Take our transformed input data and make the core logic needed to resolve this Problem
func (s *Day01) Activate(tc *lib.TestCase) {

	var result string

	switch tc.Algorithm {
	case "part01":
		result = s.part01()
		
	case "part02":
		result = s.part02()
	}

	// Assign final value to TestCase.Actual field
	tc.Actual = result
}


func (s Day01) part01() string {

	// Use regex to find digit in first position
	first := regexp.MustCompile("^(?:[a-zA-Z]*)(\\d)")

	// Use regex to find digit in last position
	last := regexp.MustCompile("(\\d)(?:[a-zA-Z]*)$")

	// Grab entries from original array
	var sum int = 0

	for _, line := range s.data {
		d0 := first.FindStringSubmatch(line)
		d1 := last.FindStringSubmatch(line)

		number, _ := strconv.Atoi(d0[1] + d1[1])

		// Sum array
		sum += number
	}

	return fmt.Sprintf("%d", sum)
}


func (s Day01) part02() string {

	return fmt.Sprintf("%d", -1)
}

