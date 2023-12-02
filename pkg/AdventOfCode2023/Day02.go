package AdventOfCode2023

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/kaitachi/go-challenges/internal/lib"
)

type Day02 struct {
	data []CubeGame
}


type CubeGame struct {
	id int
	rounds []CubeGameRound
}


type CubeGameRound struct {
	red int
	green int
	blue int
}


// 1. Assemble - How should we transform the data from our input files?
func (s *Day02) Assemble(tc *lib.TestCase) {

	s.data = []CubeGame{}
	games := strings.Split(tc.Input, "\n")

	re_game := regexp.MustCompile(`Game (\d+): `)
	re_round := regexp.MustCompile(`; `)
	re_r := regexp.MustCompile(`(\d+) red`)
	re_g := regexp.MustCompile(`(\d+) green`)
	re_b := regexp.MustCompile(`(\d+) blue`)

	for _, line := range games[:len(games)-1] {
		game := CubeGame{}

		if id, err := strconv.Atoi(re_game.FindStringSubmatch(line)[1]); err == nil {
			game.id = id
		}

		for _, v := range re_round.Split(line, -1) {
			round := CubeGameRound{}

			match_r := re_r.FindStringSubmatch(v)
			match_g := re_g.FindStringSubmatch(v)
			match_b := re_b.FindStringSubmatch(v)

			if len(match_r) > 1 {
				round.red, _ = strconv.Atoi(match_r[1])
			}

			if len(match_g) > 1 {
				round.green, _ = strconv.Atoi(match_g[1])
			}

			if len(match_b) > 1 {
				round.blue, _ = strconv.Atoi(match_b[1])
			}

			game.rounds = append(game.rounds, round)
		}

		s.data = append(s.data, game)
	}
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

	max_r := 12
	max_g := 13
	max_b := 14
	id_sum := 0

	next_game: for _, game := range s.data {
		for _, round := range game.rounds {
			if round.red > max_r || round.green > max_g || round.blue > max_b {
				continue next_game
			}
		}

		id_sum += game.id
	}

	return fmt.Sprintf("%d", id_sum)
}


func (s Day02) part02() string {

	return fmt.Sprintf("%d", -1)
}

