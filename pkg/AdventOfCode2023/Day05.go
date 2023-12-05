package AdventOfCode2023

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/kaitachi/go-challenges/internal/lib"
)

type Day05 struct {
	seeds		[]int
	maps		[]string
	transforms	map[string][]transform
}


type transform struct {
	source	int
	dest	int
	length	int
}


// 1. Assemble - How should we transform the data from our input files?
func (s *Day05) Assemble(tc *lib.TestCase) {

	s.seeds = make([]int, 0)
	s.maps = make([]string, 0)
	s.transforms = make(map[string][]transform, 0)

	re_section := regexp.MustCompile(`(^[\w\s\-]+):`)
	re_numbers := regexp.MustCompile(`\d+`)

	for _, section := range strings.Split(tc.Input, "\n\n") {
		//fmt.Println("----- NEW SECTION")
		
		sectionName := re_section.FindStringSubmatch(section)
		//fmt.Printf("%v\n", sectionName[0])

		if sectionName[1] == "seeds" {
			// build up seed array
			nums := re_numbers.FindAllStringSubmatch(section, -1)

			for _, num := range nums {
				number, _ := strconv.Atoi(num[0])
				s.seeds = append(s.seeds, number)
			}
		} else {
			// build up transform array (maps)
			lines := strings.Split(section, "\n")
			transforms := make([]transform, len(lines)-1)

			for i, line := range lines[1:] {
				//fmt.Printf("l> %v\n", line)
				
				nums := re_numbers.FindAllStringSubmatch(line, -1)

				if len(nums) == 3 {
					dest, _ := strconv.Atoi(nums[0][0])
					source, _ := strconv.Atoi(nums[1][0])
					length, _ := strconv.Atoi(nums[2][0])

					transforms[i] = transform{
						source: source,
						dest: dest,
						length: length,
					}
				}
			}

			s.maps = append(s.maps, sectionName[0])
			s.transforms[sectionName[0]] = transforms
		}
	}
}


// 2. Activate - Take our transformed input data and make the core logic needed to resolve this Problem
func (s *Day05) Activate(tc *lib.TestCase) {
	// Assign final value to TestCase.Actual field
	switch tc.Algorithm {
	case "part01":
		tc.Actual = s.part01()
		
	case "part02":
		tc.Actual = s.part02()
	}
}


func (s Day05) part01() string {

	maps := make(map[string]func(int) int, len(s.maps))

	for name, transforms := range s.transforms {
		maps[name] = mapper(transforms)
	}

	minLocation := math.MaxInt

	for _, seed := range s.seeds {
		value := seed

		for _, transform := range s.maps {
			value = maps[transform](value)
		}

		minLocation = int(math.Min(float64(minLocation), float64(value)))
	}

	return fmt.Sprintf("%d", minLocation)
}


func (s Day05) part02() string {

	maps := make(map[string]func(int) int, len(s.maps))

	for name, transforms := range s.transforms {
		maps[name] = mapper(transforms)
	}

	minLocation := math.MaxInt

	for pair := 0; pair < len(s.seeds); pair += 2 {
		for i := 0; i < s.seeds[pair+1]; i++ {
			//initial := s.seeds[pair]+i
			value := s.seeds[pair]+i

			for _, transform := range s.maps {
				value = maps[transform](value)
			}

			//fmt.Printf("> Seed number %d goes to location %d\n", initial, value)

			minLocation = int(math.Min(float64(minLocation), float64(value)))
		}
	}

	fmt.Println("> DONE")

	return fmt.Sprintf("%d", minLocation)
}


func mapper(transforms []transform) func(int) int {
	return func(key int) int {
		for _, t := range transforms {
			if t.source <= key && key < t.source + t.length {
				offset := key - t.source
				out := t.dest + offset

				//fmt.Printf("> HIT %d (offset %d) => %d \n", key, offset, out)
				return out
			}
		}

		return key
	}
}

