package AOC2022

import (
	"fmt"

	"github.com/kaitachi/go-challenges/internal/lib"
)

type AOC2022_Day01 struct {
	TestCase lib.TestCase
}


func (s AOC2022_Day01) Run() bool {
	s.Assemble(lib.TestCase{})
	s.Activate()
	return s.Assert()
}


func (s AOC2022_Day01) Assemble(tc lib.TestCase) {
	fmt.Println("AOC2022_Day01.assemble")
}


func (s AOC2022_Day01) Activate() {
	fmt.Println("AOC2022_Day01.activate")
}


func (s AOC2022_Day01) Assert() bool {
	fmt.Println("AOC2022_Day01.assert")
	return true
}

