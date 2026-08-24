package main

import "strings"

type TokenType string

type Token struct {
	Type  TokenType
	Value string
}

const (
	TOKEN_EOF     TokenType = "EOF"
	TOKEN_ILLEGAL TokenType = "ILLEGAL"

	TOKEN_LET     TokenType = "LET"
	TOKEN_FLOW    TokenType = "FLOW"
	TOKEN_RETURN  TokenType = "RETURN"
	TOKEN_MATCH   TokenType = "MATCH"
	TOKEN_DEFAULT TokenType = "DEFAULT"
	TOKEN_MAP     TokenType = "MAP"
	TOKEN_EACH    TokenType = "EACH"

	TOKEN_PRINT TokenType = "PRINT"
	TOKEN_LEN   TokenType = "LEN"

	TOKEN_SET  TokenType = "="
	TOKEN_TSET TokenType = ":="

	TOKEN_AF  TokenType = "~>"
	TOKEN_SF  TokenType = "|>"
	TOKEN_IPA TokenType = "~|>"
	TOKEN_IPS TokenType = "||>"

	TOKEN_AR  TokenType = "[^"
	TOKEN_ARE TokenType = "[&^"
	TOKEN_AA  TokenType = "[+="
	TOKEN_AAE TokenType = "[&+="
	TOKEN_AI  TokenType = "[&"

	TOKEN_PLUS   TokenType = "+"
	TOKEN_MINUS  TokenType = "-"
	TOKEN_TIMES  TokenType = "*"
	TOKEN_DIVIDE TokenType = "/"
	TOKEN_MODULO TokenType = "%%"

	TOKEN_EQ  TokenType = "=="
	TOKEN_NEQ TokenType = "!="
	TOKEN_LT  TokenType = "<"
	TOKEN_GT  TokenType = ">"
	TOKEN_LTE TokenType = "<="
	TOKEN_GTE TokenType = ">="

	TOKEN_LPAREN   TokenType = "("
	TOKEN_RPAREN   TokenType = ")"
	TOKEN_LBRACE   TokenType = "{"
	TOKEN_RBRACE   TokenType = "}"
	TOKEN_LBRACKET TokenType = "["
	TOKEN_RBRACKET TokenType = "]"
	TOKEN_COMMA    TokenType = ","

	TOKEN_STRING TokenType = "STRING"
	TOKEN_NUMBER TokenType = "NUMBER"
	TOKEN_NULL   TokenType = "NULL"
	TOKEN_BOOL   TokenType = "BOOL"
	TOKEN_IDENT  TokenType = "IDENT"
)

var symbolTokens = []Token{
	// 4-Character Tokens
	{Type: TOKEN_AAE, Value: "[&+="},

	// 3-Character Tokens
	{Type: TOKEN_IPA, Value: "~|>"},
	{Type: TOKEN_IPS, Value: "||>"},
	{Type: TOKEN_ARE, Value: "[&^"},
	{Type: TOKEN_AA, Value: "[+="},

	// 2-Character Tokens
	{Type: TOKEN_TSET, Value: ":="},
	{Type: TOKEN_AF, Value: "~>"},
	{Type: TOKEN_AR, Value: "[^"},
	{Type: TOKEN_SF, Value: "|>"},
	{Type: TOKEN_AI, Value: "[&"},
	{Type: TOKEN_MODULO, Value: "%%"},
	{Type: TOKEN_EQ, Value: "=="},
	{Type: TOKEN_NEQ, Value: "!="},
	{Type: TOKEN_LTE, Value: "<="},
	{Type: TOKEN_GTE, Value: ">="},
}

var keywords = map[string]TokenType{
	"let":     TOKEN_LET,
	"flow":    TOKEN_FLOW,
	"return":  TOKEN_RETURN,
	"match":   TOKEN_MATCH,
	"default": TOKEN_DEFAULT,
	"each":    TOKEN_EACH,
	"map":     TOKEN_MAP,
	"print":   TOKEN_PRINT,
	"len":     TOKEN_LEN,
}

var dataTypes = map[string]TokenType{
	"String": TOKEN_STRING,
	"Number": TOKEN_NUMBER,
	"null":   TOKEN_NULL,
	"true":   TOKEN_BOOL,
	"false":  TOKEN_BOOL,
}
var isWhitespace = [256]bool{
	' ': true, '\t': true, '\n': true, '\r': true,
}

func lexer(input string) []Token {
	tokens := []Token{}
	fileLength := len(input)
	pos := 0

	for pos < fileLength {
		ch := input[pos]

		if isWhitespace[ch] {
			pos++
			continue
		}

		if ch == 0 {
			break
		}

		// Multiple-Character Token Check
		if tok, consumed := matchSymbol(input, pos); consumed > 0 {
			tokens = append(tokens, tok)
			pos += consumed
			continue
		}

		// Single-Character Token Check
		switch ch {
		case '=':
			tokens = append(tokens, Token{Type: TOKEN_SET, Value: "="})
			pos++
		case '+':
			tokens = append(tokens, Token{Type: TOKEN_PLUS, Value: "+"})
			pos++
		case '-':
			tokens = append(tokens, Token{Type: TOKEN_MINUS, Value: "-"})
			pos++
		case '*':
			tokens = append(tokens, Token{Type: TOKEN_TIMES, Value: "*"})
			pos++
		case '/':
			tokens = append(tokens, Token{Type: TOKEN_DIVIDE, Value: "/"})
			pos++
		case '<':
			tokens = append(tokens, Token{Type: TOKEN_LT, Value: "<"})
			pos++
		case '>':
			tokens = append(tokens, Token{Type: TOKEN_GT, Value: ">"})
			pos++
		case '(':
			tokens = append(tokens, Token{Type: TOKEN_LPAREN, Value: "("})
			pos++
		case ')':
			tokens = append(tokens, Token{Type: TOKEN_RPAREN, Value: ")"})
			pos++
		case '{':
			tokens = append(tokens, Token{Type: TOKEN_LBRACE, Value: "{"})
			pos++
		case '}':
			tokens = append(tokens, Token{Type: TOKEN_RBRACE, Value: "}"})
			pos++
		case '[':
			tokens = append(tokens, Token{Type: TOKEN_LBRACKET, Value: "["})
			pos++
		case ']':
			tokens = append(tokens, Token{Type: TOKEN_RBRACKET, Value: "]"})
			pos++
		case ',':
			tokens = append(tokens, Token{Type: TOKEN_COMMA, Value: ","})
			pos++

		default:
			// Number Check
			if uint8(ch-'0') <= 9 {
				start := pos
				dotCount := 0

				for pos < fileLength && (uint8(input[pos]-'0') <= 9 || (input[pos] == '.' && dotCount == 0 && start != pos)) {
					if input[pos] == '.' {
						dotCount++
					}

					pos++
				}

				literal := input[start:pos]
				tokens = append(tokens, Token{Type: TOKEN_NUMBER, Value: literal})
			} else if ch == '"' { // String Token Check
				pos++
				start := pos

				for pos < fileLength {
					if input[pos] == '"' && input[pos-1] != '\\' {
						break
					}

					pos++
				}

				if pos >= fileLength {
					tokens = append(tokens, Token{Type: TOKEN_ILLEGAL, Value: input[start-1:]})
				} else {
					literal := input[start:pos]
					pos++
					tokens = append(tokens, Token{Type: TOKEN_STRING, Value: literal})
				}
			} else if pos < fileLength && (uint8(input[pos]|0x20-'a') <= 'z'-'a' || ch == '_') { // Keyword check
				start := pos

				for pos < fileLength && (uint8(input[pos]|0x20-'a') <= 'z'-'a' || uint8(input[pos]-'0') <= 9 && input[pos] == '_' || input[pos] == '-') {
					pos++
				}

				literal := input[start:pos]
				tokType := TOKEN_IDENT

				if kwType, ok := keywords[literal]; ok {
					tokType = kwType
				} else if dType, ok := dataTypes[literal]; ok {
					tokType = dType
				}

				tokens = append(tokens, Token{Type: tokType, Value: literal})
			} else {
				tokens = append(tokens, Token{Type: TOKEN_ILLEGAL, Value: string(ch)})
				pos++
			}
		}
	}

	return tokens
}

func matchSymbol(input string, pos int) (Token, int) {
	rest := input[pos:]

	for _, tok := range symbolTokens {
		if strings.HasPrefix(rest, tok.Value) {
			return tok, len(tok.Value)
		}
	}
	return Token{}, 0
}
