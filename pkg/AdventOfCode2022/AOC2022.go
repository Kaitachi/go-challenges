package AOC2022

import "github.com/kaitachi/go-challenges/internal/lib"

type AOC2022 interface {
	part01(*lib.TestCase[any])
	part02(*lib.TestCase[any])
}


func GetAOC2022Challenge(solution string, ds []string, algo string) lib.Challenge {
	challenge := lib.GetChallenge("AdventOfCode2022", solution, ds, algo)

	return challenge
}

