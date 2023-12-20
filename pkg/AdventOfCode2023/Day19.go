package AdventOfCode2023

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/kaitachi/go-challenges/internal/lib"
)

type Day19 struct {
	workflows	map[string][]rule
	parts		[]xmasPart
}


type rule struct {
	condition	string
	sendTo		string
}


type xmasPart struct {
	x int
	m int
	a int
	s int
}


var re_inequality = regexp.MustCompile(`(\d+)([>=<])(\d+)`)


// 1. Assemble - How should we transform the data from our input files?
func (s *Day19) Assemble(tc *lib.TestCase) {

	s.workflows = make(map[string][]rule, 0)
	s.parts = make([]xmasPart, 0)

	// Identify all workflows
	re_workflow := regexp.MustCompile(`(\w+)\{(.*)\}`)

	for _, match := range re_workflow.FindAllStringSubmatch(tc.Input, -1) {
		name := match[1]
		rules := make([]rule, 0)
		
		for _, rulesString := range strings.Split(match[2], ",") {
			rule := rule{}
			
			ruleParts := strings.Split(rulesString, ":")

			switch len(ruleParts) {
			case 2: // Got last step in workflow
				rule.condition = ruleParts[0]
				rule.sendTo = ruleParts[1]

			case 1: // Got intermediate step in workflow
				rule.condition = "1=1"
				rule.sendTo = ruleParts[0]
			}

			rules = append(rules, rule)
		}

		s.workflows[name] = rules
	}

	// Identify all parts
	re_xmasPart := regexp.MustCompile(`\{x=(\d+),m=(\d+),a=(\d+),s=(\d+)\}`)
	
	for _, match := range re_xmasPart.FindAllStringSubmatch(tc.Input, -1) {
		_x, _ := strconv.Atoi(match[1])
		_m, _ := strconv.Atoi(match[2])
		_a, _ := strconv.Atoi(match[3])
		_s, _ := strconv.Atoi(match[4])

		s.parts = append(s.parts, xmasPart{
			x: _x,
			m: _m,
			a: _a,
			s: _s,
		})
	}
}


// 2. Activate - Take our transformed input data and make the core logic needed to resolve this Problem
func (s *Day19) Activate(tc *lib.TestCase) {
	// Assign final value to TestCase.Actual field
	switch tc.Algorithm {
	case "part01":
		tc.Actual = s.part01()
		
	case "part02":
		tc.Actual = s.part02()
	}
}


func (s Day19) part01() string {

	accepted := 0
	rejected := 0

	// for k, v := range s.workflows {
	// 	fmt.Printf("[%s]: %v\n", k, v)
	// }

	for _, v := range s.parts {
		// fmt.Printf("[%d]: %+v\n", k, v)

		a, r := s.sort(v)

		accepted += a
		rejected += r
	}

	return fmt.Sprintf("%d", accepted)
}


func (s Day19) part02() string {

	accepted := 0

	for _x := 1; _x <= 4000; _x++ {
		for _m := 1; _m <= 4000; _m++ {
			for _a := 1; _a <= 4000; _a++ {
				for _s := 1; _s <= 4000; _s++ {
					xmas := xmasPart{
						x: _x,
						m: _m,
						a: _a,
						s: _s,
					}

					accept, _ := s.sort(xmas)

					if accept > 0 {
						accepted++
					}
				}
			}
		}
	}

	return fmt.Sprintf("%d", accepted)
}


func (s Day19) sort(xmas xmasPart) (accepted int, rejected int) {
	workflow := "in"

	for workflow != "A" && workflow != "R" {
		// fmt.Printf(">> [%s]: %v ", workflow, s.workflows[workflow])

		for _, rule := range s.workflows[workflow] {
			if evaluate(rule.condition, xmas) {
				workflow = rule.sendTo
				break
			}
		}
	}

	switch workflow {
	case "A": return xmas.sum(), 0
	case "R": return 0, xmas.sum()
	default: return 0, 0
	}
}


func evaluate(c string, xmas xmasPart) bool {
	c = strings.ReplaceAll(c, "x", strconv.Itoa(xmas.x))
	c = strings.ReplaceAll(c, "m", strconv.Itoa(xmas.m))
	c = strings.ReplaceAll(c, "a", strconv.Itoa(xmas.a))
	c = strings.ReplaceAll(c, "s", strconv.Itoa(xmas.s))

	inequality := re_inequality.FindStringSubmatch(c)
	op1, _ := strconv.Atoi(inequality[1])
	op2, _ := strconv.Atoi(inequality[3])

	switch inequality[2] {
	case "<": return op1 < op2
	case ">": return op1 > op2
	case "=": return op1 == op2
	}

	return false
}


func (xmas xmasPart) sum() int {
	return xmas.x + xmas.m + xmas.a + xmas.s
}

