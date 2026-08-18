package parser

type Node interface {
	node()
}

type Literal struct {
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
