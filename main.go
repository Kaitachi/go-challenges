package main

import (
	"fmt"

	"github.com/kaitachi/go-challenges/internal/lib"
	AOC2022 "github.com/kaitachi/go-challenges/pkg/AdventOfCode2022"
)

func main() {
	challengeName := "AOC2022"
	problemName := "Day01"
	algorithmName := "part01"
	scenarios := []string{"example"}

	problem := AOC2022.SetUp_AOC2022_Day01(scenarios, algorithmName)

	solution := lib.Solve(&problem)

	fmt.Printf("> Solution for %s, problem %s (%s): %s\n", challengeName, problemName, algorithmName, solution)
}

