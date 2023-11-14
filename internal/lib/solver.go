package lib

import (
	"fmt"
	"reflect"
)


type Solver interface {
	Assemble(*TestCase)
	Activate(*TestCase)
}


func Solve(c Challenger, s Solver, scenarios []string, algorithm string) string {

	// Calculate challenge & problem name
	challengeName := NameOf(c)
	problemName := NameOf(s)

	challenge := NewChallenge(challengeName, problemName, scenarios, algorithm)
	
	// Iterate through all provided scenarios...
	for _, scenario := range scenarios {
		fmt.Printf("> Running scenario %s...\n", scenario)
		input, output := challenge.Data(scenario)

		tc := NewTestCase(input, output, scenario, algorithm)

		// Each scenario provided must execute successfully
		s.Assemble(tc)
		s.Activate(tc)
		Assert(tc)

		fmt.Printf("> Scenario %s passed!\n", scenario)
	}

	input, output := challenge.Data("")

	tc := NewTestCase(input, output, "", algorithm)

	// Once all sample scenarios have been executed successfully,
	//	we may attempt to run the final "real data" scenario
	s.Assemble(tc)
	s.Activate(tc)
	// Assert() // We cannot assert this scenario; we don't know the actual value just yet

	// If everything is correct with the algorithm,
	//	this should be your final solution
	return tc.Actual
}


func NameOf(i interface{}) string {
	return reflect.TypeOf(i).Elem().Name()
}

