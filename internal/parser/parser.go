package parser

import "j++/internal/lexer"

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

	return ast
}
