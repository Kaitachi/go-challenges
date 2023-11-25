package lib

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"strings"
	"text/template"
)


const nextImportLine string = "\t// <<NEXT_IMPORT>>"
const nextChallengeLine string = "\t// <<NEXT_CHALLENGE>>"
const nextSolutionLine string = "\t// <<NEXT_SOLUTION>>"

type replacementMappings struct {
	text	string
	replace	string
}


func CreateChallenge(c *Challenge) {

	c.createFileFromTemplate("Challenge.tmpl", fmt.Sprintf("pkg/%s/%s.go", c.Challenge, c.Challenge))

	replacements := []replacementMappings{
		{
			text: nextImportLine,
			replace: fmt.Sprintf("\t\"github.com/kaitachi/go-challenges/pkg/%s\"\n%s", c.Challenge, nextImportLine),
		},
		{
			text: nextChallengeLine,
			replace: fmt.Sprintf("\t\"%s\": %s.Solutions,\n%s", c.Challenge, c.Challenge, nextSolutionLine),
		},
	}

	c.appendNewFileToParent("main.go", replacements)
}


func CreateSolution(c *Challenge) {

	c.createFileFromTemplate(c.getTemplateFilePath(), fmt.Sprintf("pkg/%s/%s.go", c.Challenge, c.Solution))

	replacements := []replacementMappings{
		{
			text: nextSolutionLine,
			replace: fmt.Sprintf("\t\"%s\": &%s{},\n%s", c.Solution, c.Solution, nextSolutionLine),

		},
	}

	c.appendNewFileToParent(fmt.Sprintf("pkg/%s/%s.go", c.Challenge, c.Challenge), replacements)
}


func (c *Challenge) createFileFromTemplate(src string, dest string) {
	
	// Read Template file
	file, err := c.Assets.ReadFile(src)
	if err != nil {
		panic(err)
	}

	// Convert Template file to golang's text/template
	tmpl, err := template.New("NewTemplateFile").Parse(string(file))
	if err != nil {
		panic(err)
	}

	// Create new folder location if it doesn't exist
	err = os.MkdirAll(fmt.Sprintf("pkg/%s/", c.Challenge), 0744)
	if err != nil {
		panic(err)
	}

	// New file path will be relative to our main.go location
	out, err := os.Create(dest)
	if err != nil {
		panic(err)
	}

	tmpl.Execute(out, c)
}


func (c *Challenge) appendNewFileToParent(src string, tokens []replacementMappings) {

	// Read Challenge file
	file, err := os.ReadFile(src)
	if err != nil {
		panic(err)
	}

	// Replace template string
	output := file
	for _, token := range tokens {
		output = bytes.ReplaceAll(output, []byte(token.text), []byte(token.replace))
	}

	if err = os.WriteFile(src, output, 0644); err != nil {
		panic(err)
	}
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

