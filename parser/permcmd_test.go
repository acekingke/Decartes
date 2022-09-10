package parser

import "testing"

func TestNewStep(t *testing.T) {
	str := "step f {puts a}\n"
	PushContex()
	ParserInit()
	_ = Parser(str)
	PopContex()

}

func TestPerm(t *testing.T) {
	str := `
	step a {puts do s 1}
	step b {puts "do s 2"}
	step c {puts "do s 3"}
	perm (a b c) pre {puts start} post{puts end}

	`
	PushContex()
	ParserInit()
	_ = Parser(str)
	PopContex()
}
