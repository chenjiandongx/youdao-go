package main

import (
	"github.com/docopt/docopt-go"
	"strings"
)

func main() {

	usage := `yd <words> [options]

Query words meanings via the command line.

Example:
  words could be word or sentence.

  yd hello
  yd php is the best language in the world
  yd 你好

Usage:
  yd <words>...
  yd -h | --help
  yd -v | --version

Options:
  -h --help         show this help message and exit.
  -v  --version     displays the current version of youdao-go.
  `

	args, _ := docopt.ParseDoc(usage)
	queryWords := strings.Join(args["<words>"].([]string), " ")
	Parser(queryWords)

	if args["-v"] == true {
		ShowVersion()
	}
}
