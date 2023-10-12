package lib

type TestCase struct {
	Name	string
	Input	string
	Output	string

	Actual	string
}


func (tc *TestCase) IsSolution() bool {
	return tc.Output == tc.Actual
}


func (tc *TestCase) IsUnknownTestCase() bool {
	return tc.Output == ""
}

