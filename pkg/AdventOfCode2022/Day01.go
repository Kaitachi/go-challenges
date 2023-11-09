package AOC2022

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kaitachi/go-challenges/internal/lib"
)

type AOC2022_Day01 struct {
	TestCase lib.TestCase[[][]int]

	// TODO: Could I somehow declare [][]int type at this level to be used everywhere?
}


// TODO: How can I abstract this method?
func (s AOC2022_Day01) Run(input string, output string) bool {
	testCase := lib.TestCase[[][]int]{}
	s.Assemble(&testCase, input, output)
	s.Activate(&testCase)
	return s.Assert()
}


func (s AOC2022_Day01) Assemble(tc *lib.TestCase[[][]int], input string, output string) {
	fmt.Println("AOC2022_Day01.assemble")

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

	tc.Input = elves
	tc.Output = output
}


func (s AOC2022_Day01) Activate(tc *lib.TestCase[[][]int]) {
	fmt.Println("AOC2022_Day01.activate")
}


func (s AOC2022_Day01) Assert() bool {
	fmt.Println("AOC2022_Day01.assert")
	return true
}

