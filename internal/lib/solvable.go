package lib


type Solvable interface {
	Assemble()
	Activate()
	Assert() bool
}


func Solve(s Solvable) {
	s.Assemble()
	s.Activate()
	s.Assert()
}

