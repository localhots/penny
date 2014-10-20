package test

import (
	"io/ioutil"
	"testing"

	"github.com/kr/pretty"
	"github.com/localhots/penny/lexer"
	"github.com/localhots/penny/parser"
)

func TestWorld(t *testing.T) {
	b, err := ioutil.ReadFile("example.sh")
	if err != nil {
		panic(err)
	}

	lex := lexer.NewLexer(b)
	p := parser.NewParser()
	st, err := p.Parse(lex)
	if err != nil {
		panic(err)
	}

	pretty.Println(st)
}
