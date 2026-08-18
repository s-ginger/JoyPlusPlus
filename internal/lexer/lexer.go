package lexer

import (
	"unicode"
	"unicode/utf8"
)

type Lexer struct {
	src string
	pos int
}

func NewLexer(src string) *Lexer {
	return &Lexer{
		src: src,
		pos: 0,
	}
}

func (l *Lexer) readRune() (rune, int) {
	r, size := utf8.DecodeRuneInString(l.src[l.pos:])
	return r, size
}

func (l *Lexer) skipWhitespaces() {
	for l.pos < len(l.src) {
		r, size := l.readRune()

		if !unicode.IsSpace(r) {
			break
		}

		l.pos += size
	}
}



func (l *Lexer) Tokenize() []Token {
	var tokens []Token

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
			// identifier
			l.pos += size
			continue
		}

		// unknown character
		l.pos += size
	}

	return tokens
}



