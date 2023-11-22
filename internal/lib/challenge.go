package lib

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"strings"
	"text/template"

	"github.com/kaitachi/go-challenges/assets"
)


type Challenger interface {
	GetSolution(string) (*Solver, bool)
}


type Challenge struct {
	Assets		*embed.FS
	Challenge	string
	Solution	string
	DataSet		[]string
	Algorithm	string
}


func NewChallenge(name string, solution string, ds []string, algo string) Challenge {

	challenge := Challenge{
			Assets: &assets.Assets,
			Challenge: name,
			Solution: solution,
			DataSet: ds,
			Algorithm: algo,
		}

	return challenge
}


func NewSolution(c Challenge, name string) {

	tokens := map[string]string{
		"SolutionName": name,
	}

	// Read template file
	file, err := c.Assets.ReadFile(c.getTemplateFilePath())
	if err != nil {
		panic(err)
	}

	// Convert template file to golang's text/template
	tmpl, err := template.New("NewSolutionFile").Parse(string(file))
	if err != nil {
		panic(err)
	}

	// New file path will be relative to our main.go location
	out, err := os.Create(fmt.Sprintf("pkg/%s/%s.go", c.Challenge, name))
	if err != nil {
		panic(err)
	}

	tmpl.Execute(out, tokens)
}


func (c Challenge) getTemplateFilePath() string {
	templatePath := ""

	fs.WalkDir(c.Assets, c.Challenge, func(path string, d fs.DirEntry, err error) error {
		if strings.HasSuffix(path, ".tmpl") {
			templatePath = path
		}

		return nil
	})

	return templatePath
}


func (c Challenge) getDatasetInputPath(scenario string) string {
	return fmt.Sprintf("%s/%s.%s.in", c.Challenge, c.Solution, scenario)
}


func (c Challenge) getAlgorithmOutputPath(scenario string) string {
	return fmt.Sprintf("%s/%s.%s.%s.out", c.Challenge, c.Solution, scenario, c.Algorithm)
}


func (c Challenge) getSolutionInputPath() string {
	return fmt.Sprintf("%s/%s.in", c.Challenge, c.Solution)
}


func (c Challenge) getScenarioData(scenario string) (string, string) {
	// Read sample input file
	inputPath := c.getDatasetInputPath(scenario)

	input, err := c.Assets.ReadFile(inputPath)
	if err != nil {
		panic(err)
	}

	// Read expected output file
	outputPath := c.getAlgorithmOutputPath(scenario)

	output, err := c.Assets.ReadFile(outputPath)
	if err != nil {
		panic(err)
	}

	return string(input), string(output)
}


func (c Challenge) getSolutionData() (string, string) {
	// Read input file
	solutionInput := c.getSolutionInputPath()

	input, err := c.Assets.ReadFile(solutionInput)
	if err != nil {
		panic(err)
	}

	return string(input), ""
}


func (c Challenge) Data(scenario string) (string, string) {
	switch scenario {
	case "": // Get real data
		return c.getSolutionData()
	
	default:
		return c.getScenarioData(scenario)
	}
}

