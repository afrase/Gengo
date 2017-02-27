package repl

import (
	"bufio"
	"fmt"
	"github.com/afrase/Gengo/lexer"
	"github.com/afrase/Gengo/token"
	"io"
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

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
