package main

import (
	"os"

	Parser "github.com/acekingke/Decartes/parser"
)

func main() {
	argsWithoutProg := os.Args[1:]
	data, err := os.ReadFile(argsWithoutProg[0])
	if err != nil {
		panic(err)
	}

	Parser.ParserInit()
	Parser.Parser(string(data) + "\n")
}
