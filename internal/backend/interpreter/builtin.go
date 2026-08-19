package interpreter

import (
	"fmt"
	"j++/internal/backend/types"
	"j++/internal/parser"
)

type Builtin func(*Interpreter)

var builtins = map[string]Builtin{
	"dup": func(i *Interpreter) {
		i.stack.Duplicate()
	},

	"swap": func(i *Interpreter) {
		i.stack.Swap()
	},

	"over": func(i *Interpreter) {
		i.stack.Over()
	},

	"print": func(i *Interpreter) {
		fmt.Printf("%-v", i.stack.Pop())
	},

	"len": func(i *Interpreter) {
		i.stack.Push(parser.Literal{Type: types.Int64, Value: i.stack.Len()})
	},

	"dip": func(i *Interpreter) {
		quotation, ok := i.stack.Pop().(parser.Quotation)
		value := i.stack.Pop()
		if ok {
			i.Eval(parser.AST(quotation))
		}
		i.stack.Push(value)
	},
}
