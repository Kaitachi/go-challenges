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
}


func NewChallenge(name string, solution string) *Challenge {
	return &Challenge{
		Assets: &assets.Assets,
		Challenge: name,
		Solution: solution,
	}
}


func (c *Challenge) Solve(solver Solver) string {

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


func (c Challenge) getFileData(scenario string) (string, string) {
	// Declare pattern strings
	var inputPattern = "%s/%s.%s.in"
	var outputPattern = "%s/%s.%s.%s.out"

	if scenario == "" {
		inputPattern = "%s/%s.in%.s" // Hacky way to bypass weirdness on Sprintf
	}

	// Read input file
	inputPath := fmt.Sprintf(inputPattern, c.Challenge, c.Solution, scenario)
	input, err := c.Assets.ReadFile(inputPath)
	if err != nil {
		print(input)
		panic(err)
	}

	// No scenario provided; we must be reading our unknown scenario
	if scenario == "" {
		return string(input), ""
	}

	// Read expected output file
	outputPath := fmt.Sprintf(outputPattern, c.Challenge, c.Solution, scenario, c.Algorithm)
	output, err := c.Assets.ReadFile(outputPath)
	if err != nil {
		panic(err)
	}

	return string(input), strings.TrimSpace(string(output))
}

