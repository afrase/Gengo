package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/afrase/Gengo/evaluator"
	"github.com/afrase/Gengo/lexer"
	"github.com/afrase/Gengo/object"
	"github.com/afrase/Gengo/parser"
	"github.com/afrase/Gengo/token"
)

// PROMPT what the REPL prompt shows
const PROMPT = ">> "

// Start the REPL
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

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
			_, _ = io.WriteString(out, fmt.Sprintf("%+v\n", tok))
			_, _ = io.WriteString(out, "\n")
		}

		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			_, _ = io.WriteString(out, evaluated.Inspect())
			_, _ = io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		_, _ = io.WriteString(out, "\t"+msg+"\n")
	}
}
