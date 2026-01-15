package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/IsaacHdezA/glox/ast"
	"github.com/IsaacHdezA/glox/parser"
	"github.com/IsaacHdezA/glox/scanner"
)

var hadError bool = false
var loxError error

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("== Lox Interactive Runtime (C-d to exit) ==")

	for {
		fmt.Print("> ")
		in, _ := reader.ReadString('\n')

		if in == "" {
			fmt.Println()
			break
		}

		run(in)
		hadError = false
		loxError = nil
	}
}

func runFile(filename string) {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileInfo, _ := os.Stat(filename)
	byteCount := fileInfo.Size()

	scanner := bufio.NewScanner(file)

	content := ""
	scanner.Scan()

	for {
		content += scanner.Text()

		if !scanner.Scan() {
			break
		}
		content += "\n"
	}

	fmt.Printf("File %q (%d bytes):\n", filename, byteCount)
	run(content)

	if hadError {
		os.Exit(65)
	}
}

func run(source string) {
	loxScanner := scanner.NewScanner(source)

	tokens, err := loxScanner.ScanTokens(source)
	if err != nil {
		hadError = true

		fmt.Fprintln(os.Stderr, err.Error())
	}

	loxParser := parser.NewParser(tokens)
	expression, pErr := loxParser.Parse()

	if pErr != nil {
		hadError = true

		fmt.Fprintln(os.Stderr, pErr.Error())
	}

	fmt.Printf("[EXPRESSION (RPN)]: %v\n", ast.AstPrinterRPN{}.Print(expression))
	fmt.Printf("[EXPRESSION (Pretty)]: %v\n", ast.AstPrinter{}.Print(expression))
}

func main() {
	args, argc := os.Args[1:], len(os.Args[1:])

	if argc > 1 {
		fmt.Println("Usage: jlox [script]")
	} else if argc == 1 {
		filename := args[0]
		runFile(filename)
	} else {
		runPrompt()
	}
}
