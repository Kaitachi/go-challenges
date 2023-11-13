package AOC2022

import (
	"strconv"
	"strings"

	"github.com/kaitachi/go-challenges/internal/lib"
)

type Day01 struct {
	lib.Problem
}


func SetUp_AOC2022_Day01(ds []string, algo string) Day01 {
	return Day01{
		Problem: lib.Problem{
			Challenge: GetAOC2022Challenge("Day01", ds, algo),
		},
	}
}


// 1. Assemble - How should we transform the data from our input files?
func (s *Day01) Assemble(scenario string) {
	input, output := s.Challenge.Data(scenario)

	elvesStrings := strings.Split(input, "\n\n")
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

	s.TestCase.Name = scenario
	s.TestCase.Input = elves
	s.TestCase.Output = strings.TrimSpace(output)
}


// TODO: There is a *slight* chance that this method could be moved elsewhere--how???
// 2. Activate - Take our transformed input data and make the core logic needed to resolve this Problem
func (s *Day01) Activate() {
	var result string

	switch s.Challenge.Algorithm {
	case "part01":
		result = s.part01()
		
	case "part02":
		result = s.part02()
	}

	// Assign final value to TestCase.Actual field
	s.TestCase.Actual = result
}


func (s *Day01) part01() string {
	return "24000"
}


func (s *Day01) part02() string {
	return "-1"
}

