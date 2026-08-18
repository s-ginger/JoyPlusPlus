package lexer

type TokenType int

const (
	SYMBOL TokenType = iota
	INT
	FLOAT

	LPAREN 
	RPAREN
)


type Token struct {
	Type TokenType
	Value string
	Line int
	Column int
}


