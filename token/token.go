package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"

	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN = "="
	PLUS   = "+"

	RBRACE = "}"
	LBRACE = "{"
	LPAREN = "("
	RPAREN = ")"

	//delimiters
	COMMA     = ","
	SEMICOLON = ";"

	//keywords
	LET      = "LET"
	FUNCTION = "FUNCTION"

	EOF = "EOF"
)

var keywords = map[string]TokenType{
	"let": LET,
	"fn":  FUNCTION,
}

func LookupIdent(word string) TokenType {
	if tok, ok := keywords[word]; ok {
		return tok
	}
	return IDENT
}
