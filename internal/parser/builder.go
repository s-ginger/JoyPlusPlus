package parser

import "j++/internal/backend/types"

type Builder struct {
	nodes []Node
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) Literal(value_type types.ValueType, value any) {
	b.nodes = append(b.nodes, Literal{
		Type: value_type,
		Value: value,
	})
}

func (b *Builder) Symbol(value string) {
	b.nodes = append(b.nodes, Symbol{
		Value: value,
	})
}

func (b *Builder) Quotation(fn func(*Builder)) {
	child := NewBuilder()

	fn(child)

	b.nodes = append(b.nodes, Quotation{
		Nodes: child.nodes,
	})
}

func (b *Builder) Build() AST {
	return AST{
		Nodes: b.nodes,
	}
}