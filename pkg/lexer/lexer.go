package lexer

import (
	"bufio"
	"bytes"
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
func (l *Lexer) NewTokenLitteral(typ token.Type, lit string) token.Token {
	t := l.NewToken(typ)
	t.Litteral = lit
	return t
}

// NewToken returns a token given its type
func (l *Lexer) NewToken(typ token.Type) token.Token {
	return token.Token{
		Type: typ,
	}
}

// NextToken returns the next token of the lexer
func (l *Lexer) NextToken() token.Token {
	c := l.ReadChar()

	for IsIgnoredChar(c) {
		c = l.ReadChar()
	}

	switch c {
	case 0:
		return l.NewToken(token.EOF)
	case '*':
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
			return l.NewToken(token.Equals)
		}
		return l.NewToken(token.Assign)
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
		return l.NewToken(token.GreaterThan)
	case '!': // cases !, !=
		if l.PeekNextChar() == '=' {
			l.ReadChar()
			return l.NewToken(token.NotEquals)
		}
		return l.NewToken(token.Not)
	case '&': // cases &, &&
		if l.PeekNextChar() == '&' {
			l.ReadChar()
			return l.NewToken(token.LogicalAnd)
		}
		return l.NewToken(token.BitwiseAnd)
	case '|': // cases |, ||
		if l.PeekNextChar() == '|' {
			l.ReadChar()
			return l.NewToken(token.LogicalOr)
		}
		return l.NewToken(token.BitwiseOr)
	default:
		if IsLetterExtended(c) {
			word := l.GetWord(c)

			switch word {
			case "auto":
				return l.NewToken(token.Auto)
			case "break":
				return l.NewToken(token.Break)
			case "case":
				return l.NewToken(token.Case)
			case "char":
				return l.NewToken(token.Char)
			case "const":
				return l.NewToken(token.Const)
			case "continue":
				return l.NewToken(token.Continue)
			case "default":
				return l.NewToken(token.Default)
			case "do":
				return l.NewToken(token.Do)
			case "double":
				return l.NewToken(token.Double)
			case "else":
				return l.NewToken(token.Else)
			case "enum":
				return l.NewToken(token.Enum)
			case "extern":
				return l.NewToken(token.Extern)
			case "float":
				return l.NewToken(token.Float)
			case "for":
				return l.NewToken(token.For)
			case "goto":
				return l.NewToken(token.Goto)
			case "if":
				return l.NewToken(token.If)
			case "int":
				return l.NewToken(token.Int)
			case "long":
				return l.NewToken(token.Long)
			case "register":
				return l.NewToken(token.Register)
			case "return":
				return l.NewToken(token.Return)
			case "short":
				return l.NewToken(token.Short)
			case "signed":
				return l.NewToken(token.Signed)
			case "sizeof":
				return l.NewToken(token.Sizeof)
			case "static":
				return l.NewToken(token.Static)
			case "struct":
				return l.NewToken(token.Struct)
			case "switch":
				return l.NewToken(token.Switch)
			case "typedef":
				return l.NewToken(token.Typedef)
			case "union":
				return l.NewToken(token.Union)
			case "unsigned":
				return l.NewToken(token.Unsigned)
			case "void":
				return l.NewToken(token.Void)
			case "volatile":
				return l.NewToken(token.Volatile)
			case "while":
				return l.NewToken(token.While)
			default:
				return l.NewTokenLitteral(token.Identifier, word)
			}
		} else if IsNumberExtended(c) {
			number := l.GetNumber(c)
			return l.NewTokenLitteral(token.NumberLiteral, number)
		}
		return l.NewToken(token.Invalid)
	}
}

// GetWord returns the next word
func (l *Lexer) GetWord(firstChar byte) string {
	buf := bytes.Buffer{}
	buf.WriteByte(firstChar)

	for {
		peekedBytes, err := l.reader.Peek(1)

		if err != nil || !IsAlphaNumerical(peekedBytes[0]) {
			break
		}

		c := l.ReadChar()
		buf.WriteByte(c)
	}

	return buf.String()
}

// GetNumber returns the next number (float or int)
func (l *Lexer) GetNumber(firstChar byte) string {
	buf := bytes.Buffer{}
	buf.WriteByte(firstChar)

	for {
		peekedBytes, err := l.reader.Peek(1)

		if err != nil || (!IsNumberExtended(peekedBytes[0])) {
			break
		}

		c := l.ReadChar()
		buf.WriteByte(c)
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
	return IsLetterExtended(b) || IsNumber(b)
}

// IsLetter returns if the given b char is a letter
func IsLetter(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

// IsLetterExtended returns if the given b char is a letter or an underscore
func IsLetterExtended(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || b == '_'
}

// IsNumber returns if the given b char is a letter
func IsNumber(b byte) bool {
	return (b >= '0' && b <= '9')
}

// IsNumberExtended returns if the given b char is a letter or a dot
func IsNumberExtended(b byte) bool {
	return (b >= '0' && b <= '9') || b == '.'
}
