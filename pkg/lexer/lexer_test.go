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
	l := New(strings.NewReader("* 	-+==[]"))

	table := make([]token.Token, 0, len("* 	-+=[]")+1)

	for {
		tok := l.NextToken()
		table = append(table, tok)

		if tok.Type == token.EOF {
			break
		}
	}

	expected := []token.TokenType{
		token.Asterisk,
		token.Minus,
		token.Plus,
		token.EqualsTo,
		token.SquareBracketOpen,
		token.SquareBracketClose,
		token.EOF,
	}

	for i, resToken := range table {
		if resToken.Type != expected[i] {
			t.Errorf("Expected %v got %v", expected[i], resToken.Type)
		}
	}
}

func TestIsIgnoredChar(t *testing.T) {
	serie := map[byte]bool{
		'+': false,
		'f': false,
		'R': false,
		'6': false,
		' ': true,
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
	}

	for arg, expected := range serie {
		if res := IsLetter(arg); res != expected {
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
	}

	for arg, expected := range serie {
		if res := IsNumber(arg); res != expected {
			t.Errorf("Expected %v for %c got %v", expected, arg, res)
		}
	}
}
