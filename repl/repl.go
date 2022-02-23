package repl

import (
	"bufio"
	"fmt"
	"io"

	"Gengo/evaluator"
	"Gengo/lexer"
	"Gengo/object"
	"Gengo/parser"
)

const prompt = ">> "

// Start the REPL
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Print(prompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		/*l2 := lexer.New(line)
		for tok := l2.NextToken(); tok.Type != token.EOF; tok = l2.NextToken() {
			_, _ = io.WriteString(out, fmt.Sprintf("%+v\n", tok))
		}
		_, _ = io.WriteString(out, "\n")*/

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
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		_, _ = io.WriteString(out, "\t"+msg+"\n")
	}
}
