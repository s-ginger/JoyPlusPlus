package parser

type Builder struct {
	nodes []Node
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) Literal(value any) {
	b.nodes = append(b.nodes, Literal{
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