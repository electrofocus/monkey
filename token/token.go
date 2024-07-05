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

var keywords = map[string]Type{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
