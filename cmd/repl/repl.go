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
	fmt.Print("Welcome to Joy++ repl v0.2\n")
	fmt.Print("Type help for more information")
	reader := bufio.NewReader(os.Stdin)

	i := interpreter.NewInterpreter(interpreter.Stack{})
	l := lexer.NewZeroLexer()
	p := parser.NewZeroParser()

	for {
		fmt.Print("\n> ")

		input, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		l.SetSrc(input)
		p.SetTokens(l.Tokenize())

		ast := p.Parse()

		i.Eval(ast)

		l.Flush()
		p.Flush()
	}
}


