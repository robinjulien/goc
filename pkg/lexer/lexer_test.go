package lexer

import (
	"strings"
	"testing"

	"github.com/robinjulien/goc/pkg/token"
)

const (
	prog1 = `int main(int argc, char** argv) {
	int t = 1 + 1;
	return 0;
}`
)

func TestNew(t *testing.T) {}

func TestReadChar(t *testing.T) {
	l := New(strings.NewReader(prog1))

	for i := 0; i < len(prog1); i++ {
		if res := l.ReadChar(); res != prog1[i] {
			t.Errorf("Expected %c got %c", prog1[i], res)
		}
	}

	if res := l.ReadChar(); res != 0 {
		t.Errorf("Expected %v got %c", 0, res)
	}
}

func TestPeekNextChar(t *testing.T) {
	l := New(strings.NewReader(prog1))
	var res byte

	res = l.PeekNextChar()

	if res != prog1[0] {
		t.Errorf("Expected %c got %c", prog1[0], res)
	}

	res = l.PeekNextChar()

	if res != prog1[0] {
		t.Errorf("Expected %c got %c", prog1[0], res)
	}

	for i := 0; i < len(prog1); i++ {
		l.ReadChar()
	}

	res = l.PeekNextChar()

	if res != 0 {
		t.Errorf("Expected %v got %c", 0, res)
	}
}

func TestNewTokenLitteral(t *testing.T) {
	l := New(strings.NewReader(prog1))

	tok := l.NewTokenLitteral(token.Identifier, "name")

	if tok.Type != token.Identifier || tok.Litteral != "name" {
		t.Errorf("Wrong TokenType or Litteral")
	}
}

func TestNewToken(t *testing.T) {
	l := New(strings.NewReader(prog1))

	tok := l.NewToken(token.Identifier)

	if tok.Type != token.Identifier {
		t.Errorf("Wrong TokenType or Litteral")
	}
}

func TestNextToken(t *testing.T) {
	l := New(strings.NewReader(`* 	-+==[]=<=>=<>!!=++--&&&|||(),;{}/@ var; 123.456;`))

	table := make([]token.Token, 0)

	for {
		tok := l.NextToken()
		table = append(table, tok)

		if tok.Type == token.EOF {
			break
		}
	}

	expected := []token.Type{
		token.Asterisk,
		token.Minus,
		token.Plus,
		token.Equals,
		token.SquareBracketOpen,
		token.SquareBracketClose,
		token.Assign,
		token.LowerThanEqual,
		token.GreaterThanEqual,
		token.LowerThan,
		token.GreaterThan,
		token.Not,
		token.NotEquals,
		token.Increment,
		token.Decrement,
		token.LogicalAnd,
		token.BitwiseAnd,
		token.LogicalOr,
		token.BitwiseOr,
		token.ParanthesisOpen,
		token.ParanthesisClose,
		token.Coma,
		token.SemiColon,
		token.CurlyBracketOpen,
		token.CurlyBracketClose,
		token.Divide,
		token.Invalid,
		token.Identifier,
		token.SemiColon,
		token.NumberLiteral,
		token.SemiColon,
		token.EOF,
	}

	for i, resToken := range table {
		if resToken.Type != expected[i] {
			t.Errorf("Expected %v got %v", expected[i], resToken.Type)
		}
	}
}

func TestKeywords(t *testing.T) {
	l := New(strings.NewReader(`auto break case char const continue default do
	double else enum extern float for goto if
	int long register return short signed sizeof static
	struct switch typedef union unsigned void volatile while`))

	expected := []token.Type{
		token.Auto,
		token.Break,
		token.Case,
		token.Char,
		token.Const,
		token.Continue,
		token.Default,
		token.Do,
		token.Double,
		token.Else,
		token.Enum,
		token.Extern,
		token.Float,
		token.For,
		token.Goto,
		token.If,
		token.Int,
		token.Long,
		token.Register,
		token.Return,
		token.Short,
		token.Signed,
		token.Sizeof,
		token.Static,
		token.Struct,
		token.Switch,
		token.Typedef,
		token.Union,
		token.Unsigned,
		token.Void,
		token.Volatile,
		token.While,
		token.EOF,
	}

	for i := 0; ; i++ {
		tok := l.NextToken()

		if tok.Type != expected[i] {
			t.Errorf("Expected %v got %v", expected[i], tok.Type)
		}

		if tok.Type == token.EOF {
			break
		}
	}
}

func TestGetWord(t *testing.T) {
	l := New(strings.NewReader("+=var/test--"))

	l.ReadChar()
	l.ReadChar()

	c := l.ReadChar()
	res := l.GetWord(c)

	if res != "var" {
		t.Errorf("Expected %s got %s", "var", res)
	}

	l.ReadChar()

	c = l.ReadChar()
	res = l.GetWord(c)

	if res != "test" {
		t.Errorf("Expected %s got %s", "test", res)
	}
}

func TestGetNumber(t *testing.T) {
	l := New(strings.NewReader("+=21/123.456--"))

	l.ReadChar()
	l.ReadChar()

	c := l.ReadChar()
	res := l.GetNumber(c)

	if res != "21" {
		t.Errorf("Expected %s got %s", "21", res)
	}

	l.ReadChar()

	c = l.ReadChar()
	res = l.GetNumber(c)

	if res != "123.456" {
		t.Errorf("Expected %s got %s", "123.456", res)
	}
}

func TestIsIgnoredChar(t *testing.T) {
	serie := map[byte]bool{
		'+': false,
		'f': false,
		'R': false,
		'6': false,
		' ': true,
		'_': false,
	}

	for arg, expected := range serie {
		if res := IsIgnoredChar(arg); res != expected {
			t.Errorf("Expected %v got %v", expected, res)
		}
	}
}

func TestIsAlphaNumerical(t *testing.T) {
	serie := map[byte]bool{
		'+': false,
		'f': true,
		'R': true,
		'6': true,
		' ': false,
		'_': true,
	}

	for arg, expected := range serie {
		if res := IsAlphaNumerical(arg); res != expected {
			t.Errorf("Expected %v for %c got %v", expected, arg, res)
		}
	}
}

func TestIsLetter(t *testing.T) {
	serie := map[byte]bool{
		'+': false,
		'f': true,
		'R': true,
		'6': false,
		' ': false,
		'_': false,
	}

	for arg, expected := range serie {
		if res := IsLetter(arg); res != expected {
			t.Errorf("Expected %v for %c got %v", expected, arg, res)
		}
	}
}

func TestIsLetterExtended(t *testing.T) {
	serie := map[byte]bool{
		'+': false,
		'f': true,
		'R': true,
		'6': false,
		' ': false,
		'_': true,
	}

	for arg, expected := range serie {
		if res := IsLetterExtended(arg); res != expected {
			t.Errorf("Expected %v for %c got %v", expected, arg, res)
		}
	}
}

func TestIsNumber(t *testing.T) {
	serie := map[byte]bool{
		'+': false,
		'f': false,
		'R': false,
		'6': true,
		' ': false,
		'_': false,
	}

	for arg, expected := range serie {
		if res := IsNumber(arg); res != expected {
			t.Errorf("Expected %v for %c got %v", expected, arg, res)
		}
	}
}
