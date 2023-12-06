package AdventOfCode2023

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/kaitachi/go-challenges/internal/lib"
)

type Day06 struct {
	data map[string][]int
}


// 1. Assemble - How should we transform the data from our input files?
func (s *Day06) Assemble(tc *lib.TestCase) {

	s.data = make(map[string][]int, 2)

	re_section := regexp.MustCompile(`\w+`)
	re_numbers := regexp.MustCompile(`\d+`)

	for _, line := range strings.Split(tc.Input, "\n") {
		section := re_section.FindString(line)
		numbers := re_numbers.FindAllString(line, -1)

		nums := make([]int, len(numbers))

		for i, number := range numbers {
			nums[i], _ = strconv.Atoi(number)
		}

		s.data[section] = nums
	}
}


// 2. Activate - Take our transformed input data and make the core logic needed to resolve this Problem
func (s *Day06) Activate(tc *lib.TestCase) {
	// Assign final value to TestCase.Actual field
	switch tc.Algorithm {
	case "part01":
		tc.Actual = s.part01()
		
	case "part02":
		tc.Actual = s.part02()
	}
}


func (s Day06) part01() string {

	rounds := len(s.data["Time"])
	records := make([]int, rounds)

	for i := 0; i < rounds; i++ {
		time := s.data["Time"][i]
		dist := s.data["Distance"][i]

		// How long are we going to keep the button pressed? (speed)
		for speed := 0; speed < time; speed++ {
			reach := speed * (time - speed)

			if dist < reach {
				records[i] += 1
			}
		}
	}

	product := 1

	for _, record := range records {
		if record > 0 {
			product *= record
		}
	}

	return fmt.Sprintf("%d", product)
}


func (s Day06) part02() string {

	records := 0

	t := ""
	d := ""

	for i := 0; i < len(s.data["Time"]); i++ {
		t += fmt.Sprint(s.data["Time"][i])
		d += fmt.Sprint(s.data["Distance"][i])
	}

	time, _ := strconv.Atoi(t)
	dist, _ := strconv.Atoi(d)

	// How long are we going to keep the button pressed? (speed)
	for speed := 0; speed < time; speed++ {
		reach := speed * (time - speed)

		if dist < reach {
			records += 1
		}
	}

	return fmt.Sprintf("%d", records)
}

