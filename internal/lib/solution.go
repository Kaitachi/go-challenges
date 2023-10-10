package lib

type Solution struct {
	Algorithm	string

	TestCases	[]*TestCase
}


// Read data
func (s *Solution) Assemble() {
}


// Calculate Actual output for all TestCases
func (s *Solution) Activate() {
}


// Do all TestCases pass?
func (s *Solution) Assert() {
}


// If all assertions are true, we can go ahead and solve this exercise.
func (s *Solution) Solve() {
}

