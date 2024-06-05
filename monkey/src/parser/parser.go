package parser

import (
	"fmt"
	"writing_an_interpreter_in_go/monkey/src/ast"
	"writing_an_interpreter_in_go/monkey/src/lexer"
	"writing_an_interpreter_in_go/monkey/src/token"
)

type Parser struct {
	l         *lexer.Lexer //pointer to an instance of the lexer
	curToken  token.Token
	peekToken token.Token
	errors    []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	//creates a new instance of the ast.Program struct and returns its address using the & operator.
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatment()
	default:
		return nil
	}
}

/**
* This function signature indicates that the parseLetStatement() method returns a pointer to ast.Statement.
* The * before ast.Statement denotes that the value being returned is a memory address pointing to an ast.Statement object.
* This allows the function to return nil if necessary, indicating the absence of a valid ast.Statement object.
**/

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}
	if !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseReturnStatment() *ast.ReturnStatment {
	stmt := &ast.ReturnStatment{Token: p.curToken}
	p.nextToken() //skip return keyword
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	p.peekError(t)
	return false
}

// ***** ERROR PART *****//

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}
