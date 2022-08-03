/*Generator Code , do not modify*/
// Code header part

package expr

import "fmt"

// const part
const LE = 104
const EQ = 105
const NE = 101
const GE = 102
const NUM = 100
const ERROR_ACTION = 122
const ACCEPT_ACTION = 222

// Terminal Size
const NTERMINALS = 14

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
	222, 7, 9, 0, 0, 4, 5, 6, 0, 10, 0, 8, 11, 7, 9, 0, 21, 4, 5, 6, 3, 10, 2, 8, 11, 122, 1, 7, 0, 122, 0, 0, 0, 6, 122, 3, 122, 2, 122, 122, 7, 20, 122, 0, 0, 0, 6, 122, 0, 122, -5, 7, 122, 0, -5, 4, 5, 6, -6, 7, 0, 0, -6, 4, 5, 6, -7, 7, 0, 0, -7, 4, 5, 6, -8, 7, 0, 0, -8, 4, 5, 6, 122, 122, 0, 0, 122, 122, 0, 0, 0, 122, 122, 122, 122, 122, 122, 122, 0, 122, 0, 122, 0, 0, 122, 122, 122, 0, 0, 122, 122, 3, 122, 2, 3, 122, 2, 12, 0, 3, 13, 2, 3, 0, 2, 14, 0, 3, 15, 2, 3, 0, 2, 16, 0, 3, 17, 2, 3, 0, 2, 18, 0, 0, 19,
}
var StatePackOffset = []int{
	11, -1, 26, 82, 102, 105, 110, 113, 118, 121, 126, 129, 25, 38, 83, 95, 49, 57, 65, 73, 11, 101,
}
var StackPackCheck = []int{
	1, 1, 1, -1, -1, 1, 1, 1, -1, 1, -1, 1, 1, 20, 20, -1, 20, 20, 20, 20, 0, 20, 0, 20, 20, 12, 0, 12, -1, 12, -1, -1, -1, 12, 12, 2, 12, 2, 13, 12, 13, 2, 13, -1, -1, -1, 13, 13, -1, 13, 16, 16, 13, -1, 16, 16, 16, 16, 17, 17, -1, -1, 17, 17, 17, 17, 18, 18, -1, -1, 18, 18, 18, 18, 19, 19, -1, -1, 19, 19, 19, 19, 3, 14, -1, -1, 3, 14, -1, -1, -1, 3, 14, 3, 14, 15, 3, 14, -1, 15, -1, 21, -1, -1, 15, 21, 15, -1, -1, 15, 21, 4, 21, 4, 5, 21, 5, 4, -1, 6, 5, 6, 7, -1, 7, 6, -1, 8, 7, 8, 9, -1, 9, 8, -1, 10, 9, 10, 11, -1, 11, 10, -1, -1, 11,
}
var StackPackActDef = []int{
	122, 122, 122, -10, 122, 122, 122, 122, 122, 122, 122, 122, -1, -2, -3, -4, 122, 122, 122, 122, 122, -9,
}
var StackPackGotoDef = []int{
	122,
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
		dollarDolar.YySymIndex = 15
		Dollar := StateSymStack[topIndex-3 : StackPointer]
		_ = Dollar

		/*

		   LineNo:24
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
		dollarDolar.YySymIndex = 15
		Dollar := StateSymStack[topIndex-3 : StackPointer]
		_ = Dollar

		/*

		   LineNo:27
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
		dollarDolar.YySymIndex = 15
		Dollar := StateSymStack[topIndex-3 : StackPointer]
		_ = Dollar

		/*

		   LineNo:30
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
		dollarDolar.YySymIndex = 15
		Dollar := StateSymStack[topIndex-3 : StackPointer]
		_ = Dollar

		/*

		   LineNo:33
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
		dollarDolar.YySymIndex = 15
		Dollar := StateSymStack[topIndex-3 : StackPointer]
		_ = Dollar

		/*

		   LineNo:36
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
		dollarDolar.YySymIndex = 15
		Dollar := StateSymStack[topIndex-3 : StackPointer]
		_ = Dollar

		/*

		   LineNo:43
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
		dollarDolar.YySymIndex = 15
		Dollar := StateSymStack[topIndex-3 : StackPointer]
		_ = Dollar

		/*

		   LineNo:50
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
		dollarDolar.YySymIndex = 15
		Dollar := StateSymStack[topIndex-3 : StackPointer]
		_ = Dollar

		/*

		   LineNo:57
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
		dollarDolar.YySymIndex = 15
		Dollar := StateSymStack[topIndex-3 : StackPointer]
		_ = Dollar

		/*

		   LineNo:64
		   E -> '('  E ')'
		    {
		   		$$	=	$2
		   	}
		*/
		{
			dollarDolar.val = Dollar[2].val
		}
		PopStateSym(3)
	case 10:
		dollarDolar.YySymIndex = 15
		Dollar := StateSymStack[topIndex-1 : StackPointer]
		_ = Dollar

		/*

		   LineNo:67
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
	case 47:
		conv = 2
	case 60:
		conv = 3
	case 101:
		conv = 4
	case 41:
		conv = 5
	case 43:
		conv = 6
	case 45:
		conv = 7
	case 42:
		conv = 8
	case 100:
		conv = 9
	case 102:
		conv = 10
	case 40:
		conv = 11
	case 62:
		conv = 12
	case 104:
		conv = 13
	case 105:
		conv = 14

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
		conv = "'/' "
	case 3:
		conv = "'<' "
	case 4:
		conv = "NE"
	case 5:
		conv = "')' "
	case 6:
		conv = "'+' "
	case 7:
		conv = "'-' "
	case 8:
		conv = "'*' "
	case 9:
		conv = "NUM"
	case 10:
		conv = "GE"
	case 11:
		conv = "'(' "
	case 12:
		conv = "'>' "
	case 13:
		conv = "LE"
	case 14:
		conv = "EQ"
	case 15:
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

			fmt.Printf("look ahead %s, use Reduce:E -> '('  E ')'  , go to state %d\n", look, s)
		case 10:

			fmt.Printf("look ahead %s, use Reduce:E -> NUM , go to state %d\n", look, s)

		}
	}
}

// Code Last part
