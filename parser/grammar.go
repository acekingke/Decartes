/*Generator Code , do not modify*/
// Code header part

package parser

import "fmt"
import "os"

// const part
const PERMUTATION = 21
const LPAREN = 22
const DQUOTE_STRING = 23
const CARTESIAN = 26
const EACH = 27
const POST = 28
const INTERLEAVE = 29
const SINGLEQUOTE_STRING = 30
const IF = 3
const IDENTIFIER = 4
const BREAK = 8
const RPAREN = 9
const BRACE_STRING = 10
const WHILE = 14
const PRE = 15
const NEWLINE = 16
const NORMALCMD = 17
const ELSE = 18
const STEP = 19
const SQUAREQUOTE_STRING = 20
const ERROR_ACTION = 153
const ACCEPT_ACTION = 253

// Terminal Size
const NTERMINALS = 21

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
	253, 45, 13, 0, 35, 11, -17, 12, 5, 14, 9, -17, -10, -27, -28, 7, -27, -28, 8, 32, 6, -11, 10, 31, -1, 24, -1, 2, 3, -1, 4, -1, -1, -1, -1, 14, -16, 27, 33, -1, 25, -16, -1, -2, -1, -2, 34, -6, -2, -16, -2, -2, -2, -2, 15, 26, 1, -3, -2, -3, -7, -2, -3, -2, -3, -3, -3, -3, -8, -18, 27, -30, -3, -30, -18, -3, -30, -3, -30, -30, -30, -30, -9, 27, 27, -4, -30, -4, 38, -30, -4, -30, -4, -4, -4, -4, -19, -12, -19, -19, -4, 41, 28, -4, -19, -4, -19, 18, 27, -19, 20, -14, 0, -13, -19, -5, 51, 23, 16, 0, 21, -23, 0, -23, -23, 22, 42, 0, 0, -23, 19, -23, -23, 0, -23, 17, 0, 0, -24, -23, -24, -24, -25, 0, -25, -25, -24, 0, -24, -24, -25, -24, -25, -25, 0, -25, -24, 0, 0, -26, -25, -26, -26, -19, 0, -19, -19, -26, 0, -26, -26, -19, -26, -19, 0, 0, -19, -26, 0, 0, 18, -19, 37, 20, -20, 36, -20, -20, 0, 0, 23, 0, -20, 21, -20, 0, 0, -20, 22, 0, 0, 0, -20, 19, -21, 0, -21, -21, 17, 0, 0, 0, -21, -22, -21, -22, -22, -21, 0, 0, 0, -22, -21, -22, 0, -29, -22, -29, 0, 20, -29, -22, 0, 20, -29, 0, 23, -29, -29, 21, 23, 0, 0, 21, 22, -15, 0, 30, 22, 47, -15, 0, -15, 49, 30, 0, 20, -15, -15, -15, 0, 30, 0, 23, -15, -15, 21, 29, 20, 0, 20, 22, -15, 0, 39, 23, 46, 23, 21, 20, 21, 43, 0, 22, 0, 22, 23, -17, 48, 21, 50, 0, -17, -17, 22, 0, 0, 0, -17, 52, 32, 0, 0, 0, 40, 0, 32, 0, 0, 0, 44,
}
var StatePackOffset = []int{
	23, -1, 42, 56, 25, 94, -1, 56, 69, -1, 2, 11, 224, 228, 70, 84, 105, 182, 202, 211, 119, 136, 140, 157, 23, 70, 0, 161, 240, 1, -1, 33, 41, 37, 64, 31, 178, 220, 247, 282, 50, 94, 254, 288, 58, 251, 72, 263, 87, 265, 103, 274, 101,
}
var StackPackCheck = []int{
	1, 9, 1, -1, 30, 1, 29, 1, 1, 1, 1, 29, 10, 6, 26, 1, 6, 26, 1, 29, 1, 11, 1, 29, 0, 6, 0, 1, 1, 0, 1, 0, 0, 0, 0, 4, 35, 24, 31, 0, 24, 35, 0, 2, 0, 2, 32, 33, 2, 35, 2, 2, 2, 2, 4, 24, 0, 3, 2, 3, 40, 2, 3, 2, 3, 3, 3, 3, 44, 34, 7, 14, 3, 14, 34, 3, 14, 3, 14, 14, 14, 14, 46, 8, 25, 15, 14, 15, 7, 14, 15, 14, 15, 15, 15, 15, 5, 48, 5, 5, 15, 8, 25, 15, 5, 15, 5, 16, 41, 5, 16, 52, -1, 50, 5, 16, 50, 16, 5, -1, 16, 20, -1, 20, 20, 16, 41, -1, -1, 20, 16, 20, 20, -1, 20, 16, -1, -1, 21, 20, 21, 21, 22, -1, 22, 22, 21, -1, 21, 21, 22, 21, 22, 22, -1, 22, 21, -1, -1, 23, 22, 23, 23, 27, -1, 27, 27, 23, -1, 23, 23, 27, 23, 27, -1, -1, 27, 23, -1, -1, 36, 27, 36, 36, 17, 27, 17, 17, -1, -1, 36, -1, 17, 36, 17, -1, -1, 17, 36, -1, -1, -1, 17, 36, 18, -1, 18, 18, 36, -1, -1, -1, 18, 19, 18, 19, 19, 18, -1, -1, -1, 19, 18, 19, -1, 37, 19, 37, -1, 12, 37, 19, -1, 13, 37, -1, 12, 37, 37, 12, 13, -1, -1, 13, 12, 28, -1, 28, 13, 12, 28, -1, 38, 13, 38, -1, 45, 38, 28, 42, -1, 42, -1, 45, 42, 38, 45, 28, 47, -1, 49, 45, 42, -1, 38, 47, 45, 49, 47, 51, 49, 42, -1, 47, -1, 49, 51, 39, 47, 51, 49, -1, 39, 43, 51, -1, -1, -1, 43, 51, 39, -1, -1, -1, 39, -1, 43, -1, -1, -1, 43,
}
var StackPackActDef = []int{
	153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153,
}
var StackPackGotoDef = []int{
	153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153,
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
		dollarDolar.YySymIndex = 33
		Dollar := StateSymStack[topIndex-0 : StackPointer]
		_ = Dollar

		/*

		   LineNo:46
		   COMMANDS ->

		*/

		PopStateSym(0)
	case 2:
		dollarDolar.YySymIndex = 33
		Dollar := StateSymStack[topIndex-2 : StackPointer]
		_ = Dollar

		/*

		   LineNo:46
		   COMMANDS -> COMMANDS COMMAND

		*/

		PopStateSym(2)
	case 3:
		dollarDolar.YySymIndex = 28
		Dollar := StateSymStack[topIndex-1 : StackPointer]
		_ = Dollar

		/*

		   LineNo:47
		   COMMAND -> END

		*/

		PopStateSym(1)
	case 4:
		dollarDolar.YySymIndex = 28
		Dollar := StateSymStack[topIndex-2 : StackPointer]
		_ = Dollar

		/*

		   LineNo:48
		   COMMAND -> CMD END

		*/

		PopStateSym(2)
	case 5:
		dollarDolar.YySymIndex = 31
		Dollar := StateSymStack[topIndex-2 : StackPointer]
		_ = Dollar

		/*

		   LineNo:49
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
		dollarDolar.YySymIndex = 31
		Dollar := StateSymStack[topIndex-7 : StackPointer]
		_ = Dollar

		/*

		   LineNo:55
		   CMD -> CARTESIAN ARRAYS EACH ARRAY PRE_STMT POST_STMT BRACE_STRING
		    {
		          w := NewWorkCart($2, $4,$5, $6, $7[1:len($7)-1]);
		          err := w.Execute();
		          if err != nil {
		             ParserError(err)
		          }
		       }
		*/
		{
			w := NewWorkCart(Dollar[2].Arrays, Dollar[4].Array, Dollar[5].Value, Dollar[6].Value, Dollar[7].Value[1:len(Dollar[7].Value)-1])
			err := w.Execute()
			if err != nil {
				ParserError(err)
			}
		}
		PopStateSym(7)
	case 7:
		dollarDolar.YySymIndex = 31
		Dollar := StateSymStack[topIndex-4 : StackPointer]
		_ = Dollar

		/*

		   LineNo:62
		   CMD -> PERMUTATION ARRAY PRE_STMT POST_STMT
		    {
		          p := &PermType{
		               StepNames: $2,
		               Pre: $3,
		               Post: $4,
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
				Pre:       Dollar[3].Value,
				Post:      Dollar[4].Value,
			}
			err := p.Execute()
			if err != nil {
				ParserError(err)
			}
		}
		PopStateSym(4)
	case 8:
		dollarDolar.YySymIndex = 31
		Dollar := StateSymStack[topIndex-5 : StackPointer]
		_ = Dollar

		/*

		   LineNo:73
		   CMD -> INTERLEAVE ARRAY ARRAY PRE_STMT POST_STMT
		    {
		           inter := &InterleaveType{
		               Left : $2,
		               Right: $3,
		               Pre:   $4,
		               Post:  $5,
		           }
		            err := inter.Execute();
		            if err != nil {
		                ParserError(err)
		            }
		       }
		*/
		{
			inter := &InterleaveType{
				Left:  Dollar[2].Array,
				Right: Dollar[3].Array,
				Pre:   Dollar[4].Value,
				Post:  Dollar[5].Value,
			}
			err := inter.Execute()
			if err != nil {
				ParserError(err)
			}
		}
		PopStateSym(5)
	case 9:
		dollarDolar.YySymIndex = 31
		Dollar := StateSymStack[topIndex-3 : StackPointer]
		_ = Dollar

		/*

		   LineNo:85
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
	case 10:
		dollarDolar.YySymIndex = 31
		Dollar := StateSymStack[topIndex-1 : StackPointer]
		_ = Dollar

		/*

		   LineNo:91
		   CMD -> IFCMD
		    {

		       }
		*/
		{

		}
		PopStateSym(1)
	case 11:
		dollarDolar.YySymIndex = 31
		Dollar := StateSymStack[topIndex-1 : StackPointer]
		_ = Dollar

		/*

		   LineNo:94
		   CMD -> BREAK
		    {
		           globalEnv.IsBreak = true;
		       }
		*/
		{
			globalEnv.IsBreak = true
		}
		PopStateSym(1)
	case 12:
		dollarDolar.YySymIndex = 31
		Dollar := StateSymStack[topIndex-3 : StackPointer]
		_ = Dollar

		/*

		   LineNo:97
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
	case 13:
		dollarDolar.YySymIndex = 23
		Dollar := StateSymStack[topIndex-3 : StackPointer]
		_ = Dollar

		/*

		   LineNo:102
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
	case 14:
		dollarDolar.YySymIndex = 23
		Dollar := StateSymStack[topIndex-5 : StackPointer]
		_ = Dollar

		/*

		   LineNo:106
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
	case 15:
		dollarDolar.YySymIndex = 27
		Dollar := StateSymStack[topIndex-0 : StackPointer]
		_ = Dollar

		/*

		   LineNo:111
		   PRE_STMT ->

		*/

		PopStateSym(0)
	case 16:
		dollarDolar.YySymIndex = 27
		Dollar := StateSymStack[topIndex-2 : StackPointer]
		_ = Dollar

		/*

		   LineNo:111
		   PRE_STMT -> PRE BRACE_STRING
		    {
		           $$ = $2[1:len($2)-1]
		         }
		*/
		{
			dollarDolar.Value = Dollar[2].Value[1 : len(Dollar[2].Value)-1]
		}
		PopStateSym(2)
	case 17:
		dollarDolar.YySymIndex = 22
		Dollar := StateSymStack[topIndex-0 : StackPointer]
		_ = Dollar

		/*

		   LineNo:115
		   POST_STMT ->

		*/

		PopStateSym(0)
	case 18:
		dollarDolar.YySymIndex = 22
		Dollar := StateSymStack[topIndex-2 : StackPointer]
		_ = Dollar

		/*

		   LineNo:115
		   POST_STMT -> POST BRACE_STRING
		    {
		          $$ = $2[1:len($2)-1]
		         }
		*/
		{
			dollarDolar.Value = Dollar[2].Value[1 : len(Dollar[2].Value)-1]
		}
		PopStateSym(2)
	case 19:
		dollarDolar.YySymIndex = 24
		Dollar := StateSymStack[topIndex-0 : StackPointer]
		_ = Dollar

		/*

		   LineNo:119
		   WORDS ->
		    {
		       $$= make([]string,0)
		   }
		*/
		{
			dollarDolar.Array = make([]string, 0)
		}
		PopStateSym(0)
	case 20:
		dollarDolar.YySymIndex = 24
		Dollar := StateSymStack[topIndex-2 : StackPointer]
		_ = Dollar

		/*

		   LineNo:122
		   WORDS -> WORDS WORD
		    {
		           $$ = append($1, $2)
		       }
		*/
		{
			dollarDolar.Array = append(Dollar[1].Array, Dollar[2].Value)
		}
		PopStateSym(2)
	case 21:
		dollarDolar.YySymIndex = 30
		Dollar := StateSymStack[topIndex-1 : StackPointer]
		_ = Dollar

		/*

		   LineNo:125
		   WORD -> IDENTIFIER
		    {$$ = $1 }
		*/
		{
			dollarDolar.Value = Dollar[1].Value
		}
		PopStateSym(1)
	case 22:
		dollarDolar.YySymIndex = 30
		Dollar := StateSymStack[topIndex-1 : StackPointer]
		_ = Dollar

		/*

		   LineNo:126
		   WORD -> STRING
		    { $$ = $1 }
		*/
		{
			dollarDolar.Value = Dollar[1].Value
		}
		PopStateSym(1)
	case 23:
		dollarDolar.YySymIndex = 25
		Dollar := StateSymStack[topIndex-1 : StackPointer]
		_ = Dollar

		/*

		   LineNo:127
		   STRING -> BRACE_STRING
		    {$$ = $1 }
		*/
		{
			dollarDolar.Value = Dollar[1].Value
		}
		PopStateSym(1)
	case 24:
		dollarDolar.YySymIndex = 25
		Dollar := StateSymStack[topIndex-1 : StackPointer]
		_ = Dollar

		/*

		   LineNo:128
		   STRING -> DQUOTE_STRING
		    {$$ = $1 }
		*/
		{
			dollarDolar.Value = Dollar[1].Value
		}
		PopStateSym(1)
	case 25:
		dollarDolar.YySymIndex = 25
		Dollar := StateSymStack[topIndex-1 : StackPointer]
		_ = Dollar

		/*

		   LineNo:129
		   STRING -> SINGLEQUOTE_STRING
		    {$$ = $1 }
		*/
		{
			dollarDolar.Value = Dollar[1].Value
		}
		PopStateSym(1)
	case 26:
		dollarDolar.YySymIndex = 25
		Dollar := StateSymStack[topIndex-1 : StackPointer]
		_ = Dollar

		/*

		   LineNo:130
		   STRING -> SQUAREQUOTE_STRING
		    {$$ = $1 }
		*/
		{
			dollarDolar.Value = Dollar[1].Value
		}
		PopStateSym(1)
	case 27:
		dollarDolar.YySymIndex = 26
		Dollar := StateSymStack[topIndex-0 : StackPointer]
		_ = Dollar

		/*

		   LineNo:131
		   ARRAYS ->
		    {$$ =make([][]string ,0)}
		*/
		{
			dollarDolar.Arrays = make([][]string, 0)
		}
		PopStateSym(0)
	case 28:
		dollarDolar.YySymIndex = 26
		Dollar := StateSymStack[topIndex-2 : StackPointer]
		_ = Dollar

		/*

		   LineNo:132
		   ARRAYS -> ARRAYS ARRAY
		    {
		        $$ = append($1, $2)
		        }
		*/
		{
			dollarDolar.Arrays = append(Dollar[1].Arrays, Dollar[2].Array)
		}
		PopStateSym(2)
	case 29:
		dollarDolar.YySymIndex = 32
		Dollar := StateSymStack[topIndex-3 : StackPointer]
		_ = Dollar

		/*

		   LineNo:135
		   ARRAY -> LPAREN WORDS RPAREN
		    {
		       $$ = $2
		   }
		*/
		{
			dollarDolar.Array = Dollar[2].Array
		}
		PopStateSym(3)
	case 30:
		dollarDolar.YySymIndex = 29
		Dollar := StateSymStack[topIndex-1 : StackPointer]
		_ = Dollar

		/*

		   LineNo:138
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
	case 4:
		conv = 2
	case 3:
		conv = 3
	case 9:
		conv = 4
	case 10:
		conv = 5
	case 8:
		conv = 6
	case 15:
		conv = 7
	case 14:
		conv = 8
	case 17:
		conv = 9
	case 16:
		conv = 10
	case 19:
		conv = 11
	case 20:
		conv = 12
	case 18:
		conv = 13
	case 22:
		conv = 14
	case 23:
		conv = 15
	case 21:
		conv = 16
	case 27:
		conv = 17
	case 28:
		conv = 18
	case 29:
		conv = 19
	case 30:
		conv = 20
	case 26:
		conv = 21

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
		conv = "IDENTIFIER"
	case 3:
		conv = "IF"
	case 4:
		conv = "RPAREN"
	case 5:
		conv = "BRACE_STRING"
	case 6:
		conv = "BREAK"
	case 7:
		conv = "PRE"
	case 8:
		conv = "WHILE"
	case 9:
		conv = "NORMALCMD"
	case 10:
		conv = "NEWLINE"
	case 11:
		conv = "STEP"
	case 12:
		conv = "SQUAREQUOTE_STRING"
	case 13:
		conv = "ELSE"
	case 14:
		conv = "LPAREN"
	case 15:
		conv = "DQUOTE_STRING"
	case 16:
		conv = "PERMUTATION"
	case 17:
		conv = "EACH"
	case 18:
		conv = "POST"
	case 19:
		conv = "INTERLEAVE"
	case 20:
		conv = "SINGLEQUOTE_STRING"
	case 21:
		conv = "CARTESIAN"
	case 22:
		conv = "POST_STMT"
	case 23:
		conv = "IFCMD"
	case 24:
		conv = "WORDS"
	case 25:
		conv = "STRING"
	case 26:
		conv = "ARRAYS"
	case 27:
		conv = "PRE_STMT"
	case 28:
		conv = "COMMAND"
	case 29:
		conv = "END"
	case 30:
		conv = "WORD"
	case 31:
		conv = "CMD"
	case 32:
		conv = "ARRAY"
	case 33:
		conv = "COMMANDS"

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

			fmt.Printf("look ahead %s, use Reduce:CMD -> CARTESIAN ARRAYS EACH ARRAY PRE_STMT POST_STMT BRACE_STRING , go to state %d\n", look, s)
		case 7:

			fmt.Printf("look ahead %s, use Reduce:CMD -> PERMUTATION ARRAY PRE_STMT POST_STMT , go to state %d\n", look, s)
		case 8:

			fmt.Printf("look ahead %s, use Reduce:CMD -> INTERLEAVE ARRAY ARRAY PRE_STMT POST_STMT , go to state %d\n", look, s)
		case 9:

			fmt.Printf("look ahead %s, use Reduce:CMD -> STEP IDENTIFIER STRING , go to state %d\n", look, s)
		case 10:

			fmt.Printf("look ahead %s, use Reduce:CMD -> IFCMD , go to state %d\n", look, s)
		case 11:

			fmt.Printf("look ahead %s, use Reduce:CMD -> BREAK , go to state %d\n", look, s)
		case 12:

			fmt.Printf("look ahead %s, use Reduce:CMD -> WHILE STRING STRING , go to state %d\n", look, s)
		case 13:

			fmt.Printf("look ahead %s, use Reduce:IFCMD -> IF STRING STRING , go to state %d\n", look, s)
		case 14:

			fmt.Printf("look ahead %s, use Reduce:IFCMD -> IF STRING STRING ELSE STRING , go to state %d\n", look, s)
		case 15:

			fmt.Printf("look ahead %s, use Reduce:PRE_STMT -> , go to state %d\n", look, s)
		case 16:

			fmt.Printf("look ahead %s, use Reduce:PRE_STMT -> PRE BRACE_STRING , go to state %d\n", look, s)
		case 17:

			fmt.Printf("look ahead %s, use Reduce:POST_STMT -> , go to state %d\n", look, s)
		case 18:

			fmt.Printf("look ahead %s, use Reduce:POST_STMT -> POST BRACE_STRING , go to state %d\n", look, s)
		case 19:

			fmt.Printf("look ahead %s, use Reduce:WORDS -> , go to state %d\n", look, s)
		case 20:

			fmt.Printf("look ahead %s, use Reduce:WORDS -> WORDS WORD , go to state %d\n", look, s)
		case 21:

			fmt.Printf("look ahead %s, use Reduce:WORD -> IDENTIFIER , go to state %d\n", look, s)
		case 22:

			fmt.Printf("look ahead %s, use Reduce:WORD -> STRING , go to state %d\n", look, s)
		case 23:

			fmt.Printf("look ahead %s, use Reduce:STRING -> BRACE_STRING , go to state %d\n", look, s)
		case 24:

			fmt.Printf("look ahead %s, use Reduce:STRING -> DQUOTE_STRING , go to state %d\n", look, s)
		case 25:

			fmt.Printf("look ahead %s, use Reduce:STRING -> SINGLEQUOTE_STRING , go to state %d\n", look, s)
		case 26:

			fmt.Printf("look ahead %s, use Reduce:STRING -> SQUAREQUOTE_STRING , go to state %d\n", look, s)
		case 27:

			fmt.Printf("look ahead %s, use Reduce:ARRAYS -> , go to state %d\n", look, s)
		case 28:

			fmt.Printf("look ahead %s, use Reduce:ARRAYS -> ARRAYS ARRAY , go to state %d\n", look, s)
		case 29:

			fmt.Printf("look ahead %s, use Reduce:ARRAY -> LPAREN WORDS RPAREN , go to state %d\n", look, s)
		case 30:

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
