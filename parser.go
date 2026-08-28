package main

import (
	"fmt"
	"os"
)

func parser(tokens []Token) AST {
	ast := AST{}
	pos := 0

	for tokens[pos].Type != TOKEN_EOF {
		tok := tokens[pos]

		switch tok.Type {
		case TOKEN_LET:
			stmt := parseLetStatement(tokens, &pos)
			ast.statements = append(ast.statements, stmt)
		}

		pos++
	}

	return ast
}

func parseLetStatement(tokens []Token, pos *int) LetStatement {
	*pos++

	if tokens[*pos].Type != TOKEN_IDENT {
		fmt.Println("Expected Variable Name after 'let' keyword")
		os.Exit(66)
	}

	declaration := tokens[*pos].Literal
	*pos++

	if tokens[*pos].Type != TOKEN_ASSIGN && tokens[*pos].Type != TOKEN_TYPED_ASSIGN {
		return LetStatement{
			Declaration: declaration,
			TypeSave:    false,
			Value: LiteralExpression{
				Token: Token{
					Type:    TOKEN_NULL,
					Literal: "null",
				},
				Value: "null",
			},
		}
	}

	var typeSave bool
	if tokens[*pos].Type == TOKEN_ASSIGN {
		typeSave = false
	} else {
		typeSave = true
	}
	*pos++

	value := parseExpression(tokens, pos)

	return LetStatement{
		Declaration: declaration,
		TypeSave:    typeSave,
		Value:       value,
	}
}

func parseExpression(tokens []Token, pos *int) Expression {
	tok := tokens[*pos]

	switch tok.Type {
	case TOKEN_NUMBER, TOKEN_STRING, TOKEN_BOOL, TOKEN_NULL:
		return LiteralExpression{
			Token: tok,
			Value: tok.Literal,
		}
	case TOKEN_IDENT:
		return IdentifierExpression{
			Token: tok,
			Name:  tok.Literal,
		}
	}

	fmt.Printf("Unexpected token in expression: %v\n", tok)
	os.Exit(66)
	return nil
}
