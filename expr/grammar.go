/*Generator Code , do not modify*/
// Code header part

package expr

import "fmt"

// const part
const GE = 101
const EQ = 102
const NE = 103
const AND = 104
const OR = 105
const NUM = 100
const LE = 107
const ERROR_ACTION = 126
const ACCEPT_ACTION = 226

// Terminal Size
const NTERMINALS = 17

var IsTrace bool = false
var StateSymStack = []StateSym{}
var StackPointer = 0

type Context struct {
	StackSym []StateSym
	Stackpos int
}

var globalContext = []Context{}

type ValType struct {
	// Union part

	val float64
}
type StateSym struct {
	Yystate int // state

	//sym val
	YySymIndex int
	//other
	ValType
}

// It is NeedPacked

var StatePackAction = []int{
	126, -9, -9, 11, 126, 10, 126, 126, 8, 9, 126, -9, -9, 4, 7, 126, 5, 6, -10, -10, 11, 0, 10, 0, 0, 8, 9, 0, 12, -10, 4, 7, 0, 5, 6, 226, 0, 11, 0, 10, 0, 0, 8, 9, 0, 12, 13, 4, 7, 2, 5, 6, 25, 11, 0, 10, 0, 3, 8, 9, 1, 12, 13, 4, 7, 126, 5, 6, 0, 126, 0, 126, 126, 0, 0, 126, 0, 126, 0, 7, 126, 126, 6, 126, 126, -5, -5, 126, 0, 0, 0, 7, 126, 0, 6, -5, -5, 4, 7, 0, 5, 6, -6, -6, 0, 0, 0, 0, 0, -7, -7, 0, -6, -6, 4, 7, 0, 5, 6, -7, -7, 4, 7, 2, 5, 6, -8, -8, 0, 0, 0, 3, 0, 0, 24, 0, -8, -8, 4, 7, 126, 5, 6, 2, 126, 0, 126, 126, 0, 2, 126, 3, 126, 0, 14, 126, 126, 3, 126, 126, 15, 2, 126, 0, 126, 2, 0, 126, 126, 3, 126, 126, 16, 3, 126, 2, 17, 2, 0, 126, 0, 2, 2, 3, 0, 3, 18, 2, 19, 3, 3, 2, 20, 21, 0, 3, 0, 0, 22, 3, 0, 0, 23,
}
var StatePackOffset = []int{
	42, 34, 116, 0, 136, 142, 154, 158, 168, 170, 174, 175, 180, 184, 65, 77, 140, 152, 84, 101, 108, 125, 0, 17, 50, 164,
}
var StackPackCheck = []int{
	3, 22, 22, 22, 3, 22, 3, 3, 22, 22, 3, 22, 22, 22, 22, 3, 22, 22, 23, 23, 23, -1, 23, -1, -1, 23, 23, -1, 23, 23, 23, 23, -1, 23, 23, 1, -1, 1, -1, 1, -1, -1, 1, 1, -1, 1, 1, 1, 1, 0, 1, 1, 24, 24, -1, 24, -1, 0, 24, 24, 0, 24, 24, 24, 24, 14, 24, 24, -1, 14, -1, 14, 14, -1, -1, 14, -1, 15, -1, 14, 14, 15, 14, 15, 15, 18, 18, 15, -1, -1, -1, 15, 15, -1, 15, 18, 18, 18, 18, -1, 18, 18, 19, 19, -1, -1, -1, -1, -1, 20, 20, -1, 19, 19, 19, 19, -1, 19, 19, 20, 20, 20, 20, 2, 20, 20, 21, 21, -1, -1, -1, 2, -1, -1, 2, -1, 21, 21, 21, 21, 16, 21, 21, 4, 16, -1, 16, 16, -1, 5, 16, 4, 17, -1, 4, 16, 17, 5, 17, 17, 5, 6, 17, -1, 25, 7, -1, 17, 25, 6, 25, 25, 6, 7, 25, 8, 7, 9, -1, 25, -1, 10, 11, 8, -1, 9, 8, 12, 9, 10, 11, 13, 10, 11, -1, 12, -1, -1, 12, 13, -1, -1, 13,
}
var StackPackActDef = []int{
	126, 126, 126, -12, 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, -1, -2, -3, -4, 126, 126, 126, 126, 126, 126, 126, -11,
}
var StackPackGotoDef = []int{
	126,
}

func (s *StateSym) Action(a int) int {
	if StatePackOffset[s.Yystate]+a < 0 {
		return ERROR_ACTION
	}
	if StatePackOffset[s.Yystate]+a >= len(StackPackCheck) ||
		StackPackCheck[StatePackOffset[s.Yystate]+a] != s.Yystate {
		if a > NTERMINALS {
			return StackPackGotoDef[a-NTERMINALS-1]
		} else {
			return StackPackActDef[s.Yystate]
		}
	} else {
		return StatePackAction[StatePackOffset[s.Yystate]+a]
	}
}

func TraceShift(s *StateSym) {
	if IsTrace {
		fmt.Printf("Shift %s, push state %d\n", TraceTranslate(s.YySymIndex), s.Yystate)
	}
}

// Reduce function
func ReduceFunc(reduceIndex int) *StateSym {
	dollarDolar := &StateSym{}
	topIndex := StackPointer - 1
	switch reduceIndex {
	case 1:
		dollarDolar.YySymIndex = 18
		Dollar := StateSymStack[topIndex-3 : StackPointer]
		_ = Dollar

		/*

		   LineNo:27
		   E -> E '+'  E
		    {
		   		$$	=	$1 + $3
		   	}
		*/
		{
			dollarDolar.val = Dollar[1].val + Dollar[3].val
		}
		PopStateSym(3)
	case 2:
		dollarDolar.YySymIndex = 18
		Dollar := StateSymStack[topIndex-3 : StackPointer]
		_ = Dollar

		/*

		   LineNo:30
		   E -> E '-'  E
		    {
		   		$$	=	$1 - $3
		   	}
		*/
		{
			dollarDolar.val = Dollar[1].val - Dollar[3].val
		}
		PopStateSym(3)
	case 3:
		dollarDolar.YySymIndex = 18
		Dollar := StateSymStack[topIndex-3 : StackPointer]
		_ = Dollar

		/*

		   LineNo:33
		   E -> E '*'  E
		    {
		   		$$	=	$1 * $3
		   	}
		*/
		{
			dollarDolar.val = Dollar[1].val * Dollar[3].val
		}
		PopStateSym(3)
	case 4:
		dollarDolar.YySymIndex = 18
		Dollar := StateSymStack[topIndex-3 : StackPointer]
		_ = Dollar

		/*

		   LineNo:36
		   E -> E '/'  E
		    {
		   		$$	=	$1 / $3
		   	}
		*/
		{
			dollarDolar.val = Dollar[1].val / Dollar[3].val
		}
		PopStateSym(3)
	case 5:
		dollarDolar.YySymIndex = 18
		Dollar := StateSymStack[topIndex-3 : StackPointer]
		_ = Dollar

		/*

		   LineNo:39
		   E -> E '>'  E
		    {
		   		if $1 > $3 {
		   			$$	=	1
		   		} else {
		   			$$	=	0
		   		}
		   	}
		*/
		{
			if Dollar[1].val > Dollar[3].val {
				dollarDolar.val = 1
			} else {
				dollarDolar.val = 0
			}
		}
		PopStateSym(3)
	case 6:
		dollarDolar.YySymIndex = 18
		Dollar := StateSymStack[topIndex-3 : StackPointer]
		_ = Dollar

		/*

		   LineNo:46
		   E -> E '<'  E
		    {
		   		if $1 < $3 {
		   			$$	=	1
		   		} else {
		   			$$	=	0
		   		}
		   	}
		*/
		{
			if Dollar[1].val < Dollar[3].val {
				dollarDolar.val = 1
			} else {
				dollarDolar.val = 0
			}
		}
		PopStateSym(3)
	case 7:
		dollarDolar.YySymIndex = 18
		Dollar := StateSymStack[topIndex-3 : StackPointer]
		_ = Dollar

		/*

		   LineNo:53
		   E -> E GE E
		    {
		   		if $1 >= $3 {
		   			$$	=	1
		   		} else {
		   			$$	=	0
		   		}
		   	}
		*/
		{
			if Dollar[1].val >= Dollar[3].val {
				dollarDolar.val = 1
			} else {
				dollarDolar.val = 0
			}
		}
		PopStateSym(3)
	case 8:
		dollarDolar.YySymIndex = 18
		Dollar := StateSymStack[topIndex-3 : StackPointer]
		_ = Dollar

		/*

		   LineNo:60
		   E -> E LE E
		    {
		   		if $1 <= $3 {
		   			$$	=	1
		   		} else {
		   			$$	=	0
		   		}
		   	}
		*/
		{
			if Dollar[1].val <= Dollar[3].val {
				dollarDolar.val = 1
			} else {
				dollarDolar.val = 0
			}
		}
		PopStateSym(3)
	case 9:
		dollarDolar.YySymIndex = 18
		Dollar := StateSymStack[topIndex-3 : StackPointer]
		_ = Dollar

		/*

		   LineNo:67
		   E -> E AND E
		    {
		   		if $1 >0 && $3 >0 {
		   			$$	=	1
		   		} else {
		   			$$	=	0
		   		}
		   	}
		*/
		{
			if Dollar[1].val > 0 && Dollar[3].val > 0 {
				dollarDolar.val = 1
			} else {
				dollarDolar.val = 0
			}
		}
		PopStateSym(3)
	case 10:
		dollarDolar.YySymIndex = 18
		Dollar := StateSymStack[topIndex-3 : StackPointer]
		_ = Dollar

		/*

		   LineNo:74
		   E -> E OR E
		    {
		   		if $1 > 0 || $3 > 0{
		   			$$	=	1
		   		} else {
		   			$$	=	0
		   		}
		   	}
		*/
		{
			if Dollar[1].val > 0 || Dollar[3].val > 0 {
				dollarDolar.val = 1
			} else {
				dollarDolar.val = 0
			}
		}
		PopStateSym(3)
	case 11:
		dollarDolar.YySymIndex = 18
		Dollar := StateSymStack[topIndex-3 : StackPointer]
		_ = Dollar

		/*

		   LineNo:81
		   E -> '('  E ')'
		    {
		   		$$	=	$2
		   	}
		*/
		{
			dollarDolar.val = Dollar[2].val
		}
		PopStateSym(3)
	case 12:
		dollarDolar.YySymIndex = 18
		Dollar := StateSymStack[topIndex-1 : StackPointer]
		_ = Dollar

		/*

		   LineNo:84
		   E -> NUM
		    {
		   		$$	=	$1
		   	}
		*/
		{
			dollarDolar.val = Dollar[1].val
		}
		PopStateSym(1)

	}
	return dollarDolar
}

// Push StateSym
func PushStateSym(state *StateSym) {
	TraceShift(state)
	if StackPointer >= len(StateSymStack) {
		StateSymStack = append(StateSymStack, *state)
	} else {
		StateSymStack[StackPointer] = *state
	}
	StackPointer++
}

// Pop num StateSym
func PopStateSym(num int) {
	StackPointer -= num
}

func init() {
	ParserInit()
}
func PushContex() {
	globalContext = append(globalContext, Context{
		StackSym: StateSymStack,
		Stackpos: StackPointer,
	})
}
func PopContex() {
	StackPointer = globalContext[len(globalContext)-1].Stackpos
	StateSymStack = globalContext[len(globalContext)-1].StackSym
	globalContext = globalContext[:len(globalContext)-1]
}
func ParserInit() {
	StateSymStack = []StateSym{
		{
			Yystate:    0,
			YySymIndex: 1, //$
		},
	}
	StackPointer = 1
}

func Parser(input string) *ValType {
	var currentPos int = 0
	var val ValType
	lookAhead := fetchLookAhead(input, &val, &currentPos)
	for {

		if StackPointer == 0 {
			break
		}
		if StackPointer > len(StateSymStack) {
			break
		}
		s := &StateSymStack[StackPointer-1]
		a := s.Action(lookAhead)

		if a == ERROR_ACTION {
			panic(fmt.Sprintf("Grammar error near pos %d", currentPos) + ":" + TraceTranslate(lookAhead))
		} else if a == ACCEPT_ACTION {
			return &s.ValType
		} else {
			if a > 0 {
				// shift
				PushStateSym(&StateSym{
					Yystate:    a,
					YySymIndex: lookAhead,
					ValType:    val,
				})
				lookAhead = fetchLookAhead(input, &val, &currentPos)
			} else {
				reduceIndex := -a
				SymTy := ReduceFunc(reduceIndex)
				s := &StateSymStack[StackPointer-1]
				gotoState := s.Action(SymTy.YySymIndex)
				SymTy.Yystate = gotoState
				TraceReduce(reduceIndex, gotoState, TraceTranslate(lookAhead))
				PushStateSym(SymTy)
			}
		}
	}
	return nil
}
func fetchLookAhead(input string, val *ValType, pos *int) int {
	token := GetToken(input, val, pos)
	return translate(token)
}
func translate(c int) int {
	var conv int = 0
	switch c {
	case -1:
		conv = 1
	case 41:
		conv = 2
	case 107:
		conv = 3
	case 33:
		conv = 4
	case 101:
		conv = 5
	case 102:
		conv = 6
	case 40:
		conv = 7
	case 62:
		conv = 8
	case 60:
		conv = 9
	case 103:
		conv = 10
	case 104:
		conv = 11
	case 105:
		conv = 12
	case 43:
		conv = 13
	case 47:
		conv = 14
	case 100:
		conv = 15
	case 45:
		conv = 16
	case 42:
		conv = 17

	}
	return conv
}

// Trace function for translate
func TraceTranslate(c int) string {
	var conv string = ""
	switch c {
	case 0:
		conv = "start"
	case 1:
		conv = "$"
	case 2:
		conv = "')' "
	case 3:
		conv = "LE"
	case 4:
		conv = "'!' "
	case 5:
		conv = "GE"
	case 6:
		conv = "EQ"
	case 7:
		conv = "'(' "
	case 8:
		conv = "'>' "
	case 9:
		conv = "'<' "
	case 10:
		conv = "NE"
	case 11:
		conv = "AND"
	case 12:
		conv = "OR"
	case 13:
		conv = "'+' "
	case 14:
		conv = "'/' "
	case 15:
		conv = "NUM"
	case 16:
		conv = "'-' "
	case 17:
		conv = "'*' "
	case 18:
		conv = "E"

	}
	return conv
}

// Trace function for reduce
func TraceReduce(reduceIndex, s int, look string) {
	if IsTrace {
		switch reduceIndex {
		case 1:

			fmt.Printf("look ahead %s, use Reduce:E -> E '+'  E , go to state %d\n", look, s)
		case 2:

			fmt.Printf("look ahead %s, use Reduce:E -> E '-'  E , go to state %d\n", look, s)
		case 3:

			fmt.Printf("look ahead %s, use Reduce:E -> E '*'  E , go to state %d\n", look, s)
		case 4:

			fmt.Printf("look ahead %s, use Reduce:E -> E '/'  E , go to state %d\n", look, s)
		case 5:

			fmt.Printf("look ahead %s, use Reduce:E -> E '>'  E , go to state %d\n", look, s)
		case 6:

			fmt.Printf("look ahead %s, use Reduce:E -> E '<'  E , go to state %d\n", look, s)
		case 7:

			fmt.Printf("look ahead %s, use Reduce:E -> E GE E , go to state %d\n", look, s)
		case 8:

			fmt.Printf("look ahead %s, use Reduce:E -> E LE E , go to state %d\n", look, s)
		case 9:

			fmt.Printf("look ahead %s, use Reduce:E -> E AND E , go to state %d\n", look, s)
		case 10:

			fmt.Printf("look ahead %s, use Reduce:E -> E OR E , go to state %d\n", look, s)
		case 11:

			fmt.Printf("look ahead %s, use Reduce:E -> '('  E ')'  , go to state %d\n", look, s)
		case 12:

			fmt.Printf("look ahead %s, use Reduce:E -> NUM , go to state %d\n", look, s)

		}
	}
}

// Code Last part
