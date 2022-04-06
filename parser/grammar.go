package Parser

import (
	"fmt"
	"os"
	"strings"
)

// const part
const BREAK = 10
const EACH = 11
const IDENTIFIER = 12
const SINGLEQUOTE_STRING = 13
const SQUAREQUOTE_STRING = 14
const CARTESIAN = 15
const NORMALCMD = 16
const RPAREN = 17
const BRACE_STRING = 18
const NEWLINE = 3
const LPAREN = 4
const DQUOTE_STRING = 7
const ERROR_ACTION = 0
const ACCEPT_ACTION = 125

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

var StatePackAction = []int{
	0, 0, -7, -15, -6, 7, 20, -7, 9, 8, 0, -15, -7, -7, -7, 0, -5, 0, 12, -7, -7, 14, 17, 10, 21, -7, 11, 15, 16, 22, -7, 23, -16, 0, 13, -7, -7, -7, 0, 0, -16, 12, -7, -7, 14, 0, 10, 125, 7, 11, 15, 16, 3, 0, 0, -8, 24, 13, 0, 0, -8, 2, 4, 6, 5, -8, -8, -8, 0, -9, 0, 0, -8, -8, -9, 0, 20, 0, -10, -9, -9, -9, 0, -10, 18, 0, -9, -9, -10, -10, -10, -17, -11, 0, 19, -10, -10, -11, 0, -17, 0, -12, -11, -11, -11, 0, -12, 0, -17, -11, -11, -12, -12, -12, 0, -13, 0, 0, -12, -12, -13, -2, -2, 0, -14, -13, -13, -13, 0, -14, -3, -3, -13, -13, -14, -14, -14, -2, -2, -1, -1, -14, -14, 1, -18, -18, -3, -3, -4, -4, 0, 0, 0, 0, 0, -1, -1, 0, 0, 0, -18, -18, 0, 0, -4, -4,
}
var StatePackOffset = []int{
	138, 46, 120, 129, 3, 0, 0, 143, 147, 14, 53, 67, 76, 90, 99, 113, 122, 73, 3, 29, 23, 9, 2, 37, 88,
}
var StackPackCheck = []int{
	-1, -1, 5, 6, 22, 4, 18, 5, 5, 4, -1, 6, 5, 5, 5, -1, 9, -1, 9, 5, 5, 9, 6, 9, 18, 20, 9, 9, 9, 21, 20, 20, 19, -1, 9, 20, 20, 20, -1, -1, 19, 23, 20, 20, 23, -1, 23, 1, 1, 23, 23, 23, 1, -1, -1, 10, 23, 23, -1, -1, 10, 1, 1, 1, 1, 10, 10, 10, -1, 11, -1, -1, 10, 10, 11, -1, 17, -1, 12, 11, 11, 11, -1, 12, 17, -1, 11, 11, 12, 12, 12, 24, 13, -1, 17, 12, 12, 13, -1, 24, -1, 14, 13, 13, 13, -1, 14, -1, 24, 13, 13, 14, 14, 14, -1, 15, -1, -1, 14, 14, 15, 2, 2, -1, 16, 15, 15, 15, -1, 16, 3, 3, 15, 15, 16, 16, 16, 2, 2, 0, 0, 16, 16, 0, 7, 7, 3, 3, 8, 8, -1, -1, -1, -1, -1, 0, 0, -1, -1, -1, 7, 7, -1, -1, 8, 8,
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
	} else {
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

func ReduceFunc(reduceIndex int) *StateSym {
	dollarDolar := &StateSym{}
	topIndex := StackPointer - 1
	switch reduceIndex {
	case 1:
		dollarDolar.YySymIndex = 5
		Dollar := StateSymStack[topIndex-0 : StackPointer]
		_ = Dollar

		/*

		   LineNo:39
		   COMMANDS ->

		*/

		PopStateSym(0)
	case 2:
		dollarDolar.YySymIndex = 5
		Dollar := StateSymStack[topIndex-2 : StackPointer]
		_ = Dollar

		/*

		   LineNo:39
		   COMMANDS -> COMMANDS COMMAND

		*/

		PopStateSym(2)
	case 3:
		dollarDolar.YySymIndex = 15
		Dollar := StateSymStack[topIndex-1 : StackPointer]
		_ = Dollar

		/*

		   LineNo:40
		   COMMAND -> END

		*/

		PopStateSym(1)
	case 4:
		dollarDolar.YySymIndex = 15
		Dollar := StateSymStack[topIndex-2 : StackPointer]
		_ = Dollar

		/*

		   LineNo:41
		   COMMAND -> CMD END

		*/

		PopStateSym(2)
	case 5:
		dollarDolar.YySymIndex = 16
		Dollar := StateSymStack[topIndex-2 : StackPointer]
		_ = Dollar

		/*

		   LineNo:42
		   CMD -> NORMALCMD WORDS
		    {
		          err := ExecuteNormalCmd($1, $2);
		          if err != nil {
		               ParserError(err)
		          }
		       }
		*/
		{
			err := ExecuteNormalCmd(Dollar[1].Value, Dollar[2].Array)
			if err != nil {
				ParserError(err)
			}
		}
		PopStateSym(2)
	case 6:
		dollarDolar.YySymIndex = 16
		Dollar := StateSymStack[topIndex-5 : StackPointer]
		_ = Dollar

		/*

		   LineNo:48
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
			w := NewWorkCart(Dollar[2].Arrays, Dollar[4].Array, Dollar[5].Value[1:len(Dollar[5].Value)-1])
			err := w.Execute()
			if err != nil {
				ParserError(err)
			}
		}
		PopStateSym(5)
	case 7:
		dollarDolar.YySymIndex = 8
		Dollar := StateSymStack[topIndex-0 : StackPointer]
		_ = Dollar

		/*

		   LineNo:55
		   WORDS ->
		    {
		       $$= make([]string,0)
		   }
		*/
		{
			dollarDolar.Array = make([]string, 0)
		}
		PopStateSym(0)
	case 8:
		dollarDolar.YySymIndex = 8
		Dollar := StateSymStack[topIndex-2 : StackPointer]
		_ = Dollar

		/*

		   LineNo:58
		   WORDS -> WORDS WORD
		    {
		           $$ = append($1, $2)
		       }
		*/
		{
			dollarDolar.Array = append(Dollar[1].Array, Dollar[2].Value)
		}
		PopStateSym(2)
	case 9:
		dollarDolar.YySymIndex = 9
		Dollar := StateSymStack[topIndex-1 : StackPointer]
		_ = Dollar

		/*

		   LineNo:61
		   WORD -> IDENTIFIER
		    {$$ = $1 }
		*/
		{
			dollarDolar.Value = Dollar[1].Value
		}
		PopStateSym(1)
	case 10:
		dollarDolar.YySymIndex = 9
		Dollar := StateSymStack[topIndex-1 : StackPointer]
		_ = Dollar

		/*

		   LineNo:62
		   WORD -> STRING
		    { $$ = $1 }
		*/
		{
			dollarDolar.Value = Dollar[1].Value
		}
		PopStateSym(1)
	case 11:
		dollarDolar.YySymIndex = 4
		Dollar := StateSymStack[topIndex-1 : StackPointer]
		_ = Dollar

		/*

		   LineNo:63
		   STRING -> BRACE_STRING
		    {$$ = $1 }
		*/
		{
			dollarDolar.Value = Dollar[1].Value
		}
		PopStateSym(1)
	case 12:
		dollarDolar.YySymIndex = 4
		Dollar := StateSymStack[topIndex-1 : StackPointer]
		_ = Dollar

		/*

		   LineNo:64
		   STRING -> DQUOTE_STRING
		    {$$ = $1 }
		*/
		{
			dollarDolar.Value = Dollar[1].Value
		}
		PopStateSym(1)
	case 13:
		dollarDolar.YySymIndex = 4
		Dollar := StateSymStack[topIndex-1 : StackPointer]
		_ = Dollar

		/*

		   LineNo:65
		   STRING -> SINGLEQUOTE_STRING
		    {$$ = $1 }
		*/
		{
			dollarDolar.Value = Dollar[1].Value
		}
		PopStateSym(1)
	case 14:
		dollarDolar.YySymIndex = 4
		Dollar := StateSymStack[topIndex-1 : StackPointer]
		_ = Dollar

		/*

		   LineNo:66
		   STRING -> SQUAREQUOTE_STRING
		    {$$ = $1 }
		*/
		{
			dollarDolar.Value = Dollar[1].Value
		}
		PopStateSym(1)
	case 15:
		dollarDolar.YySymIndex = 22
		Dollar := StateSymStack[topIndex-0 : StackPointer]
		_ = Dollar

		/*

		   LineNo:67
		   ARRAYS ->
		    {$$ =make([][]string ,0)}
		*/
		{
			dollarDolar.Arrays = make([][]string, 0)
		}
		PopStateSym(0)
	case 16:
		dollarDolar.YySymIndex = 22
		Dollar := StateSymStack[topIndex-2 : StackPointer]
		_ = Dollar

		/*

		   LineNo:68
		   ARRAYS -> ARRAYS ARRAY
		    {
		        $$ = append($1, $2)
		        }
		*/
		{
			dollarDolar.Arrays = append(Dollar[1].Arrays, Dollar[2].Array)
		}
		PopStateSym(2)
	case 17:
		dollarDolar.YySymIndex = 21
		Dollar := StateSymStack[topIndex-3 : StackPointer]
		_ = Dollar

		/*

		   LineNo:71
		   ARRAY -> LPAREN WORDS RPAREN
		    {
		       $$ = $2
		   }
		*/
		{
			dollarDolar.Array = Dollar[2].Array
		}
		PopStateSym(3)
	case 18:
		dollarDolar.YySymIndex = 6
		Dollar := StateSymStack[topIndex-1 : StackPointer]
		_ = Dollar

		/*

		   LineNo:74
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
	case 3:
		conv = 2
	case 4:
		conv = 3
	case 7:
		conv = 7
	case 10:
		conv = 10
	case 11:
		conv = 11
	case 12:
		conv = 12
	case 13:
		conv = 13
	case 14:
		conv = 14
	case 15:
		conv = 17
	case 16:
		conv = 18
	case 17:
		conv = 19
	case 18:
		conv = 20

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
		conv = "NEWLINE"
	case 3:
		conv = "LPAREN"
	case 4:
		conv = "STRING"
	case 5:
		conv = "COMMANDS"
	case 6:
		conv = "END"
	case 7:
		conv = "DQUOTE_STRING"
	case 8:
		conv = "WORDS"
	case 9:
		conv = "WORD"
	case 10:
		conv = "BREAK"
	case 11:
		conv = "EACH"
	case 12:
		conv = "IDENTIFIER"
	case 13:
		conv = "SINGLEQUOTE_STRING"
	case 14:
		conv = "SQUAREQUOTE_STRING"
	case 15:
		conv = "COMMAND"
	case 16:
		conv = "CMD"
	case 17:
		conv = "CARTESIAN"
	case 18:
		conv = "NORMALCMD"
	case 19:
		conv = "RPAREN"
	case 20:
		conv = "BRACE_STRING"
	case 21:
		conv = "ARRAY"
	case 22:
		conv = "ARRAYS"

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

			fmt.Printf("look ahead %s, use Reduce:WORDS -> , go to state %d\n", look, s)
		case 8:

			fmt.Printf("look ahead %s, use Reduce:WORDS -> WORDS WORD , go to state %d\n", look, s)
		case 9:

			fmt.Printf("look ahead %s, use Reduce:WORD -> IDENTIFIER , go to state %d\n", look, s)
		case 10:

			fmt.Printf("look ahead %s, use Reduce:WORD -> STRING , go to state %d\n", look, s)
		case 11:

			fmt.Printf("look ahead %s, use Reduce:STRING -> BRACE_STRING , go to state %d\n", look, s)
		case 12:

			fmt.Printf("look ahead %s, use Reduce:STRING -> DQUOTE_STRING , go to state %d\n", look, s)
		case 13:

			fmt.Printf("look ahead %s, use Reduce:STRING -> SINGLEQUOTE_STRING , go to state %d\n", look, s)
		case 14:

			fmt.Printf("look ahead %s, use Reduce:STRING -> SQUAREQUOTE_STRING , go to state %d\n", look, s)
		case 15:

			fmt.Printf("look ahead %s, use Reduce:ARRAYS -> , go to state %d\n", look, s)
		case 16:

			fmt.Printf("look ahead %s, use Reduce:ARRAYS -> ARRAYS ARRAY , go to state %d\n", look, s)
		case 17:

			fmt.Printf("look ahead %s, use Reduce:ARRAY -> LPAREN WORDS RPAREN , go to state %d\n", look, s)
		case 18:

			fmt.Printf("look ahead %s, use Reduce:END -> NEWLINE , go to state %d\n", look, s)

		}
	}
}
