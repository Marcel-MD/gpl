package ast

import (
	"bytes"

	"github.com/Marcel-MD/gpl/token"
)

type PrefixExpression struct {
	Token    token.Token // Prefix token ! -
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (e *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("{\n")
	tab++
	out.WriteString(printTab() + "type: PREFIX " + string(e.Token.Type) + "\n")
	out.WriteString(printTab() + "operator: " + e.Operator + "\n")
	out.WriteString(printTab() + "right: " + e.Right.String())
	tab--
	out.WriteString(printTab() + "}\n")
	return out.String()
}

type InfixExpression struct {
	Token    token.Token // Operator token + * /
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (e *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("{\n")
	tab++
	out.WriteString(printTab() + "type: INFIX " + string(e.Token.Type) + "\n")
	out.WriteString(printTab() + "left: " + e.Left.String())
	out.WriteString(printTab() + "operator: " + e.Operator + "\n")
	out.WriteString(printTab() + "right: " + e.Right.String())
	tab--
	out.WriteString(printTab() + "}\n")
	return out.String()
}

type AssignExpression struct {
	Token    token.Token
	Name     *Identifier
	Operator string
	Value    Expression
}

func (as *AssignExpression) TokenLiteral() string { return as.Token.Literal }
func (e *AssignExpression) String() string {
	var out bytes.Buffer
	out.WriteString("{\n")
	tab++
	out.WriteString(printTab() + "type: ASSIGN\n")
	out.WriteString(printTab() + "identifier: " + e.Name.String())
	out.WriteString(printTab() + "operator: " + e.Operator + "\n")
	out.WriteString(printTab() + "value: " + e.Value.String())
	tab--
	out.WriteString(printTab() + "}\n")
	return out.String()
}

type IfExpression struct {
	Token       token.Token // The 'if' token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (e *IfExpression) String() string {
	var out bytes.Buffer
	out.WriteString("{\n")
	tab++
	out.WriteString(printTab() + "type: " + string(e.Token.Type) + "\n")
	out.WriteString(printTab() + "condition: " + e.Condition.String())
	out.WriteString(printTab() + "consequence: " + e.Consequence.String())
	if e.Alternative != nil {
		out.WriteString(printTab() + "alternative: " + e.Alternative.String())
	}
	tab--
	out.WriteString(printTab() + "}\n")
	return out.String()
}

type ForExpression struct {
	Token     token.Token // for token
	Condition Expression
	Block     *BlockStatement
}

func (ie *ForExpression) TokenLiteral() string { return ie.Token.Literal }
func (e *ForExpression) String() string {
	var out bytes.Buffer
	out.WriteString("{\n")
	tab++
	out.WriteString(printTab() + "type: " + string(e.Token.Type) + "\n")
	out.WriteString(printTab() + "condition: " + e.Condition.String())
	out.WriteString(printTab() + "body: " + e.Block.String())
	tab--
	out.WriteString(printTab() + "}\n")
	return out.String()
}

type CallExpression struct {
	Token     token.Token // The '(' token
	Function  Expression  // Identifier or FunctionLiteral
	Arguments []Expression
}

func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (e *CallExpression) String() string {
	var out bytes.Buffer
	out.WriteString("{\n")
	tab++
	out.WriteString(printTab() + "type: CALL\n")
	out.WriteString(printTab() + "function: " + e.Function.String())
	out.WriteString(printTab() + "arguments: [\n")
	tab++
	for _, a := range e.Arguments {
		out.WriteString(printTab() + a.String())
	}
	tab--
	out.WriteString(printTab() + "]\n")
	tab--
	out.WriteString(printTab() + "}\n")
	return out.String()
}

type IndexExpression struct {
	Token token.Token // The [ token
	Left  Expression
	Index Expression
}

func (ie *IndexExpression) TokenLiteral() string { return ie.Token.Literal }
func (e *IndexExpression) String() string {
	var out bytes.Buffer
	out.WriteString("{\n")
	tab++
	out.WriteString(printTab() + "type: INDEX\n")
	out.WriteString(printTab() + "left: " + e.Left.String())
	out.WriteString(printTab() + "index: " + e.Index.String())
	tab--
	out.WriteString(printTab() + "}\n")
	return out.String()
}
