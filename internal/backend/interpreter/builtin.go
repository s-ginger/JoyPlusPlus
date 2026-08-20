package interpreter

import (
	"fmt"
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
		val := i.stack.Pop()
		fmt.Printf("%v", val)
		i.stack.Push(val)
	},

	"len": func(i *Interpreter) {
		i.stack.Push(i.stack.Len())
	},

	"i": func(i *Interpreter) {
		quotation, ok := i.stack.Pop().(parser.Quotation)
		if !ok {
			panic("i: expected quotation")
		}
		i.Eval(parser.AST{
			Nodes: quotation.Nodes,
		})
	},

	"dip": func(i *Interpreter) {
		quotation, ok := i.stack.Pop().(parser.Quotation)
		if !ok {
			panic("dip: expected quotation")
		}

		value := i.stack.Pop()

		i.Eval(parser.AST{
			Nodes: quotation.Nodes,
		})

		i.stack.Push(value)
	},

	"free!" : func(i *Interpreter) {
		i.stack.values = []any{}
	},
}