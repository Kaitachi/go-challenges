package main

import (
	"embed"

	"github.com/kaitachi/go-challenges/internal/lib"
	AOC2022 "github.com/kaitachi/go-challenges/pkg/AdventOfCode2022"
)

//go:embed assets/*
var assets embed.FS

func main() {
	var soln lib.Solution
	soln = AOC2022.AOC2022_Day01{}

	thing := lib.Challenge{
		Assets:		&assets,
		Challenge:	"AdventOfCode2022",
		Solution:	"Day01",
		DataSet:	"example",
		Algorithm:	"part01",
		Soln:		soln,
	}

	thing.Execute()
}

