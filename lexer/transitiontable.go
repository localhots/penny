
package lexer



/*
Let s be the current state
Let r be the current input rune
transitionTable[s](r) returns the next state.
*/
type TransitionTable [NumStates] func(rune) int

var TransTab = TransitionTable{
	
		// S0
		func(r rune) int {
			switch {
			case r == 9 : // ['\t','\t']
				return 1
			case r == 10 : // ['\n','\n']
				return 1
			case r == 13 : // ['\r','\r']
				return 1
			case r == 32 : // [' ',' ']
				return 1
			case r == 33 : // ['!','!']
				return 2
			case r == 38 : // ['&','&']
				return 3
			case r == 40 : // ['(','(']
				return 4
			case r == 41 : // [')',')']
				return 5
			case 48 <= r && r <= 57 : // ['0','9']
				return 6
			case r == 59 : // [';',';']
				return 7
			case r == 60 : // ['<','<']
				return 8
			case r == 61 : // ['=','=']
				return 9
			case r == 62 : // ['>','>']
				return 10
			case r == 92 : // ['\','\']
				return 11
			case r == 95 : // ['_','_']
				return 12
			case 97 <= r && r <= 98 : // ['a','b']
				return 13
			case r == 99 : // ['c','c']
				return 14
			case r == 100 : // ['d','d']
				return 15
			case r == 101 : // ['e','e']
				return 16
			case r == 102 : // ['f','f']
				return 17
			case 103 <= r && r <= 104 : // ['g','h']
				return 13
			case r == 105 : // ['i','i']
				return 18
			case 106 <= r && r <= 115 : // ['j','s']
				return 13
			case r == 116 : // ['t','t']
				return 19
			case r == 117 : // ['u','u']
				return 20
			case r == 118 : // ['v','v']
				return 13
			case r == 119 : // ['w','w']
				return 21
			case 120 <= r && r <= 122 : // ['x','z']
				return 13
			case r == 123 : // ['{','{']
				return 22
			case r == 124 : // ['|','|']
				return 23
			case r == 125 : // ['}','}']
				return 24
			
			
			default:
				return 25
			}
			
		},
	
		// S1
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S2
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S3
		func(r rune) int {
			switch {
			case r == 38 : // ['&','&']
				return 26
			
			
			
			}
			return NoState
			
		},
	
		// S4
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S5
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S6
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 6
			
			
			
			}
			return NoState
			
		},
	
		// S7
		func(r rune) int {
			switch {
			case r == 59 : // [';',';']
				return 27
			
			
			
			}
			return NoState
			
		},
	
		// S8
		func(r rune) int {
			switch {
			case r == 38 : // ['&','&']
				return 28
			case r == 60 : // ['<','<']
				return 29
			case r == 62 : // ['>','>']
				return 30
			
			
			
			}
			return NoState
			
		},
	
		// S9
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S10
		func(r rune) int {
			switch {
			case r == 38 : // ['&','&']
				return 31
			case r == 62 : // ['>','>']
				return 32
			case r == 124 : // ['|','|']
				return 33
			
			
			
			}
			return NoState
			
		},
	
		// S11
		func(r rune) int {
			switch {
			case r == 110 : // ['n','n']
				return 34
			
			
			
			}
			return NoState
			
		},
	
		// S12
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 35
			case r == 95 : // ['_','_']
				return 12
			case 97 <= r && r <= 122 : // ['a','z']
				return 36
			
			
			
			}
			return NoState
			
		},
	
		// S13
		func(r rune) int {
			switch {
			case 65 <= r && r <= 90 : // ['A','Z']
				return 37
			
			
			
			}
			return NoState
			
		},
	
		// S14
		func(r rune) int {
			switch {
			case 65 <= r && r <= 90 : // ['A','Z']
				return 37
			case r == 97 : // ['a','a']
				return 38
			
			
			
			}
			return NoState
			
		},
	
		// S15
		func(r rune) int {
			switch {
			case 65 <= r && r <= 90 : // ['A','Z']
				return 37
			case r == 111 : // ['o','o']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S16
		func(r rune) int {
			switch {
			case 65 <= r && r <= 90 : // ['A','Z']
				return 37
			case r == 108 : // ['l','l']
				return 40
			case r == 115 : // ['s','s']
				return 41
			
			
			
			}
			return NoState
			
		},
	
		// S17
		func(r rune) int {
			switch {
			case 65 <= r && r <= 90 : // ['A','Z']
				return 37
			case r == 105 : // ['i','i']
				return 42
			case r == 111 : // ['o','o']
				return 43
			
			
			
			}
			return NoState
			
		},
	
		// S18
		func(r rune) int {
			switch {
			case 65 <= r && r <= 90 : // ['A','Z']
				return 37
			case r == 102 : // ['f','f']
				return 44
			case r == 110 : // ['n','n']
				return 45
			
			
			
			}
			return NoState
			
		},
	
		// S19
		func(r rune) int {
			switch {
			case 65 <= r && r <= 90 : // ['A','Z']
				return 37
			case r == 104 : // ['h','h']
				return 46
			
			
			
			}
			return NoState
			
		},
	
		// S20
		func(r rune) int {
			switch {
			case 65 <= r && r <= 90 : // ['A','Z']
				return 37
			case r == 110 : // ['n','n']
				return 47
			
			
			
			}
			return NoState
			
		},
	
		// S21
		func(r rune) int {
			switch {
			case 65 <= r && r <= 90 : // ['A','Z']
				return 37
			case r == 104 : // ['h','h']
				return 48
			
			
			
			}
			return NoState
			
		},
	
		// S22
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S23
		func(r rune) int {
			switch {
			case r == 124 : // ['|','|']
				return 49
			
			
			
			}
			return NoState
			
		},
	
		// S24
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S25
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S26
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S27
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S28
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S29
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 50
			
			
			
			}
			return NoState
			
		},
	
		// S30
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S31
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S32
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S33
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S34
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S35
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 35
			case r == 95 : // ['_','_']
				return 12
			case 97 <= r && r <= 122 : // ['a','z']
				return 36
			
			
			
			}
			return NoState
			
		},
	
		// S36
		func(r rune) int {
			switch {
			case 65 <= r && r <= 90 : // ['A','Z']
				return 51
			
			
			
			}
			return NoState
			
		},
	
		// S37
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 35
			case r == 95 : // ['_','_']
				return 52
			case 97 <= r && r <= 122 : // ['a','z']
				return 53
			
			
			
			}
			return NoState
			
		},
	
		// S38
		func(r rune) int {
			switch {
			case r == 115 : // ['s','s']
				return 54
			
			
			
			}
			return NoState
			
		},
	
		// S39
		func(r rune) int {
			switch {
			case r == 110 : // ['n','n']
				return 55
			
			
			
			}
			return NoState
			
		},
	
		// S40
		func(r rune) int {
			switch {
			case r == 105 : // ['i','i']
				return 56
			case r == 115 : // ['s','s']
				return 57
			
			
			
			}
			return NoState
			
		},
	
		// S41
		func(r rune) int {
			switch {
			case r == 97 : // ['a','a']
				return 58
			
			
			
			}
			return NoState
			
		},
	
		// S42
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S43
		func(r rune) int {
			switch {
			case r == 114 : // ['r','r']
				return 59
			
			
			
			}
			return NoState
			
		},
	
		// S44
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S45
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S46
		func(r rune) int {
			switch {
			case r == 101 : // ['e','e']
				return 60
			
			
			
			}
			return NoState
			
		},
	
		// S47
		func(r rune) int {
			switch {
			case r == 116 : // ['t','t']
				return 61
			
			
			
			}
			return NoState
			
		},
	
		// S48
		func(r rune) int {
			switch {
			case r == 105 : // ['i','i']
				return 62
			
			
			
			}
			return NoState
			
		},
	
		// S49
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S50
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S51
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 35
			case r == 95 : // ['_','_']
				return 12
			case 97 <= r && r <= 122 : // ['a','z']
				return 36
			
			
			
			}
			return NoState
			
		},
	
		// S52
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 35
			case r == 95 : // ['_','_']
				return 52
			case 97 <= r && r <= 122 : // ['a','z']
				return 53
			
			
			
			}
			return NoState
			
		},
	
		// S53
		func(r rune) int {
			switch {
			case 65 <= r && r <= 90 : // ['A','Z']
				return 37
			
			
			
			}
			return NoState
			
		},
	
		// S54
		func(r rune) int {
			switch {
			case r == 101 : // ['e','e']
				return 63
			
			
			
			}
			return NoState
			
		},
	
		// S55
		func(r rune) int {
			switch {
			case r == 101 : // ['e','e']
				return 64
			
			
			
			}
			return NoState
			
		},
	
		// S56
		func(r rune) int {
			switch {
			case r == 102 : // ['f','f']
				return 65
			
			
			
			}
			return NoState
			
		},
	
		// S57
		func(r rune) int {
			switch {
			case r == 101 : // ['e','e']
				return 66
			
			
			
			}
			return NoState
			
		},
	
		// S58
		func(r rune) int {
			switch {
			case r == 99 : // ['c','c']
				return 67
			
			
			
			}
			return NoState
			
		},
	
		// S59
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S60
		func(r rune) int {
			switch {
			case r == 110 : // ['n','n']
				return 68
			
			
			
			}
			return NoState
			
		},
	
		// S61
		func(r rune) int {
			switch {
			case r == 105 : // ['i','i']
				return 69
			
			
			
			}
			return NoState
			
		},
	
		// S62
		func(r rune) int {
			switch {
			case r == 108 : // ['l','l']
				return 70
			
			
			
			}
			return NoState
			
		},
	
		// S63
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S64
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S65
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S66
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S67
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S68
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S69
		func(r rune) int {
			switch {
			case r == 108 : // ['l','l']
				return 71
			
			
			
			}
			return NoState
			
		},
	
		// S70
		func(r rune) int {
			switch {
			case r == 101 : // ['e','e']
				return 72
			
			
			
			}
			return NoState
			
		},
	
		// S71
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S72
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
}
