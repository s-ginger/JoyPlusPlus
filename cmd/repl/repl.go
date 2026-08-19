package repl

import (
	"bufio"
	"fmt"
	"j++/internal/backend/interpreter"
	"j++/internal/lexer"
	"j++/internal/parser"
	"os"
)

func StartRepl() {
	reader := bufio.NewReader(os.Stdin)

	i := interpreter.NewInterpreter(interpreter.Stack{})
	l := lexer.NewZeroLexer()
	p := parser.NewZeroParser()

	for {
		fmt.Print("> ")

		input, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		l.SetSrc(input)
		p.SetTokens(l.Tokenize())

		ast := p.Parse()

		i.Eval(ast)
		i.Print()

		l.Flush()
		p.Flush()
	}
}


