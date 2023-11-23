package AOC2022

import "github.com/kaitachi/go-challenges/internal/lib"


var solutions = map[string]lib.Solver{
	"Day01": &Day01{},
}


func GetChallenge() *lib.Challenge {
	return lib.NewChallenge("AdventOfCode2022", solutions)
}

