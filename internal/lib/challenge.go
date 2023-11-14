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
	GetProblem(string) (Solvable, bool)
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
	if err := fs.WalkDir(c.Assets, ".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
 
		if !strings.HasPrefix(path, c.getAssetPath()) {
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


func (c Challenge) getAssetPath() string {
	return fmt.Sprintf("%s", c.Challenge)
}


func (c Challenge) getDatasetInputPath(scenario string) string {
	return fmt.Sprintf("%s/%s.%s.in", c.getAssetPath(), c.Solution, scenario)
}


func (c Challenge) getAlgorithmOutputPath(scenario string) string {
	return fmt.Sprintf("%s/%s.%s.%s.out", c.getAssetPath(), c.Solution, scenario, c.Algorithm)
}


func (c Challenge) getSolutionInputPath() string {
	return fmt.Sprintf("%s/%s.in", c.getAssetPath(), c.Solution)
}


func (c Challenge) getDatasetInputData(scenario string) (string, error) {
	data, err := c.Assets.ReadFile(c.getDatasetInputPath(scenario))

	return string(data), err
}


func (c Challenge) getAlgorithmOutputData(scenario string) (string, error) {
	data, err := c.Assets.ReadFile(c.getAlgorithmOutputPath(scenario))

	return string(data), err
}


func (c Challenge) getSolutionInputData() (string, error) {
	data, err := c.Assets.ReadFile(c.getSolutionInputPath())

	return string(data), err
}


func (c Challenge) getScenarioData(scenario string) (string, string) {
	input, err := c.getDatasetInputData(scenario)
	if err != nil {
		panic(err)
	}

	output, err := c.getAlgorithmOutputData(scenario)
	if err != nil {
		panic(err)
	}

	return input, output
}


func (c Challenge) getSolutionData() (string, string) {
	input, err := c.getSolutionInputData()
	if err != nil {
		panic(err)
	}

	return input, ""
}


func (c Challenge) Data(scenario string) (string, string) {
	switch scenario {
	case "": // Get real data
		return c.getSolutionData()
	
	default:
		return c.getScenarioData(scenario)
	}
}

