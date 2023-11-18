package AOC2022

import "github.com/kaitachi/go-challenges/internal/lib"


type AdventOfCode2022 struct {}


func (c *AdventOfCode2022) GetSolution(name string) (*lib.Solver, bool) {
	var solutions = map[string]lib.Solver{
		"Day01": &Day01{},
	}

	value, ok := solutions[name]
	return &value, ok
}

