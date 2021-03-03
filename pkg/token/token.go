package token

// Type is token type
type Type int

// Token includes type of a token and all informations related to int
type Token struct {
	Type     Type
	File     string
	Line     int
	Column   int
	Litteral string
}

// Token types list
const (
	CurlyBracketOpen Type = iota
	CurlyBracketClose

	SquareBracketOpen
	SquareBracketClose

	ParanthesisOpen
	ParanthesisClose

	Auto
	Break
	Case
	Char
	Const
	Continue
	Default
	Do
	Double
	Else
	Enum
	Extern
	Float
	For
	Goto
	If
	Int
	Long
	Register
	Return
	Short
	Signed
	Sizeof
	Static
	Struct
	Switch
	Typedef
	Union
	Unsigned
	Void
	Volatile
	While

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

	StringLiteral
	ASCIILiteral
	NumberLiteral

	Coma
	SemiColon

	Increment
	Decrement

	Identifier
	EOF
	Invalid
)
