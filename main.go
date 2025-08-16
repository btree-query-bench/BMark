package main

import (
	"bufio"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	parser "github.com/btree-query-bench/bmark/antlr"
	"github.com/btree-query-bench/bmark/query"
	"os"
	"strings"
)

func main() {
	fmt.Println("\n ______ _______ _______ ______ __  __ \n|   __ \\   |   |   _   |   __ \\  |/  |\n|   __ <       |       |      <     < \n|______/__|_|__|___|___|___|__|__|\\__|")
	fmt.Println("-----------------------------------------")
	fmt.Println("Please enter your query.")

	visitor := &query.Executor{}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("BMARK - type 'exit' to quit")

	for {
		fmt.Print(">> ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if line == "exit" || line == "quit" {
			break
		}
		if line == "" {
			continue
		}

		input := antlr.NewInputStream(line)
		lexer := parser.NewQueryLangLexer(input)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
		langParser := parser.NewQueryLangParser(stream)

		tree := langParser.Query()
		visitor.VisitQuery(tree.(*parser.QueryContext))
	}

}
