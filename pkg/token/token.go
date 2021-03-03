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
	TypeDef
	Struct
	Union
	Register
	Static

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
	Equals
	NotEquals

	LogicalAnd
	LogicalOr

	Assign

	Not

	StringLitteral
	ASCIILitteral
	NumberConstant

	Coma
	SemiColon

	Increment
	Decrement

	Identifier
	EOF
	Invalid
)
