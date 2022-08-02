package parser

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"syscall"

	"github.com/acekingke/Decartes/cartesian"
	"github.com/acekingke/Decartes/expr"
)

const (
	TypeValue = 1
	TypeArray = 2
)

type WorkValue struct {
	Type   int
	Value  string
	Array  []string
	Arrays [][]string
}

type SymbolTable struct {
	Symbols map[string]*WorkValue
	Parent  *SymbolTable
}

type Evironment struct {
	Result  string
	IsBreak bool
}

var globalEnv *Evironment = &Evironment{}

type NormalCmdFun func(env *Evironment, symbols *SymbolTable, argv []string) error

var CommandMap = map[string]NormalCmdFun{}

var CurrentSymbolTable *SymbolTable = nil
var GlobalSymbolTable *SymbolTable = MakeSymbolTable()

func MakeSymbolTable() *SymbolTable {
	CurrentSymbolTable = &SymbolTable{
		Symbols: make(map[string]*WorkValue),
		Parent:  CurrentSymbolTable,
	}
	return CurrentSymbolTable
}

func (s *SymbolTable) Add(name string, value *WorkValue) {
	s.Symbols[name] = value
}

func (s *SymbolTable) Get(name string) *WorkValue {
	for start := s; start != nil; start = start.Parent {
		if v, ok := start.Symbols[name]; ok {
			return v
		}
	}
	return nil
}

func PopSymbolTable() {
	if CurrentSymbolTable.Parent != nil {
		CurrentSymbolTable = CurrentSymbolTable.Parent
	}
}

// work cart for execute
// cart () ... () each (a b c)  {commands}
type WorkCart struct {
	CartArray [][]string
	Parameter []string
	Commands  string
}

func NewWorkCart(cartArr [][]string, p []string, c string) *WorkCart {
	return &WorkCart{
		CartArray: cartArr,
		Parameter: p,
		Commands:  c,
	}

}

func (w *WorkCart) Execute() error {
	if CheckBreak() {
		return nil
	}
	if len(w.CartArray) != len(w.Parameter) {
		return errors.New("cart: parameter count not match")
	}
	res := cartesian.CartN(w.CartArray...)
	localSymbolTable := MakeSymbolTable()
	for _, v := range res {
		for i, p := range w.Parameter {
			localSymbolTable.Add(p, &WorkValue{
				Type:  TypeValue,
				Value: v[i],
			})
		}
		cmdStr := fmt.Sprintf("%s\n", w.Commands)
		PushContex()
		ParserInit()
		_ = Parser(cmdStr)
		PopContex()

	}
	globalEnv.Result = "0"
	PopSymbolTable()
	return nil
}

// IF Cmd
type IfCmdStruct struct {
	expr      string
	ifTrueCmd string
	elseCmd   string
}

func CheckBreak() bool {
	return globalEnv.IsBreak
}

func ResetBreak() {
	globalEnv.IsBreak = false
}

func NewIfCmd(expr string, ifcmd string, elseString string) *IfCmdStruct {
	return &IfCmdStruct{
		expr:      expr,
		ifTrueCmd: ifcmd,
		elseCmd:   elseString,
	}
}

func (ifcmd *IfCmdStruct) Execute() error {
	if CheckBreak() {
		return nil
	}
	newExprStr := fmt.Sprintf("expr %s\n", ifcmd.expr)
	Runstring(newExprStr)
	if globalEnv.Result != "0" {
		arg_, err := processString(globalEnv, GlobalSymbolTable, ifcmd.ifTrueCmd)
		if err != nil {
			return err
		}
		Runstring(arg_)
	} else {
		if ifcmd.elseCmd != "" {
			arg_, err := processString(globalEnv, GlobalSymbolTable, ifcmd.elseCmd)
			if err != nil {
				return err
			}
			Runstring(arg_)
		}
	}
	return nil
}

type WhileCmdStruct struct {
	expr         string
	WhileTrueCmd string
}

func NewWhile(expr string, whileTrue string) *WhileCmdStruct {
	return &WhileCmdStruct{
		expr:         expr,
		WhileTrueCmd: whileTrue,
	}
}

func (whileCmd *WhileCmdStruct) Execute() error {
	if CheckBreak() {
		return nil
	}
	newExprStr := fmt.Sprintf("expr %s\n", whileCmd.expr)
	Runstring(newExprStr)
	for globalEnv.Result != "0" && !globalEnv.IsBreak {
		arg_, err := processString(globalEnv, GlobalSymbolTable, whileCmd.WhileTrueCmd)
		if err != nil {
			return err
		}
		Runstring(arg_)
		Runstring(newExprStr)
	}
	return nil
}

func GetToken(input string, valTy *ValType, pos *int) int {
	if *pos >= len(input) {
		return -1
	}
	res, Tok, err := Scan(input[*pos:])
	if err != nil {
		return -1
	}
	*pos += len(input[*pos:]) - len(res)
	valTy.Type = TypeValue
	valTy.Value = Tok.TokenString
	return Tok.TokenType
}

// Runstring
func Runstring(cmdStr string) error {
	PushContex()
	ParserInit()
	_ = Parser(cmdStr)
	PopContex()
	return nil
}

//  process string
func processString(env *Evironment, symbols *SymbolTable, arg string) (string, error) {
	newarg := arg
	if strings.HasPrefix(arg, "{") || strings.HasPrefix(arg, "\"") {
		newarg = arg[1 : len(arg)-1]

	}
	// Single quote string do not replace variables.
	if strings.HasPrefix(arg, "'") && strings.HasSuffix(arg, "'") {
		newarg = arg[1 : len(arg)-1]
		return newarg, nil
	}

	if strings.HasPrefix(arg, "[") {
		newarg = arg[1 : len(arg)-1]
		cmdStr := fmt.Sprintf("%s\n", newarg)
		Runstring(cmdStr)
		return env.Result, nil
	}
	// fetch newargs $var
	var err error = nil
	reg := regexp.MustCompile(`\$[a-zA-Z0-9_]+`)
	newarg = reg.ReplaceAllStringFunc(newarg, func(s string) string {
		if sy := symbols.Get(s[1:]); sy != nil {
			return sy.Value
		} else {
			err = fmt.Errorf("not found symbol %s", s)
		}
		return s
	})
	return newarg, err
}

// Puts cmd result
func PutsCMD(env *Evironment, symbols *SymbolTable, argv []string) error {
	// for {} string or "" string, first do the replace string
	replace_args := make([]string, len(argv))
	result := ""
	for index, arg := range argv {
		arg_, err := processString(env, symbols, arg)
		if err != nil {
			return err
		}
		replace_args[index] = arg_
	}
	for _, arg := range replace_args {
		result += (arg + " ")
	}
	env.Result = result
	fmt.Println(result)

	return nil
}

func SetCMD(env *Evironment, symbols *SymbolTable, argv []string) error {
	if len(argv) != 2 {
		return errors.New("set: parameter count not match")
	}
	arg_, err := processString(env, symbols, argv[1])
	if err != nil {
		return err
	}
	symbols.Add(argv[0], &WorkValue{
		Type:  TypeValue,
		Value: arg_,
	})
	env.Result = arg_
	return nil
}

// func TearDownCmd(env *Evironment, symbols *SymbolTable, argv []string) error {
// 	return nil
// }

func ShellCmd(env *Evironment, symbols *SymbolTable, argv []string) error {
	// run shell command
	if len(argv) == 0 {
		return errors.New("shell: should run with strings")
	}
	arg_, err := processString(env, symbols, argv[0])
	if err != nil {
		return err
	}
	cmd := exec.Command("sh", "-c", arg_)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	ws := cmd.ProcessState.Sys().(syscall.WaitStatus)
	exitCode := ws.ExitStatus()
	env.Result = fmt.Sprintf("%d", exitCode)
	return nil
}

func ExprCMD(env *Evironment, symbols *SymbolTable, argv []string) error {
	// run shell command
	if len(argv) == 0 {
		return errors.New("expr: should run with strings")
	}
	arg_, err := processString(env, symbols, argv[0])
	if err != nil {
		return err
	}
	// run expr
	res := expr.RunParser(arg_)
	env.Result = fmt.Sprintf("%v", res)
	return nil
}

func ExecuteNormalCmd(name string, argv []string) error {
	if f, ok := CommandMap[name]; ok {
		if !CheckBreak() {
			return f(globalEnv, CurrentSymbolTable, argv)
		}
	} else {
		return errors.New("command not found")
	}
	return nil
}

func RegisterCommand(name string, f NormalCmdFun) {
	CommandMap[name] = f
}

func init() {
	RegisterCommand("puts", PutsCMD)
	RegisterCommand("set", SetCMD)
	RegisterCommand("shell", ShellCmd)
	RegisterCommand("expr", ExprCMD)
}
