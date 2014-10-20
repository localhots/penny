
package parser

type(
	actionTable [numStates]actionRow
	actionRow struct {
		canRecover bool
		actions [numSymbols]action
	}
)

var actionTab = actionTable{
	actionRow{ // S0
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(2),		/* word */
			nil,		/* number */
			nil,		/* = */
			nil,		/* && */
			nil,		/* || */
			nil,		/* ! */
			nil,		/* | */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* for */
			nil,		/* in */
			nil,		/* name */
			nil,		/* case */
			nil,		/* esac */
			nil,		/* ;; */
			nil,		/* if */
			nil,		/* then */
			nil,		/* fi */
			nil,		/* elif */
			nil,		/* else */
			nil,		/* while */
			nil,		/* until */
			nil,		/* { */
			nil,		/* } */
			nil,		/* do */
			nil,		/* done */
			nil,		/* < */
			nil,		/* <& */
			nil,		/* > */
			nil,		/* >& */
			nil,		/* >> */
			nil,		/* <> */
			nil,		/* >| */
			nil,		/* << */
			nil,		/* <<- */
			nil,		/* \n */
			nil,		/* nothing */
			nil,		/* & */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S1
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			accept(true),		/* $ */
			nil,		/* word */
			nil,		/* number */
			nil,		/* = */
			nil,		/* && */
			nil,		/* || */
			nil,		/* ! */
			nil,		/* | */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* for */
			nil,		/* in */
			nil,		/* name */
			nil,		/* case */
			nil,		/* esac */
			nil,		/* ;; */
			nil,		/* if */
			nil,		/* then */
			nil,		/* fi */
			nil,		/* elif */
			nil,		/* else */
			nil,		/* while */
			nil,		/* until */
			nil,		/* { */
			nil,		/* } */
			nil,		/* do */
			nil,		/* done */
			nil,		/* < */
			nil,		/* <& */
			nil,		/* > */
			nil,		/* >& */
			nil,		/* >> */
			nil,		/* <> */
			nil,		/* >| */
			nil,		/* << */
			nil,		/* <<- */
			nil,		/* \n */
			nil,		/* nothing */
			nil,		/* & */
			nil,		/* ; */
			
		},

	},
	actionRow{ // S2
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(1),		/* $, reduce: Word */
			nil,		/* word */
			nil,		/* number */
			nil,		/* = */
			nil,		/* && */
			nil,		/* || */
			nil,		/* ! */
			nil,		/* | */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* for */
			nil,		/* in */
			nil,		/* name */
			nil,		/* case */
			nil,		/* esac */
			nil,		/* ;; */
			nil,		/* if */
			nil,		/* then */
			nil,		/* fi */
			nil,		/* elif */
			nil,		/* else */
			nil,		/* while */
			nil,		/* until */
			nil,		/* { */
			nil,		/* } */
			nil,		/* do */
			nil,		/* done */
			nil,		/* < */
			nil,		/* <& */
			nil,		/* > */
			nil,		/* >& */
			nil,		/* >> */
			nil,		/* <> */
			nil,		/* >| */
			nil,		/* << */
			nil,		/* <<- */
			nil,		/* \n */
			nil,		/* nothing */
			nil,		/* & */
			nil,		/* ; */
			
		},

	},
	
}

