package lib

type Solver interface {
	Assemble(*TestCase)
	Activate(*TestCase)
}

