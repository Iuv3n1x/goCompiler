package main

// ==========================================
// 1. Grundlegende Interfaces
// ==========================================

// Jeder Knoten im AST muss TokenLiteral() implementieren.
type Node interface {
	TokenLiteral() string
}

// Jede Anweisung (z. B. let x = 5) implementiert Statement.
type Statement interface {
	Node
	statementNode()
}

// Jeder Ausdruck, der einen Wert liefert (z. B. 5 + 5, x, match { ... }), implementiert Expression.
type Expression interface {
	Node
	expressionNode()
}

// Wurzelknoten des gesamten Programms
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// ==========================================
// 2. Literale & Basis-Ausdrücke (Expressions)
// ==========================================

// Bezeichner/Variablennamen (z. B. status, exampleVariable)
type Identifier struct {
	Token Token  // TOKEN_IDENT
	Value string // "status"
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Value }

// Zahlen-Literal (z. B. 200, 3.14)
type NumberLiteral struct {
	Token Token
	Value float64
}

func (nl *NumberLiteral) expressionNode()      {}
func (nl *NumberLiteral) TokenLiteral() string { return nl.Token.Value }

// String-Literal (z. B. "Hallo Welt")
type StringLiteral struct {
	Token Token
	Value string
}

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Value }

// Boolean-Literal (true / false)
type BooleanLiteral struct {
	Token Token
	Value bool
}

func (bl *BooleanLiteral) expressionNode()      {}
func (bl *BooleanLiteral) TokenLiteral() string { return bl.Token.Value }

// Array-Literal (z. B. [1, "Hallo", null])
type ArrayLiteral struct {
	Token    Token        // Das '[' Token
	Elements []Expression // Die Elemente im Array
}

func (al *ArrayLiteral) expressionNode()      {}
func (al *ArrayLiteral) TokenLiteral() string { return al.Token.Value }

// ExpressionStatement kapselt eine Expression, die alleine als Anweisung steht (z. B. Pipeline-Flows)
type ExpressionStatement struct {
	Token      Token // Das erste Token der Expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Value }

// ==========================================
// 3. Statements (Anweisungen)
// ==========================================

// Variable Deklaration: let x = 8 oder let x := [1, 2, 3]
type LetStatement struct {
	Token   Token       // Das 'let' Token
	Name    *Identifier // Name der Variable
	IsTyped bool        // true wenn ':=', false wenn '='
	Value   Expression  // Wert der zugewiesen wird (kann nil sein bei "let x")
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Value }

// Return Anweisung: return result oder return
type ReturnStatement struct {
	Token       Token      // Das 'return' Token
	ReturnValue Expression // Kann nil sein bei leerem return
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Value }

// Block von Anweisungen (z. B. innerhalb von flow { ... })
type BlockStatement struct {
	Token      Token // Das '{' Token
	Statements []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Value }

// ==========================================
// 4. Pipelines & Operatoren (~>, |>, ~|>, ||>)
// ==========================================

// Repräsentiert Pipeline-Ketten: status ~> processSuccess() ~> print
type PipelineExpression struct {
	Token    Token      // Das Pipeline-Token (~>, |>, ~|>, ||>)
	Left     Expression // Was links vom Operator steht
	Operator string     // "~>", "|>", "~|>", "||>"
	Right    Expression // Was rechts vom Operator steht
}

func (pe *PipelineExpression) expressionNode()      {}
func (pe *PipelineExpression) TokenLiteral() string { return pe.Token.Value }

// Parallele Ausführung in Pipelines: [function1(), function2()] ~> variable
type ParallelGroupExpression struct {
	Token       Token        // Das '[' Token
	Expressions []Expression // Aufrufe, die zeitgleich ausgeführt werden
}

func (pg *ParallelGroupExpression) expressionNode()      {}
func (pg *ParallelGroupExpression) TokenLiteral() string { return pg.Token.Value }

// Operatoren für Mathe & Vergleiche (z. B. counter + 1, x == 200)
type BinaryExpression struct {
	Token    Token // Das Operator-Token (+, -, ==, etc.)
	Left     Expression
	Operator string
	Right    Expression
}

func (be *BinaryExpression) expressionNode()      {}
func (be *BinaryExpression) TokenLiteral() string { return be.Token.Value }

// ==========================================
// 5. Array-Operationen (Zugriff, Entfernen, Hinzufügen)
// ==========================================

// Repräsentiert Operations-Formen wie x[0], x[^0], x[&^0], x[+="Wert"], x[0="Wert"]
type ArrayOpExpression struct {
	Token     Token      // Das '[' Token
	Array     Expression // Die Array-Variable (z. B. Identifier "x")
	Index     Expression // Index oder neu hinzuzufügender Wert (kann nil sein)
	Operation string     // "INDEX", "REMOVE", "REMOVE_MUTATE", "APPEND", "APPEND_MUTATE", "SET", "SET_MUTATE"
	Value     Expression // Der neu zuzuweisende Wert bei Überschreibung [0="Hallo"]
}

func (aoe *ArrayOpExpression) expressionNode()      {}
func (aoe *ArrayOpExpression) TokenLiteral() string { return aoe.Token.Value }

// ==========================================
// 6. Funktionen (flow) & Aufrufe
// ==========================================

// Parameterspezifikation für flow-Funktionen
type FunctionParameter struct {
	Name         string // Name des Parameters
	Type         string // Parametertyp (z. B. "Number", "Array")
	IsPipeTarget bool   // true wenn mit '{}' markiert (param{})
}

// Funktions-Definition: flow flowName(param paramType{}) { ... }
type FlowDefinitionStatement struct {
	Token      Token                // Das 'flow' Token
	Name       *Identifier          // Name des flows
	Parameters []*FunctionParameter // Parameterliste
	Body       *BlockStatement      // Funktionsinhalt
}

func (fds *FlowDefinitionStatement) statementNode()       {}
func (fds *FlowDefinitionStatement) TokenLiteral() string { return fds.Token.Value }

// Funktionsaufruf: processSuccess() oder keepRunning(param{})
type CallExpression struct {
	Token     Token        // Das '(' Token oder Name
	Function  Expression   // Funktions-Identifier
	Arguments []Expression // Argumente
}

func (ce *CallExpression) expressionNode()      {}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Value }

// ==========================================
// 7. Control Flow (match)
// ==========================================

// Ein einzelner Fall im match-Block: [== 200] ~> processSuccess()
type MatchCase struct {
	Token     Token      // Das '[' Token
	Condition Expression // Bedingung (z. B. BinaryExpression [== 200], ArrayOpExpression [list&^0] oder nil bei default)
	IsDefault bool       // true wenn [default]
	Body      Expression // Aktion, die ausgeführt wird (~> ...)
}

// Der gesamte match-Block: status ~> match { ... }
type MatchExpression struct {
	Token Token        // Das 'match' Token
	Cases []*MatchCase // Liste aller Fälle
}

func (me *MatchExpression) expressionNode()      {}
func (me *MatchExpression) TokenLiteral() string { return me.Token.Value }

// ==========================================
// 8. Built-in Funktionen (print & len)
// ==========================================

// Built-in len Aufruf: varName, varLen ~> len
type LenExpression struct {
	Token   Token       // Das 'len' Token
	Target  Expression  // Die eingehende Variable
	LenVar  *Identifier // Zweite Variable für die Länge (kann nil sein)
	InPlace bool        // true bei ~|> len
}

func (le *LenExpression) expressionNode()      {}
func (le *LenExpression) TokenLiteral() string { return le.Token.Value }

// Built-in print Aufruf (pass-through)
type PrintExpression struct {
	Token Token // Das 'print' Token
}

func (pe *PrintExpression) expressionNode()      {}
func (pe *PrintExpression) TokenLiteral() string { return pe.Token.Value }
