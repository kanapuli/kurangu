package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// Token Type constants

	// Exceptionals
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers and Literals
	IDENT = "IDENT" // variables
	INT   = "INT"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {

	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
