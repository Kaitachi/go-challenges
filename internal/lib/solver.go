package lib

import "fmt"

type Solver interface {
	// 1. Assemble - How should we transform the data from our input files?
	Assemble(*TestCase)
	
	// 2. Activate - Take our transformed input data and make the core logic needed to resolve this Problem
	Activate(*TestCase)
}

// 3. Assert - Every Scenario should be verified
func Assert(tc *TestCase) {
	if tc.Output != tc.Actual {
		panic(fmt.Sprintf("> Sample scenario %s failed! Expected: %s; actual: %s.", tc.Name, tc.Output, tc.Actual))
	}
}

