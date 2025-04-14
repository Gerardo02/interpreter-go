package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/Gerardo02/interpreter-go/lexer"
	"github.com/Gerardo02/interpreter-go/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		EvalLine(line, out)
	}
}

func EvalLine(line string, out io.Writer) {
	l := lexer.New(line)

	for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		fmt.Fprintf(out, "%+v\n", tok)
	}
}
