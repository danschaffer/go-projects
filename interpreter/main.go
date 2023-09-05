package main

import (
	"fmt"
	"log"
	"monkey/evaluator"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"monkey/repl"
	"os"
	"os/user"
	"time"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 {
		runRepl()
	} else {
		runFromFile(argsWithoutProg[0])
	}
}

func runRepl() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)

}

func runFromFile(filename string) {
	fmt.Printf("Interpreting file %s\n", filename)
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	var duration time.Duration
	start := time.Now()
	l := lexer.New(string(fileContents))
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()
	result := evaluator.Eval(program, env)
	fmt.Printf("%s\n", result.Inspect())
	duration = time.Since(start)
	fmt.Printf("%s\n", duration)
}
