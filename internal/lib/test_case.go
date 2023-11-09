package lib

type TestCase[T any] struct {
	Name	string
	Input	T
	Output	string

	Actual	string
}


func (tc *TestCase[any]) IsSolution() bool {
	return tc.Output == tc.Actual
}


func (tc *TestCase[any]) IsUnknownTestCase() bool {
	return tc.Output == ""
}

