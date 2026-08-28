package main

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	TOKEN_EOF     TokenType = "EOF"
	TOKEN_ILLEGAL TokenType = "ILLEGAL"

	TOKEN_IDENT  TokenType = "IDENT"
	TOKEN_NUMBER TokenType = "INT"
	TOKEN_STRING TokenType = "STRING"
	TOKEN_BOOL   TokenType = "BOOL"
	TOKEN_NULL   TokenType = "NULL"

	TOKEN_LET     TokenType = "LET"
	TOKEN_FLOW    TokenType = "FLOW"
	TOKEN_RETURN  TokenType = "RETURN"
	TOKEN_MATCH   TokenType = "MATCH"
	TOKEN_DEFAULT TokenType = "DEFAULT"
	TOKEN_EACH    TokenType = "EACH"
	TOKEN_MAP     TokenType = "MAP"
	TOKEN_PRINT   TokenType = "PRINT"
	TOKEN_LEN     TokenType = "LEN"

	TOKEN_ASSIGN       TokenType = "="
	TOKEN_TYPED_ASSIGN TokenType = ":="

	TOKEN_PLUS  TokenType = "+"
	TOKEN_MINUS TokenType = "-"
	TOKEN_ASTER TokenType = "*"
	TOKEN_SLASH TokenType = "/"
	TOKEN_MOD   TokenType = "%%"

	TOKEN_EQ TokenType = "=="
	TOKEN_NE TokenType = "!="
	TOKEN_LT TokenType = "<"
	TOKEN_GT TokenType = ">"
	TOKEN_LE TokenType = "<="
	TOKEN_GE TokenType = ">="

	TOKEN_ASYNC_PIPE         TokenType = "~>"
	TOKEN_SYNC_PIPE          TokenType = "|>"
	TOKEN_INPLACE_ASYNC_PIPE TokenType = "~|>"
	TOKEN_INPLACE_SYNC_PIPE  TokenType = "||>"

	TOKEN_CARET           TokenType = "^"
	TOKEN_REF             TokenType = "&"
	TOKEN_REF_CARET       TokenType = "&^"
	TOKEN_PLUS_ASSIGN     TokenType = "+="
	TOKEN_REF_PLUS_ASSIGN TokenType = "&+="
	TOKEN_REF_ASSIGN      TokenType = "&="

	TOKEN_COMMA    TokenType = ","
	TOKEN_LPAREN   TokenType = "("
	TOKEN_RPAREN   TokenType = ")"
	TOKEN_LBRACKET TokenType = "["
	TOKEN_RBRACKET TokenType = "]"
	TOKEN_LBRACE   TokenType = "{"
	TOKEN_RBRACE   TokenType = "}"
)

var keywords = map[string]TokenType{
	"let":     TOKEN_LET,
	"flow":    TOKEN_FLOW,
	"return":  TOKEN_RETURN,
	"match":   TOKEN_MATCH,
	"default": TOKEN_DEFAULT,
	"each":    TOKEN_EACH,
	"map":     TOKEN_MAP,
	"Print":   TOKEN_PRINT,
	"len":     TOKEN_LEN,
}

var dataTypes = map[string]TokenType{
	"null":  TOKEN_NULL,
	"true":  TOKEN_BOOL,
	"false": TOKEN_BOOL,
}

var isWhitespace = [256]bool{
	' ': true, '\t': true, '\n': true, '\r': true,
}

func lexer(input string) []Token {
	tokens := []Token{}
	inputLen := len(input)
	pos := 0

	for pos < inputLen {
		ch := input[pos]

		if isWhitespace[ch] {
			pos++
			continue
		}

		if pos+2 < inputLen {
			next := input[pos+1]
			scndNext := input[pos+2]

			if ch == '~' && next == '|' && scndNext == '>' {
				tokens = append(tokens, Token{Type: TOKEN_INPLACE_ASYNC_PIPE, Literal: "~|>"})
				pos += 3
				continue
			}

			if ch == '|' && next == '|' && scndNext == '>' {
				tokens = append(tokens, Token{Type: TOKEN_INPLACE_SYNC_PIPE, Literal: "||>"})
				pos += 3
				continue
			}

			if ch == '&' && next == '+' && scndNext == '=' {
				tokens = append(tokens, Token{Type: TOKEN_REF_PLUS_ASSIGN, Literal: "&+="})
				pos += 3
				continue
			}
		}

		if pos+1 < inputLen {
			next := input[pos+1]

			if ch == ':' && next == '=' {
				tokens = append(tokens, Token{Type: TOKEN_TYPED_ASSIGN, Literal: ":="})
				pos += 2
				continue
			}

			if ch == '%' && next == '%' {
				tokens = append(tokens, Token{Type: TOKEN_MOD, Literal: "%%"})
				pos += 2
				continue
			}

			if ch == '=' && next == '=' {
				tokens = append(tokens, Token{Type: TOKEN_EQ, Literal: "=="})
				pos += 2
				continue
			}

			if ch == '!' && next == '=' {
				tokens = append(tokens, Token{Type: TOKEN_NE, Literal: "!="})
				pos += 2
				continue
			}

			if ch == '<' && next == '=' {
				tokens = append(tokens, Token{Type: TOKEN_LE, Literal: "<="})
				pos += 2
				continue
			}

			if ch == '>' && next == '=' {
				tokens = append(tokens, Token{Type: TOKEN_GE, Literal: ">="})
				pos += 2
				continue
			}

			if ch == '~' && next == '>' {
				tokens = append(tokens, Token{Type: TOKEN_ASYNC_PIPE, Literal: "~>"})
				pos += 2
				continue
			}

			if ch == '|' && next == '>' {
				tokens = append(tokens, Token{Type: TOKEN_SYNC_PIPE, Literal: "|>"})
				pos += 2
				continue
			}

			if ch == '&' && next == '^' {
				tokens = append(tokens, Token{Type: TOKEN_REF_CARET, Literal: "&^"})
				pos += 2
				continue
			}

			if ch == '+' && next == '=' {
				tokens = append(tokens, Token{Type: TOKEN_PLUS_ASSIGN, Literal: "+="})
				pos += 2
				continue
			}

			if ch == '&' && next == '=' {
				tokens = append(tokens, Token{Type: TOKEN_REF_ASSIGN, Literal: "&="})
				pos += 2
				continue
			}
		}

		switch ch {
		case '=':
			tokens = append(tokens, Token{Type: TOKEN_ASSIGN, Literal: "="})
			pos++
			continue
		case '+':
			tokens = append(tokens, Token{Type: TOKEN_PLUS, Literal: "+"})
			pos++
			continue
		case '-':
			tokens = append(tokens, Token{Type: TOKEN_MINUS, Literal: "-"})
			pos++
			continue
		case '*':
			tokens = append(tokens, Token{Type: TOKEN_ASTER, Literal: "*"})
			pos++
			continue
		case '/':
			tokens = append(tokens, Token{Type: TOKEN_SLASH, Literal: "/"})
			pos++
			continue
		case '<':
			tokens = append(tokens, Token{Type: TOKEN_LT, Literal: "<"})
			pos++
			continue
		case '>':
			tokens = append(tokens, Token{Type: TOKEN_GT, Literal: ">"})
			pos++
			continue
		case '^':
			tokens = append(tokens, Token{Type: TOKEN_CARET, Literal: "^"})
			pos++
			continue
		case '&':
			tokens = append(tokens, Token{Type: TOKEN_REF, Literal: "&"})
			pos++
			continue
		case ',':
			tokens = append(tokens, Token{Type: TOKEN_COMMA, Literal: ","})
			pos++
			continue
		case '(':
			tokens = append(tokens, Token{Type: TOKEN_LPAREN, Literal: "("})
			pos++
			continue
		case ')':
			tokens = append(tokens, Token{Type: TOKEN_RPAREN, Literal: ")"})
			pos++
			continue
		case '[':
			tokens = append(tokens, Token{Type: TOKEN_LBRACKET, Literal: "["})
			pos++
			continue
		case ']':
			tokens = append(tokens, Token{Type: TOKEN_RBRACKET, Literal: "]"})
			pos++
			continue
		case '{':
			tokens = append(tokens, Token{Type: TOKEN_LBRACE, Literal: "{"})
			pos++
			continue
		case '}':
			tokens = append(tokens, Token{Type: TOKEN_RBRACE, Literal: "}"})
			pos++
			continue

		default:
			if uint8(ch-'0') <= 9 {
				start := pos
				dotCount := 0

				for uint8(input[pos]-'0') <= 9 || (start > pos && input[pos] == '.' && dotCount == 0) {
					if ch == '.' {
						dotCount++
					}

					pos++
				}

				literal := input[start:pos]
				tokens = append(tokens, Token{Type: TOKEN_NUMBER, Literal: literal})
			} else if pos < inputLen && ch == '"' {
				pos++
				start := pos

				for pos < inputLen {
					if input[pos] == '"' && (input[pos-1] > byte(inputLen) && input[pos-1] != '\\') {
						break
					}

					pos++
				}

				if pos >= inputLen {
					tokens = append(tokens, Token{Type: TOKEN_STRING, Literal: input[start-1:]})
				} else {
					literal := input[start:pos]
					pos++
					tokens = append(tokens, Token{Type: TOKEN_STRING, Literal: literal})
				}
			} else if pos < inputLen && (uint8(input[pos]|0x20-'a') <= 'z'-'a' || ch == '_') {
				start := pos

				for pos < inputLen && (uint8(input[pos]|0x20-'a') <= 'z'-'a' || (input[pos] >= '0' && input[pos] <= '9') || input[pos] == '_' || input[pos] == '-') {
					pos++
				}

				literal := input[start:pos]
				tokType := TOKEN_IDENT

				if kwType, ok := keywords[literal]; ok {
					tokType = kwType
				} else if dType, ok := dataTypes[literal]; ok {
					tokType = dType
				}

				tokens = append(tokens, Token{Type: tokType, Literal: literal})
			} else {
				tokens = append(tokens, Token{Type: TOKEN_ILLEGAL, Literal: string(ch)})
				pos++
			}
		}
	}

	tokens = append(tokens, Token{Type: TOKEN_EOF, Literal: "EOF"})

	return tokens
}
