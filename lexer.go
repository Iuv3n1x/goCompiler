package main

type TokenType string

type Position struct {
	Offset int
	Line   int
	Column int
}

type Token struct {
	Type     TokenType
	Literal  string
	Position Position
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
				pos += 2
				continue
			}

			if ch == '|' && next == '|' && scndNext == '>' {
				tokens = append(tokens, Token{Type: TOKEN_INPLACE_SYNC_PIPE, Literal: "||>"})
				pos += 2
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
				tokens = append(tokens, Token{Type: TOKEN_ASYNC_PIPE, Literal: ":="})
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
				tokens = append(tokens, Token{Type: TOKEN_REF_ASSIGN, Literal: ":="})
				pos += 2
				continue
			}
		}

		switch ch {
		case '=':
			tokens = append(tokens, Token{Type: TOKEN_ASSIGN, Literal: "="})
		}
	}

	return tokens
}
