package token

type TokenType string
type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	// Identifiers + literals
	IDENT = "IDENT" // add, foo , bar , x, y
	INT   = "INT"
	//Operators
	ASSIGN = "="
	PLUS   = "+"
	//Delimters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	//keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIden(ident string) TokenType {
	//This means that if the keywords(ident) is found in keywords map, it returns tok with true(ok).
	// tok, ok := keywords[ident]
	// if ok {
	// 	return tok
	// }
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
