package ast

import "writing_an_interpreter_in_go/monkey/src/token"

type Node interface {
	TokenLiteral() string
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

func (i *Identifier) expressionNode() {

}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (rs *ReturnStatment) statementNode() {

}

func (rs *ReturnStatment) TokenLiteral() string {
	return rs.Token.Literal
}
