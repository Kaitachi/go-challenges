package lib

import (
	"fmt"
	"reflect"
	"strings"
)


type Solvable[T any] interface {
	Assemble(*TestCase[T])
	Activate(*TestCase[T])
}


func Solve(s Solvable[any], scenarios []string, algorithm string) string {

	// Calculate challenge & problem name
	challengeName := ChallengeOf(s)
	problemName := NameOf(s)

	c := NewChallenge(challengeName, problemName, scenarios, algorithm)
	
	// Iterate through all provided scenarios...
	for _, scenario := range scenarios {
		fmt.Printf("> Running scenario %s...\n", scenario)
		input, output := c.Data(scenario)

		tc := NewTestCase[any](input, output, scenario, algorithm)

		// Each scenario provided must execute successfully
		s.Assemble(tc)
		s.Activate(tc)
		Assert(tc)

		fmt.Printf("> Scenario %s passed!\n", scenario)
	}

	input, output := c.Data("")

	tc := NewTestCase[any](input, output, "", algorithm)

	// Once all sample scenarios have been executed successfully,
	//	we may attempt to run the final "real data" scenario
	s.Assemble(tc)
	s.Activate(tc)
	// Assert() // We cannot assert this scenario; we don't know the actual value just yet

	// If everything is correct with the algorithm,
	//	this should be your final solution
	return tc.Actual
}


func ChallengeOf(s Solvable[any]) string {
	challengePath := strings.Split(reflect.TypeOf(s).PkgPath(), "/")
	return challengePath[len(challengePath)-1]
}


func NameOf(s Solvable[any]) string {
	return reflect.TypeOf(s).Name()
}

