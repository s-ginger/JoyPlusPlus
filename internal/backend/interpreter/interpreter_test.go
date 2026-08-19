package interpreter

import (
	"j++/internal/lexer"
	"j++/internal/parser"
	"testing"
)

func TestInterpreter_Eval(t *testing.T) {

	l := lexer.NewLexer("1 dup dup len print")
	p := parser.NewParser(l.Tokenize())
	ast := p.Parse()
	i := NewInterpreter(Stack{})

	i.Eval(ast)

}

