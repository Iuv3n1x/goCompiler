package main

type Node interface {
	node()
}

type Expression interface {
	Node
	expressionNode()
}

type LiteralExpression struct {
	Token Token
	Value string
}

func (LiteralExpression) node()           {}
func (LiteralExpression) expressionNode() {}

type IdentifierExpression struct {
	Token Token
	Name  string
}

func (IdentifierExpression) node()           {}
func (IdentifierExpression) expressionNode() {}

type BinaryExpression struct {
	Left     Expression
	Operator string
	Right    Expression
}

func (BinaryExpression) node()           {}
func (BinaryExpression) expressionNode() {}

type AST struct {
	statements []Node
}

type LetStatement struct {
	Declaration string
	TypeSave    bool
	Value       Expression
}

func (LetStatement) node() {}

// [{LET let} {IDENT status} {= =} {INT 200} {EOF EOF}]
