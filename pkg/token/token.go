package token

type TokenType int

type Token struct {
	Type     TokenType
	File     string
	Line     int
	Column   int
	Litteral string
}

const (
	CurlyBracketOpen TokenType = iota
	CurlyBracketClose

	SquareBracketOpen
	SquareBracketClose

	ParanthesisOpen
	ParanthesisClose

	Return

	Plus
	Minus
	Asterisk
	Divide

	BitwiseAnd
	BitwiseOr

	LowerThan
	GreaterThan
	LowerThanEqual
	GreaterThanEqual
	EqualsTo
	NotEqualsTo

	LogicalAnd
	LogicalOr

	Equals

	Not

	StringLitteral
	ASCIILitteral

	Coma
	SemiColon

	Increment
	Decrement

	Identifier
	EOF
	Invalid
)
