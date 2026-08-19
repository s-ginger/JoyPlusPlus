package parser

import (
	"j++/internal/backend/types"
	"j++/internal/lexer"
	"strconv"
)

type Parser struct {
	tokens  []lexer.Token
	builder *Builder
	pos     int
}

func NewParser(tokens []lexer.Token) *Parser {
	return &Parser{
		tokens:  tokens,
		builder: NewBuilder(),
	}
}

func NewZeroParser() *Parser {
	return &Parser{
		tokens: []lexer.Token{},
		builder:  NewBuilder(),
	}
}

func (p *Parser) SetTokens(tokens []lexer.Token) {
	p.tokens = tokens
}

func (p *Parser) Flush() {
	p.tokens = []lexer.Token{}
	p.builder = NewBuilder()
	p.pos = 0
}

func (p *Parser) current() lexer.Token {
	return p.tokens[p.pos]
}

func (p *Parser) peek() (lexer.Token, bool) {
	if p.pos+1 >= len(p.tokens) {
		return lexer.Token{}, false
	}
	
	return p.tokens[p.pos+1], true
}

func (p *Parser) advance() {
	if p.pos < len(p.tokens) {
		p.pos++
	}
}

func (p *Parser) parseInteger(tok lexer.Token) {
	value, err := strconv.ParseInt(tok.Value, 10, 64)
	if err != nil {
		// обработка ошибки
		return
	}

	p.builder.Literal(types.Int64, value)
}

func (p *Parser) parseFloat(tok lexer.Token) {
	value, err := strconv.ParseFloat(tok.Value, 64)
	if err != nil {
		// пока просто выходим
		return
	}

	p.builder.Literal(types.Float64, value)
}

func (p *Parser) parseQuotation() {
	p.advance() // пропускаем [

	p.builder.Quotation(func(q *Builder) {
		for p.pos < len(p.tokens) {
			tok := p.current()

			switch tok.Type {
			case lexer.RBRACKET:
				p.advance()
				return

			case lexer.SYMBOL:
				q.Symbol(tok.Value)
				p.advance()

			case lexer.INT:
				value, err := strconv.ParseInt(tok.Value, 10, 64)
				if err != nil {
					p.advance()
					continue
				}

				q.Literal(types.Int64, value)
				p.advance()

			case lexer.FLOAT:
				value, err := strconv.ParseFloat(tok.Value, 64)
				if err != nil {
					p.advance()
					continue
				}

				q.Literal(types.Float64, value)
				p.advance()

			case lexer.LBRACKET:
				// вложенная quotation
				// здесь нужен отдельный механизм
				p.parseQuotation()

			default:
				p.advance()
			}
		}
	})
}

func (p *Parser) Parse() AST {
	for p.pos < len(p.tokens) {
		tok := p.current()

		switch tok.Type {
		case lexer.SYMBOL:
			p.builder.Symbol(tok.Value)
			p.advance()

		case lexer.INT:
			p.parseInteger(tok)
			p.advance()

		case lexer.FLOAT:
			p.parseFloat(tok)
			p.advance()

		case lexer.LBRACKET:
			p.parseQuotation()

		default:
			p.advance()
		}
	}

	return p.builder.Build()
}