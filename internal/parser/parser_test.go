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

	nodes := p.Parse()

	if len(nodes) != 1 {
		t.Fatalf("expected 1 node, got %d", len(nodes))
	}

	literal, ok := nodes[0].(Literal)
	if !ok {
		t.Fatalf("expected Literal, got %T", nodes[0])
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

	nodes := p.Parse()

	if len(nodes) != 1 {
		t.Fatalf("expected 1 node, got %d", len(nodes))
	}

	literal, ok := nodes[0].(Literal)
	if !ok {
		t.Fatalf("expected Literal, got %T", nodes[0])
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

	nodes := p.Parse()

	if len(nodes) != 1 {
		t.Fatalf("expected 1 node, got %d", len(nodes))
	}

	symbol, ok := nodes[0].(Symbol)
	if !ok {
		t.Fatalf("expected Symbol, got %T", nodes[0])
	}

	if symbol.Value != "dup" {
		t.Errorf(
			"expected symbol %q, got %q",
			"dup",
			symbol.Value,
		)
	}
}

func TestParser_ParseMultipleNodes(t *testing.T) {
	tokens := []lexer.Token{
		{Type: lexer.INT, Value: "42"},
		{Type: lexer.SYMBOL, Value: "dup"},
		{Type: lexer.FLOAT, Value: "3.14"},
	}

	p := NewParser(tokens)

	nodes := p.Parse()

	if len(nodes) != 3 {
		t.Fatalf("expected 3 nodes, got %d", len(nodes))
	}

	literal, ok := nodes[0].(Literal)
	if !ok {
		t.Fatalf("expected first node Literal, got %T", nodes[0])
	}

	if literal.Value != int64(42) {
		t.Errorf("expected 42, got %v", literal.Value)
	}

	symbol, ok := nodes[1].(Symbol)
	if !ok {
		t.Fatalf("expected second node Symbol, got %T", nodes[1])
	}

	if symbol.Value != "dup" {
		t.Errorf("expected dup, got %q", symbol.Value)
	}

	floatLiteral, ok := nodes[2].(Literal)
	if !ok {
		t.Fatalf("expected third node Literal, got %T", nodes[2])
	}

	if floatLiteral.Value != float64(3.14) {
		t.Errorf("expected 3.14, got %v", floatLiteral.Value)
	}
}

func TestParser_Quotation(t *testing.T) {
	tokens := []lexer.Token{
		{Type: lexer.INT, Value: "42"},
		{Type: lexer.LBRACKET, Value: "["},
		{Type: lexer.SYMBOL, Value: "dup"},
		{Type: lexer.FLOAT, Value: "3.14"},
		{Type: lexer.RBRACKET, Value: "]"},
	}

	p := NewParser(tokens)

	nodes := p.Parse()

	if len(nodes) != 2 {
		t.Fatalf("expected 2 root nodes, got %d", len(nodes))
	}

	// 42
	literal, ok := nodes[0].(Literal)
	if !ok {
		t.Fatalf("expected first node Literal, got %T", nodes[0])
	}

	if literal.Value != int64(42) {
		t.Errorf("expected value 42, got %v", literal.Value)
	}

	if literal.Type != types.Int64 {
		t.Errorf("expected type Int64, got %v", literal.Type)
	}

	// [ ... ]
	quotation, ok := nodes[1].(Quotation)
	if !ok {
		t.Fatalf("expected second node Quotation, got %T", nodes[1])
	}

	if len(quotation.Nodes) != 2 {
		t.Fatalf(
			"expected quotation to contain 2 nodes, got %d",
			len(quotation.Nodes),
		)
	}

	// dup
	symbol, ok := quotation.Nodes[0].(Symbol)
	if !ok {
		t.Fatalf(
			"expected first quotation node Symbol, got %T",
			quotation.Nodes[0],
		)
	}

	if symbol.Value != "dup" {
		t.Errorf(
			"expected symbol %q, got %q",
			"dup",
			symbol.Value,
		)
	}

	// 3.14
	floatLiteral, ok := quotation.Nodes[1].(Literal)
	if !ok {
		t.Fatalf(
			"expected second quotation node Literal, got %T",
			quotation.Nodes[1],
		)
	}

	if floatLiteral.Value != float64(3.14) {
		t.Errorf(
			"expected value 3.14, got %v",
			floatLiteral.Value,
		)
	}

	if floatLiteral.Type != types.Float64 {
		t.Errorf(
			"expected type Float64, got %v",
			floatLiteral.Type,
		)
	}
}
