package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

const (
	exitCodeError = 1
	exitCodeUsage = 2
)

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "usage: %s <FILE>\n", os.Args[0])
		os.Exit(exitCodeUsage)
	}

	f, err := parser.ParseFile(token.NewFileSet(), args[0], nil, parser.AllErrors)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(exitCodeError)
	}

	ast.Print(token.NewFileSet(), f)
}
