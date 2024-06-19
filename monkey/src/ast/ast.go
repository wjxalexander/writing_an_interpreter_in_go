package ast

import (
	"bytes"
	"writing_an_interpreter_in_go/monkey/src/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

type Identifier struct {
	Token token.Token
	Value string
}

type ReturnStatment struct {
	Token       token.Token
	ReturnValue Expression
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		/**
		In Go, when embedding an interface within another interface,
		you don't access the embedded methods through the embedded interface name.
		Instead, you directly access the methods as if they were defined in the embedding interface itself.

		In the code you provided,
		the Node interface is embedded within the Statement interface using the line Node.
		This means that any type implementing the Statement interface will also implicitly implement the Node interface and its methods.
		*/

		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// satisfy Statement
func (ls *LetStatement) statementNode() {

}

// satisfy Node
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

func (i *Identifier) expressionNode() {

}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}

func (rs *ReturnStatment) statementNode() {

}

func (rs *ReturnStatment) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatment) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

func (es *ExpressionStatement) statementNode() {

}
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

func (il *IntegerLiteral) expressionNode() {

}

func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}
func (il *IntegerLiteral) String() string { return il.Token.Literal }
