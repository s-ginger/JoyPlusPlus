package parser


type Parser struct {
	builder *Builder
}

func NewParser() *Parser {
	return &Parser{
		builder: NewBuilder(),
	}
}

func (p *Parser) Parse() AST {
	ast := AST{} 

	return ast
}
