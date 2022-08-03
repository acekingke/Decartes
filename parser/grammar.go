/*Generator Code , do not modify*/
// Code header part

package parser

import "fmt"
import "os"

// const part
const PERMUTATION = 24
const BREAK = 3
const SINGLEQUOTE_STRING = 5
const STEP = 13
const EACH = 14
const DQUOTE_STRING = 7
const NORMALCMD = 17
const LPAREN = 18
const RPAREN = 19
const SQUAREQUOTE_STRING = 20
const ELSE = 23
const BRACE_STRING = 25
const IDENTIFIER = 4
const NEWLINE = 9
const IF = 10
const WHILE = 11
const CARTESIAN = 12
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
	240, 6, 13, 12, 11, 7, 10, 21, 8, -22, 20, 22, 5, -22, 30, -9, 19, 17, -23, 18, 23, 4, -23, 16, 26, 2, 9, 3, -1, -1, -1, -1, -1, -1, -1, 31, -1, 32, -12, 13, -1, -2, -2, -2, -2, -2, -2, -2, 24, -2, 1, 38, 26, -2, -3, -3, -3, -3, -3, -3, -3, -10, -3, 25, 14, 26, -3, -25, -25, -25, -25, -25, -25, -25, 28, -25, 27, -6, -7, -25, -4, -4, -4, -4, -4, -4, -4, 21, -4, -14, 20, 22, -4, -8, -14, -11, 19, -14, -14, 34, -13, -14, -5, -14, -14, 15, 0, 21, 0, 0, 20, 22, 0, -18, 0, 0, 19, 17, -18, 18, 0, -18, -18, 16, -19, -18, -18, -18, -18, -19, 0, 0, -19, -19, 0, -20, -19, -19, -19, -19, -20, 0, 0, -20, -20, 0, -21, -20, -20, -20, -20, -21, 0, 0, -21, -21, 0, -14, -21, -21, -21, -21, -14, 0, 0, -14, -14, 0, 0, -14, -15, -14, -14, 29, 0, -15, 0, 0, -15, -15, -24, -16, -15, 0, -15, -15, -16, -24, 0, -16, -16, -24, -17, -16, -24, -16, -16, -17, 0, 21, -17, -17, 20, 22, -17, 0, -17, -17, 19, 21, 0, 36, 20, 22, 0, 0, 21, 0, 19, 20, 22, 33, 0, 21, 0, 19, 20, 22, 35, 0, 21, 0, 19, 20, 22, 37, 0, 0, 0, 19, 0, 0, 39,
}
var StatePackOffset = []int{
	27, -1, 40, 53, 36, 86, -1, 10, 19, 12, 58, 79, 191, 66, 79, 99, 167, 178, 189, 110, 121, 132, 143, 38, 51, 8, 154, 57, 74, -1, 177, 75, 201, 90, 208, 92, 215, 35, 222, 97,
}
var StackPackCheck = []int{
	1, 1, 1, 1, 1, 1, 1, 29, 1, 6, 29, 29, 1, 6, 29, 9, 29, 29, 25, 29, 6, 1, 25, 29, 7, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 7, 0, 8, 37, 4, 0, 2, 2, 2, 2, 2, 2, 2, 23, 2, 0, 37, 23, 2, 3, 3, 3, 3, 3, 3, 3, 10, 3, 23, 4, 24, 3, 13, 13, 13, 13, 13, 13, 13, 27, 13, 24, 28, 31, 13, 14, 14, 14, 14, 14, 14, 14, 11, 14, 5, 11, 11, 14, 33, 5, 35, 11, 5, 5, 11, 39, 5, 15, 5, 5, 5, -1, 15, -1, -1, 15, 15, -1, 19, -1, -1, 15, 15, 19, 15, -1, 19, 19, 15, 20, 19, 19, 19, 19, 20, -1, -1, 20, 20, -1, 21, 20, 20, 20, 20, 21, -1, -1, 21, 21, -1, 22, 21, 21, 21, 21, 22, -1, -1, 22, 22, -1, 26, 22, 22, 22, 22, 26, -1, -1, 26, 26, -1, -1, 26, 16, 26, 26, 26, -1, 16, -1, -1, 16, 16, 30, 17, 16, -1, 16, 16, 17, 30, -1, 17, 17, 30, 18, 17, 30, 17, 17, 18, -1, 12, 18, 18, 12, 12, 18, -1, 18, 18, 12, 32, -1, 12, 32, 32, -1, -1, 34, -1, 32, 34, 34, 32, -1, 36, -1, 34, 36, 36, 34, -1, 38, -1, 36, 38, 38, 36, -1, -1, -1, 38, -1, -1, 38,
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
		dollarDolar.YySymIndex = 23
		Dollar := StateSymStack[topIndex-0 : StackPointer]
		_ = Dollar

		/*

		   LineNo:42
		   COMMANDS ->

		*/

		PopStateSym(0)
	case 2:
		dollarDolar.YySymIndex = 23
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
		dollarDolar.YySymIndex = 22
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
		dollarDolar.YySymIndex = 22
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
		dollarDolar.YySymIndex = 22
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
		dollarDolar.YySymIndex = 22
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
		dollarDolar.YySymIndex = 22
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
		dollarDolar.YySymIndex = 22
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
		dollarDolar.YySymIndex = 22
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
		dollarDolar.YySymIndex = 27
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
		dollarDolar.YySymIndex = 27
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
		dollarDolar.YySymIndex = 24
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
		dollarDolar.YySymIndex = 24
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
		dollarDolar.YySymIndex = 20
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
		dollarDolar.YySymIndex = 20
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
		dollarDolar.YySymIndex = 20
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
		dollarDolar.YySymIndex = 20
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
		dollarDolar.YySymIndex = 21
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
		dollarDolar.YySymIndex = 21
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
		dollarDolar.YySymIndex = 25
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
		dollarDolar.YySymIndex = 28
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
	case 12:
		conv = 2
	case 9:
		conv = 3
	case 10:
		conv = 4
	case 11:
		conv = 5
	case 24:
		conv = 6
	case 3:
		conv = 7
	case 5:
		conv = 8
	case 13:
		conv = 9
	case 14:
		conv = 10
	case 7:
		conv = 11
	case 20:
		conv = 12
	case 17:
		conv = 13
	case 18:
		conv = 14
	case 19:
		conv = 15
	case 23:
		conv = 16
	case 25:
		conv = 17
	case 4:
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
		conv = "CARTESIAN"
	case 3:
		conv = "NEWLINE"
	case 4:
		conv = "IF"
	case 5:
		conv = "WHILE"
	case 6:
		conv = "PERMUTATION"
	case 7:
		conv = "BREAK"
	case 8:
		conv = "SINGLEQUOTE_STRING"
	case 9:
		conv = "STEP"
	case 10:
		conv = "EACH"
	case 11:
		conv = "DQUOTE_STRING"
	case 12:
		conv = "SQUAREQUOTE_STRING"
	case 13:
		conv = "NORMALCMD"
	case 14:
		conv = "LPAREN"
	case 15:
		conv = "RPAREN"
	case 16:
		conv = "ELSE"
	case 17:
		conv = "BRACE_STRING"
	case 18:
		conv = "IDENTIFIER"
	case 19:
		conv = "WORDS"
	case 20:
		conv = "STRING"
	case 21:
		conv = "ARRAYS"
	case 22:
		conv = "CMD"
	case 23:
		conv = "COMMANDS"
	case 24:
		conv = "WORD"
	case 25:
		conv = "ARRAY"
	case 26:
		conv = "COMMAND"
	case 27:
		conv = "IFCMD"
	case 28:
		conv = "END"

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
