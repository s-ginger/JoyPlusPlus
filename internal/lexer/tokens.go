package lexer

type TokenType int

const (
	SYMBOL TokenType = iota // 'anything!-><?+~

	STRING // "character"
	RUNE   // "c

	INT   //2
	FLOAT //2.5

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
