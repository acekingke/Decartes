/*Generator Code , do not modify*/
// Code header part

package parser

import "fmt"
import "os"

// const part
const LPAREN = 12
const SINGLEQUOTE_STRING = 14
const ELSE = 17
const BRACE_STRING = 6
const DQUOTE_STRING = 7
const WHILE = 8
const CARTESIAN = 9
const RPAREN = 13
const NEWLINE = 15
const IF = 16
const PERMUTATION = 18
const STEP = 11
const BREAK = 22
const EACH = 23
const NORMALCMD = 4
const IDENTIFIER = 5
const SQUAREQUOTE_STRING = 3
const ERROR_ACTION = 140
const ACCEPT_ACTION = 240

// Terminal Size
const NTERMINALS = 18

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

var StatePackAction = []int{
	240, 22, 5, 32, 26, 21, 0, 19, 20, 11, 6, 8, 28, 13, 12, 7, 10, -22, -1, 4, -1, 9, 31, -23, 2, 3, 34, -1, -1, -1, -22, -1, -1, -1, -1, -2, -23, -2, 26, -9, 23, 1, 26, -10, -2, -2, -2, 13, -2, -2, -2, -2, -3, -6, -3, 24, 27, 38, -7, 14, 25, -3, -3, -3, -12, -3, -3, -3, -3, -25, -8, -25, -11, -13, 0, 0, 0, 0, -25, -25, -25, 0, -25, -25, -25, -25, -4, 22, -4, 17, 0, 21, 0, 19, 20, -4, -4, -4, 30, -4, -4, -4, -4, -14, 0, -14, 16, -14, 0, -14, -14, 0, 18, 0, -14, -14, 0, 22, 0, 17, 15, 21, 0, 19, 20, 0, -18, 0, -18, -5, -18, -18, -18, -18, 0, 0, 16, -18, -18, -19, 0, -19, 18, -19, -19, -19, -19, 0, 0, 0, -19, -19, -20, 0, -20, 0, -20, -20, -20, -20, 0, 0, 0, -20, -20, -21, 0, -21, 0, -21, -21, -21, -21, 0, 0, 0, -21, -21, -14, 0, -14, 0, -14, 0, -14, -14, 0, 0, -24, -14, -14, -24, -15, 0, -15, 29, -15, -24, -15, -15, 0, -24, 0, -15, -15, -16, 0, -16, 0, -16, 0, -16, -16, 0, 0, 0, -16, -16, -17, 0, -17, 0, -17, 0, -17, -17, 0, 22, 0, -17, -17, 21, 22, 19, 20, 0, 21, 22, 19, 20, 0, 21, 22, 19, 20, 0, 21, 22, 19, 20, 0, 21, 36, 19, 20, 0, 0, 33, 0, 0, 0, 0, 35, 0, 0, 0, 0, 37, 0, 0, 0, 0, 39,
}
var StatePackOffset = []int{
	17, -1, 34, 51, 33, 101, 12, -1, -1, 25, 29, -1, 225, 68, 85, 115, 190, 203, 216, 124, 137, 150, 163, 37, 33, 18, 176, 4, 39, 85, 183, 44, 230, 56, 235, 58, 240, 50, 245, 59,
}
var StackPackCheck = []int{
	1, 11, 1, 8, 7, 11, -1, 11, 11, 1, 1, 1, 27, 1, 1, 1, 1, 6, 0, 1, 0, 1, 7, 25, 1, 1, 11, 0, 0, 0, 6, 0, 0, 0, 0, 2, 25, 2, 24, 9, 6, 0, 23, 10, 2, 2, 2, 4, 2, 2, 2, 2, 3, 28, 3, 23, 24, 37, 31, 4, 23, 3, 3, 3, 37, 3, 3, 3, 3, 13, 33, 13, 35, 39, -1, -1, -1, -1, 13, 13, 13, -1, 13, 13, 13, 13, 14, 29, 14, 29, -1, 29, -1, 29, 29, 14, 14, 14, 29, 14, 14, 14, 14, 5, -1, 5, 29, 5, -1, 5, 5, -1, 29, -1, 5, 5, -1, 15, -1, 15, 5, 15, -1, 15, 15, -1, 19, -1, 19, 15, 19, 19, 19, 19, -1, -1, 15, 19, 19, 20, -1, 20, 15, 20, 20, 20, 20, -1, -1, -1, 20, 20, 21, -1, 21, -1, 21, 21, 21, 21, -1, -1, -1, 21, 21, 22, -1, 22, -1, 22, 22, 22, 22, -1, -1, -1, 22, 22, 26, -1, 26, -1, 26, -1, 26, 26, -1, -1, 30, 26, 26, 30, 16, -1, 16, 26, 16, 30, 16, 16, -1, 30, -1, 16, 16, 17, -1, 17, -1, 17, -1, 17, 17, -1, -1, -1, 17, 17, 18, -1, 18, -1, 18, -1, 18, 18, -1, 12, -1, 18, 18, 12, 32, 12, 12, -1, 32, 34, 32, 32, -1, 34, 36, 34, 34, -1, 36, 38, 36, 36, -1, 38, 12, 38, 38, -1, -1, 32, -1, -1, -1, -1, 34, -1, -1, -1, -1, 36, -1, -1, -1, -1, 38,
}
var StackPackActDef = []int{
	140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140,
}
var StackPackGotoDef = []int{
	140, 140, 140, 140, 140, 140, 140, 140, 140, 140,
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
		dollarDolar.YySymIndex = 24
		Dollar := StateSymStack[topIndex-0 : StackPointer]
		_ = Dollar

		/*

		   LineNo:42
		   COMMANDS ->

		*/

		PopStateSym(0)
	case 2:
		dollarDolar.YySymIndex = 24
		Dollar := StateSymStack[topIndex-2 : StackPointer]
		_ = Dollar

		/*

		   LineNo:42
		   COMMANDS -> COMMANDS COMMAND

		*/

		PopStateSym(2)
	case 3:
		dollarDolar.YySymIndex = 25
		Dollar := StateSymStack[topIndex-1 : StackPointer]
		_ = Dollar

		/*

		   LineNo:43
		   COMMAND -> END

		*/

		PopStateSym(1)
	case 4:
		dollarDolar.YySymIndex = 25
		Dollar := StateSymStack[topIndex-2 : StackPointer]
		_ = Dollar

		/*

		   LineNo:44
		   COMMAND -> CMD END

		*/

		PopStateSym(2)
	case 5:
		dollarDolar.YySymIndex = 20
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
			err := ExecuteNormalCmd(Dollar[1].Value, Dollar[2].Array)
			if err != nil {
				ParserError(err)
			}
		}
		PopStateSym(2)
	case 6:
		dollarDolar.YySymIndex = 20
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
			w := NewWorkCart(Dollar[2].Arrays, Dollar[4].Array, Dollar[5].Value[1:len(Dollar[5].Value)-1])
			err := w.Execute()
			if err != nil {
				ParserError(err)
			}
		}
		PopStateSym(5)
	case 7:
		dollarDolar.YySymIndex = 20
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
			err := p.Execute()
			if err != nil {
				ParserError(err)
			}
		}
		PopStateSym(2)
	case 8:
		dollarDolar.YySymIndex = 20
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
		dollarDolar.YySymIndex = 20
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
		dollarDolar.YySymIndex = 20
		Dollar := StateSymStack[topIndex-1 : StackPointer]
		_ = Dollar

		/*

		   LineNo:76
		   CMD -> BREAK
		    {
		           globalEnv.IsBreak = true;
		       }
		*/
		{
			globalEnv.IsBreak = true
		}
		PopStateSym(1)
	case 11:
		dollarDolar.YySymIndex = 20
		Dollar := StateSymStack[topIndex-3 : StackPointer]
		_ = Dollar

		/*

		   LineNo:79
		   CMD -> WHILE STRING STRING
		    {
		           w := NewWhile($2, $3)
		           w.Execute()
		           globalEnv.IsBreak = false;
		       }
		*/
		{
			w := NewWhile(Dollar[2].Value, Dollar[3].Value)
			w.Execute()
			globalEnv.IsBreak = false
		}
		PopStateSym(3)
	case 12:
		dollarDolar.YySymIndex = 22
		Dollar := StateSymStack[topIndex-3 : StackPointer]
		_ = Dollar

		/*

		   LineNo:84
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
	case 13:
		dollarDolar.YySymIndex = 22
		Dollar := StateSymStack[topIndex-5 : StackPointer]
		_ = Dollar

		/*

		   LineNo:88
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
	case 14:
		dollarDolar.YySymIndex = 19
		Dollar := StateSymStack[topIndex-0 : StackPointer]
		_ = Dollar

		/*

		   LineNo:93
		   WORDS ->
		    {
		       $$= make([]string,0)
		   }
		*/
		{
			dollarDolar.Array = make([]string, 0)
		}
		PopStateSym(0)
	case 15:
		dollarDolar.YySymIndex = 19
		Dollar := StateSymStack[topIndex-2 : StackPointer]
		_ = Dollar

		/*

		   LineNo:96
		   WORDS -> WORDS WORD
		    {
		           $$ = append($1, $2)
		       }
		*/
		{
			dollarDolar.Array = append(Dollar[1].Array, Dollar[2].Value)
		}
		PopStateSym(2)
	case 16:
		dollarDolar.YySymIndex = 21
		Dollar := StateSymStack[topIndex-1 : StackPointer]
		_ = Dollar

		/*

		   LineNo:99
		   WORD -> IDENTIFIER
		    {$$ = $1 }
		*/
		{
			dollarDolar.Value = Dollar[1].Value
		}
		PopStateSym(1)
	case 17:
		dollarDolar.YySymIndex = 21
		Dollar := StateSymStack[topIndex-1 : StackPointer]
		_ = Dollar

		/*

		   LineNo:100
		   WORD -> STRING
		    { $$ = $1 }
		*/
		{
			dollarDolar.Value = Dollar[1].Value
		}
		PopStateSym(1)
	case 18:
		dollarDolar.YySymIndex = 27
		Dollar := StateSymStack[topIndex-1 : StackPointer]
		_ = Dollar

		/*

		   LineNo:101
		   STRING -> BRACE_STRING
		    {$$ = $1 }
		*/
		{
			dollarDolar.Value = Dollar[1].Value
		}
		PopStateSym(1)
	case 19:
		dollarDolar.YySymIndex = 27
		Dollar := StateSymStack[topIndex-1 : StackPointer]
		_ = Dollar

		/*

		   LineNo:102
		   STRING -> DQUOTE_STRING
		    {$$ = $1 }
		*/
		{
			dollarDolar.Value = Dollar[1].Value
		}
		PopStateSym(1)
	case 20:
		dollarDolar.YySymIndex = 27
		Dollar := StateSymStack[topIndex-1 : StackPointer]
		_ = Dollar

		/*

		   LineNo:103
		   STRING -> SINGLEQUOTE_STRING
		    {$$ = $1 }
		*/
		{
			dollarDolar.Value = Dollar[1].Value
		}
		PopStateSym(1)
	case 21:
		dollarDolar.YySymIndex = 27
		Dollar := StateSymStack[topIndex-1 : StackPointer]
		_ = Dollar

		/*

		   LineNo:104
		   STRING -> SQUAREQUOTE_STRING
		    {$$ = $1 }
		*/
		{
			dollarDolar.Value = Dollar[1].Value
		}
		PopStateSym(1)
	case 22:
		dollarDolar.YySymIndex = 28
		Dollar := StateSymStack[topIndex-0 : StackPointer]
		_ = Dollar

		/*

		   LineNo:105
		   ARRAYS ->
		    {$$ =make([][]string ,0)}
		*/
		{
			dollarDolar.Arrays = make([][]string, 0)
		}
		PopStateSym(0)
	case 23:
		dollarDolar.YySymIndex = 28
		Dollar := StateSymStack[topIndex-2 : StackPointer]
		_ = Dollar

		/*

		   LineNo:106
		   ARRAYS -> ARRAYS ARRAY
		    {
		        $$ = append($1, $2)
		        }
		*/
		{
			dollarDolar.Arrays = append(Dollar[1].Arrays, Dollar[2].Array)
		}
		PopStateSym(2)
	case 24:
		dollarDolar.YySymIndex = 23
		Dollar := StateSymStack[topIndex-3 : StackPointer]
		_ = Dollar

		/*

		   LineNo:109
		   ARRAY -> LPAREN WORDS RPAREN
		    {
		       $$ = $2
		   }
		*/
		{
			dollarDolar.Array = Dollar[2].Array
		}
		PopStateSym(3)
	case 25:
		dollarDolar.YySymIndex = 26
		Dollar := StateSymStack[topIndex-1 : StackPointer]
		_ = Dollar

		/*

		   LineNo:112
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
	case 3:
		conv = 2
	case 4:
		conv = 3
	case 5:
		conv = 4
	case 12:
		conv = 5
	case 14:
		conv = 6
	case 17:
		conv = 7
	case 6:
		conv = 8
	case 7:
		conv = 9
	case 8:
		conv = 10
	case 9:
		conv = 11
	case 11:
		conv = 12
	case 13:
		conv = 13
	case 15:
		conv = 14
	case 16:
		conv = 15
	case 18:
		conv = 16
	case 22:
		conv = 17
	case 23:
		conv = 18

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
		conv = "SQUAREQUOTE_STRING"
	case 3:
		conv = "NORMALCMD"
	case 4:
		conv = "IDENTIFIER"
	case 5:
		conv = "LPAREN"
	case 6:
		conv = "SINGLEQUOTE_STRING"
	case 7:
		conv = "ELSE"
	case 8:
		conv = "BRACE_STRING"
	case 9:
		conv = "DQUOTE_STRING"
	case 10:
		conv = "WHILE"
	case 11:
		conv = "CARTESIAN"
	case 12:
		conv = "STEP"
	case 13:
		conv = "RPAREN"
	case 14:
		conv = "NEWLINE"
	case 15:
		conv = "IF"
	case 16:
		conv = "PERMUTATION"
	case 17:
		conv = "BREAK"
	case 18:
		conv = "EACH"
	case 19:
		conv = "WORDS"
	case 20:
		conv = "CMD"
	case 21:
		conv = "WORD"
	case 22:
		conv = "IFCMD"
	case 23:
		conv = "ARRAY"
	case 24:
		conv = "COMMANDS"
	case 25:
		conv = "COMMAND"
	case 26:
		conv = "END"
	case 27:
		conv = "STRING"
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

			fmt.Printf("look ahead %s, use Reduce:CMD -> BREAK , go to state %d\n", look, s)
		case 11:

			fmt.Printf("look ahead %s, use Reduce:CMD -> WHILE STRING STRING , go to state %d\n", look, s)
		case 12:

			fmt.Printf("look ahead %s, use Reduce:IFCMD -> IF STRING STRING , go to state %d\n", look, s)
		case 13:

			fmt.Printf("look ahead %s, use Reduce:IFCMD -> IF STRING STRING ELSE STRING , go to state %d\n", look, s)
		case 14:

			fmt.Printf("look ahead %s, use Reduce:WORDS -> , go to state %d\n", look, s)
		case 15:

			fmt.Printf("look ahead %s, use Reduce:WORDS -> WORDS WORD , go to state %d\n", look, s)
		case 16:

			fmt.Printf("look ahead %s, use Reduce:WORD -> IDENTIFIER , go to state %d\n", look, s)
		case 17:

			fmt.Printf("look ahead %s, use Reduce:WORD -> STRING , go to state %d\n", look, s)
		case 18:

			fmt.Printf("look ahead %s, use Reduce:STRING -> BRACE_STRING , go to state %d\n", look, s)
		case 19:

			fmt.Printf("look ahead %s, use Reduce:STRING -> DQUOTE_STRING , go to state %d\n", look, s)
		case 20:

			fmt.Printf("look ahead %s, use Reduce:STRING -> SINGLEQUOTE_STRING , go to state %d\n", look, s)
		case 21:

			fmt.Printf("look ahead %s, use Reduce:STRING -> SQUAREQUOTE_STRING , go to state %d\n", look, s)
		case 22:

			fmt.Printf("look ahead %s, use Reduce:ARRAYS -> , go to state %d\n", look, s)
		case 23:

			fmt.Printf("look ahead %s, use Reduce:ARRAYS -> ARRAYS ARRAY , go to state %d\n", look, s)
		case 24:

			fmt.Printf("look ahead %s, use Reduce:ARRAY -> LPAREN WORDS RPAREN , go to state %d\n", look, s)
		case 25:

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
