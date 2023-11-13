package lib

import "fmt"

type TestCase[T any] struct {
	Name		string
	Algorithm	string
	Input		string
	Output		string

	Data		T
	Actual		string
}


// Create Test Case with scenario data
func NewTestCase[T any](c Challenge, scenario string, algorithm string) *TestCase[T] {

	var input, output string

	switch scenario {
	case "":
		input, output = c.getSolutionData()
		break

	default:
		input, output = c.getScenarioData(scenario)
	}

	return &TestCase[T]{
		Name: scenario,
		Algorithm: algorithm,
		Input: input,
		Output: output,
	}
}


// 3. Assert - Every Scenario should be verified
func Assert(tc *TestCase[any]) {
	if tc.Output != tc.Actual {
		panic(fmt.Sprintf("> Sample scenario %s failed! Expected: %s; actual: %s.", tc.Name, tc.Output, tc.Actual))
	}
}

