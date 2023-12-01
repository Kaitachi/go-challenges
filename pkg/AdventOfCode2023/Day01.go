package AdventOfCode2023

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kaitachi/go-challenges/internal/lib"
)

type Day01 struct {
	data [][]int
}


// 1. Assemble - How should we transform the data from our input files?
func (s *Day01) Assemble(tc *lib.TestCase) {

	elvesStrings := strings.Split(tc.Input, "\n\n")
	elves := [][]int{}

	for _, elf := range elvesStrings {
		items := strings.Split(elf, "\n")
		collection := []int{}

		for _, item := range items {
			item = strings.TrimSpace(item)
			if item == "" { continue }

			i, err := strconv.Atoi(item)
			if err != nil {
				panic(err)
			}

			collection = append(collection, i)
		}

		elves = append(elves, collection)
	}

	s.data = elves
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

	return fmt.Sprintf("%d", -1)
}


func (s Day01) part02() string {

	return fmt.Sprintf("%d", -1)
}

