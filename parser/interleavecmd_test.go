package parser

import "testing"

func TestInterleaveType_Execute(t *testing.T) {
	str := `
	step a {puts do a}
	step b {puts "do b"}
	step c {puts "do c"}
	step x {puts do x}
	step y {puts "do y"}
	step z {puts "do z"}
	interleave (a b c) (x y z) pre {puts start} post{puts end}

	`
	ParserInit()
	_ = Parser(str)
}
