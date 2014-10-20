package ast

import (
	"strconv"

	"github.com/localhots/penny/token"
)

// Basics

type (
	IoNumber int
	Word     string
	Attrib   interface{}
)

func NewWord(word Attrib) (Word, error) {
	lit := word.(*token.Token).Lit
	return Word(lit), nil
}

func NewIoNumber(number Attrib) (IoNumber, error) {
	lit := number.(*token.Token).Lit
	n, err := strconv.Atoi(string(lit))
	return IoNumber(n), err
}

// Assignment

type (
	Assignment struct {
		a, b Word
	}
)

func NewAssignment(a, b Attrib) (*Assignment, error) {
	return &Assignment{a.(Word), b.(Word)}, nil
}

// List

type (
	List []*AndOr
)

func NewList(ao Attrib) (List, error) {
	return List{ao.(*AndOr)}, nil
}

func AppendToList(list Attrib, el Attrib) (List, error) {
	l := list.(List)
	l = append(l, el.(*AndOr))
	return l, nil
}

// AndOr

type (
	Logic     int
	AndOrStmt []*AndOr
	AndOr     struct {
		op Logic
		p  *Pipeline
	}
)

const (
	L_FIRST Logic = iota
	L_AND
	L_OR
)

func NewAndOr(p Attrib) (AndOrStmt, error) {
	return AndOrStmt{&AndOr{L_FIRST, p.(*Pipeline)}}, nil
}

func AppendAnd(stmt Attrib, p Attrib) (AndOrStmt, error) {
	pp := stmt.(AndOrStmt)
	pp = append(pp, &AndOr{L_AND, p.(*Pipeline)})
	return pp, nil
}

func AppendOr(stmt Attrib, p Attrib) (AndOrStmt, error) {
	pp := stmt.(AndOrStmt)
	pp = append(pp, &AndOr{L_OR, p.(*Pipeline)})
	return pp, nil
}

// Pipeline

type (
	Pipeline struct {
		seq     PipeSequence
		inverse bool
	}
)

func NewPipeline(seq Attrib, inverse bool) (*Pipeline, error) {
	return &Pipeline{seq.(PipeSequence), inverse}, nil
}

// PipeSequence

type (
	PipeSequence []Command
)

func NewPipeSequence(cmd Attrib) (PipeSequence, error) {
	return PipeSequence{cmd.(*Command)}, nil
}

func AppendToPipeSequence(ps Attrib, cmd Attrib) (PipeSequence, error) {
	ps1 := ps.(PipeSequence)
	ps1 = append(ps1, cmd.(Command))
	return ps1, nil
}

// Command

type (
	Command interface{}
)

func NewCommand(cmd Attrib, rl Attrib) (Command, error) {
	return Command(cmd), nil
}

// CompoundCommand

type (
	CompoundCommand interface{}
)

func NewCompoundCommand(cmd Attrib) (CompoundCommand, error) {
	return CompoundCommand(cmd), nil
}

// Subshell

type (
	Subshell struct {
		cc CompoundCommand
	}
)

func NewSubshell(cc Attrib) (*Subshell, error) {
	return &Subshell{cc.(CompoundCommand)}, nil
}

// CompoundList

type (
	CompoundList struct {
		term *Term
		sep  *Separator
	}
)

func NewCompoundList(term Attrib, sep Attrib) (*CompoundList, error) {
	return &CompoundList{term.(*Term), sep.(*Separator)}, nil
}

// Term

type (
	Term     []*TermItem
	TermItem struct {
		ao  AndOrStmt
		sep *Separator
	}
)

func NewTerm(ao Attrib) (Term, error) {
	return Term{&TermItem{ao.(AndOrStmt), nil}}, nil
}

func AppendToTerm(term Attrib, ao Attrib, sep Attrib) (Term, error) {
	t := term.(Term)
	t = append(t, &TermItem{ao.(AndOrStmt), sep.(*Separator)})
	return t, nil
}

// ForClause

type (
	ForClause struct {
		name Name
		wl   Wordlist
		dg   *DoGroup
	}
)

func NewForClause(name Attrib, wl Attrib, dg Attrib) (*ForClause, error) {
	return &ForClause{
		name: name.(Name),
		wl:   wl.(Wordlist),
		dg:   dg.(*DoGroup),
	}, nil
}

// Name

type (
	Name Word
)

func NewName(w Attrib) (Name, error) {
	return Name(w.(Word)), nil
}

// Wordlist

type (
	Wordlist []Word
)

func NewWordlist(w Attrib) (Wordlist, error) {
	return Wordlist{w.(Word)}, nil
}

func AppendToWordlist(wl Attrib, w Attrib) (Wordlist, error) {
	wl1 := wl.(Wordlist)
	wl1 = append(wl1, w.(Word))
	return wl1, nil
}

// CaseClause

type (
	CaseClause struct {
		word Word
		cl   CaseList
	}
)

func NewCaseClause(word Attrib, cl Attrib) (*CaseClause, error) {
	return &CaseClause{
		word: word.(Word),
		cl:   cl.(CaseList),
	}, nil
}

// CaseListNs
// CaseList

type (
	CaseList []*CaseItem
)

func NewCaseList(ci Attrib) (CaseList, error) {
	return CaseList{ci.(*CaseItem)}, nil
}

func AppendToCaseList(cl Attrib, ci Attrib) (CaseList, error) {
	cl1 := cl.(CaseList)
	cl1 = append(cl1, ci.(*CaseItem))
	return cl1, nil
}

// CaseItemNs
// CaseItem

type (
	CaseItem struct {
		p  Pattern
		cl *CompoundList
	}
)

func NewCaseItem(p Attrib, cl Attrib) (*CaseItem, error) {
	return &CaseItem{
		p:  p.(Pattern),
		cl: cl.(*CompoundList),
	}, nil
}

// Pattern

type (
	Pattern []Word
)

func NewPattern(w Attrib) (Pattern, error) {
	return Pattern{w.(Word)}, nil
}

func AppendToPattern(p Attrib, w Attrib) (Pattern, error) {
	p1 := p.(Pattern)
	p1 = append(p1, w.(Word))
	return p1, nil
}

// IfClause
// ElsePart

type (
	IfClause struct {
		cond   *CompoundList
		action *CompoundList
		elsep  *IfClause
	}
)

func NewIfClause(cond Attrib, action Attrib, elsep Attrib) (*IfClause, error) {
	return &IfClause{
		cond:   cond.(*CompoundList),
		action: action.(*CompoundList),
		elsep:  elsep.(*IfClause),
	}, nil
}

// WhileClause

type (
	WhileClause struct {
		cond *CompoundList
		dg   *DoGroup
	}
)

func NewWhileClause(cond Attrib, dg Attrib) (*WhileClause, error) {
	return &WhileClause{
		cond: cond.(*CompoundList),
		dg:   dg.(*DoGroup),
	}, nil
}

// UntilClause

type (
	UntilClause struct {
		cond *CompoundList
		dg   *DoGroup
	}
)

func NewUntilClause(cond Attrib, dg Attrib) (*UntilClause, error) {
	return &UntilClause{
		cond: cond.(*CompoundList),
		dg:   dg.(*DoGroup),
	}, nil
}

// FunctionDefinition

type (
	FunctionDefinition struct {
		name FunctionName
		body *FunctionBody
	}
)

func NewFunctionDefinition(name Attrib, body Attrib) (*FunctionDefinition, error) {
	return &FunctionDefinition{
		name: name.(FunctionName),
		body: body.(*FunctionBody),
	}, nil
}

// FunctionBody

type (
	FunctionBody struct {
		cc CompoundCommand
		rl RedirectList
	}
)

func NewFunctionBody(cc Attrib, rl Attrib) (*FunctionBody, error) {
	return &FunctionBody{
		cc: cc.(CompoundCommand),
		rl: rl.(RedirectList),
	}, nil
}

// FunctionName

type (
	FunctionName Word
)

func NewFunctionName(w Attrib) (FunctionName, error) {
	return FunctionName(w.(Word)), nil
}

// BraceGroup

type (
	BraceGroup struct {
		cl *CompoundList
	}
)

func NewBraceGroup(cl Attrib) (*BraceGroup, error) {
	return &BraceGroup{
		cl: cl.(*CompoundList),
	}, nil
}

// DoGroup

type (
	DoGroup struct {
		cl *CompoundList
	}
)

func NewDoGroup(cl Attrib) (*DoGroup, error) {
	return &DoGroup{cl.(*CompoundList)}, nil
}

// SimpleCommand

type (
	SimpleCommand struct {
		prefix *CmdPrefix
		name   Name
		word   Word
		suffix *CmdSuffix
	}
)

func NewSimpleCommand(prefix, name, word, suffix Attrib) (*SimpleCommand, error) {
	return &SimpleCommand{
		prefix: prefix.(*CmdPrefix),
		name:   name.(Name),
		word:   word.(Word),
		suffix: suffix.(*CmdSuffix),
	}, nil
}

// CmdName

type (
	CmdName Word
)

func NewCmdName(w Attrib) (CmdName, error) {
	return CmdName(w.(Word)), nil
}

// CmdWord

type (
	CmdWord Word
)

func NewCmdWord(w Attrib) (CmdWord, error) {
	return CmdWord(w.(Word)), nil
}

// CmdPrefix

type (
	CmdPrefix struct {
		assign *Assignment
		redir  *IoRedirect
		prefix *CmdPrefix
	}
)

func NewCmdPrefix(assign, redir, prefix Attrib) (*CmdPrefix, error) {
	return &CmdPrefix{
		assign: assign.(*Assignment),
		redir:  assign.(*IoRedirect),
		prefix: prefix.(*CmdPrefix),
	}, nil
}

// CmdSuffix

type (
	CmdSuffix struct {
		word   Word
		redir  *IoRedirect
		suffix *CmdSuffix
	}
)

func NewCmdSuffix(word, redir, suffix Attrib) (*CmdSuffix, error) {
	return &CmdSuffix{
		word:   word.(Word),
		redir:  redir.(*IoRedirect),
		suffix: suffix.(*CmdSuffix),
	}, nil
}

// RedirectList

type (
	RedirectList []*IoRedirect
)

func NewRedirectList(redir Attrib) (RedirectList, error) {
	return RedirectList{redir.(*IoRedirect)}, nil
}

func AppendToRedirectList(list Attrib, el Attrib) (RedirectList, error) {
	l := list.(RedirectList)
	l = append(l, el.(*IoRedirect))
	return l, nil
}

// IoRedirect

type (
	IoRedirect struct {
		file *IoFile
		num  IoNumber
		here *IoHere
	}
)

func NewIoRedirect(file Attrib, num Attrib, here Attrib) (*IoRedirect, error) {
	return &IoRedirect{
		file: file.(*IoFile),
		num:  num.(IoNumber),
		here: here.(*IoHere),
	}, nil
}

// IoFile

type (
	Redirection int
	IoFile      struct {
		file  Filename
		redir Redirection
	}
)

const (
	R_STDIN Redirection = iota
	R_INFD
	R_STDOUT
	R_OUTFD
	R_APPEND
	R_ORWFD
	R_OUTSP
)

func NewIoFile(file Attrib, redir Redirection) (*IoFile, error) {
	return &IoFile{
		file:  file.(Filename),
		redir: redir,
	}, nil
}

// Filename

type (
	Filename Word
)

func NewFilename(w Attrib) (Filename, error) {
	return Filename(w.(Word)), nil
}

// IoHere

type (
	IoHere struct {
		word         Word
		suppressTabs bool
	}
)

func NewIoHere(w Attrib, st bool) (*IoHere, error) {
	return &IoHere{
		word:         w.(Word),
		suppressTabs: st,
	}, nil
}

// HereEnd

type (
	HereEnd Word
)

func NewHereEnd(w Attrib) (HereEnd, error) {
	return HereEnd(w.(Word)), nil
}

// NewlineList

// Linebreak

// SeparatorOp

type (
	SeparatorOp int
)

const (
	S_AMP SeparatorOp = iota
	S_SEMICOLON
)

// Separartor

type (
	Separator struct {
		s SeparatorOp
	}
)

func NewSeparator(op Attrib) (*Separator, error) {
	return &Separator{op.(SeparatorOp)}, nil
}

// SequentialSep
