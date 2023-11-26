package lib

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"github.com/joho/godotenv"
	"github.com/mattn/godown"
)

type scraper interface {
	Scrape(Challenge)
}


type AdventOfCodeScraper struct {
	url			string
	input		string
	selector	string
	header		string
	token		string
}


func GetAdventOfCodeScraper() *AdventOfCodeScraper {
	return &AdventOfCodeScraper{
		url: "https://adventofcode.com/%d/day/%d",
		input: "https://adventofcode.com/%d/day/%d/input",
		selector: "body > main",
		header: "session=%s",
		token: "AOC_SESSION_TOKEN",
	}
}


func (s *AdventOfCodeScraper) Scrape(c *Challenge) {
	s.scrapeExplanation(c)
	s.scrapeInput(c)
}


func (s *AdventOfCodeScraper) scrapeInput(c *Challenge) {
	digits := regexp.MustCompile("[0-9]+")

	year, err := strconv.Atoi(digits.FindString(c.Challenge))
	if err != nil {
		panic(err)
	}

	day, err := strconv.Atoi(digits.FindString(c.Solution))
	if err != nil {
		panic(err)
	}


	coll := colly.NewCollector()

	err = godotenv.Load()
	if err != nil {
		panic(err)
	}

	coll.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)

		r.Headers.Set("cookie", fmt.Sprintf(s.header, os.Getenv(s.token)))
	})

	coll.OnError(func(r *colly.Response, err error) {
		fmt.Println("Something went wrong: ", err)
	})

	coll.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)
		var b bytes.Buffer
		b.Write(r.Body)

		createFile(fmt.Sprintf("assets/%s/%s.in", c.Challenge, c.Solution), b)
	})

	coll.OnScraped(func(r *colly.Response) {
		fmt.Println(r.Request.URL, " scraped!")
	})

	coll.Visit(fmt.Sprintf(s.input, year, day))
}


func (s *AdventOfCodeScraper) scrapeExplanation(c *Challenge) {
	digits := regexp.MustCompile("[0-9]+")

	year, err := strconv.Atoi(digits.FindString(c.Challenge))
	if err != nil {
		panic(err)
	}

	day, err := strconv.Atoi(digits.FindString(c.Solution))
	if err != nil {
		panic(err)
	}


	coll := colly.NewCollector()

	err = godotenv.Load()
	if err != nil {
		panic(err)
	}

	coll.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)

		r.Headers.Set("cookie", fmt.Sprintf(s.header, os.Getenv(s.token)))
	})

	coll.OnError(func(r *colly.Response, err error) {
		fmt.Println("Something went wrong: ", err)
	})

	coll.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)
	})

	coll.OnHTML(s.selector, func(e *colly.HTMLElement) {
		html, err := e.DOM.Html()
		if err != nil {
			panic(err)
		}

		md := Markdownify(html)

		createFile(fmt.Sprintf("assets/%s/%s.md", c.Challenge, c.Solution), md)
	})

	coll.OnScraped(func(r *colly.Response) {
		fmt.Println(r.Request.URL, " scraped!")
	})

	coll.Visit(fmt.Sprintf(s.url, year, day))
}


func Markdownify(html string) bytes.Buffer {
	var buf bytes.Buffer

	err := godown.Convert(&buf, strings.NewReader(html), nil)
	if err != nil {
		panic(err)
	}

	return buf
}


func createFile(dest string, text bytes.Buffer) {
	// Create new folder location if it doesn't exist
	err := os.MkdirAll(path.Dir(dest), 0744)
	if err != nil {
		panic(err)
	}

	// New file path will be relative to our main.go location
	_, err = os.Create(dest)
	if err != nil {
		panic(err)
	}

	os.WriteFile(dest, text.Bytes(), 0744)
}

