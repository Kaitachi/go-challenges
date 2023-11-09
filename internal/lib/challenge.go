package lib

import (
	"embed"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

type Challenge[T any] struct {
	Assets		*embed.FS
	Challenge	string
	Solution	string
	DataSet		string
	Algorithm	string

	Soln		Solution[T]
}


// Kudos to @mrsoftware for the function below!
// Code adapted for use case
// https://gist.github.com/clarkmcc/1fdab4472283bb68464d066d6b4169bc?permalink_comment_id=4405804#gistcomment-4405804
func (c Challenge[any]) GetFiles() (files []string, err error) {
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


// Run current configuration
func (c Challenge[any]) Execute() {
	c.solve()
}


func (c Challenge[any]) getAssetPath() string {
	return fmt.Sprintf("assets/%s", c.Challenge)
}


func (c Challenge[any]) getDatasetInputPath() string {
	return fmt.Sprintf("%s/%s.%s.in", c.getAssetPath(), c.Solution, c.DataSet)
}


func (c Challenge[any]) getAlgorithmOutputPath() string {
	return fmt.Sprintf("%s/%s.%s.%s.out", c.getAssetPath(), c.Solution, c.DataSet, c.Algorithm)
}


func (c Challenge[any]) getSolutionInputPath() string {
	return fmt.Sprintf("%s/%s.in", c.getAssetPath(), c.Solution)
}


func (c Challenge[any]) GetDatasetInputData() (string, error) {
	data, err := c.Assets.ReadFile(c.getDatasetInputPath())

	return string(data), err
}


func (c Challenge[any]) GetAlgorithmOutputData() (string, error) {
	data, err := c.Assets.ReadFile(c.getAlgorithmOutputPath())

	return string(data), err
}


func (c Challenge[any]) GetSolutionInputData() (string, error) {
	data, err := c.Assets.ReadFile(c.getSolutionInputPath())

	return string(data), err
}



func (c Challenge[any]) solve() {

	inputData, err := c.GetDatasetInputData()
	if err != nil {
		panic("> Something went wrong while reading dataset input!")
	}

	outputData, err := c.GetAlgorithmOutputData()
	if err != nil {
		panic("> Something went wrong while reading dataset output!")
	}

	if c.Soln.Run(inputData, outputData) {
		fmt.Println("Found solution!")
	}

	fmt.Println("Finished execution.")
}

