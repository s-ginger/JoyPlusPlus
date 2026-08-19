package parser


import (
	"j++/internal/backend/types"
	"testing"
)

func TestBuilderLiteral(t *testing.T) {
	b := NewBuilder()

	b.Literal(types.Int32, 42)

	ast := b.Build()

	if len(ast.Nodes) != 1 {
		t.Fatalf("expected 1 node, got %d", len(ast.Nodes))
	}

	literal, ok := ast.Nodes[0].(Literal)
	if !ok {
		t.Fatalf("expected Literal, got %T", ast.Nodes[0])
	}

	if literal.Value != 42 {
		t.Errorf("expected value 42, got %v", literal.Value)
	}

	if literal.Type != types.Int32 {
		t.Errorf("expected type Int32, got %v", literal.Type)
	}
}

func TestBuilderSymbol(t *testing.T) {
	b := NewBuilder()

	b.Symbol("dup")

	ast := b.Build()

	if len(ast.Nodes) != 1 {
		t.Fatalf("expected 1 node, got %d", len(ast.Nodes))
	}

	symbol, ok := ast.Nodes[0].(Symbol)
	if !ok {
		t.Fatalf("expected Symbol, got %T", ast.Nodes[0])
	}

	if symbol.Value != "dup" {
		t.Errorf("expected symbol dup, got %q", symbol.Value)
	}
}

func TestBuilderQuotation(t *testing.T) {
	b := NewBuilder()

	b.Quotation(func(q *Builder) {
		q.Literal(types.Int32, 1)
		q.Symbol("+")
	})

	ast := b.Build()

	if len(ast.Nodes) != 1 {
		t.Fatalf("expected 1 node, got %d", len(ast.Nodes))
	}

	quotation, ok := ast.Nodes[0].(Quotation)
	if !ok {
		t.Fatalf("expected Quotation, got %T", ast.Nodes[0])
	}

	if len(quotation.Nodes) != 2 {
		t.Fatalf(
			"expected quotation to contain 2 nodes, got %d",
			len(quotation.Nodes),
		)
	}

	if literal, ok := quotation.Nodes[0].(Literal); !ok {
		t.Errorf("expected first node to be Literal, got %T", quotation.Nodes[0])
	} else if literal.Value != 1 {
		t.Errorf("expected value 1, got %v", literal.Value)
	}

	if symbol, ok := quotation.Nodes[1].(Symbol); !ok {
		t.Errorf("expected second node to be Symbol, got %T", quotation.Nodes[1])
	} else if symbol.Value != "+" {
		t.Errorf("expected symbol +, got %q", symbol.Value)
	}
}

func TestBuilderNestedQuotation(t *testing.T) {
	b := NewBuilder()

	b.Quotation(func(q *Builder) {
		q.Literal(types.Int32, 10)

		q.Quotation(func(q *Builder) {
			q.Literal(types.Int32, 20)
			q.Symbol("+")
		})

		q.Symbol("call")
	})

	ast := b.Build()

	if len(ast.Nodes) != 1 {
		t.Fatalf("expected 1 root node, got %d", len(ast.Nodes))
	}

	outer, ok := ast.Nodes[0].(Quotation)
	if !ok {
		t.Fatalf("expected outer Quotation, got %T", ast.Nodes[0])
	}

	if len(outer.Nodes) != 3 {
		t.Fatalf("expected 3 nodes in outer quotation, got %d", len(outer.Nodes))
	}

	inner, ok := outer.Nodes[1].(Quotation)
	if !ok {
		t.Fatalf("expected inner Quotation, got %T", outer.Nodes[1])
	}

	if len(inner.Nodes) != 2 {
		t.Fatalf("expected 2 nodes in inner quotation, got %d", len(inner.Nodes))
	}

	if literal, ok := inner.Nodes[0].(Literal); !ok {
		t.Errorf("expected Literal, got %T", inner.Nodes[0])
	} else if literal.Value != 20 {
		t.Errorf("expected 20, got %v", literal.Value)
	}

	if symbol, ok := inner.Nodes[1].(Symbol); !ok {
		t.Errorf("expected Symbol, got %T", inner.Nodes[1])
	} else if symbol.Value != "+" {
		t.Errorf("expected +, got %q", symbol.Value)
	}
}