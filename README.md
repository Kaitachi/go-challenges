![GitHub go.mod Go version (branch & subdirectory of monorepo)][badge-go-version]
![GitHub License][badge-license]

# Go-Challenges
Welcome! Over here, you'll find a collection of various Challenges I've solved using Golang. This repository has been worked out to be as generic and flexible as possible, allowing for most Coding Challenges to have a place in here (thus avoiding multiple repos per website).

Some examples of compatible Coding Challenge websites include, but are not limited to:
* [Advent of Code](https://adventofcode.com)
* [Google Coding Competitions](https://github.com/google/coding-competitions-archive)
* [Project Euler](https://projecteuler.net)
* [Hacker Rank](https://www.hackerrank.com)
* ...and potentially many others!


# Project Listing
<details>
	<summary>List of Projects</summary>
	<ul>
		<li><a href="pkg/AdventOfCode2022/">Advent of Code 2022</a></li>
		<li><a href="pkg/AdventOfCode2023/">Advent of Code 2023</a></li>
	</ul>
</details>


# How It Works
This repository has a [thin layer](internal/lib/README.md) that has been designed to work in a simple fashion:
1. Read input files from the `assets` folder
2. Example/scenario files are read and executed
	* These example files consist of a single input file (with the sample input) and potentially multiple output files (with expected outputs)
3. If all examples run successfully, the program then proceeds to execute the "real" scenario (a scenario for which the expected output is not known)


# Getting Started
TODO: getting started


# Contributing
Since this is a repository designed to mainly contain all algorithms and progress I have over various Coding Challenges, contributions are not expected. However, if you would like to contribute to the core logic orchestrating code executions (anything within the `internal` folder), feel free to do so! Contributions to the `pkg` folder would most likely be ignored, since I'd like to maintain this a project that holds my own solution attempts as much as possible.


<!-- Links -->
[badge-license]: https://img.shields.io/github/license/Kaitachi/go-challenges
[badge-go-version]: https://img.shields.io/github/go-mod/go-version/Kaitachi/go-challenges/main

