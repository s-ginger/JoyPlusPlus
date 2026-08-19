package interpreter

import (
	"fmt"
	"j++/internal/parser"
)

type Interpreter struct {
	stack Stack
}

func NewInterpreter(stack Stack) *Interpreter {
	return &Interpreter{
		stack: stack,
	}
}

func (i *Interpreter) Eval(ast parser.AST) {
	for _,  el := range ast.Nodes {
		switch node := el.(type) {
		case parser.Symbol:
			if node.Value == "print" {
				val := i.stack.Pop()
				fmt.Print(val)
			} else if node.Value == "len" {
				i.stack.Push(i.stack.Len())
			}
		case parser.Literal:
			i.stack.Push(node.Value)
		case parser.Quotation:
			i.Eval(parser.AST(node))
		}
	}
}
