package AdventOfCode2023

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/kaitachi/go-challenges/internal/lib"
)

type Day15 struct {
	data []string
}


// 1. Assemble - How should we transform the data from our input files?
func (s *Day15) Assemble(tc *lib.TestCase) {

	s.data = make([]string, 0)

	for _, line := range strings.Split(tc.Input, ",") {
		s.data = append(s.data, line)
	}
}


// 2. Activate - Take our transformed input data and make the core logic needed to resolve this Problem
func (s *Day15) Activate(tc *lib.TestCase) {
	// Assign final value to TestCase.Actual field
	switch tc.Algorithm {
	case "part01":
		tc.Actual = s.part01()
		
	case "part02":
		tc.Actual = s.part02()
	}
}


func (s Day15) part01() string {

	mult := 17
	mod := 256
	sum := 0

	for _, line := range s.data {
		sum += hash(line, mult, mod)
	}

	return fmt.Sprintf("%d", sum)
}


func (s Day15) part02() string {

	mult := 17
	mod := 256
	power := 0

	boxes := make(map[int][][]string, 0)

	re_parts := regexp.MustCompile(`(\w+)[=|\-](\d+)?`)

	for _, line := range s.data {

		instruction := re_parts.FindStringSubmatch(line)

		label := instruction[1]
		hash := hash(label, mult, mod)

		fmt.Printf("> line %s [%d]: %v\n", line, hash, instruction)


		if instruction[2] == "" {
			s.remove(&boxes, hash, label)
		} else {
			s.append(&boxes, hash, instruction[1:])
		}
	}

	for box, lenses := range boxes {
		// fmt.Printf("[%d]: %+v\n", box, lenses)
		for slot, lens := range lenses {
			if len(lens) == 0 { continue }
			focalLength, _ := strconv.Atoi(lens[1])
			power += (box+1) * (slot+1) * (focalLength)
		}
	}

	return fmt.Sprintf("%d", power)
}


func hash(s string, mult int, mod int) int {
	hash := 0
	for _, c := range s {
		value := (hash+int(c)) * mult
		// fmt.Printf("> char %c = %d, value %d\n", c, int(c), value)
		hash = (value) % mod
		// fmt.Printf(">> hash %d\n", hash)
	}

	return hash
}


func (s Day15) remove(boxes *map[int][][]string, hash int, label string) {
	lenses := (*boxes)[hash]
	index, ok := s.lensIndex(&lenses, label)

	if ok {
		(*boxes)[hash] = append(lenses[:index], lenses[index+1:]...)
	}
}


func (s Day15) append(boxes *map[int][][]string, hash int, lens []string) {
	lenses := (*boxes)[hash]
	index, ok := s.lensIndex(&lenses, lens[0])

	if ok {
		(*boxes)[hash][index] = lens
	} else {
		(*boxes)[hash] = append((*boxes)[hash], lens)
	}
}


func (s Day15) lensIndex(box *[][]string, label string) (int, bool) {
	index := 0
	ok := false

	for i := 0; i < len(*box); i++ {
		if (*box)[i][0] == label {
			index, ok = i, true
		}
	}

	return index, ok
}

