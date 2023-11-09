package lib

type Solution[T any] interface {
	Run(string, string)	bool

	Assemble(*TestCase[T], string, string)
	Activate(*TestCase[T])
	Assert()	bool
}

