package lib

import (
	"embed"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

type Challenge struct {
	Assets		*embed.FS
	Challenge	string
	Solution	string
	DataSet		string
	Algorithm	string
}


func GetChallenge(fs *embed.FS, name string, solution string, ds string, algo string) Challenge {
	challenge := Challenge{
			Assets: fs,
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
	return fmt.Sprintf("assets/%s", c.Challenge)
}


func (c Challenge) getDatasetInputPath() string {
	return fmt.Sprintf("%s/%s.%s.in", c.getAssetPath(), c.Solution, c.DataSet)
}


func (c Challenge) getAlgorithmOutputPath() string {
	return fmt.Sprintf("%s/%s.%s.%s.out", c.getAssetPath(), c.Solution, c.DataSet, c.Algorithm)
}


func (c Challenge) getSolutionInputPath() string {
	return fmt.Sprintf("%s/%s.in", c.getAssetPath(), c.Solution)
}


func (c Challenge) GetDatasetInputData() (string, error) {
	data, err := c.Assets.ReadFile(c.getDatasetInputPath())

	return string(data), err
}


func (c Challenge) GetAlgorithmOutputData() (string, error) {
	data, err := c.Assets.ReadFile(c.getAlgorithmOutputPath())

	return string(data), err
}


func (c Challenge) GetSolutionInputData() (string, error) {
	data, err := c.Assets.ReadFile(c.getSolutionInputPath())

	return string(data), err
}


func (c Challenge) GetScenarioData() (string, string) {
	input, err := c.GetDatasetInputData()
	if err != nil {
		panic(err)
	}

	output, err := c.GetAlgorithmOutputData()
	if err != nil {
		panic(err)
	}

	return input, output
}

