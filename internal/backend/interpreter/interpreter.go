package interpreter

import (
	"j++/internal/backend/types"
	"j++/internal/parser"
)

type Interpreter struct {
	stack    Stack
	builtins map[string]Builtin
}

func NewInterpreter(stack Stack) *Interpreter {
	return &Interpreter{
		builtins: builtins,
		stack:    stack,
	}
}

func (i *Interpreter) Eval(nodes parser.Nodes) {
	for _, el := range nodes {
		switch node := el.(type) {
		case parser.Symbol:
			builtin, ok := i.builtins[node.Value]
			if !ok {
				panic("unknown word: " + node.Value)
			}
			builtin(i)
		case parser.Literal:
			i.evalLiteral(node)
		case parser.Quotation:
			i.stack.Push(node)
		}
	}
}

func (i *Interpreter) evalLiteral(literal parser.Literal) {
	switch literal.Type {
	case types.Int32:
		value, ok := literal.Value.(int32)
		if !ok {
			panic("invalid int32 literal")
		}
		i.stack.Push(value)
	case types.Int64:
		value, ok := literal.Value.(int64)
		if !ok {
			panic("invalid int64 literal")
		}
		i.stack.Push(value)
	case types.Float64:
		value, ok := literal.Value.(float64)
		if !ok {
			panic("invalid int64 literal")
		}
		i.stack.Push(value)
	case types.Float32:
		value, ok := literal.Value.(float32)
		if !ok {
			panic("invalid int64 literal")
		}
		i.stack.Push(value)
	default:
		panic("A unknown or nil type")
	}
}

