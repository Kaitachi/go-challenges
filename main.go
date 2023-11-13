package main

import (
	"fmt"

	"github.com/kaitachi/go-challenges/internal/lib"
	AOC2022 "github.com/kaitachi/go-challenges/pkg/AdventOfCode2022"
)

func main() {
	problem := AOC2022.Day01{}
	algorithmName := "part01"
	scenarios := []string{"example"}

	solution := lib.Solve(problem, scenarios, algorithmName)

	fmt.Printf("> Solution for %s, problem %s (%s): %s\n", lib.ChallengeOf(problem), lib.NameOf(problem), algorithmName, solution)
}

