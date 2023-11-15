package lib

import (
	"embed"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

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


// Kudos to @mrsoftware for the function below!
// Code adapted for use case
// https://gist.github.com/clarkmcc/1fdab4472283bb68464d066d6b4169bc?permalink_comment_id=4405804#gistcomment-4405804
func (c Challenge) GetFiles() (files []string, err error) {
	if err := fs.WalkDir(c.Assets, c.Challenge, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
 
		if !strings.HasPrefix(filepath.Base(path), c.Solution) {
			return nil
		}

		files = append(files, path)

		return nil
	}); err != nil {
		return nil, err
	}

	return files, nil
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

