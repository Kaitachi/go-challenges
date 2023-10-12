package lib

type Solution interface {
	Run()	bool

	Assemble(TestCase)	
	Activate()
	Assert()	bool
}

