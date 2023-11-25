package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/kaitachi/go-challenges/internal/lib"
	AOC2022 "github.com/kaitachi/go-challenges/pkg/AdventOfCode2022"
)


var challenges = map[string]map[string]lib.Solver{
	"AdventOfCode2022": *&AOC2022.Solutions,
}


func main() {
	args := os.Args[1:]

	challenge, action := retrieveChallenge(args)

	switch action {
	case "create:solution": // Creates new Solver with given Solution name
		createSolution(challenge, args[1:])
		break

	case "solve": // Run Solver for given parameters
		solve(challenge, args[1:])
		break
	}

	os.Exit(0)
}


func retrieveChallenge(args []string) (*lib.Challenge, string) {

	// Validate user's choice
	switch strings.ToLower(args[0]) {
	case "create:solution":
		if len(args) < 3 {
			panic("Usage: create CHALLENGE SOLUTION")
		}

		break

	case "solve":
		if len(args) < 5 {
			panic("Usage: solve CHALLENGE SOLUTION ALGORITHM SCENARIO [...SCENARIO]")
		}

		break

	default:
		panic("Try adding `create:solution` or `solve` arguments")
	}

	// Create Challenge
	challengeName := args[1]
	solutionName := args[2]
	challenge := lib.NewChallenge(challengeName, solutionName)

	return challenge, strings.ToLower(args[0])
}


func createSolution(c *lib.Challenge, args []string) {
	lib.CreateSolution(c, args[1])
}


func solve(c *lib.Challenge, args []string) {

	// Retrieve Solution
	solver, ok := challenges[c.Challenge][c.Solution]
	if !ok {
		panic(fmt.Sprintf("Solution %s not found!", c.Solution))
	}

	c.Algorithm = args[2]
	c.Scenarios = args[3:]

	solution := c.Solve(solver)

	fmt.Printf("> Solution for %s, problem %s (%s): %s\n", c.Challenge, c.Solution, c.Algorithm, solution)
}

