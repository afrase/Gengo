package repl

import (
	"bufio"
	"fmt"
	"github.com/afrase/Gengo/evaluator"
	"github.com/afrase/Gengo/lexer"
	"github.com/afrase/Gengo/parser"
	"io"
	"github.com/afrase/Gengo/evaluator"
)

const PROMPT = ">> "

//noinspection GoUnusedParameter
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
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

		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
