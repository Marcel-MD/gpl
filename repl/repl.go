package repl

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"

	"github.com/Marcel-MD/gpl/evaluator"
	"github.com/Marcel-MD/gpl/lexer"
	"github.com/Marcel-MD/gpl/object"
	"github.com/Marcel-MD/gpl/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)

		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func ExecuteFile(path string, out io.Writer) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	code := string(content)
	l := lexer.New(code)
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		printParserErrors(out, p.Errors())
		return
	}

	env := object.NewEnvironment()
	evaluator.Eval(program, env)
}

func PrintAst(path string, out io.Writer) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	code := string(content)
	l := lexer.New(code)
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		printParserErrors(out, p.Errors())
		return
	}

	out.Write([]byte(program.String()))
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
