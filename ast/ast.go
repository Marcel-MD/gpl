package ast

import (
	"github.com/Marcel-MD/gpl/token"
)

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode() // dummy method
}

type Expression interface {
	Node
	expressionNode() // dummy method
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) statementNode()       {}

type Identifier struct {
	Token token.Token // IDENT token
	Value string
}

func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) expressionNode()      {}
