package main

import (
	"embed"
	"fmt"

	"github.com/kaitachi/go-challenges/internal/lib"
	AOC2022 "github.com/kaitachi/go-challenges/pkg/AdventOfCode2022"
)

//go:embed assets/*
var assets embed.FS

func main() {
	scenarios := []string{"example"}

	problem := AOC2022.SetUp_AOC2022_Day01(&assets, scenarios, "part01")

	solution := lib.Solve(&problem)

	fmt.Printf("> Solution for this problem: %s\n", solution)
}

