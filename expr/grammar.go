
/*Generator Code , do not modify*/
// Code header part 

package expr
import "fmt"




// const part
const GE = 101
const NUM = 100
const LE = 102
const ERROR_ACTION = 120
const ACCEPT_ACTION = 220


// Terminal Size
const NTERMINALS = 11


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
 
var StatePackAction = []int {
	120,	220,	120,	10,	4,	5,	120,	6,	7,	8,	0,	9,	10,	4,	5,	0,	6,	7,	8,	19,	9,	120,	0,	120,	0,	0,	5,	120,	6,	-4,	0,	0,	4,	5,	0,	6,	-5,	0,	-4,	4,	5,	0,	6,	-6,	0,	-5,	4,	5,	0,	6,	-7,	3,	-6,	4,	5,	2,	6,	0,	3,	-7,	3,	1,	2,	3,	2,	3,	0,	2,	18,	2,	11,	0,	3,	12,	3,	13,	2,	3,	2,	3,	0,	2,	14,	2,	15,	0,	120,	16,	120,	17,	0,	120,	120,	120,	120,	0,	120,	120,	0,	0,	120,	 
}
var StatePackOffset = []int {
	49,	0,	56,	0,	58,	61,	63,	70,	72,	75,	77,	21,	86,	91,	28,	35,	42,	49,	9,	94,	
}
var StackPackCheck = []int {
	3,	1,	3,	1,	1,	1,	3,	1,	1,	1,	-1,	1,	18,	18,	18,	-1,	18,	18,	18,	18,	18,	11,	-1,	11,	-1,	-1,	11,	11,	11,	14,	-1,	-1,	14,	14,	-1,	14,	15,	-1,	14,	15,	15,	-1,	15,	16,	-1,	15,	16,	16,	-1,	16,	17,	0,	16,	17,	17,	0,	17,	-1,	2,	17,	4,	0,	2,	5,	4,	6,	-1,	5,	2,	6,	4,	-1,	7,	5,	8,	6,	7,	9,	8,	10,	-1,	9,	7,	10,	8,	-1,	12,	9,	12,	10,	-1,	13,	12,	13,	19,	-1,	19,	13,	-1,	-1,	19,	
}
var StackPackActDef = []int {
	120,	120,	120,	-9,	120,	120,	120,	120,	120,	120,	120,	-1,	-2,	-3,	120,	120,	120,	120,	120,	-8,	
}
var StackPackGotoDef = []int {
	120,	
}

 func (s *StateSym) Action(a int) int {
	if StatePackOffset[s.Yystate]+a  < 0 {
		 return ERROR_ACTION
	}
	if StatePackOffset[s.Yystate]+a >= len(StackPackCheck) || 
		StackPackCheck[StatePackOffset[s.Yystate]+a] != s.Yystate {
		 if a > NTERMINALS {
			 return StackPackGotoDef[a - NTERMINALS - 1]
		 }else {
			 return StackPackActDef[s.Yystate]
		 }
	}else{
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
	dollarDolar.YySymIndex = 12
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
		dollarDolar.val	=	Dollar[1].val + Dollar[3].val
	}
	PopStateSym(3)
case 2: 
	dollarDolar.YySymIndex = 12
	Dollar := StateSymStack[topIndex-3 : StackPointer]
	_ = Dollar

/*

LineNo:27
E -> E '*'  E 
 {
		$$	=	$1 * $3
	}
*/
{
		dollarDolar.val	=	Dollar[1].val * Dollar[3].val
	}
	PopStateSym(3)
case 3: 
	dollarDolar.YySymIndex = 12
	Dollar := StateSymStack[topIndex-3 : StackPointer]
	_ = Dollar

/*

LineNo:30
E -> E '/'  E 
 {
		$$	=	$1 / $3
	}
*/
{
		dollarDolar.val	=	Dollar[1].val / Dollar[3].val
	}
	PopStateSym(3)
case 4: 
	dollarDolar.YySymIndex = 12
	Dollar := StateSymStack[topIndex-3 : StackPointer]
	_ = Dollar

/*

LineNo:33
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
			dollarDolar.val	=	1
		} else {
			dollarDolar.val	=	0
		}
	}
	PopStateSym(3)
case 5: 
	dollarDolar.YySymIndex = 12
	Dollar := StateSymStack[topIndex-3 : StackPointer]
	_ = Dollar

/*

LineNo:40
E -> E '<'  E 
 {
		if $1 > $3 {
			$$	=	0
		} else {
			$$	=	1
		}
	}
*/
{
		if Dollar[1].val > Dollar[3].val {
			dollarDolar.val	=	0
		} else {
			dollarDolar.val	=	1
		}
	}
	PopStateSym(3)
case 6: 
	dollarDolar.YySymIndex = 12
	Dollar := StateSymStack[topIndex-3 : StackPointer]
	_ = Dollar

/*

LineNo:47
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
			dollarDolar.val	=	1
		} else {
			dollarDolar.val	=	0
		}
	}
	PopStateSym(3)
case 7: 
	dollarDolar.YySymIndex = 12
	Dollar := StateSymStack[topIndex-3 : StackPointer]
	_ = Dollar

/*

LineNo:54
E -> E LE E 
 {
		if $1 > $3 {
			$$	=	0
		} else {
			$$	=	1
		}
	}
*/
{
		if Dollar[1].val > Dollar[3].val {
			dollarDolar.val	=	0
		} else {
			dollarDolar.val	=	1
		}
	}
	PopStateSym(3)
case 8: 
	dollarDolar.YySymIndex = 12
	Dollar := StateSymStack[topIndex-3 : StackPointer]
	_ = Dollar

/*

LineNo:61
E -> '('  E ')'  
 {
		$$	=	$2
	}
*/
{
		dollarDolar.val	=	Dollar[2].val
	}
	PopStateSym(3)
case 9: 
	dollarDolar.YySymIndex = 12
	Dollar := StateSymStack[topIndex-1 : StackPointer]
	_ = Dollar

/*

LineNo:64
E -> NUM 
 {
		$$	=	$1
	}
*/
{
		dollarDolar.val	=	Dollar[1].val
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
					ValType:      val,
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
	case 100:
 	conv = 2
	case 102:
 	conv = 3
	case 43:
 	conv = 4
	case 42:
 	conv = 5
	case 40:
 	conv = 6
	case 47:
 	conv = 7
	case 62:
 	conv = 8
	case 60:
 	conv = 9
	case 41:
 	conv = 10
	case 101:
 	conv = 11

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
 	conv = "NUM"
	case 3:
 	conv = "LE"
	case 4:
 	conv = "'+' "
	case 5:
 	conv = "'*' "
	case 6:
 	conv = "'(' "
	case 7:
 	conv = "'/' "
	case 8:
 	conv = "'>' "
	case 9:
 	conv = "'<' "
	case 10:
 	conv = "')' "
	case 11:
 	conv = "GE"
	case 12:
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

		fmt.Printf("look ahead %s, use Reduce:E -> E '*'  E , go to state %d\n", look, s)
		case 3: 

		fmt.Printf("look ahead %s, use Reduce:E -> E '/'  E , go to state %d\n", look, s)
		case 4: 

		fmt.Printf("look ahead %s, use Reduce:E -> E '>'  E , go to state %d\n", look, s)
		case 5: 

		fmt.Printf("look ahead %s, use Reduce:E -> E '<'  E , go to state %d\n", look, s)
		case 6: 

		fmt.Printf("look ahead %s, use Reduce:E -> E GE E , go to state %d\n", look, s)
		case 7: 

		fmt.Printf("look ahead %s, use Reduce:E -> E LE E , go to state %d\n", look, s)
		case 8: 

		fmt.Printf("look ahead %s, use Reduce:E -> '('  E ')'  , go to state %d\n", look, s)
		case 9: 

		fmt.Printf("look ahead %s, use Reduce:E -> NUM , go to state %d\n", look, s)

		}
	}
}

// Code Last part

