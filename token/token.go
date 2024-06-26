package token

const (
	ILLEGAL Type = "ILLEGAL"
	EOF     Type = "EOF"

	IDENT Type = "IDENT"
	INT   Type = "INT"

	ASSIGN Type = "="
	PLUS   Type = "+"

	COMMA     Type = ","
	SEMICOLON Type = ";"

	LPAREN Type = "("
	RPAREN Type = ")"
	LBRACE Type = "{"
	RBRACE Type = "}"

	FUNCTION Type = "FUNCTION"
	LET      Type = "LET"
)

type Type string

type Token struct {
	Type    Type
	Literal string
}
