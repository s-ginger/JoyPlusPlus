package parser

import "j++/internal/backend/types"

// AST 
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

// type AST struct {
// 	Nodes []Node
// }

type Nodes []Node

// Stmt 
type Stmt interface {
	stmt()
}

type Mod struct {
	ModName string
}

func (Mod) stmt() {}

type Define struct {
	Name string
	Arrity int
	Body Nodes
}

func (Define) stmt() {}

type Import struct {
	ImportStr string
}

type Stmts []Stmt
