package AOC2022

import (
	"embed"
	"strconv"
	"strings"

	"github.com/kaitachi/go-challenges/internal/lib"
)

type AOC2022_Day01 struct {
	Challenge lib.Challenge
	TestCase lib.TestCase[[][]int]
}


func SetUp_AOC2022_Day01(fs *embed.FS, ds []string, algo string) AOC2022_Day01 {
	return AOC2022_Day01{
		Challenge: GetAOC2022Challenge(fs, "Day01", ds, algo),
	}
}


func (s *AOC2022_Day01) Assemble(scenario string) {
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


func (s *AOC2022_Day01) Activate() {

	// Assign final value to TestCase.Actual field
	s.TestCase.Actual = "24000"
}


func (s *AOC2022_Day01) Assert() {

	s.TestCase.Verify()
}



// TODO: Is there any way to move these metods elsewhere???
func (s *AOC2022_Day01) Scenarios() []string {
	return s.Challenge.DataSet
}


func (s *AOC2022_Day01) Solution() string {
	return s.TestCase.Actual
}

