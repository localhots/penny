/*
 */
package parser

const numNTSymbols = 47

type (
	gotoTable [numStates]gotoRow
	gotoRow   [numNTSymbols]int
)

var gotoTab = gotoTable{
	gotoRow{ // S0

		-1, // S'
		1,  // Word
		-1, // IoNumber
		-1, // AssignmentWord
		-1, // List
		-1, // AndOr
		-1, // Pipeline
		-1, // PipeSequence
		-1, // Command
		-1, // CompoundCommand
		-1, // Subshell
		-1, // CompoundList
		-1, // Term
		-1, // ForClause
		-1, // Name
		-1, // Wordlist
		-1, // CaseClause
		-1, // CaseListNs
		-1, // CaseList
		-1, // CaseItemNs
		-1, // CaseItem
		-1, // Pattern
		-1, // IfClause
		-1, // ElsePart
		-1, // WhileClause
		-1, // UntilClause
		-1, // FunctionDefinition
		-1, // FunctionBody
		-1, // FunctionName
		-1, // BraceGroup
		-1, // DoGroup
		-1, // SimpleCommand
		-1, // CmdName
		-1, // CmdWord
		-1, // CmdPrefix
		-1, // CmdSuffix
		-1, // RedirectList
		-1, // IoRedirect
		-1, // IoFile
		-1, // Filename
		-1, // IoHere
		-1, // HereEnd
		-1, // NewlineList
		-1, // Linebreak
		-1, // SeparatorOp
		-1, // Separator
		-1, // SequentialSep

	},
	gotoRow{ // S1

		-1, // S'
		-1, // Word
		-1, // IoNumber
		-1, // AssignmentWord
		-1, // List
		-1, // AndOr
		-1, // Pipeline
		-1, // PipeSequence
		-1, // Command
		-1, // CompoundCommand
		-1, // Subshell
		-1, // CompoundList
		-1, // Term
		-1, // ForClause
		-1, // Name
		-1, // Wordlist
		-1, // CaseClause
		-1, // CaseListNs
		-1, // CaseList
		-1, // CaseItemNs
		-1, // CaseItem
		-1, // Pattern
		-1, // IfClause
		-1, // ElsePart
		-1, // WhileClause
		-1, // UntilClause
		-1, // FunctionDefinition
		-1, // FunctionBody
		-1, // FunctionName
		-1, // BraceGroup
		-1, // DoGroup
		-1, // SimpleCommand
		-1, // CmdName
		-1, // CmdWord
		-1, // CmdPrefix
		-1, // CmdSuffix
		-1, // RedirectList
		-1, // IoRedirect
		-1, // IoFile
		-1, // Filename
		-1, // IoHere
		-1, // HereEnd
		-1, // NewlineList
		-1, // Linebreak
		-1, // SeparatorOp
		-1, // Separator
		-1, // SequentialSep

	},
	gotoRow{ // S2

		-1, // S'
		-1, // Word
		-1, // IoNumber
		-1, // AssignmentWord
		-1, // List
		-1, // AndOr
		-1, // Pipeline
		-1, // PipeSequence
		-1, // Command
		-1, // CompoundCommand
		-1, // Subshell
		-1, // CompoundList
		-1, // Term
		-1, // ForClause
		-1, // Name
		-1, // Wordlist
		-1, // CaseClause
		-1, // CaseListNs
		-1, // CaseList
		-1, // CaseItemNs
		-1, // CaseItem
		-1, // Pattern
		-1, // IfClause
		-1, // ElsePart
		-1, // WhileClause
		-1, // UntilClause
		-1, // FunctionDefinition
		-1, // FunctionBody
		-1, // FunctionName
		-1, // BraceGroup
		-1, // DoGroup
		-1, // SimpleCommand
		-1, // CmdName
		-1, // CmdWord
		-1, // CmdPrefix
		-1, // CmdSuffix
		-1, // RedirectList
		-1, // IoRedirect
		-1, // IoFile
		-1, // Filename
		-1, // IoHere
		-1, // HereEnd
		-1, // NewlineList
		-1, // Linebreak
		-1, // SeparatorOp
		-1, // Separator
		-1, // SequentialSep

	},
}
