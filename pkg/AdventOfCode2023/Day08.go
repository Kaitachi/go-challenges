package AdventOfCode2023

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/kaitachi/go-challenges/internal/lib"
)

type Day08 struct {
	instructions []string
	tree map[string]*node
}


type node struct {
	name string
	left *node
	right *node
}


// 1. Assemble - How should we transform the data from our input files?
func (s *Day08) Assemble(tc *lib.TestCase) {

	s.instructions = make([]string, 0)
	s.tree = make(map[string]*node, 0)

	re_instructions := regexp.MustCompile(`[R|L]`)
	re_nodes := regexp.MustCompile(`(\w{3}) = \((\w{3}), (\w{3})\)`)

	first_line := strings.Split(tc.Input, "\n\n")[0]
	s.instructions = re_instructions.FindAllString(first_line, -1)

	for _, match := range re_nodes.FindAllStringSubmatch(tc.Input, -1) {
		name := match[1]
		left := match[2]
		right := match[3]

		// fmt.Printf("name: >%s<; left: >%s<; right: >%s<\n", name, left, right)

		// Create current node (if not present)
		node_current, ok := s.tree[name]
		if !ok {
			node_current = &node{
				name: name,
			}

			s.tree[name] = node_current
		}

		// Create left node (if not present)
		node_left, ok := s.tree[left]
		if !ok {
			node_left = &node{
				name: left,
			}

			s.tree[left] = node_left
		}


		// Create right node (if not present)
		node_right, ok := s.tree[right]
		if !ok {
			node_right = &node{
				name: right,
			}

			s.tree[right] = node_right
		}

		// Append left and right nodes
		node_current.left = node_left
		node_current.right = node_right
	}

}


// 2. Activate - Take our transformed input data and make the core logic needed to resolve this Problem
func (s *Day08) Activate(tc *lib.TestCase) {
	// Assign final value to TestCase.Actual field
	switch tc.Algorithm {
	case "part01":
		tc.Actual = s.part01()
		
	case "part02":
		tc.Actual = s.part02()
	}
}


func (s Day08) part01() string {

	// fmt.Printf("%+v\n", s)
	//
	// for k, v := range s.tree {
	// 	fmt.Printf("[%s]: %+v\n", k, v)
	// }

	start, _ := s.tree["AAA"]

	cursor := start

	instructions := len(s.instructions)
	steps := 0

	for steps = 0; cursor.name != "ZZZ"; steps++ {
		current := steps % instructions

		switch s.instructions[current] {
		case "L":
			cursor = cursor.left

		case "R":
			cursor = cursor.right
		}
	}

	return fmt.Sprintf("%d", steps)
}


func (s Day08) part02() string {

	re_A := regexp.MustCompile(`A$`)
	re_Z := regexp.MustCompile(`Z$`)

	cursors := make([]*node, 0)

	solves := make(map[string]int, 0)

	// Identify all starting positions
	for k, v := range s.tree {
		if re_A.MatchString(k) {
			cursors = append(cursors, v)
		}
	}

	instructions := len(s.instructions)
	steps := 0

	for steps = 0; len(cursors) != len(solves); steps++ {
		current := steps % instructions

		switch s.instructions[current] {
		case "L":
			for i, cursor := range cursors {
				cursors[i] = cursor.left
			}

		case "R":
			for i, cursor := range cursors {
				cursors[i] = cursor.right
			}
		}

		for _, cursor := range cursors {
			if _, ok := solves[cursor.name]; !ok {
				if re_Z.MatchString(cursor.name) {
					solves[cursor.name] = steps+1
				}
			}
		}
	}

	fmt.Printf(">>> %v\n", solves)

	lcm := 1

	for _, v := range solves {
		lcm = LCM(lcm, v)
	}

	steps = lcm

	return fmt.Sprintf("%d", steps)
}


// Thanks to @siongui from GitHub for the functions below!
// https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/ 
// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
      for b != 0 {
              t := b
              b = a % b
              a = t
      }
      return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
      result := a * b / GCD(a, b)

      for i := 0; i < len(integers); i++ {
              result = LCM(result, integers[i])
      }

      return result
}


