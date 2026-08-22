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

type Nodes []Node

// Stmt 
type Stmt interface {
	stmt()
}

type Stmts []Stmt

type Mod struct {
	ModName string
	Stmts Stmts
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

func (Import) stmt() {}

