package main

import (
	"fmt"
	"os"

	"github.com/kaitachi/go-challenges/internal/lib"
	AOC2022 "github.com/kaitachi/go-challenges/pkg/AdventOfCode2022"
)


var challenges = map[string]lib.Challenger{
	"AdventOfCode2022": &AOC2022.AdventOfCode2022{},
}


func main() {
	args := os.Args[1:]

	if len(args) < 4 {
		panic("Usage: CHALLENGE PROBLEM ALGORITHM SCENARIO [...SCENARIO]")
	}
	fmt.Println(args)

	challengeName := args[0]
	problemName := args[1]
	algorithmName := args[2]
	scenarios := args[3:]

	challenge, ok := challenges[challengeName]
	if !ok {
		panic(fmt.Sprintf("Challenge %s not found!", challengeName))
	}

	problem, ok := challenge.GetProblem(problemName)
	if !ok {
		panic(fmt.Sprintf("Invalid Problem name given for %s: %s", challengeName, problemName))
	}

	solution := lib.Solve(problem, scenarios, algorithmName)

	fmt.Printf("> Solution for %s, problem %s (%s): %s\n", lib.ChallengeOf(problem), lib.NameOf(problem), algorithmName, solution)
}

