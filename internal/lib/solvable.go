package lib

import "fmt"


type Solvable interface {
	Assemble(string)
	Activate()
	Assert()

	Scenarios() []string
	Solution() string
}


func Solve(s Solvable) string {
	
	// Iterate through all provided scenarios...
	for _, scenario := range s.Scenarios() {
		fmt.Printf("> Running scenario %s...\n", scenario)

		// Each scenario provided must execute successfully
		s.Assemble(scenario)
		s.Activate()
		s.Assert()

		fmt.Printf("> Scenario %s passed!\n", scenario)
	}

	// Once all sample scenarios have been executed successfully,
	//	we may attempt to run the final "real data" scenario
	s.Assemble("")
	s.Activate()
	// s.Assert() // We cannot assert this scenario; we don't know the actual value just yet

	// If everything is correct with the algorithm,
	//	this should be your final solution
	return s.Solution()
}

