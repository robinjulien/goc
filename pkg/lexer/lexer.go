package lexer

import (
	"bufio"
	"bytes"
	"fmt"
	"io"

	"github.com/robinjulien/goc/pkg/token"
)

// Lexer lexer
type Lexer struct {
	reader *bufio.Reader
}

// New creates a new lexer given an io.Reader
func New(r io.Reader) *Lexer {
	lexer := &Lexer{
		reader: bufio.NewReader(r),
	}

	return lexer
}

// ReadChar returns next char of a lexer
func (l *Lexer) ReadChar() byte {
	b, err := l.reader.ReadByte()

	if err != nil {
		return 0
	}

	return b
}

// PeekNextChar returns next char without advancing the reader
func (l *Lexer) PeekNextChar() byte {
	peekedBytes, err := l.reader.Peek(1)

	if err != nil {
		return 0
	}

	return peekedBytes[0]
}

// NewTokenLitteral returns a token given its type and its litteral
func (l *Lexer) NewTokenLitteral(typ token.TokenType, lit string) token.Token {
	t := l.NewToken(typ)
	t.Litteral = lit
	return t
}

// NewToken returns a token given its type
func (l *Lexer) NewToken(typ token.TokenType) token.Token {
	return token.Token{
		Type: typ,
	}
}

// NextToken returns the next token of the lexer
func (l *Lexer) NextToken() token.Token {
	c := l.ReadChar()

	fmt.Println(c)

	for IsIgnoredChar(c) {
		c = l.ReadChar()
		fmt.Println("on est la 1", c)
	}

	fmt.Println("on est la 2")

	switch c {
	case 0:
		return l.NewToken(token.EOF)
	case '*':
		fmt.Println("astrerixx")
		return l.NewToken(token.Asterisk)
	case '/':
		return l.NewToken(token.Divide)
	case '{':
		return l.NewToken(token.CurlyBracketOpen)
	case '}':
		return l.NewToken(token.CurlyBracketClose)
	case '[':
		return l.NewToken(token.SquareBracketOpen)
	case ']':
		return l.NewToken(token.SquareBracketClose)
	case '(':
		return l.NewToken(token.ParanthesisOpen)
	case ')':
		return l.NewToken(token.ParanthesisClose)
	case ',':
		return l.NewToken(token.Coma)
	case ';':
		return l.NewToken(token.SemiColon)
	case '+': // cases +, ++
		if l.PeekNextChar() == '+' {
			l.ReadChar()
			return l.NewToken(token.Increment)
		}
		return l.NewToken(token.Plus)
	case '-': // cases -, --
		if l.PeekNextChar() == '-' {
			l.ReadChar()
			return l.NewToken(token.Decrement)
		}
		return l.NewToken(token.Minus)
	case '=': // cases =, ==
		if l.PeekNextChar() == '=' {
			l.ReadChar()
			return l.NewToken(token.EqualsTo)
		}
		return l.NewToken(token.Equals)
	case '<': // cases <, <=
		if l.PeekNextChar() == '=' {
			l.ReadChar()
			return l.NewToken(token.LowerThanEqual)
		}
		return l.NewToken(token.LowerThan)
	case '>': // cases >, >=
		if l.PeekNextChar() == '=' {
			l.ReadChar()
			return l.NewToken(token.GreaterThanEqual)
		}
		return l.NewToken(token.LowerThan)
	case '!': // cases !, !=
		if l.PeekNextChar() == '=' {
			l.ReadChar()
			return l.NewToken(token.NotEqualsTo)
		}
		return l.NewToken(token.Not)
	case '&': // cases &, &&
		if l.PeekNextChar() == '&' {
			l.ReadChar()
			return l.NewToken(token.LogicalAnd)
		}
		return l.NewToken(token.BitwiseAnd)
	case '|': // cases |, ||
		if l.PeekNextChar() == '+' {
			l.ReadChar()
			return l.NewToken(token.Increment)
		}
		return l.NewToken(token.BitwiseOr)
	default:
		return l.NewToken(token.Invalid)
	}
}

// GetWord returns the next word
func (l *Lexer) GetWord(firstChar byte) string {
	buf := bytes.Buffer{}
	buf.WriteByte(firstChar)

	for peekedBytes, err := l.reader.Peek(1); err == nil && IsAlphaNumerical(peekedBytes[0]); {
		buf.WriteByte(peekedBytes[0])
	}

	return buf.String()
}

// IsIgnoredChar returns if the given b char is an ignored char or not
func IsIgnoredChar(b byte) bool {
	switch b {
	case ' ', '\t', '\n', '\r':
		return true
	default:
		return false
	}
}

// IsAlphaNumerical returns if the given b char is a letter or a number
func IsAlphaNumerical(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9')
}

// IsLetter returns if the given b char is a letter
func IsLetter(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

// IsNumber returns if the given b char is a letter
func IsNumber(b byte) bool {
	return (b >= '0' && b <= '9')
}
