package lib

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
	"text/template"
)


func CreateSolution(c *Challenge) {

	tokens := map[string]string{
		"SolutionName": c.Solution,
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
	out, err := os.Create(fmt.Sprintf("pkg/%s/%s.go", c.Challenge, c.Solution))
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

