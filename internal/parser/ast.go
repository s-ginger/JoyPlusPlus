package parser

import "j++/internal/backend/types"

type Node interface {
	node()
}

type Literal struct {
	Type types.ValueType
	Value any
}

func (Literal) node() {}

type Symbol struct {
	Value string
}

func (Symbol) node() {}

type Quotation struct {
	Nodes []Node
}

func (Quotation) node() {}

type AST struct {
	Nodes []Node
}
