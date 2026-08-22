package lexer

type TokenType int

const (
	SYMBOL TokenType = iota

	STRING
	CHAR

	INT
	FLOAT

	LBRACKET // [
	RBRACKET // ]

	MOD       // mod
	DEFINE    // define
	EQEQ      // ==
	SEMICOLON // ;
	IMPORT    // import
	SEPARATOR // /
)

type Token struct {
	Type   TokenType
	Value  string
	Line   int
	Column int
}
