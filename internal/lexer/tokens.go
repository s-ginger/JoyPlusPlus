package lexer

type TokenType int

const (
	SYMBOL TokenType = iota
	INT
	FLOAT

	LBRACKET 
	RBRACKET
)


type Token struct {
	Type TokenType
	Value string
	Line int
	Column int
}


