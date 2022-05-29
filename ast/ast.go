package ast

import (
	"bytes"

	"github.com/Marcel-MD/gpl/token"
)

var tab int = 0

func printTab() string {
	var out bytes.Buffer
	for i := 0; i < tab; i++ {
		out.WriteString("  ")
	}
	return out.String()
}

type Expression interface {
	TokenLiteral() string
	String() string // for testing
}

type Program struct {
	Statements []Expression
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer
	out.WriteString("{\n")
	tab++
	out.WriteString(printTab() + "type: PROGRAM\n")
	out.WriteString(printTab() + "statements: [\n")
	tab++
	for _, s := range p.Statements {
		out.WriteString(printTab() + s.String())
	}
	tab--
	out.WriteString(printTab() + "]\n")
	tab--
	out.WriteString("}\n")
	return out.String()
}

type ReturnStatement struct {
	Token token.Token // RETURN token
	Value Expression
}

func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (s *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString("{\n")
	tab++
	out.WriteString(printTab() + "type: " + string(s.Token.Type) + "\n")
	out.WriteString(printTab() + "value: " + s.Value.String())
	tab--
	out.WriteString(printTab() + "}\n")
	return out.String()
}

type ExpressionStatement struct {
	Token token.Token // the first token of the expression
	Value Expression
}

func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (s *ExpressionStatement) String() string {
	return s.Value.String()
}

type Identifier struct {
	Token token.Token // IDENT token
	Value string
}

func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string {
	var out bytes.Buffer
	out.WriteString("{\n")
	tab++
	out.WriteString(printTab() + "type: " + string(i.Token.Type) + "\n")
	out.WriteString(printTab() + "value: " + i.Value + "\n")
	tab--
	out.WriteString(printTab() + "}\n")
	return out.String()
}

type BlockStatement struct {
	Token      token.Token // the { token
	Statements []Expression
}

func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (b *BlockStatement) String() string {
	var out bytes.Buffer
	out.WriteString("{\n")
	tab++
	out.WriteString(printTab() + "type: BLOCK\n")
	out.WriteString(printTab() + "statements: [\n")
	tab++
	for _, s := range b.Statements {
		out.WriteString(printTab() + s.String())
	}
	tab--
	out.WriteString(printTab() + "]\n")
	tab--
	out.WriteString(printTab() + "}\n")
	return out.String()
}
