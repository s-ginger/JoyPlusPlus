package parser

import (
	"j++/internal/backend/types"
	"j++/internal/lexer"
)

type Parser struct {
	tokens  []lexer.Token
	builder *Builder
}

func NewParser(tokens []lexer.Token) *Parser {
	return &Parser{
		tokens: tokens,
		builder: NewBuilder(),
	}
}

func (p *Parser) Parse() AST {
	ast := AST{}

	for _, tok := range p.tokens {

		if tok.Type == lexer.SYMBOL {
			p.builder.Symbol(tok.Value)
		}

		if tok.Type == lexer.INT {
			p.builder.Literal(types.Int64, tok.Value)
		}
		
		if tok.Type == lexer.FLOAT {
			p.builder.Literal(types.Float64, tok.Value)
		}

		if tok.Type == lexer.LBRACKET {

		}

	}

	return ast
}
