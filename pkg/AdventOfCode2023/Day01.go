package AdventOfCode2023

import (
	"fmt"
	"regexp"
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

	digits := map[string]int{
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

	return fmt.Sprintf("%d", calculateSum(s.data, digits))
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

	return fmt.Sprintf("%d", calculateSum(s.data, digits))
}


func concatenateKeys(m map[string]int, s string) string {
	words := ""

	for key, _ := range m {
		words += "|" + key
	}

	return words[len(s):]
}


func calculateSum(array []string, digits map[string]int) int {
	re := regexp.MustCompile(fmt.Sprintf("%s", concatenateKeys(digits, "|")))

	// Grab entries from original array
	var sum int = 0

	for _, line := range array {

		d0 := digits[search(line, re)]
		d1 := digits[reverseSearch(line, re)]

		// Sum array
		number := d0 * 10 + d1

		sum += number
	}

	return sum
}


func search(line string, re *regexp.Regexp) string {
	idx := re.FindAllStringSubmatchIndex(line, -1)

	return line[idx[0][0]:idx[0][1]]
}


func reverseSearch(line string, re *regexp.Regexp) string {
	for i := len(line); i >= 0; i-- {
		idx := re.FindAllStringSubmatchIndex(line[i:], -1)

		if len(idx) > 0 {
			return line[idx[0][0]+i:idx[0][1]+i]
		}
	}

	return ""
}

