package main

import (
	"embed"

	"github.com/kaitachi/go-challenges/internal/lib"
	AOC2022 "github.com/kaitachi/go-challenges/pkg/AdventOfCode2022"
)

//go:embed assets/*
var assets embed.FS

func main() {
	problem := AOC2022.SetUp_AOC2022_Day01(&assets, "example", "part01")


	lib.Solve(&problem)
}

