package interpreter

import (
	"j++/internal/parser"
)

type Interpreter struct {
	stack Stack
	builtins  map[string]Builtin
}

func NewInterpreter(stack Stack) *Interpreter {
	return &Interpreter{
		builtins: builtins,
		stack: stack,
	}
}

func (i *Interpreter) Eval(ast parser.AST) {
	for _,  el := range ast.Nodes {
		switch node := el.(type) {
		case parser.Symbol:
			i.builtins[node.Value](i)
		case parser.Literal:
			i.stack.Push(node)
		case parser.Quotation:
			i.Eval(parser.AST(node))
		}
	}
}
