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

func NewLexer(src string) *Lexer {
	return &Lexer{
		src: src,
		pos: 0,
		line: 0,
		column: 0,
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

		r, size := l.readRune()

		if unicode.IsDigit(r) {
			// number
			l.pos += size
			continue
		}

		if unicode.IsLetter(r) {
			value := l.readSymbol()

			tokens = append(tokens, Token{
				Type:  SYMBOL,
				Value: value,
				Line: l.line,
				Column: l.column,
			})

			continue
		}

		// unknown character
		l.advance()
	}

	return tokens
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


