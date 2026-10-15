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
		if s.match('=') {
			s.Tokens = append(s.Tokens, NewToken(BangEqual, "", "", s.line))
		} else {
			s.Tokens = append(s.Tokens, NewToken(Bang, "", "", s.line))
		}
	case '=':
		if s.match('=') {
			s.Tokens = append(s.Tokens, NewToken(EqualEqual, "", "", s.line))
		} else {
			s.Tokens = append(s.Tokens, NewToken(Equal, "", "", s.line))
		}
	case '>':
		if s.match('=') {
			s.Tokens = append(s.Tokens, NewToken(GreaterEqual, "", "", s.line))
		} else {
			s.Tokens = append(s.Tokens, NewToken(Greater, "", "", s.line))
		}
	case '<':
		if s.match('=') {
			s.Tokens = append(s.Tokens, NewToken(LessEqual, "", "", s.line))
		} else {
			s.Tokens = append(s.Tokens, NewToken(Less, "", "", s.line))
		}
	case '/':
		if s.peekchar() == '/' {
			for {
				cha := s.nextchar()
				if cha == '\n' || cha == 0 {
					break
				}
			}

			// this is a comment just forward until new line or eof
		} else {
			s.Tokens = append(s.Tokens, NewToken(Slash, "", "", s.line))
		}
	default:
		fmt.Println(fmt.Errorf("unexpected char %q", ch))
	}

	s.Tokens = append(s.Tokens, NewToken(EOF, "", "", s.line))

	return nil
}

func (s *Scanner) peekchar() rune {
	if s.current >= len(s.input) {
		return 0
	}

	return rune(s.input[s.current])
}

func (s *Scanner) match(r rune) bool {
	r2 := s.peekchar()
	if r2 == 0 || r2 != r {
		return false
	}

	s.current++

	return true
}

func (s *Scanner) nextchar() rune {
	if s.current >= len(s.input) {
		return 0
	}

	r := s.input[s.current]

	s.current++

	return rune(r)
}
