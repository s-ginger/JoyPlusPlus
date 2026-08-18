package lexer

import (
	"reflect"
	"testing"
)

func TestLexer_Tokenize(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []Token
	}{
		{
			name:  "Whitespace",
			input: "\n\r\t ",
			want: []Token{},
		},
		{
			name: "Symbol",
			input: "print",
			want: []Token{
				{Type: SYMBOL, Value:"print", Line: 0, Column: 5},
			},
		},
		{
			name: "Integer",
			input: "123",
			want: []Token{
				{Type: INT, Value:"123", Line: 0, Column: 3},
			},
		},
		{
			name: "LPAREN RPAREN",
			input: "[ ]",
			want: []Token{
				{Type: LPAREN, Value:"[", Line: 0, Column: 0},
				{Type: RPAREN, Value:"]", Line: 0, Column: 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLexer(tt.input)

			got := l.Tokenize()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tokenize() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

