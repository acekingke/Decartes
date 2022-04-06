package Parser

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
