package parser

import (
	"testing"

	"j++/internal/backend/types"
	"j++/internal/lexer"
)

func TestParser_ParseInteger(t *testing.T) {
	tokens := []lexer.Token{
		{
			Type:  lexer.INT,
			Value: "42",
		},
	}

	p := NewParser(tokens)

	ast := p.Parse()

	if len(ast.Nodes) != 1 {
		t.Fatalf("expected 1 node, got %d", len(ast.Nodes))
	}

	literal, ok := ast.Nodes[0].(Literal)
	if !ok {
		t.Fatalf("expected Literal, got %T", ast.Nodes[0])
	}

	if literal.Type != types.Int64 {
		t.Errorf(
			"expected type %v, got %v",
			types.Int64,
			literal.Type,
		)
	}

	if literal.Value != int64(42) {
		t.Errorf(
			"expected value %v, got %v",
			int64(42),
			literal.Value,
		)
	}
}

func TestParser_ParseFloat(t *testing.T) {
	tokens := []lexer.Token{
		{
			Type:  lexer.FLOAT,
			Value: "3.14",
		},
	}

	p := NewParser(tokens)

	ast := p.Parse()

	if len(ast.Nodes) != 1 {
		t.Fatalf("expected 1 node, got %d", len(ast.Nodes))
	}

	literal, ok := ast.Nodes[0].(Literal)
	if !ok {
		t.Fatalf("expected Literal, got %T", ast.Nodes[0])
	}

	if literal.Type != types.Float64 {
		t.Errorf(
			"expected type %v, got %v",
			types.Float64,
			literal.Type,
		)
	}

	if literal.Value != float64(3.14) {
		t.Errorf(
			"expected value %v, got %v",
			float64(3.14),
			literal.Value,
		)
	}
}

func TestParser_ParseSymbol(t *testing.T) {
	tokens := []lexer.Token{
		{
			Type:  lexer.SYMBOL,
			Value: "dup",
		},
	}

	p := NewParser(tokens)

	ast := p.Parse()

	if len(ast.Nodes) != 1 {
		t.Fatalf("expected 1 node, got %d", len(ast.Nodes))
	}

	symbol, ok := ast.Nodes[0].(Symbol)
	if !ok {
		t.Fatalf("expected Symbol, got %T", ast.Nodes[0])
	}

	if symbol.Value != "dup" {
		t.Errorf(
			"expected symbol %q, got %q",
			"dup",
			symbol.Value,
		)
	}
}