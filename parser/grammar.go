
package Parser
import "fmt"
import "os"

 import "strings"
// const part 
const NEWLINE = 17
const LPAREN = 18
const RPAREN = 19
const BRACE_STRING = 14
const SINGLEQUOTE_STRING = 15
const PERMUTATION = 21
const STEP = 22
const EACH = 23
const CARTESIAN = 7
const NORMALCMD = 8
const IDENTIFIER = 9
const DQUOTE_STRING = 3
const BREAK = 11
const SQUAREQUOTE_STRING = 12
const ERROR_ACTION = 0
const ACCEPT_ACTION = 130

var IsTrace bool = false
var StateSymStack = []StateSym{}
var StackPointer = 0
type Context struct {
	StackSym []StateSym
	Stackpos int
}

var globalContext = []Context{}

type ValType struct {
	
    WorkValue

}
type StateSym struct {
	Yystate int // state
	
 	//sym val
	YySymIndex int
	//other
	ValType
}
var StatePackAction = []int {
	0,	130,	7,	8,	0,	20,	0,	6,	5,	28,	-9,	-9,	11,	0,	0,	-9,	21,	2,	4,	3,	9,	-9,	22,	-9,	-9,	-9,	13,	16,	-17,	12,	19,	18,	24,	-9,	-9,	25,	-18,	-5,	-9,	-19,	15,	17,	14,	-6,	-9,	-17,	-9,	-9,	-9,	13,	16,	27,	12,	-18,	18,	-19,	-19,	22,	-19,	10,	9,	-7,	26,	15,	17,	14,	-1,	-1,	-1,	-10,	-10,	23,	-1,	-1,	-10,	16,	-8,	22,	0,	18,	-10,	1,	-10,	-10,	-10,	-1,	-11,	-11,	15,	17,	29,	-11,	0,	-12,	-12,	0,	0,	-11,	-12,	-11,	-11,	-11,	0,	0,	-12,	0,	-12,	-12,	-12,	-13,	-13,	0,	0,	0,	-13,	0,	-14,	-14,	0,	0,	-13,	-14,	-13,	-13,	-13,	0,	0,	-14,	0,	-14,	-14,	-14,	-15,	-15,	0,	0,	0,	-15,	0,	-16,	-16,	0,	0,	-15,	-16,	-15,	-15,	-15,	0,	0,	-16,	0,	-16,	-16,	-16,	-2,	-2,	-2,	-3,	-3,	-3,	-2,	-2,	0,	-3,	-3,	-20,	-20,	-20,	-4,	-4,	-4,	-20,	-20,	-2,	-4,	-4,	-3,	0,	0,	0,	0,	0,	0,	0,	-20,	0,	0,	-4,	 
}
var StatePackOffset = []int {
	65,	0,	154,	157,	40,	1,	24,	36,	0,	165,	168,	17,	60,	77,	84,	100,	107,	123,	130,	1,	56,	32,	24,	9,	23,	40,	35,	41,	65,	56,	
}
var StackPackCheck = []int {
	-1,	1,	1,	1,	-1,	19,	-1,	1,	1,	8,	5,	5,	5,	-1,	-1,	5,	19,	1,	1,	1,	1,	5,	19,	5,	5,	5,	11,	11,	6,	11,	6,	11,	23,	22,	22,	22,	21,	11,	22,	26,	11,	11,	11,	24,	22,	6,	22,	22,	22,	25,	25,	7,	25,	21,	25,	26,	26,	7,	26,	4,	4,	27,	25,	25,	25,	25,	0,	0,	0,	12,	12,	20,	0,	0,	12,	28,	29,	20,	-1,	28,	12,	0,	12,	12,	12,	0,	13,	13,	28,	28,	28,	13,	-1,	14,	14,	-1,	-1,	13,	14,	13,	13,	13,	-1,	-1,	14,	-1,	14,	14,	14,	15,	15,	-1,	-1,	-1,	15,	-1,	16,	16,	-1,	-1,	15,	16,	15,	15,	15,	-1,	-1,	16,	-1,	16,	16,	16,	17,	17,	-1,	-1,	-1,	17,	-1,	18,	18,	-1,	-1,	17,	18,	17,	17,	17,	-1,	-1,	18,	-1,	18,	18,	18,	2,	2,	2,	3,	3,	3,	2,	2,	-1,	3,	3,	9,	9,	9,	10,	10,	10,	9,	9,	2,	10,	10,	3,	-1,	-1,	-1,	-1,	-1,	-1,	-1,	9,	-1,	-1,	10,	
}	

func ParserError(err error) {
    fmt.Print("ParserError: ")
    fmt.Println(err)
    os.Exit(1)
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

func (s *StateSym) Action(a int) int {
	if StatePackOffset[s.Yystate]+a >= len(StackPackCheck) || StackPackCheck[StatePackOffset[s.Yystate]+a] != s.Yystate {
		return 0
	}else{
		return StatePackAction[StatePackOffset[s.Yystate]+a]
	}
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
func init() {
	ParserInit()
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
			lines := strings.Split(input[:currentPos], "\n")
			panic("Grammar parse error near :" + lines[len(lines)-1])
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

func ReduceFunc(reduceIndex int) *StateSym {
	dollarDolar := &StateSym{}
	topIndex := StackPointer - 1
	switch reduceIndex {
		case 1: 
	dollarDolar.YySymIndex = 16
	Dollar := StateSymStack[topIndex-0 : StackPointer]
	_ = Dollar

/*

LineNo:40
COMMANDS -> 
 
*/

	PopStateSym(0)
case 2: 
	dollarDolar.YySymIndex = 16
	Dollar := StateSymStack[topIndex-2 : StackPointer]
	_ = Dollar

/*

LineNo:40
COMMANDS -> COMMANDS COMMAND 
 
*/

	PopStateSym(2)
case 3: 
	dollarDolar.YySymIndex = 17
	Dollar := StateSymStack[topIndex-1 : StackPointer]
	_ = Dollar

/*

LineNo:41
COMMAND -> END 
 
*/

	PopStateSym(1)
case 4: 
	dollarDolar.YySymIndex = 17
	Dollar := StateSymStack[topIndex-2 : StackPointer]
	_ = Dollar

/*

LineNo:42
COMMAND -> CMD END 
 
*/

	PopStateSym(2)
case 5: 
	dollarDolar.YySymIndex = 18
	Dollar := StateSymStack[topIndex-2 : StackPointer]
	_ = Dollar

/*

LineNo:43
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
	dollarDolar.YySymIndex = 18
	Dollar := StateSymStack[topIndex-5 : StackPointer]
	_ = Dollar

/*

LineNo:49
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
	dollarDolar.YySymIndex = 18
	Dollar := StateSymStack[topIndex-2 : StackPointer]
	_ = Dollar

/*

LineNo:56
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
	dollarDolar.YySymIndex = 18
	Dollar := StateSymStack[topIndex-3 : StackPointer]
	_ = Dollar

/*

LineNo:65
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
	dollarDolar.YySymIndex = 11
	Dollar := StateSymStack[topIndex-0 : StackPointer]
	_ = Dollar

/*

LineNo:72
WORDS -> 
 {
    $$= make([]string,0)
}
*/
{
    dollarDolar.Array= make([]string,0)
}
	PopStateSym(0)
case 10: 
	dollarDolar.YySymIndex = 11
	Dollar := StateSymStack[topIndex-2 : StackPointer]
	_ = Dollar

/*

LineNo:75
WORDS -> WORDS WORD 
 {
        $$ = append($1, $2)
    }
*/
{
        dollarDolar.Array = append(Dollar[1].Array, Dollar[2].Value)
    }
	PopStateSym(2)
case 11: 
	dollarDolar.YySymIndex = 12
	Dollar := StateSymStack[topIndex-1 : StackPointer]
	_ = Dollar

/*

LineNo:78
WORD -> IDENTIFIER 
 {$$ = $1 }
*/
{dollarDolar.Value = Dollar[1].Value }
	PopStateSym(1)
case 12: 
	dollarDolar.YySymIndex = 12
	Dollar := StateSymStack[topIndex-1 : StackPointer]
	_ = Dollar

/*

LineNo:79
WORD -> STRING 
 { $$ = $1 }
*/
{ dollarDolar.Value = Dollar[1].Value }
	PopStateSym(1)
case 13: 
	dollarDolar.YySymIndex = 25
	Dollar := StateSymStack[topIndex-1 : StackPointer]
	_ = Dollar

/*

LineNo:80
STRING -> BRACE_STRING 
 {$$ = $1 }
*/
{dollarDolar.Value = Dollar[1].Value }
	PopStateSym(1)
case 14: 
	dollarDolar.YySymIndex = 25
	Dollar := StateSymStack[topIndex-1 : StackPointer]
	_ = Dollar

/*

LineNo:81
STRING -> DQUOTE_STRING 
 {$$ = $1 }
*/
{dollarDolar.Value = Dollar[1].Value }
	PopStateSym(1)
case 15: 
	dollarDolar.YySymIndex = 25
	Dollar := StateSymStack[topIndex-1 : StackPointer]
	_ = Dollar

/*

LineNo:82
STRING -> SINGLEQUOTE_STRING 
 {$$ = $1 }
*/
{dollarDolar.Value = Dollar[1].Value }
	PopStateSym(1)
case 16: 
	dollarDolar.YySymIndex = 25
	Dollar := StateSymStack[topIndex-1 : StackPointer]
	_ = Dollar

/*

LineNo:83
STRING -> SQUAREQUOTE_STRING 
 {$$ = $1 }
*/
{dollarDolar.Value = Dollar[1].Value }
	PopStateSym(1)
case 17: 
	dollarDolar.YySymIndex = 6
	Dollar := StateSymStack[topIndex-0 : StackPointer]
	_ = Dollar

/*

LineNo:84
ARRAYS -> 
 {$$ =make([][]string ,0)}
*/
{dollarDolar.Arrays =make([][]string ,0)}
	PopStateSym(0)
case 18: 
	dollarDolar.YySymIndex = 6
	Dollar := StateSymStack[topIndex-2 : StackPointer]
	_ = Dollar

/*

LineNo:85
ARRAYS -> ARRAYS ARRAY 
 {
     $$ = append($1, $2) 
     }
*/
{
     dollarDolar.Arrays = append(Dollar[1].Arrays, Dollar[2].Array) 
     }
	PopStateSym(2)
case 19: 
	dollarDolar.YySymIndex = 15
	Dollar := StateSymStack[topIndex-3 : StackPointer]
	_ = Dollar

/*

LineNo:88
ARRAY -> LPAREN WORDS RPAREN 
 {
    $$ = $2
}
*/
{
    dollarDolar.Array = Dollar[2].Array
}
	PopStateSym(3)
case 20: 
	dollarDolar.YySymIndex = 19
	Dollar := StateSymStack[topIndex-1 : StackPointer]
	_ = Dollar

/*

LineNo:91
END -> NEWLINE 
 
*/

	PopStateSym(1)

	}
	return dollarDolar
}

func translate(c int) int {
	var conv int = 0
	switch c {
	case -1:
 	conv = 1
	case 21:
 	conv = 2
	case 22:
 	conv = 3
	case 23:
 	conv = 4
	case 7:
 	conv = 7
	case 8:
 	conv = 8
	case 9:
 	conv = 9
	case 3:
 	conv = 10
	case 11:
 	conv = 13
	case 12:
 	conv = 14
	case 17:
 	conv = 20
	case 18:
 	conv = 21
	case 19:
 	conv = 22
	case 14:
 	conv = 23
	case 15:
 	conv = 24
	
	}
	return conv
}

func TraceShift(s *StateSym) {
	if IsTrace {
		fmt.Printf("Shift %s, push state %d\n", TraceTranslate(s.YySymIndex), s.Yystate)
	}
}
func TraceTranslate(c int) string {
	var conv string = ""
	switch c {
	case 0:
 	conv = "start"
	case 1:
 	conv = "$"
	case 2:
 	conv = "PERMUTATION"
	case 3:
 	conv = "STEP"
	case 4:
 	conv = "EACH"
	case 5:
 	conv = "STEPCMDS"
	case 6:
 	conv = "ARRAYS"
	case 7:
 	conv = "CARTESIAN"
	case 8:
 	conv = "NORMALCMD"
	case 9:
 	conv = "IDENTIFIER"
	case 10:
 	conv = "DQUOTE_STRING"
	case 11:
 	conv = "WORDS"
	case 12:
 	conv = "WORD"
	case 13:
 	conv = "BREAK"
	case 14:
 	conv = "SQUAREQUOTE_STRING"
	case 15:
 	conv = "ARRAY"
	case 16:
 	conv = "COMMANDS"
	case 17:
 	conv = "COMMAND"
	case 18:
 	conv = "CMD"
	case 19:
 	conv = "END"
	case 20:
 	conv = "NEWLINE"
	case 21:
 	conv = "LPAREN"
	case 22:
 	conv = "RPAREN"
	case 23:
 	conv = "BRACE_STRING"
	case 24:
 	conv = "SINGLEQUOTE_STRING"
	case 25:
 	conv = "STRING"

	}
	return conv
}


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

		fmt.Printf("look ahead %s, use Reduce:WORDS -> , go to state %d\n", look, s)
		case 10: 

		fmt.Printf("look ahead %s, use Reduce:WORDS -> WORDS WORD , go to state %d\n", look, s)
		case 11: 

		fmt.Printf("look ahead %s, use Reduce:WORD -> IDENTIFIER , go to state %d\n", look, s)
		case 12: 

		fmt.Printf("look ahead %s, use Reduce:WORD -> STRING , go to state %d\n", look, s)
		case 13: 

		fmt.Printf("look ahead %s, use Reduce:STRING -> BRACE_STRING , go to state %d\n", look, s)
		case 14: 

		fmt.Printf("look ahead %s, use Reduce:STRING -> DQUOTE_STRING , go to state %d\n", look, s)
		case 15: 

		fmt.Printf("look ahead %s, use Reduce:STRING -> SINGLEQUOTE_STRING , go to state %d\n", look, s)
		case 16: 

		fmt.Printf("look ahead %s, use Reduce:STRING -> SQUAREQUOTE_STRING , go to state %d\n", look, s)
		case 17: 

		fmt.Printf("look ahead %s, use Reduce:ARRAYS -> , go to state %d\n", look, s)
		case 18: 

		fmt.Printf("look ahead %s, use Reduce:ARRAYS -> ARRAYS ARRAY , go to state %d\n", look, s)
		case 19: 

		fmt.Printf("look ahead %s, use Reduce:ARRAY -> LPAREN WORDS RPAREN , go to state %d\n", look, s)
		case 20: 

		fmt.Printf("look ahead %s, use Reduce:END -> NEWLINE , go to state %d\n", look, s)

		}
	}
}
