package lib

type Problem struct {
	Challenge Challenge
	TestCase TestCase[any]
}


// 3. Assert - Every Scenario should be verified
func (p *Problem) Assert() {
	p.TestCase.Verify()
}


// Get collection of scenarios defined for this Problem
func (p *Problem) Scenarios() []string {
	return p.Challenge.DataSet
}


// Get solution for current Problem and TestCase data being used
func (p *Problem) Solution() string {
	return p.TestCase.Actual
}

