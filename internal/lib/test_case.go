package lib

import "fmt"

type TestCase[T any] struct {
	Name	string
	Input	T
	Output	string

	Actual	string
}


func (tc *TestCase[any]) Verify() {
	if tc.Output != tc.Actual {
		panic(fmt.Sprintf("> Sample scenario %s failed! Expected: %s; actual: %s.", tc.Name, tc.Output, tc.Actual))
	}
}

