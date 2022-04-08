package Parser

import "testing"

func TestNewStep(t *testing.T) {
	str := "step a {puts a}\n"
	ParserInit()
	_ = Parser(str)

}

func TestPerm(t *testing.T) {
	str := `
	step a {puts do s 1}
	step b {puts "do s 2"}
	step c {puts "do s 3"}
	perm (a b c)

	`
	ParserInit()
	_ = Parser(str)

}
