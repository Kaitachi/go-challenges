package AOC2022

import"github.com/kaitachi/go-challenges/internal/lib"


type AdventOfCode2022 struct {}


var problems = map[string]lib.Solver{
	"Day01": &Day01{},
}


func (c *AdventOfCode2022) GetProblem(name string) (*lib.Solver, bool) {
	value, ok := problems[name]
	return &value, ok
}

