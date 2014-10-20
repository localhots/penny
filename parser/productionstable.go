
package parser

import "github.com/localhots/penny/ast"

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab {
	ProdTabEntry{
		String: `S' : Word	<<  >>`,
		Id: "S'",
		NTType: 0,
		Index: 0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Word : word	<< ast.NewWord(X[0]) >>`,
		Id: "Word",
		NTType: 1,
		Index: 1,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewWord(X[0])
		},
	},
	ProdTabEntry{
		String: `IoNumber : number	<< ast.NewIoNumber(X[0]) >>`,
		Id: "IoNumber",
		NTType: 2,
		Index: 2,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIoNumber(X[0])
		},
	},
	ProdTabEntry{
		String: `AssignmentWord : Word "=" Word	<< ast.NewAssignment(X[0], X[2]) >>`,
		Id: "AssignmentWord",
		NTType: 3,
		Index: 3,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewAssignment(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `List : List SeparatorOp AndOr	<< ast.AppendToList(X[0], X[2]) >>`,
		Id: "List",
		NTType: 4,
		Index: 4,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendToList(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `List : AndOr	<< ast.NewList(X[0]) >>`,
		Id: "List",
		NTType: 4,
		Index: 5,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewList(X[0])
		},
	},
	ProdTabEntry{
		String: `AndOr : Pipeline	<< ast.NewAndOr(X[0]) >>`,
		Id: "AndOr",
		NTType: 5,
		Index: 6,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewAndOr(X[0])
		},
	},
	ProdTabEntry{
		String: `AndOr : AndOr "&&" Linebreak Pipeline	<< ast.AppendAnd(X[0], X[3]) >>`,
		Id: "AndOr",
		NTType: 5,
		Index: 7,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendAnd(X[0], X[3])
		},
	},
	ProdTabEntry{
		String: `AndOr : AndOr "||" Linebreak Pipeline	<< ast.AppendOr(X[0], X[3]) >>`,
		Id: "AndOr",
		NTType: 5,
		Index: 8,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendOr(X[0], X[3])
		},
	},
	ProdTabEntry{
		String: `Pipeline : PipeSequence	<< ast.NewPipeline(X[1], false) >>`,
		Id: "Pipeline",
		NTType: 6,
		Index: 9,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewPipeline(X[1], false)
		},
	},
	ProdTabEntry{
		String: `Pipeline : "!" PipeSequence	<< ast.NewPipeline(X[1], true) >>`,
		Id: "Pipeline",
		NTType: 6,
		Index: 10,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewPipeline(X[1], true)
		},
	},
	ProdTabEntry{
		String: `PipeSequence : Command	<< ast.NewPipeSequence(X[0]) >>`,
		Id: "PipeSequence",
		NTType: 7,
		Index: 11,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewPipeSequence(X[0])
		},
	},
	ProdTabEntry{
		String: `PipeSequence : PipeSequence "|" Linebreak Command	<< ast.AppendToPipeSequence(X[0], X[3]) >>`,
		Id: "PipeSequence",
		NTType: 7,
		Index: 12,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendToPipeSequence(X[0], X[3])
		},
	},
	ProdTabEntry{
		String: `Command : SimpleCommand	<< ast.NewCommand(X[0], nil) >>`,
		Id: "Command",
		NTType: 8,
		Index: 13,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCommand(X[0], nil)
		},
	},
	ProdTabEntry{
		String: `Command : CompoundCommand	<< ast.NewCommand(X[0], nil) >>`,
		Id: "Command",
		NTType: 8,
		Index: 14,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCommand(X[0], nil)
		},
	},
	ProdTabEntry{
		String: `Command : CompoundCommand RedirectList	<< ast.NewCommand(X[0], X[1]) >>`,
		Id: "Command",
		NTType: 8,
		Index: 15,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCommand(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `Command : FunctionDefinition	<< ast.NewCommand(X[0], nil) >>`,
		Id: "Command",
		NTType: 8,
		Index: 16,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCommand(X[0], nil)
		},
	},
	ProdTabEntry{
		String: `CompoundCommand : BraceGroup	<< ast.NewCompoundCommand(X[0]) >>`,
		Id: "CompoundCommand",
		NTType: 9,
		Index: 17,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCompoundCommand(X[0])
		},
	},
	ProdTabEntry{
		String: `CompoundCommand : Subshell	<< ast.NewCompoundCommand(X[0]) >>`,
		Id: "CompoundCommand",
		NTType: 9,
		Index: 18,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCompoundCommand(X[0])
		},
	},
	ProdTabEntry{
		String: `CompoundCommand : ForClause	<< ast.NewCompoundCommand(X[0]) >>`,
		Id: "CompoundCommand",
		NTType: 9,
		Index: 19,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCompoundCommand(X[0])
		},
	},
	ProdTabEntry{
		String: `CompoundCommand : CaseClause	<< ast.NewCompoundCommand(X[0]) >>`,
		Id: "CompoundCommand",
		NTType: 9,
		Index: 20,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCompoundCommand(X[0])
		},
	},
	ProdTabEntry{
		String: `CompoundCommand : IfClause	<< ast.NewCompoundCommand(X[0]) >>`,
		Id: "CompoundCommand",
		NTType: 9,
		Index: 21,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCompoundCommand(X[0])
		},
	},
	ProdTabEntry{
		String: `CompoundCommand : WhileClause	<< ast.NewCompoundCommand(X[0]) >>`,
		Id: "CompoundCommand",
		NTType: 9,
		Index: 22,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCompoundCommand(X[0])
		},
	},
	ProdTabEntry{
		String: `CompoundCommand : UntilClause	<< ast.NewCompoundCommand(X[0]) >>`,
		Id: "CompoundCommand",
		NTType: 9,
		Index: 23,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCompoundCommand(X[0])
		},
	},
	ProdTabEntry{
		String: `Subshell : "(" CompoundList ")"	<< ast.NewSubshell(X[1]) >>`,
		Id: "Subshell",
		NTType: 10,
		Index: 24,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewSubshell(X[1])
		},
	},
	ProdTabEntry{
		String: `CompoundList : Term	<< ast.NewCompoundList(X[0], nil) >>`,
		Id: "CompoundList",
		NTType: 11,
		Index: 25,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCompoundList(X[0], nil)
		},
	},
	ProdTabEntry{
		String: `CompoundList : NewlineList Term	<< ast.NewCompoundList(X[1], X[2]) >>`,
		Id: "CompoundList",
		NTType: 11,
		Index: 26,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCompoundList(X[1], X[2])
		},
	},
	ProdTabEntry{
		String: `CompoundList : Term Separator	<< ast.NewCompoundList(X[0], X[1]) >>`,
		Id: "CompoundList",
		NTType: 11,
		Index: 27,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCompoundList(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `CompoundList : NewlineList Term Separator	<< ast.NewCompoundList(X[1], X[2]) >>`,
		Id: "CompoundList",
		NTType: 11,
		Index: 28,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCompoundList(X[1], X[2])
		},
	},
	ProdTabEntry{
		String: `Term : Term Separator AndOr	<< ast.AppendToTerm(X[0], X[2], X[1]) >>`,
		Id: "Term",
		NTType: 12,
		Index: 29,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendToTerm(X[0], X[2], X[1])
		},
	},
	ProdTabEntry{
		String: `Term : AndOr	<< ast.NewTerm(X[0]) >>`,
		Id: "Term",
		NTType: 12,
		Index: 30,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewTerm(X[0])
		},
	},
	ProdTabEntry{
		String: `ForClause : "for" Name Linebreak DoGroup	<< ast.NewForClause(X[1], ast.Wordlist{}, X[3]) >>`,
		Id: "ForClause",
		NTType: 13,
		Index: 31,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewForClause(X[1], ast.Wordlist{}, X[3])
		},
	},
	ProdTabEntry{
		String: `ForClause : "for" Name Linebreak "in" SequentialSep DoGroup	<< ast.NewForClause(X[1], ast.Wordlist{}, X[5]) >>`,
		Id: "ForClause",
		NTType: 13,
		Index: 32,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewForClause(X[1], ast.Wordlist{}, X[5])
		},
	},
	ProdTabEntry{
		String: `ForClause : "for" Name Linebreak "in" Wordlist SequentialSep DoGroup	<< ast.NewForClause(X[1], X[4], X[6]) >>`,
		Id: "ForClause",
		NTType: 13,
		Index: 33,
		NumSymbols: 7,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewForClause(X[1], X[4], X[6])
		},
	},
	ProdTabEntry{
		String: `Name : name	<< ast.NewName(X[0]) >>`,
		Id: "Name",
		NTType: 14,
		Index: 34,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewName(X[0])
		},
	},
	ProdTabEntry{
		String: `Wordlist : Wordlist Word	<< ast.AppendToWordlist(X[0], X[1]) >>`,
		Id: "Wordlist",
		NTType: 15,
		Index: 35,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendToWordlist(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `Wordlist : Word	<< ast.NewWordlist(X[0]) >>`,
		Id: "Wordlist",
		NTType: 15,
		Index: 36,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewWordlist(X[0])
		},
	},
	ProdTabEntry{
		String: `CaseClause : "case" Word Linebreak "in" Linebreak CaseList "esac"	<< ast.NewCaseClause(X[1], X[5]) >>`,
		Id: "CaseClause",
		NTType: 16,
		Index: 37,
		NumSymbols: 7,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCaseClause(X[1], X[5])
		},
	},
	ProdTabEntry{
		String: `CaseClause : "case" Word Linebreak "in" Linebreak CaseListNs "esac"	<< ast.NewCaseClause(X[1], X[5]) >>`,
		Id: "CaseClause",
		NTType: 16,
		Index: 38,
		NumSymbols: 7,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCaseClause(X[1], X[5])
		},
	},
	ProdTabEntry{
		String: `CaseClause : "case" Word Linebreak "in" Linebreak "esac"	<< ast.NewCaseClause(X[1], ast.CaseList{}) >>`,
		Id: "CaseClause",
		NTType: 16,
		Index: 39,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCaseClause(X[1], ast.CaseList{})
		},
	},
	ProdTabEntry{
		String: `CaseListNs : CaseList CaseItemNs	<< ast.AppendToCaseList(X[0], X[1]) >>`,
		Id: "CaseListNs",
		NTType: 17,
		Index: 40,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendToCaseList(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `CaseListNs : CaseItemNs	<< ast.NewCaseList(X[0]) >>`,
		Id: "CaseListNs",
		NTType: 17,
		Index: 41,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCaseList(X[0])
		},
	},
	ProdTabEntry{
		String: `CaseList : CaseList CaseItem	<< ast.AppendToCaseList(X[0], X[1]) >>`,
		Id: "CaseList",
		NTType: 18,
		Index: 42,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendToCaseList(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `CaseList : CaseItem	<< ast.NewCaseList(X[0]) >>`,
		Id: "CaseList",
		NTType: 18,
		Index: 43,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCaseList(X[0])
		},
	},
	ProdTabEntry{
		String: `CaseItemNs : Pattern ")" Linebreak	<< ast.NewCaseItem(X[0], nil) >>`,
		Id: "CaseItemNs",
		NTType: 19,
		Index: 44,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCaseItem(X[0], nil)
		},
	},
	ProdTabEntry{
		String: `CaseItemNs : Pattern ")" CompoundList Linebreak	<< ast.NewCaseItem(X[0], X[2]) >>`,
		Id: "CaseItemNs",
		NTType: 19,
		Index: 45,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCaseItem(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `CaseItemNs : "(" Pattern ")" Linebreak	<< ast.NewCaseItem(X[1], nil) >>`,
		Id: "CaseItemNs",
		NTType: 19,
		Index: 46,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCaseItem(X[1], nil)
		},
	},
	ProdTabEntry{
		String: `CaseItemNs : "(" Pattern ")" CompoundList Linebreak	<< ast.NewCaseItem(X[1], X[3]) >>`,
		Id: "CaseItemNs",
		NTType: 19,
		Index: 47,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCaseItem(X[1], X[3])
		},
	},
	ProdTabEntry{
		String: `CaseItem : Pattern ")" Linebreak ";;" Linebreak	<< ast.NewCaseItem(X[0], nil) >>`,
		Id: "CaseItem",
		NTType: 20,
		Index: 48,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCaseItem(X[0], nil)
		},
	},
	ProdTabEntry{
		String: `CaseItem : Pattern ")" CompoundList ";;" Linebreak	<< ast.NewCaseItem(X[0], X[2]) >>`,
		Id: "CaseItem",
		NTType: 20,
		Index: 49,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCaseItem(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `CaseItem : "(" Pattern ")" Linebreak ";;" Linebreak	<< ast.NewCaseItem(X[1], nil) >>`,
		Id: "CaseItem",
		NTType: 20,
		Index: 50,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCaseItem(X[1], nil)
		},
	},
	ProdTabEntry{
		String: `CaseItem : "(" Pattern ")" CompoundList ";;" Linebreak	<< ast.NewCaseItem(X[1], X[3]) >>`,
		Id: "CaseItem",
		NTType: 20,
		Index: 51,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCaseItem(X[1], X[3])
		},
	},
	ProdTabEntry{
		String: `Pattern : Word	<< ast.NewPattern(X[0]) >>`,
		Id: "Pattern",
		NTType: 21,
		Index: 52,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewPattern(X[0])
		},
	},
	ProdTabEntry{
		String: `Pattern : Pattern "|" Word	<< ast.AppendToPattern(X[0], X[2]) >>`,
		Id: "Pattern",
		NTType: 21,
		Index: 53,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendToPattern(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `IfClause : "if" CompoundList "then" CompoundList ElsePart "fi"	<< ast.NewIfClause(X[1], X[3], X[4]) >>`,
		Id: "IfClause",
		NTType: 22,
		Index: 54,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIfClause(X[1], X[3], X[4])
		},
	},
	ProdTabEntry{
		String: `IfClause : "if" CompoundList "then" CompoundList "fi"	<< ast.NewIfClause(X[1], X[3], nil) >>`,
		Id: "IfClause",
		NTType: 22,
		Index: 55,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIfClause(X[1], X[3], nil)
		},
	},
	ProdTabEntry{
		String: `ElsePart : "elif" CompoundList "then" ElsePart	<< ast.NewIfClause(X[1], nil, X[3]) >>`,
		Id: "ElsePart",
		NTType: 23,
		Index: 56,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIfClause(X[1], nil, X[3])
		},
	},
	ProdTabEntry{
		String: `ElsePart : "else" CompoundList	<< ast.NewIfClause(nil, X[1], nil) >>`,
		Id: "ElsePart",
		NTType: 23,
		Index: 57,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIfClause(nil, X[1], nil)
		},
	},
	ProdTabEntry{
		String: `WhileClause : "while" CompoundList DoGroup	<< ast.NewWhileClause(X[1], X[2]) >>`,
		Id: "WhileClause",
		NTType: 24,
		Index: 58,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewWhileClause(X[1], X[2])
		},
	},
	ProdTabEntry{
		String: `UntilClause : "until" CompoundList DoGroup	<< ast.NewUntilClause(X[1], X[2]) >>`,
		Id: "UntilClause",
		NTType: 25,
		Index: 59,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewUntilClause(X[1], X[2])
		},
	},
	ProdTabEntry{
		String: `FunctionDefinition : FunctionName "(" ")" Linebreak FunctionBody	<< ast.NewFunctionDefinition(X[0], X[4]) >>`,
		Id: "FunctionDefinition",
		NTType: 26,
		Index: 60,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewFunctionDefinition(X[0], X[4])
		},
	},
	ProdTabEntry{
		String: `FunctionBody : CompoundCommand	<< ast.NewFunctionBody(X[0], ast.RedirectList{}) >>`,
		Id: "FunctionBody",
		NTType: 27,
		Index: 61,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewFunctionBody(X[0], ast.RedirectList{})
		},
	},
	ProdTabEntry{
		String: `FunctionBody : CompoundCommand RedirectList	<< ast.NewFunctionBody(X[0], X[1]) >>`,
		Id: "FunctionBody",
		NTType: 27,
		Index: 62,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewFunctionBody(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `FunctionName : Word	<< ast.NewFunctionName(X[0]) >>`,
		Id: "FunctionName",
		NTType: 28,
		Index: 63,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewFunctionName(X[0])
		},
	},
	ProdTabEntry{
		String: `BraceGroup : "{" CompoundList "}"	<< ast.NewBraceGroup(X[1]) >>`,
		Id: "BraceGroup",
		NTType: 29,
		Index: 64,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBraceGroup(X[1])
		},
	},
	ProdTabEntry{
		String: `DoGroup : "do" CompoundList "done"	<< ast.NewDoGroup(X[1]) >>`,
		Id: "DoGroup",
		NTType: 30,
		Index: 65,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewDoGroup(X[1])
		},
	},
	ProdTabEntry{
		String: `SimpleCommand : CmdPrefix CmdWord CmdSuffix	<< ast.NewSimpleCommand(X[0],  nil, X[1],  X[2]) >>`,
		Id: "SimpleCommand",
		NTType: 31,
		Index: 66,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewSimpleCommand(X[0],  nil, X[1],  X[2])
		},
	},
	ProdTabEntry{
		String: `SimpleCommand : CmdPrefix CmdWord	<< ast.NewSimpleCommand(X[0],  nil, X[1],  nil) >>`,
		Id: "SimpleCommand",
		NTType: 31,
		Index: 67,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewSimpleCommand(X[0],  nil, X[1],  nil)
		},
	},
	ProdTabEntry{
		String: `SimpleCommand : CmdPrefix	<< ast.NewSimpleCommand(X[0],  nil, nil, nil) >>`,
		Id: "SimpleCommand",
		NTType: 31,
		Index: 68,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewSimpleCommand(X[0],  nil, nil, nil)
		},
	},
	ProdTabEntry{
		String: `SimpleCommand : CmdName CmdSuffix	<< ast.NewSimpleCommand(nil, X[0],  nil, X[1]) >>`,
		Id: "SimpleCommand",
		NTType: 31,
		Index: 69,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewSimpleCommand(nil, X[0],  nil, X[1])
		},
	},
	ProdTabEntry{
		String: `SimpleCommand : CmdName	<< ast.NewSimpleCommand(nil, X[0],  nil, nil) >>`,
		Id: "SimpleCommand",
		NTType: 31,
		Index: 70,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewSimpleCommand(nil, X[0],  nil, nil)
		},
	},
	ProdTabEntry{
		String: `CmdName : Word	<< ast.NewCmdName(X[0]) >>`,
		Id: "CmdName",
		NTType: 32,
		Index: 71,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCmdName(X[0])
		},
	},
	ProdTabEntry{
		String: `CmdWord : Word	<< ast.NewCmdWord(X[0]) >>`,
		Id: "CmdWord",
		NTType: 33,
		Index: 72,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCmdWord(X[0])
		},
	},
	ProdTabEntry{
		String: `CmdPrefix : IoRedirect	<< ast.NewCmdPrefix(nil, X[0], nil) >>`,
		Id: "CmdPrefix",
		NTType: 34,
		Index: 73,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCmdPrefix(nil, X[0], nil)
		},
	},
	ProdTabEntry{
		String: `CmdPrefix : CmdPrefix IoRedirect	<< ast.NewCmdPrefix(nil, X[1], X[0]) >>`,
		Id: "CmdPrefix",
		NTType: 34,
		Index: 74,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCmdPrefix(nil, X[1], X[0])
		},
	},
	ProdTabEntry{
		String: `CmdPrefix : AssignmentWord	<< ast.NewCmdPrefix(X[0], nil, nil) >>`,
		Id: "CmdPrefix",
		NTType: 34,
		Index: 75,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCmdPrefix(X[0], nil, nil)
		},
	},
	ProdTabEntry{
		String: `CmdPrefix : CmdPrefix AssignmentWord	<< ast.NewCmdPrefix(X[1], nil, X[0]) >>`,
		Id: "CmdPrefix",
		NTType: 34,
		Index: 76,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCmdPrefix(X[1], nil, X[0])
		},
	},
	ProdTabEntry{
		String: `CmdSuffix : IoRedirect	<< ast.NewCmdSuffix(ast.Word(""), X[0], nil) >>`,
		Id: "CmdSuffix",
		NTType: 35,
		Index: 77,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCmdSuffix(ast.Word(""), X[0], nil)
		},
	},
	ProdTabEntry{
		String: `CmdSuffix : CmdSuffix IoRedirect	<< ast.NewCmdSuffix(ast.Word(""), X[1], X[0]) >>`,
		Id: "CmdSuffix",
		NTType: 35,
		Index: 78,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCmdSuffix(ast.Word(""), X[1], X[0])
		},
	},
	ProdTabEntry{
		String: `CmdSuffix : Word	<< ast.NewCmdSuffix(X[0], nil, nil) >>`,
		Id: "CmdSuffix",
		NTType: 35,
		Index: 79,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCmdSuffix(X[0], nil, nil)
		},
	},
	ProdTabEntry{
		String: `CmdSuffix : CmdSuffix Word	<< ast.NewCmdSuffix(X[1], nil, X[0]) >>`,
		Id: "CmdSuffix",
		NTType: 35,
		Index: 80,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCmdSuffix(X[1], nil, X[0])
		},
	},
	ProdTabEntry{
		String: `RedirectList : IoRedirect	<< ast.NewRedirectList(X[0]) >>`,
		Id: "RedirectList",
		NTType: 36,
		Index: 81,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewRedirectList(X[0])
		},
	},
	ProdTabEntry{
		String: `RedirectList : RedirectList IoRedirect	<< ast.AppendToRedirectList(X[0], X[1]) >>`,
		Id: "RedirectList",
		NTType: 36,
		Index: 82,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendToRedirectList(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `IoRedirect : IoFile	<< ast.NewIoRedirect(X[0], ast.IoNumber(0), nil) >>`,
		Id: "IoRedirect",
		NTType: 37,
		Index: 83,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIoRedirect(X[0], ast.IoNumber(0), nil)
		},
	},
	ProdTabEntry{
		String: `IoRedirect : IoNumber IoFile	<< ast.NewIoRedirect(X[1], X[0], nil) >>`,
		Id: "IoRedirect",
		NTType: 37,
		Index: 84,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIoRedirect(X[1], X[0], nil)
		},
	},
	ProdTabEntry{
		String: `IoRedirect : IoHere	<< ast.NewIoRedirect(nil, ast.IoNumber(0), X[0]) >>`,
		Id: "IoRedirect",
		NTType: 37,
		Index: 85,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIoRedirect(nil, ast.IoNumber(0), X[0])
		},
	},
	ProdTabEntry{
		String: `IoRedirect : IoNumber IoHere	<< ast.NewIoRedirect(nil, X[0], X[1]) >>`,
		Id: "IoRedirect",
		NTType: 37,
		Index: 86,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIoRedirect(nil, X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `IoFile : "<" Filename	<< ast.NewIoFile(X[1], ast.R_STDIN) >>`,
		Id: "IoFile",
		NTType: 38,
		Index: 87,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIoFile(X[1], ast.R_STDIN)
		},
	},
	ProdTabEntry{
		String: `IoFile : "<&" Filename	<< ast.NewIoFile(X[1], ast.R_INFD) >>`,
		Id: "IoFile",
		NTType: 38,
		Index: 88,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIoFile(X[1], ast.R_INFD)
		},
	},
	ProdTabEntry{
		String: `IoFile : ">" Filename	<< ast.NewIoFile(X[1], ast.R_STDOUT) >>`,
		Id: "IoFile",
		NTType: 38,
		Index: 89,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIoFile(X[1], ast.R_STDOUT)
		},
	},
	ProdTabEntry{
		String: `IoFile : ">&" Filename	<< ast.NewIoFile(X[1], ast.R_OUTFD) >>`,
		Id: "IoFile",
		NTType: 38,
		Index: 90,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIoFile(X[1], ast.R_OUTFD)
		},
	},
	ProdTabEntry{
		String: `IoFile : ">>" Filename	<< ast.NewIoFile(X[1], ast.R_APPEND) >>`,
		Id: "IoFile",
		NTType: 38,
		Index: 91,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIoFile(X[1], ast.R_APPEND)
		},
	},
	ProdTabEntry{
		String: `IoFile : "<>" Filename	<< ast.NewIoFile(X[1], ast.R_ORWFD) >>`,
		Id: "IoFile",
		NTType: 38,
		Index: 92,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIoFile(X[1], ast.R_ORWFD)
		},
	},
	ProdTabEntry{
		String: `IoFile : ">|" Filename	<< ast.NewIoFile(X[1], ast.R_OUTSP) >>`,
		Id: "IoFile",
		NTType: 38,
		Index: 93,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIoFile(X[1], ast.R_OUTSP)
		},
	},
	ProdTabEntry{
		String: `Filename : Word	<< ast.NewFilename(X[0]) >>`,
		Id: "Filename",
		NTType: 39,
		Index: 94,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewFilename(X[0])
		},
	},
	ProdTabEntry{
		String: `IoHere : "<<" HereEnd	<< ast.NewIoHere(X[1], false) >>`,
		Id: "IoHere",
		NTType: 40,
		Index: 95,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIoHere(X[1], false)
		},
	},
	ProdTabEntry{
		String: `IoHere : "<<-" HereEnd	<< ast.NewIoHere(X[1], true) >>`,
		Id: "IoHere",
		NTType: 40,
		Index: 96,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIoHere(X[1], true)
		},
	},
	ProdTabEntry{
		String: `HereEnd : Word	<< ast.NewHereEnd(X[0]) >>`,
		Id: "HereEnd",
		NTType: 41,
		Index: 97,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewHereEnd(X[0])
		},
	},
	ProdTabEntry{
		String: `NewlineList : "\n"	<<  >>`,
		Id: "NewlineList",
		NTType: 42,
		Index: 98,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `NewlineList : NewlineList "\n"	<<  >>`,
		Id: "NewlineList",
		NTType: 42,
		Index: 99,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Linebreak : NewlineList	<<  >>`,
		Id: "Linebreak",
		NTType: 43,
		Index: 100,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Linebreak : nothing	<<  >>`,
		Id: "Linebreak",
		NTType: 43,
		Index: 101,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `SeparatorOp : "&"	<< ast.S_AMP, nil >>`,
		Id: "SeparatorOp",
		NTType: 44,
		Index: 102,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.S_AMP, nil
		},
	},
	ProdTabEntry{
		String: `SeparatorOp : ";"	<< ast.S_SEMICOLON, nil >>`,
		Id: "SeparatorOp",
		NTType: 44,
		Index: 103,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.S_SEMICOLON, nil
		},
	},
	ProdTabEntry{
		String: `Separator : SeparatorOp Linebreak	<< ast.NewSeparator(X[0]) >>`,
		Id: "Separator",
		NTType: 45,
		Index: 104,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewSeparator(X[0])
		},
	},
	ProdTabEntry{
		String: `Separator : NewlineList	<<  >>`,
		Id: "Separator",
		NTType: 45,
		Index: 105,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `SequentialSep : ";" Linebreak	<<  >>`,
		Id: "SequentialSep",
		NTType: 46,
		Index: 106,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `SequentialSep : NewlineList	<<  >>`,
		Id: "SequentialSep",
		NTType: 46,
		Index: 107,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	
}
