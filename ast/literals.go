package ast

import (
	"bytes"
	"fmt"

	"github.com/Marcel-MD/gpl/token"
)

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (l *IntegerLiteral) String() string {
	var out bytes.Buffer
	out.WriteString("{\n")
	tab++
	out.WriteString(printTab() + "type: " + string(l.Token.Type) + "\n")
	out.WriteString(printTab() + fmt.Sprintf("value: %d\n", l.Value))
	tab--
	out.WriteString(printTab() + "}\n")
	return out.String()
}

type FloatLiteral struct {
	Token token.Token
	Value float64
}

func (fl *FloatLiteral) TokenLiteral() string { return fl.Token.Literal }
func (l *FloatLiteral) String() string {
	var out bytes.Buffer
	out.WriteString("{\n")
	tab++
	out.WriteString(printTab() + "type: " + string(l.Token.Type) + "\n")
	out.WriteString(printTab() + fmt.Sprintf("value: %f\n", l.Value))
	tab--
	out.WriteString(printTab() + "}\n")
	return out.String()
}

type StringLiteral struct {
	Token token.Token
	Value string
}

func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (l *StringLiteral) String() string {
	var out bytes.Buffer
	out.WriteString("{\n")
	tab++
	out.WriteString(printTab() + "type: " + string(l.Token.Type) + "\n")
	out.WriteString(printTab() + "value: " + l.Token.Literal + "\n")
	tab--
	out.WriteString(printTab() + "}\n")
	return out.String()
}

type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (l *Boolean) String() string {
	var out bytes.Buffer
	out.WriteString("{\n")
	tab++
	out.WriteString(printTab() + "type: " + string(l.Token.Type) + "\n")
	out.WriteString(printTab() + "value: " + l.Token.Literal + "\n")
	tab--
	out.WriteString(printTab() + "}\n")
	return out.String()
}

type FunctionLiteral struct {
	Token      token.Token // The 'fn' token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (s *FunctionLiteral) String() string {
	var out bytes.Buffer
	out.WriteString("{\n")
	tab++
	out.WriteString(printTab() + "type: FUNCTION\n")
	out.WriteString(printTab() + "parameters: [\n")
	tab++
	for _, p := range s.Parameters {
		out.WriteString(printTab() + p.String())
	}
	tab--
	out.WriteString(printTab() + "]\n")
	out.WriteString(printTab() + "body: " + s.Body.String())
	tab--
	out.WriteString(printTab() + "}\n")
	return out.String()
}

type ArrayLiteral struct {
	Token    token.Token // the '[' token
	Elements []Expression
}

func (al *ArrayLiteral) TokenLiteral() string { return al.Token.Literal }
func (b *ArrayLiteral) String() string {
	var out bytes.Buffer
	out.WriteString("{\n")
	tab++
	out.WriteString(printTab() + "type: ARRAY\n")
	out.WriteString(printTab() + "elements: [\n")
	tab++
	for _, s := range b.Elements {
		out.WriteString(printTab() + s.String())
	}
	tab--
	out.WriteString(printTab() + "]\n")
	tab--
	out.WriteString(printTab() + "}\n")
	return out.String()
}
