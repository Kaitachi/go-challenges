package AdventOfCode2023

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/kaitachi/go-challenges/internal/lib"
)

type Day01 struct {
	data []string
}


// 1. Assemble - How should we transform the data from our input files?
func (s *Day01) Assemble(tc *lib.TestCase) {

	lines := strings.Split(tc.Input, "\n")
	
	s.data = lines[:len(lines)-1]
}


// 2. Activate - Take our transformed input data and make the core logic needed to resolve this Problem
func (s *Day01) Activate(tc *lib.TestCase) {

	var result string

	switch tc.Algorithm {
	case "part01":
		result = s.part01()
		
	case "part02":
		result = s.part02()
	}

	// Assign final value to TestCase.Actual field
	tc.Actual = result
}


func (s Day01) part01() string {

	// Use regex to find digit in first position
	first := regexp.MustCompile("^(?:[a-zA-Z]*)(\\d)")

	// Use regex to find digit in last position
	last := regexp.MustCompile("(\\d)(?:[a-zA-Z]*)$")

	sum := findSubstrings(s.data, first, last)

	return fmt.Sprintf("%d", sum)
}


func (s Day01) part02() string {

	digits := map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
		"four": 4,
		"five": 5,
		"six": 6,
		"seven": 7,
		"eight": 8,
		"nine": 9,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
	}

	words := ""

	for key, _ := range digits {
		words += "|" + key
	}

	rewords := regexp.MustCompile(words[1:])

	// Grab entries from original array
	var sum int = 0

	for _, line := range s.data {
		fmt.Println("---", line)

		idx := rewords.FindAllStringSubmatchIndex(line, -1)

		fmt.Println(idx)

		for _, match := range idx {
			fmt.Println(line[match[0]:match[1]])
			fmt.Println("------")
		}

		d0 := digits[line[idx[0][0]:idx[0][1]]]
		d1 := digits[line[idx[len(idx)-1][0]:idx[len(idx)-1][1]]]

		// Sum array
		number := d0 * 10 + d1

		fmt.Println(">> ", number)
		sum += number
	}

	return fmt.Sprintf("%d", sum)
}


func findSubstrings(array []string, re0 *regexp.Regexp, re1 *regexp.Regexp) int {
	// Grab entries from original array
	var sum int = 0

	for _, line := range array {
		d0 := re0.FindStringSubmatch(line)
		d1 := re1.FindStringSubmatch(line)

		fmt.Println(d0)

		number, _ := strconv.Atoi(d0[1] + d1[1])

		// Sum array
		sum += number
	}

	return sum
}

