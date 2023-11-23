package lib

import "fmt"

type TestCase struct {
	Name		string
	Algorithm	string
	Input		string
	Output		string

	Actual		string
}


// 3. Assert - Every Scenario should be verified
func Assert(tc *TestCase) {
	if tc.Output != tc.Actual {
		panic(fmt.Sprintf("> Sample scenario %s failed! Expected: %s; actual: %s.", tc.Name, tc.Output, tc.Actual))
	}
}

