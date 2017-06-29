package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/afrase/Gengo/evaluator"
	"github.com/afrase/Gengo/lexer"
	"github.com/afrase/Gengo/parser"
	"github.com/afrase/Gengo/token"
)

const PROMPT = ">> "

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
		l2 := lexer.New(line)

		for tok := l2.NextToken(); tok.Type != token.EOF; tok = l2.NextToken() {
			io.WriteString(out, fmt.Sprintf("%+v\n", tok))
			io.WriteString(out, "\n")
		}

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
