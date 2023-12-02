package AdventOfCode2023

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/kaitachi/go-challenges/internal/lib"
)

type Day02 struct {
	data []string
}


// 1. Assemble - How should we transform the data from our input files?
func (s *Day02) Assemble(tc *lib.TestCase) {
	s.data = strings.Split(tc.Input, "\n")
}


// 2. Activate - Take our transformed input data and make the core logic needed to resolve this Problem
func (s *Day02) Activate(tc *lib.TestCase) {

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


func (s Day02) part01() string {
	id_sum := 0

	max_allowed := map[string]float64{
		"red": 12,
		"green": 13,
		"blue": 14,
	}

	next_game: for id, game := range maxCubesPerGame(s.data) {
		for color, max := range max_allowed {
			if game[color] > max {
				continue next_game
			}
		}

		id_sum += id+1
	}

	return fmt.Sprintf("%d", id_sum)
}


func (s Day02) part02() string {
	power_sum := float64(0)

	for _, game := range maxCubesPerGame(s.data) {
		game_product := float64(1)

		for _, max := range game {
			game_product *= max
		}

		power_sum += game_product
	}

	return fmt.Sprintf("%.0f", power_sum)
}


func maxCubesPerGame(games []string) []map[string]float64 {
	re_cubes := regexp.MustCompile(`(\d+) (red|green|blue)`)
	game_cubes := make([]map[string]float64, 0)

	for _, game := range games {
		cubes := re_cubes.FindAllStringSubmatch(game, -1)

		max_cubes := map[string]float64{
			"red": 0,
			"green": 0,
			"blue": 0,
		}

		for _, match := range cubes {
			if count, err := strconv.ParseFloat(match[1], 64); err == nil {
				max_cubes[match[2]] = math.Max(max_cubes[match[2]], count)
			}
		}

		game_cubes = append(game_cubes, max_cubes)
	}

	return game_cubes
}

