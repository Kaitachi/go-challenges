package AdventOfCode2023

import (
	"fmt"
	"strings"

	"github.com/kaitachi/go-challenges/internal/lib"
)

type Day15 struct {
	data []string
}


// 1. Assemble - How should we transform the data from our input files?
func (s *Day15) Assemble(tc *lib.TestCase) {

	s.data = make([]string, 0)

	for _, line := range strings.Split(tc.Input, ",") {
		s.data = append(s.data, line)
	}
}


// 2. Activate - Take our transformed input data and make the core logic needed to resolve this Problem
func (s *Day15) Activate(tc *lib.TestCase) {
	// Assign final value to TestCase.Actual field
	switch tc.Algorithm {
	case "part01":
		tc.Actual = s.part01()
		
	case "part02":
		tc.Actual = s.part02()
	}
}


func (s Day15) part01() string {

	mult := 17
	mod := 256
	sum := 0

	for _, line := range s.data {
		sum += hash(line, mult, mod)
	}

	return fmt.Sprintf("%d", sum)
}


func (s Day15) part02() string {

	return fmt.Sprintf("%d", -1)
}


func hash(s string, mult int, mod int) int {
	hash := 0
	for _, c := range s {
		value := (hash+int(c)) * mult
		// fmt.Printf("> char %c = %d, value %d\n", c, int(c), value)
		hash = (value) % mod
		// fmt.Printf(">> hash %d\n", hash)
	}

	return hash
}

