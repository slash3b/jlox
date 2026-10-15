package main

import "fmt"

type Scanner struct {
	input   string
	Tokens  []Token
	start   int
	current int
	line    int
}

func NewScanner(s string) *Scanner {
	return &Scanner{
		input:  s,
		Tokens: make([]Token, 0),
	}
}

// func NewToken(t TokenType, lex string, lit string, line int ) Token {
func (s *Scanner) addToken(tt TokenType) {
	lit := s.input[s.start:s.current]

	// why do I need lex while we have a lit already?
	t := NewToken(tt, "", lit, s.line)

	s.Tokens = append(s.Tokens, t)
}

func (s *Scanner) Scan() error {
	ch := s.nextchar()
	switch ch {
	case '(':
		s.Tokens = append(s.Tokens, NewToken(LeftParen, "", "", s.line))
	case ')':
		s.Tokens = append(s.Tokens, NewToken(RightParen, "", "", s.line))
	case '{':
		s.Tokens = append(s.Tokens, NewToken(LeftBrace, "", "", s.line))
	case '}':
		s.Tokens = append(s.Tokens, NewToken(RightBrace, "", "", s.line))
	case ',':
		s.Tokens = append(s.Tokens, NewToken(Comma, "", "", s.line))
	case '.':
		s.Tokens = append(s.Tokens, NewToken(Dot, "", "", s.line))
	case '-':
		s.Tokens = append(s.Tokens, NewToken(Minus, "", "", s.line))
	case '+':
		s.Tokens = append(s.Tokens, NewToken(Plus, "", "", s.line))
	case ';':
		s.Tokens = append(s.Tokens, NewToken(Semicolon, "", "", s.line))
	case '*':
		s.Tokens = append(s.Tokens, NewToken(Star, "", "", s.line))
	// operators
	case '!':
		s.Tokens = append(s.Tokens, NewToken(Bang, "", "", s.line))
	case '=':
		s.Tokens = append(s.Tokens, NewToken(Equal, "", "", s.line))
	case '>':
		s.Tokens = append(s.Tokens, NewToken(Greater, "", "", s.line))
	case '<':
		s.Tokens = append(s.Tokens, NewToken(Less, "", "", s.line))
	default:
		fmt.Println(fmt.Errorf("unexpected char %q", ch))
	}

	s.Tokens = append(s.Tokens, NewToken(EOF, "", "", s.line))

	return nil
}

func (s *Scanner) peekchar() rune {
	if s.current+1 >= len(s.input) {
		return 0
	}

	return rune(s.input[s.current+1])
}

func (s *Scanner) match(r rune) bool {
	r2 := s.peekchar()

	return r2 != 0 && rune(r2) == r
}

func (s *Scanner) nextchar() rune {
	if s.current >= len(s.input) {
		return 0
	}

	r := s.input[s.current]

	s.current++

	return rune(r)
}
