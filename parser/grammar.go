
/*Generator Code , do not modify*/
// Code header part 

package parser
import "fmt"
import "os"



// const part
const BREAK = 15
const CARTESIAN = 16
const IDENTIFIER = 13
const RPAREN = 18
const BRACE_STRING = 19
const SINGLEQUOTE_STRING = 20
const PERMUTATION = 23
const NORMALCMD = 24
const STEP = 10
const DQUOTE_STRING = 11
const NEWLINE = 14
const IF = 22
const EACH = 17
const ELSE = 5
const LPAREN = 6
const SQUAREQUOTE_STRING = 12
const ERROR_ACTION = 136
const ACCEPT_ACTION = 236


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
	
    WorkValue

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
	236,	8,	0,	11,	10,	-10,	-12,	-12,	34,	11,	6,	30,	-12,	7,	5,	-12,	-12,	-9,	9,	-12,	-12,	26,	-6,	4,	13,	2,	3,	-1,	-1,	24,	-1,	-1,	12,	18,	-5,	-7,	-20,	-1,	-20,	20,	-1,	-1,	15,	19,	29,	-8,	1,	17,	14,	-16,	-16,	-11,	16,	-16,	-21,	-16,	-21,	0,	-16,	-16,	-17,	-17,	-16,	-16,	-17,	0,	-17,	0,	0,	-17,	-17,	-18,	-18,	-17,	-17,	-18,	0,	-18,	0,	0,	-18,	-18,	-19,	-19,	-18,	-18,	-19,	0,	-19,	0,	24,	-19,	-19,	-12,	-12,	-19,	-19,	0,	22,	-12,	24,	0,	-12,	-12,	18,	25,	-12,	-12,	0,	0,	20,	27,	0,	15,	19,	23,	0,	28,	17,	14,	0,	-2,	-2,	16,	-2,	-2,	-3,	-3,	0,	-3,	-3,	-2,	0,	0,	-2,	-2,	-3,	0,	0,	-3,	-3,	-23,	-23,	0,	-23,	-23,	-4,	-4,	0,	-4,	-4,	-23,	-13,	-13,	-23,	-23,	-4,	0,	-13,	-4,	-4,	-13,	-13,	-14,	-14,	-13,	-13,	0,	0,	-14,	0,	0,	-14,	-14,	-15,	-15,	-14,	-14,	0,	18,	-15,	0,	0,	-15,	-15,	20,	18,	-15,	-15,	19,	0,	0,	20,	17,	0,	0,	19,	18,	32,	0,	17,	0,	0,	20,	18,	31,	0,	19,	0,	0,	20,	17,	0,	-22,	19,	-22,	33,	-22,	17,	0,	0,	0,	0,	35,	0,	0,	-22,	 
}
var StatePackOffset = []int {
	26,	-1,	120,	125,	5,	3,	30,	21,	-1,	13,	176,	140,	145,	30,	149,	160,	171,	46,	57,	68,	79,	92,	82,	48,	90,	4,	18,	101,	209,	31,	183,	41,	194,	1,	201,	47,	
}
var StackPackCheck = []int {
	1,	1,	-1,	1,	1,	33,	5,	5,	33,	4,	1,	8,	5,	1,	1,	5,	5,	9,	1,	5,	5,	25,	26,	1,	5,	1,	1,	0,	0,	7,	0,	0,	4,	13,	13,	29,	6,	0,	6,	13,	0,	0,	13,	13,	7,	31,	0,	13,	13,	17,	17,	35,	13,	17,	23,	17,	23,	-1,	17,	17,	18,	18,	17,	17,	18,	-1,	18,	-1,	-1,	18,	18,	19,	19,	18,	18,	19,	-1,	19,	-1,	-1,	19,	19,	20,	20,	19,	19,	20,	-1,	20,	-1,	22,	20,	20,	24,	24,	20,	20,	-1,	21,	24,	21,	-1,	24,	24,	27,	22,	24,	24,	-1,	-1,	27,	24,	-1,	27,	27,	21,	-1,	27,	27,	27,	-1,	2,	2,	27,	2,	2,	3,	3,	-1,	3,	3,	2,	-1,	-1,	2,	2,	3,	-1,	-1,	3,	3,	11,	11,	-1,	11,	11,	12,	12,	-1,	12,	12,	11,	14,	14,	11,	11,	12,	-1,	14,	12,	12,	14,	14,	15,	15,	14,	14,	-1,	-1,	15,	-1,	-1,	15,	15,	16,	16,	15,	15,	-1,	10,	16,	-1,	-1,	16,	16,	10,	30,	16,	16,	10,	-1,	-1,	30,	10,	-1,	-1,	30,	32,	10,	-1,	30,	-1,	-1,	32,	34,	30,	-1,	32,	-1,	-1,	34,	32,	-1,	28,	34,	28,	32,	28,	34,	-1,	-1,	-1,	-1,	34,	-1,	-1,	28,	
}
var StackPackActDef = []int {
	136,	136,	136,	136,	136,	136,	136,	136,	136,	136,	136,	136,	136,	136,	136,	136,	136,	136,	136,	136,	136,	136,	136,	136,	136,	136,	136,	136,	136,	136,	136,	136,	136,	136,	136,	136,	
}
var StackPackGotoDef = []int {
	136,	136,	136,	136,	136,	136,	136,	136,	136,	136,	
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
	dollarDolar.YySymIndex = 20
	Dollar := StateSymStack[topIndex-0 : StackPointer]
	_ = Dollar

/*

LineNo:42
COMMANDS -> 
 
*/

	PopStateSym(0)
case 2: 
	dollarDolar.YySymIndex = 20
	Dollar := StateSymStack[topIndex-2 : StackPointer]
	_ = Dollar

/*

LineNo:42
COMMANDS -> COMMANDS COMMAND 
 
*/

	PopStateSym(2)
case 3: 
	dollarDolar.YySymIndex = 26
	Dollar := StateSymStack[topIndex-1 : StackPointer]
	_ = Dollar

/*

LineNo:43
COMMAND -> END 
 
*/

	PopStateSym(1)
case 4: 
	dollarDolar.YySymIndex = 26
	Dollar := StateSymStack[topIndex-2 : StackPointer]
	_ = Dollar

/*

LineNo:44
COMMAND -> CMD END 
 
*/

	PopStateSym(2)
case 5: 
	dollarDolar.YySymIndex = 24
	Dollar := StateSymStack[topIndex-2 : StackPointer]
	_ = Dollar

/*

LineNo:45
CMD -> NORMALCMD WORDS 
 {
       err := ExecuteNormalCmd($1, $2);
       if err != nil {
            ParserError(err)
       }
    }
*/
{
       err := ExecuteNormalCmd(Dollar[1].Value, Dollar[2].Array);
       if err != nil {
            ParserError(err)
       }
    }
	PopStateSym(2)
case 6: 
	dollarDolar.YySymIndex = 24
	Dollar := StateSymStack[topIndex-5 : StackPointer]
	_ = Dollar

/*

LineNo:51
CMD -> CARTESIAN ARRAYS EACH ARRAY BRACE_STRING 
 {
       w := NewWorkCart($2, $4, $5[1:len($5)-1]);
       err := w.Execute();
       if err != nil {
          ParserError(err)
       }
    }
*/
{
       w := NewWorkCart(Dollar[2].Arrays, Dollar[4].Array, Dollar[5].Value[1:len(Dollar[5].Value)-1]);
       err := w.Execute();
       if err != nil {
          ParserError(err)
       }
    }
	PopStateSym(5)
case 7: 
	dollarDolar.YySymIndex = 24
	Dollar := StateSymStack[topIndex-2 : StackPointer]
	_ = Dollar

/*

LineNo:58
CMD -> PERMUTATION ARRAY 
 {
       p := &PermType{
            StepNames: $2,
       }
       err := p.Execute();
         if err != nil {
             ParserError(err)
         }
    }
*/
{
       p := &PermType{
            StepNames: Dollar[2].Array,
       }
       err := p.Execute();
         if err != nil {
             ParserError(err)
         }
    }
	PopStateSym(2)
case 8: 
	dollarDolar.YySymIndex = 24
	Dollar := StateSymStack[topIndex-3 : StackPointer]
	_ = Dollar

/*

LineNo:67
CMD -> STEP IDENTIFIER STRING 
 {
       err := NewStep($2, $3)
       if err != nil {
          ParserError(err)
       }
    }
*/
{
       err := NewStep(Dollar[2].Value, Dollar[3].Value)
       if err != nil {
          ParserError(err)
       }
    }
	PopStateSym(3)
case 9: 
	dollarDolar.YySymIndex = 24
	Dollar := StateSymStack[topIndex-1 : StackPointer]
	_ = Dollar

/*

LineNo:73
CMD -> IFCMD 
 {
       
    }
*/
{
       
    }
	PopStateSym(1)
case 10: 
	dollarDolar.YySymIndex = 19
	Dollar := StateSymStack[topIndex-3 : StackPointer]
	_ = Dollar

/*

LineNo:76
IFCMD -> IF STRING STRING 
 {
         ifcmd := NewIfCmd($2, $3, "")
          ifcmd.Execute()
        }
*/
{
         ifcmd := NewIfCmd(Dollar[2].Value, Dollar[3].Value, "")
          ifcmd.Execute()
        }
	PopStateSym(3)
case 11: 
	dollarDolar.YySymIndex = 19
	Dollar := StateSymStack[topIndex-5 : StackPointer]
	_ = Dollar

/*

LineNo:80
IFCMD -> IF STRING STRING ELSE STRING 
 {
                ifcmd := NewIfCmd($2, $3, $5)
          ifcmd.Execute()
       }
*/
{
                ifcmd := NewIfCmd(Dollar[2].Value, Dollar[3].Value, Dollar[5].Value)
          ifcmd.Execute()
       }
	PopStateSym(5)
case 12: 
	dollarDolar.YySymIndex = 21
	Dollar := StateSymStack[topIndex-0 : StackPointer]
	_ = Dollar

/*

LineNo:85
WORDS -> 
 {
    $$= make([]string,0)
}
*/
{
    dollarDolar.Array= make([]string,0)
}
	PopStateSym(0)
case 13: 
	dollarDolar.YySymIndex = 21
	Dollar := StateSymStack[topIndex-2 : StackPointer]
	_ = Dollar

/*

LineNo:88
WORDS -> WORDS WORD 
 {
        $$ = append($1, $2)
    }
*/
{
        dollarDolar.Array = append(Dollar[1].Array, Dollar[2].Value)
    }
	PopStateSym(2)
case 14: 
	dollarDolar.YySymIndex = 18
	Dollar := StateSymStack[topIndex-1 : StackPointer]
	_ = Dollar

/*

LineNo:91
WORD -> IDENTIFIER 
 {$$ = $1 }
*/
{dollarDolar.Value = Dollar[1].Value }
	PopStateSym(1)
case 15: 
	dollarDolar.YySymIndex = 18
	Dollar := StateSymStack[topIndex-1 : StackPointer]
	_ = Dollar

/*

LineNo:92
WORD -> STRING 
 { $$ = $1 }
*/
{ dollarDolar.Value = Dollar[1].Value }
	PopStateSym(1)
case 16: 
	dollarDolar.YySymIndex = 22
	Dollar := StateSymStack[topIndex-1 : StackPointer]
	_ = Dollar

/*

LineNo:93
STRING -> BRACE_STRING 
 {$$ = $1 }
*/
{dollarDolar.Value = Dollar[1].Value }
	PopStateSym(1)
case 17: 
	dollarDolar.YySymIndex = 22
	Dollar := StateSymStack[topIndex-1 : StackPointer]
	_ = Dollar

/*

LineNo:94
STRING -> DQUOTE_STRING 
 {$$ = $1 }
*/
{dollarDolar.Value = Dollar[1].Value }
	PopStateSym(1)
case 18: 
	dollarDolar.YySymIndex = 22
	Dollar := StateSymStack[topIndex-1 : StackPointer]
	_ = Dollar

/*

LineNo:95
STRING -> SINGLEQUOTE_STRING 
 {$$ = $1 }
*/
{dollarDolar.Value = Dollar[1].Value }
	PopStateSym(1)
case 19: 
	dollarDolar.YySymIndex = 22
	Dollar := StateSymStack[topIndex-1 : StackPointer]
	_ = Dollar

/*

LineNo:96
STRING -> SQUAREQUOTE_STRING 
 {$$ = $1 }
*/
{dollarDolar.Value = Dollar[1].Value }
	PopStateSym(1)
case 20: 
	dollarDolar.YySymIndex = 28
	Dollar := StateSymStack[topIndex-0 : StackPointer]
	_ = Dollar

/*

LineNo:97
ARRAYS -> 
 {$$ =make([][]string ,0)}
*/
{dollarDolar.Arrays =make([][]string ,0)}
	PopStateSym(0)
case 21: 
	dollarDolar.YySymIndex = 28
	Dollar := StateSymStack[topIndex-2 : StackPointer]
	_ = Dollar

/*

LineNo:98
ARRAYS -> ARRAYS ARRAY 
 {
     $$ = append($1, $2) 
     }
*/
{
     dollarDolar.Arrays = append(Dollar[1].Arrays, Dollar[2].Array) 
     }
	PopStateSym(2)
case 22: 
	dollarDolar.YySymIndex = 23
	Dollar := StateSymStack[topIndex-3 : StackPointer]
	_ = Dollar

/*

LineNo:101
ARRAY -> LPAREN WORDS RPAREN 
 {
    $$ = $2
}
*/
{
    dollarDolar.Array = Dollar[2].Array
}
	PopStateSym(3)
case 23: 
	dollarDolar.YySymIndex = 27
	Dollar := StateSymStack[topIndex-1 : StackPointer]
	_ = Dollar

/*

LineNo:104
END -> NEWLINE 
 
*/

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
	case 10:
 	conv = 2
	case 11:
 	conv = 3
	case 14:
 	conv = 4
	case 22:
 	conv = 5
	case 17:
 	conv = 6
	case 5:
 	conv = 7
	case 6:
 	conv = 8
	case 12:
 	conv = 9
	case 15:
 	conv = 10
	case 16:
 	conv = 11
	case 13:
 	conv = 12
	case 20:
 	conv = 13
	case 23:
 	conv = 14
	case 24:
 	conv = 15
	case 18:
 	conv = 16
	case 19:
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
 	conv = "STEP"
	case 3:
 	conv = "DQUOTE_STRING"
	case 4:
 	conv = "NEWLINE"
	case 5:
 	conv = "IF"
	case 6:
 	conv = "EACH"
	case 7:
 	conv = "ELSE"
	case 8:
 	conv = "LPAREN"
	case 9:
 	conv = "SQUAREQUOTE_STRING"
	case 10:
 	conv = "BREAK"
	case 11:
 	conv = "CARTESIAN"
	case 12:
 	conv = "IDENTIFIER"
	case 13:
 	conv = "SINGLEQUOTE_STRING"
	case 14:
 	conv = "PERMUTATION"
	case 15:
 	conv = "NORMALCMD"
	case 16:
 	conv = "RPAREN"
	case 17:
 	conv = "BRACE_STRING"
	case 18:
 	conv = "WORD"
	case 19:
 	conv = "IFCMD"
	case 20:
 	conv = "COMMANDS"
	case 21:
 	conv = "WORDS"
	case 22:
 	conv = "STRING"
	case 23:
 	conv = "ARRAY"
	case 24:
 	conv = "CMD"
	case 25:
 	conv = "STEPCMDS"
	case 26:
 	conv = "COMMAND"
	case 27:
 	conv = "END"
	case 28:
 	conv = "ARRAYS"

	}
	return conv
}
// Trace function for reduce
func TraceReduce(reduceIndex, s int, look string) {
	if IsTrace {
		switch reduceIndex {
		case 1: 

		fmt.Printf("look ahead %s, use Reduce:COMMANDS -> , go to state %d\n", look, s)
		case 2: 

		fmt.Printf("look ahead %s, use Reduce:COMMANDS -> COMMANDS COMMAND , go to state %d\n", look, s)
		case 3: 

		fmt.Printf("look ahead %s, use Reduce:COMMAND -> END , go to state %d\n", look, s)
		case 4: 

		fmt.Printf("look ahead %s, use Reduce:COMMAND -> CMD END , go to state %d\n", look, s)
		case 5: 

		fmt.Printf("look ahead %s, use Reduce:CMD -> NORMALCMD WORDS , go to state %d\n", look, s)
		case 6: 

		fmt.Printf("look ahead %s, use Reduce:CMD -> CARTESIAN ARRAYS EACH ARRAY BRACE_STRING , go to state %d\n", look, s)
		case 7: 

		fmt.Printf("look ahead %s, use Reduce:CMD -> PERMUTATION ARRAY , go to state %d\n", look, s)
		case 8: 

		fmt.Printf("look ahead %s, use Reduce:CMD -> STEP IDENTIFIER STRING , go to state %d\n", look, s)
		case 9: 

		fmt.Printf("look ahead %s, use Reduce:CMD -> IFCMD , go to state %d\n", look, s)
		case 10: 

		fmt.Printf("look ahead %s, use Reduce:IFCMD -> IF STRING STRING , go to state %d\n", look, s)
		case 11: 

		fmt.Printf("look ahead %s, use Reduce:IFCMD -> IF STRING STRING ELSE STRING , go to state %d\n", look, s)
		case 12: 

		fmt.Printf("look ahead %s, use Reduce:WORDS -> , go to state %d\n", look, s)
		case 13: 

		fmt.Printf("look ahead %s, use Reduce:WORDS -> WORDS WORD , go to state %d\n", look, s)
		case 14: 

		fmt.Printf("look ahead %s, use Reduce:WORD -> IDENTIFIER , go to state %d\n", look, s)
		case 15: 

		fmt.Printf("look ahead %s, use Reduce:WORD -> STRING , go to state %d\n", look, s)
		case 16: 

		fmt.Printf("look ahead %s, use Reduce:STRING -> BRACE_STRING , go to state %d\n", look, s)
		case 17: 

		fmt.Printf("look ahead %s, use Reduce:STRING -> DQUOTE_STRING , go to state %d\n", look, s)
		case 18: 

		fmt.Printf("look ahead %s, use Reduce:STRING -> SINGLEQUOTE_STRING , go to state %d\n", look, s)
		case 19: 

		fmt.Printf("look ahead %s, use Reduce:STRING -> SQUAREQUOTE_STRING , go to state %d\n", look, s)
		case 20: 

		fmt.Printf("look ahead %s, use Reduce:ARRAYS -> , go to state %d\n", look, s)
		case 21: 

		fmt.Printf("look ahead %s, use Reduce:ARRAYS -> ARRAYS ARRAY , go to state %d\n", look, s)
		case 22: 

		fmt.Printf("look ahead %s, use Reduce:ARRAY -> LPAREN WORDS RPAREN , go to state %d\n", look, s)
		case 23: 

		fmt.Printf("look ahead %s, use Reduce:END -> NEWLINE , go to state %d\n", look, s)

		}
	}
}

// Code Last part

func ParserError(err error) {
    fmt.Print("ParserError: ")
    fmt.Println(err)
    os.Exit(1)
}
