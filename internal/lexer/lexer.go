package lexer

import (
	"unicode"
	"unicode/utf8"
)

type Lexer struct {
	src string
	pos int

	line   int
	column int
}

func NewZeroLexer() *Lexer {
	return &Lexer{
		src: "",
		line:   1,
		column: 1,
	}
}

func (l *Lexer) SetSrc(src string) {
	l.src = src
}

func (l *Lexer) Flush() {
	l.src = ""
	l.pos = 0
	l.column = 1
	l.line = 1
}

func NewLexer(src string) *Lexer {
	return &Lexer{
		src:    src,
		line:   1,
		column: 1,
	}
}

func (l *Lexer) readRune() (rune, int) {
	r, size := utf8.DecodeRuneInString(l.src[l.pos:])
	return r, size
}

func (l *Lexer) advance() {
	r, size := l.readRune()

	l.pos += size

	if r == '\n' {
		l.line++
		l.column = 1
		return
	}

	l.column++
}

func (l *Lexer) skipWhitespaces() {
	for l.pos < len(l.src) {
		r, _ := l.readRune()

		if !unicode.IsSpace(r) {
			break
		}

		l.advance()
	}
}

func (l *Lexer) Tokenize() []Token {
	tokens := []Token{}

	for l.pos < len(l.src) {
		l.skipWhitespaces()

		if l.pos >= len(l.src) {
			break
		}

		r, _ := l.readRune()

		if r == '[' {
			tokens = append(tokens, Token{
				Type:   LBRACKET,
				Value:  string(r),
				Line:   l.line,
				Column: l.column,
			})
			l.advance()
			continue
		}

		if r == ']' {
			tokens = append(tokens, Token{
				Type:   RBRACKET,
				Value:  string(r),
				Line:   l.line,
				Column: l.column,
			})
			l.advance()
			continue
		}

		if unicode.IsDigit(r) {
			line := l.line
			column := l.column

			value, hasDot := l.readNumber()

			if hasDot {
				tokens = append(tokens, Token{
					Type:   FLOAT,
					Value:  value,
					Line:   line,
					Column: column,
				})
			} else {
				tokens = append(tokens, Token{
					Type:   INT,
					Value:  value,
					Line:   line,
					Column: column,
				})
			}

			continue
		}

		if unicode.IsLetter(r) {
			line := l.line
			column := l.column

			value := l.readSymbol()

			tokens = append(tokens, Token{
				Type:   SYMBOL,
				Value:  value,
				Line:   line,
				Column: column,
			})

			continue
		}

		// unknown character
		l.advance()
	}

	return tokens
}

func (l *Lexer) readNumber() (string, bool) {
	start := l.pos
	hasDot := false

	for l.pos < len(l.src) {
		r, _ := l.readRune()

		if unicode.IsDigit(r) {
			l.advance()
			continue
		}

		if r == '.' && !hasDot {
			hasDot = true
			l.advance()
			continue
		}

		break
	}

	return l.src[start:l.pos], hasDot
}

func (l *Lexer) readSymbol() string {
	start := l.pos

	for l.pos < len(l.src) {
		r, _ := l.readRune()

		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			break
		}

		l.advance()
	}

	return l.src[start:l.pos]
}
