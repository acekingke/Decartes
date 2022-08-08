package parser

import (
	"fmt"
	"testing"
)

func TestWorkCart_Execute(t *testing.T) {
	type fields struct {
		CartArray [][]string
		Parameter []string
		Commands  string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			fields: fields{
				CartArray: [][]string{
					{"1", "2", "3"},

					{"4", "5", "6"},
					{"7", "8", "9"},
				},
				Parameter: []string{"a", "b", "c"},
				Commands:  "puts \"hello\"",
			},
			wantErr: false,
		},
		{
			name: "test1",
			fields: fields{
				CartArray: [][]string{
					{"1", "2", "3"},

					{"4", "5", "6"},
				},
				Parameter: []string{"a", "b", "c"},
				Commands:  "shell 'echo \"hello\"'",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WorkCart{
				CartArray: tt.fields.CartArray,
				Parameter: tt.fields.Parameter,
				Commands:  tt.fields.Commands,
			}
			if err := w.Execute(); (err != nil) != tt.wantErr {
				t.Errorf("WorkCart.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_Parser(t *testing.T) {

	ParserInit()
	val2 := Parser("set x world\n\n\t\tputs hello $x\n\t\n")
	fmt.Println(val2)
	val := Parser("set x world\n\n\t\tputs hello $a $b $c\n\t\n")
	fmt.Println(val)
}

func Test_CarNParser(t *testing.T) {
	ParserInit()
	str := `cart (1 2 3) (4 5 6) (7 8 9) each (a b c) {
		puts hello $a $b $c
	}
	`
	val := Parser(str)

	fmt.Println(val)
}

func Test_ShellCode(t *testing.T) {
	ParserInit()
	str := `shell {
		echo hello  
	}
	`
	val := Parser(str)

	fmt.Println(val)
}

func Test_braceCode(t *testing.T) {
	ParserInit()
	str := `
	set y [set x world]
	puts $y
	`
	_ = Parser(str)
}

func Test_err1(t *testing.T) {
	ParserInit()
	str := `
	set ax 1
	puts $ax
	puts now test the cartesian product for test
	cart (1 2 3) (4 5 6) (7 8 9) each (a b c) {
			shell {echo  hello $a $b $c}
	}
	`
	_ = Parser(str)
}

func Test_SingleQuote(t *testing.T) {
	ParserInit()
	str := `
	set x world
	puts 'hello $x'
	puts "hello $x"
	`
	_ = Parser(str)
}

func Test_Expr(t *testing.T) {
	ParserInit()
	str := ` 
	puts [expr {1 + 2 * 3 }]
	`
	_ = Parser(str)
}

func Test_Expr2(t *testing.T) {
	ParserInit()
	str := ` 
	puts [expr {1+2>3}]
	`
	_ = Runstring(str)
}

func Test_Expr3(t *testing.T) {
	ParserInit()
	str := ` 
	set x 3
	puts [expr {1 + 2 * $x }]
	`
	_ = Parser(str)
}
func Test_Ifcmd(t *testing.T) {
	ParserInit()
	str := `
	if {1+2>2} {
		puts hello
	}
	`
	_ = Parser(str)
}

func Test_Ifcmd2(t *testing.T) {
	ParserInit()
	str := `
	set x 3
	if {1+2>$x} {
		puts hello
	}else{
		puts 'it not over 3'
	}
	`
	_ = Parser(str)
}

func Test_Ifcmd3(t *testing.T) {
	ParserInit()
	str := `
	if {1+2>3} {
		puts hello
	}else{
		puts 'it not over 3'
	}
	`
	_ = Parser(str)
}

func Test_WhileCmd(t *testing.T) {
	ParserInit()
	str := `
	set x 3
	while {$x>0} {
		puts $x
		set x [expr {$x-1}]
	}
	`
	_ = Parser(str)
}
func Test_WhileCmd2(t *testing.T) {
	ParserInit()
	str := `
	set x 1
	while {$x<2} \ 
	{
		set x [expr {$x+1}]
		puts $x 
	}
	`
	_ = Parser(str)
}

func Test_WhileCmd3(t *testing.T) {
	ParserInit()
	str := `
	set x 1
	puts $x
	while {$x} \ 
	{
		set x [expr {$x+1}]
		puts $x
		if {$x>2} {
			break
		}
	}
	`
	_ = Parser(str)
}

func Test_BreakLine(t *testing.T) {
	ParserInit()
	str := `
	puts hello \
	world \ 
	nice
	`
	_ = Parser(str)
}
