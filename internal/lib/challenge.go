package lib

import (
	"embed"
	"fmt"
	"strings"

	"github.com/kaitachi/go-challenges/assets"
)


type Challenge struct {
	Assets		*embed.FS
	Challenge	string
	Solution	string
	Scenarios	[]string
	Algorithm	string

	Solutions	map[string]Solver
}


func NewChallenge(name string, solutions map[string]Solver) *Challenge {
	return &Challenge{
		Assets: &assets.Assets,
		Challenge: name,
		Solutions: solutions,
	}
}


func (c *Challenge) Solve() string {

	solver, ok := c.Solutions[c.Solution]
	if !ok {
		panic(fmt.Sprintf("Invalid Solution name given for %s: %s", c.Challenge, c.Solution))
	}

	// Iterate through all provided scenarios...
	for _, scenario := range c.Scenarios {
		fmt.Printf("> Running scenario %s...\n", scenario)
		tc := NewTestCase(c, scenario)

		// Each scenario provided must execute successfully
		solver.Assemble(tc)
		solver.Activate(tc)
		Assert(tc)

		fmt.Printf("> Scenario %s passed!\n", scenario)
	}

	// Once all sample scenarios have been executed successfully,
	//	we may attempt to run the final "real data" scenario

	tc := NewTestCase(c, "")

	solver.Assemble(tc)
	solver.Activate(tc)
	// Assert() // We cannot assert this scenario; we don't know what the actual value will be!

	// If everything is correct with the algorithm,
	//	this should be your final solution
	return tc.Actual
}


func (c Challenge) getScenarioData(scenario string) (string, string) {
	// Read sample input file
	inputPath := fmt.Sprintf("%s/%s.%s.in", c.Challenge, c.Solution, scenario)

	input, err := c.Assets.ReadFile(inputPath)
	if err != nil {
		panic(err)
	}

	// Read expected output file
	outputPath := fmt.Sprintf("%s/%s.%s.%s.out", c.Challenge, c.Solution, scenario, c.Algorithm)

	output, err := c.Assets.ReadFile(outputPath)
	if err != nil {
		panic(err)
	}

	return string(input), strings.TrimSpace(string(output))
}


func (c Challenge) getSolutionData() (string, string) {
	// Read input file
	solutionInput := fmt.Sprintf("%s/%s.in", c.Challenge, c.Solution)

	input, err := c.Assets.ReadFile(solutionInput)
	if err != nil {
		panic(err)
	}

	return string(input), ""
}

