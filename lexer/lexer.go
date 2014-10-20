
package lexer

import (
	
	// "fmt"
	// "github.com/localhots/penny/util"
	
	"io/ioutil"
	"unicode/utf8"
	"github.com/localhots/penny/token"
)

const(
	NoState = -1
	NumStates = 73
	NumSymbols = 90
) 

type Lexer struct {
	src             []byte
	pos             int
	line            int
	column          int
}

func NewLexer(src []byte) *Lexer {
	lexer := &Lexer{
		src:    src,
		pos:    0,
		line:   1,
		column: 1,
	}
	return lexer
}

func NewLexerFile(fpath string) (*Lexer, error) {
	src, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	return NewLexer(src), nil
}

func (this *Lexer) Scan() (tok *token.Token) {
	
	// fmt.Printf("Lexer.Scan() pos=%d\n", this.pos)
	
	tok = new(token.Token)
	if this.pos >= len(this.src) {
		tok.Type = token.EOF
		tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = this.pos, this.line, this.column
		return
	}
	start, end := this.pos, 0
	tok.Type = token.INVALID
	state, rune1, size := 0, rune(-1), 0
	for state != -1 {
	
		// fmt.Printf("\tpos=%d, line=%d, col=%d, state=%d\n", this.pos, this.line, this.column, state)
	
		if this.pos >= len(this.src) {
			rune1 = -1
		} else {
			rune1, size = utf8.DecodeRune(this.src[this.pos:])
			this.pos += size
		}
		switch rune1 {
		case '\n':
			this.line++
			this.column = 1
		case '\r':
			this.column = 1
		case '\t':
			this.column += 4
		default:
			this.column++
		}

	
		// Production start
		if rune1 != -1 {
			state = TransTab[state](rune1)
		} else {
			state = -1
		}
		// Production end

		// Debug start
		// nextState := -1
		// if rune1 != -1 {
		// 	nextState = TransTab[state](rune1)
		// }
		// fmt.Printf("\tS%d, : tok=%s, rune == %s(%x), next state == %d\n", state, token.TokMap.Id(tok.Type), util.RuneToString(rune1), rune1, nextState)
		// fmt.Printf("\t\tpos=%d, size=%d, start=%d, end=%d\n", this.pos, size, start, end)
		// if nextState != -1 {
		// 	fmt.Printf("\t\taction:%s\n", ActTab[nextState].String())
		// }
		// state = nextState
		// Debug end
	

		if state != -1 {
			switch {
			case ActTab[state].Accept != -1:
				tok.Type = ActTab[state].Accept
				// fmt.Printf("\t Accept(%s), %s(%d)\n", string(act), token.TokMap.Id(tok), tok)
				end = this.pos
			case ActTab[state].Ignore != "":
				// fmt.Printf("\t Ignore(%s)\n", string(act))
				start = this.pos
				state = 0
				if start >= len(this.src) {
					tok.Type = token.EOF
				}

			}
		} else {
			if tok.Type == token.INVALID {
				end = this.pos
			}
		}
	}
	if end > start {
		this.pos = end
		tok.Lit = this.src[start:end]
	} else {
		tok.Lit = []byte{}
	}
	tok.Pos.Offset = start
	tok.Pos.Column = this.column
	tok.Pos.Line = this.line
	return
}

func (this *Lexer) Reset() {
	this.pos = 0
}

/*
Lexer symbols:
0: '_'
1: '_'
2: '_'
3: '='
4: '&'
5: '&'
6: '|'
7: '|'
8: '!'
9: '|'
10: '('
11: ')'
12: 'f'
13: 'o'
14: 'r'
15: 'i'
16: 'n'
17: 'c'
18: 'a'
19: 's'
20: 'e'
21: 'e'
22: 's'
23: 'a'
24: 'c'
25: ';'
26: ';'
27: 'i'
28: 'f'
29: 't'
30: 'h'
31: 'e'
32: 'n'
33: 'f'
34: 'i'
35: 'e'
36: 'l'
37: 'i'
38: 'f'
39: 'e'
40: 'l'
41: 's'
42: 'e'
43: 'w'
44: 'h'
45: 'i'
46: 'l'
47: 'e'
48: 'u'
49: 'n'
50: 't'
51: 'i'
52: 'l'
53: '{'
54: '}'
55: 'd'
56: 'o'
57: 'd'
58: 'o'
59: 'n'
60: 'e'
61: '<'
62: '<'
63: '&'
64: '>'
65: '>'
66: '&'
67: '>'
68: '>'
69: '<'
70: '>'
71: '>'
72: '|'
73: '<'
74: '<'
75: '<'
76: '<'
77: '-'
78: '\'
79: 'n'
80: '&'
81: ';'
82: ' '
83: '\t'
84: '\n'
85: '\r'
86: 'a'-'z'
87: 'A'-'Z'
88: '0'-'9'
89: .

*/
