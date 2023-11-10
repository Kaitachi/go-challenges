package AOC2022

import (
	"embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/kaitachi/go-challenges/internal/lib"
)

type AOC2022_Day01 struct {
	Challenge lib.Challenge
	TestCase lib.TestCase[[][]int]
}


func SetUp_AOC2022_Day01(fs *embed.FS, ds string, algo string) AOC2022_Day01 {
	return AOC2022_Day01{
		Challenge: GetAOC2022Challenge(fs, "Day01", ds, algo),
		TestCase: lib.TestCase[[][]int]{},
	}
}


func (s *AOC2022_Day01) Assemble() {
	fmt.Println("AOC2022_Day01.assemble")

	input, output := s.Challenge.GetScenarioData()

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

	s.TestCase.Input = elves
	s.TestCase.Output = strings.TrimSpace(output)

	fmt.Println("> output: " + s.TestCase.Output)
}


func (s *AOC2022_Day01) Activate() {
	fmt.Println("AOC2022_Day01.activate")
	fmt.Println("> output: " + s.TestCase.Output)
}


func (s *AOC2022_Day01) Assert() bool {
	fmt.Println("AOC2022_Day01.assert")
	fmt.Println("> output: " + s.TestCase.Output)

	return true
}

