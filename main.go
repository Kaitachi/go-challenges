package main

import (
	"embed"
	"fmt"

	"github.com/kaitachi/go-challenges/internal/lib"
)

//go:embed assets/*
var assets embed.FS

func main() {
	thing := lib.Challenge{
		Assets:		&assets,
		AssetPath:	"assets/AdventOfCode2022",
		Solution:	"Day01",
	}

	files, err := thing.GetFiles()
	if err != nil {
		panic(err)
	}

	fmt.Println("FILES FOR THIS SOLUTION:")
	for _, file := range files {
		fmt.Println(file)
	}
	fmt.Println("^^^")

	test, err := assets.ReadFile(files[0])
	if err != nil {
		panic(err)
	}

	fmt.Println(string(test))
}

