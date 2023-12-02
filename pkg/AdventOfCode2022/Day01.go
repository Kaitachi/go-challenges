package AdventOfCode2022

import (
	"fmt"
	"sort"
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
	// Assign final value to TestCase.Actual field
	switch tc.Algorithm {
	case "part01":
		tc.Actual = s.part01()
		
	case "part02":
		tc.Actual = s.part02()
	}
}


func (s Day01) part01() string {
	
	elves := getElfWeights(s.data)
	sort.Sort(sort.Reverse(sort.IntSlice(elves)))

	return fmt.Sprintf("%d", elves[0])
}


func (s Day01) part02() string {

	// Select Top N elves with the most Calories
	topN := 3

	elves := getElfWeights(s.data)
	sort.Sort(sort.Reverse(sort.IntSlice(elves)))

	sum := 0
	for i := 0; i < topN; i++ {
		sum += elves[i]
	}

	return fmt.Sprintf("%d", sum)
}


func getElfWeights(elves [][]int) []int {
	weights := make([]int, len(elves))

	// Add up all items per elf
	for i, elf := range elves {

		for _, item := range elf {
			weights[i] += item
		}
	}

	return weights
}

