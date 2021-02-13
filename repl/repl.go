package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/kanapuli/kurangu/lexer"
)

const PROMPT = ">> "

func Start(in io.Reader) {

	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != "EOF"; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
