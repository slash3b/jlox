package main

type TokenType int

const (
	// single char token
	LeftParen TokenType = iota
	RightParen
	LeftBrace
	RightBrace
	Comma
	Dot
	Minus
	Plus
	Semicolon
	Slash
	Star

	Bang
	BangEqual
	Equal
	EqualEqual
	Greater
	GreaterEqual
	Less
	LessEqual

	// literals
	Identifier
	String
	Number

	// keywords
	Class
	And
	Else
	False
	Fun
	For
	If
	Nil
	Or
	Print
	Return
	Super
	This
	True
	Var
	While

	EOF
)

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal string
	Line    int
}

func NewToken(t TokenType, lex, lit string, line int) Token {
	return Token{
		Type:    t,
		Lexeme:  lex,
		Literal: lit,
		Line:    line,
	}
}
