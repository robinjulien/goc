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
	CurlyBracketOpen  Type = iota // {
	CurlyBracketClose             // }

	SquareBracketOpen  // [
	SquareBracketClose // ]

	ParanthesisOpen  // (
	ParanthesisClose // )

	// Keywords
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

	// Arithmetic Operators
	Plus    // +
	Minus   // -
	Divide  // /
	Modulus // %

	// Bitwise Operators
	BitwiseAnd // &
	BitwiseOr  // |
	LeftShift  // <<
	RightShift // >>

	LowerThan        // <
	GreaterThan      // >
	LowerThanEqual   // <=
	GreaterThanEqual // >=
	Equals           // ==
	NotEquals        // !=

	// Logical Operators
	LogicalAnd // &&
	LogicalOr  // ||
	Not        // !

	// Assign Operators
	Assign        // =
	AssignPlus    // +=
	AssignMinus   // -=
	AssignTimes   // *=
	AssignDivide  // /=
	AssignModulus // %=
	AssignAnd     // &=
	AssignOr      // |=
	AssignXor     // ^=

	// Literals
	StringLiteral // "string"
	ASCIILiteral  // 'o'
	NumberLiteral // 123, 123.456

	// Punctuation
	Dot       // .
	Coma      // ,
	SemiColon // ;
	Colon     // :
	Ternary   // ?

	// Increment/Decrement operators
	Increment
	Decrement

	// Special operators
	Asterisk  // *
	Ampersand // &

	Identifier // var name, function name, struct member
	EOF        // End of file
	Invalid    // Invalid token
)
