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
const nextSolutionLine string = "\t// <<NEXT>>" 


func CreateSolution(c *Challenge) {

	c.createSolutionFile()

	c.appendSolutionToChallenge()
}


func CreateChallenge(c *Challenge) {

	c.createChallengeFile()

	c.appendChallengeToMain()
}


func (c *Challenge) createSolutionFile() {

	tokens := map[string]string{
		"SolutionName": c.Solution,
	}

	// Read Template file
	file, err := c.Assets.ReadFile(c.getTemplateFilePath())
	if err != nil {
		panic(err)
	}

	// Convert Template file to golang's text/template
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


func (c *Challenge) appendSolutionToChallenge() {

	// File name to be read/written
	filePath := fmt.Sprintf("pkg/%s/%s.go", c.Challenge, c.Challenge)
	
	// Create string that will be added to our Challenge file
	newSolutionLine := fmt.Sprintf("\t\"%s\": &%s{},\n%s", c.Solution, c.Solution, nextSolutionLine)

	// Read Challenge file
	file, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	// Replace template string
	output := bytes.ReplaceAll(file, []byte(nextSolutionLine), []byte(newSolutionLine))

	if err = os.WriteFile(filePath, output, 0644); err != nil {
		panic(err)
	}
}


func (c *Challenge) createChallengeFile() {

	tokens := map[string]string{
		"ChallengeName": c.Challenge,
	}

	// Read Template file
	file, err := c.Assets.ReadFile("Challenge.tmpl")
	if err != nil {
		panic(err)
	}

	// Convert Template file to golang's text/template
	tmpl, err := template.New("NewChallengeFile").Parse(string(file))
	if err != nil {
		panic(err)
	}

	// Create new folder location if it doesn't exist
	err = os.MkdirAll(fmt.Sprintf("pkg/%s/", c.Challenge), 0744)
	if err != nil {
		panic(err)
	}

	// New file path will be relative to our main.go location
	out, err := os.Create(fmt.Sprintf("pkg/%s/%s.go", c.Challenge, c.Challenge))
	if err != nil {
		panic(err)
	}

	tmpl.Execute(out, tokens)
}


func (c *Challenge) appendChallengeToMain() {

	// File name to be read/written
	filePath := "main.go"
	
	// Create string that will be added to our Challenge file
	newImportLine := fmt.Sprintf("\t\"github.com/kaitachi/go-challenges/pkg/%s\"\n%s", c.Challenge, nextImportLine)
	newChallengeLine := fmt.Sprintf("\t\"%s\": %s.Solutions,\n%s", c.Challenge, c.Challenge, nextSolutionLine)

	// Read Challenge file
	file, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	// Replace template string
	output := file
	output = bytes.ReplaceAll(output, []byte(nextImportLine), []byte(newImportLine))
	output = bytes.ReplaceAll(output, []byte(nextSolutionLine), []byte(newChallengeLine))

	if err = os.WriteFile(filePath, output, 0644); err != nil {
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

