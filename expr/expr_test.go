package expr

import (
	"testing"
)

func Test_Parser(t *testing.T) {
	v := Parser("1+2*31/3").val
	//21.666667
	if v != (1 + 2*31.0/3.0) {
		t.Error("Parser error")
	}
}

func Test_Parser2(t *testing.T) {
	v := Parser("1>2").val
	//21.666667
	if v != 0 {
		t.Error("Parser error")
	}
}

func Test_Parser3(t *testing.T) {
	v := Parser("1+2*31/3>3").val
	//21.666667
	if v != 1 {
		t.Error("Parser error")
	}
}
func Test_Parser4(t *testing.T) {
	v := Parser("1+2*31/3<3 ").val
	//21.666667
	if v == 1 {
		t.Error("Parser error")
	}
}
func Test_Parser5(t *testing.T) {
	v := Parser("1+2*31/3<=3 ").val
	//21.666667
	if v == 1 {
		t.Error("Parser error")
	}
}
func Test_Parser6(t *testing.T) {
	v := Parser("1+2*31/3>=3 ").val
	//21.666667
	if v != 1 {
		t.Error("Parser error")
	}
}
func Test_ParserOR(t *testing.T) {
	v := Parser("1&&3 ").val
	//21.666667
	if v != 1 {
		t.Error("Parser error")
	}
}
func Test_ParserOR2(t *testing.T) {
	v := Parser("1-1&&3 ").val
	//21.666667
	if v != 0 {
		t.Error("Parser error")
	}
}
func TestGetToken(t *testing.T) {
	type args struct {
		input string
		valTy *ValType
		pos   *int
	}
	Opos := 0
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				input: "1.2e+3 + 2",
				valTy: &ValType{},
				pos:   &Opos,
			},
			want: NUM,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetToken(tt.args.input, tt.args.valTy, tt.args.pos); got != tt.want {
				t.Errorf("GetToken() = %v, want %v", got, tt.want)
			}
			if got := GetToken(tt.args.input, tt.args.valTy, tt.args.pos); got != int('+') {
				t.Errorf("GetToken() = %v, want %v", got, tt.want)
			}
			if got := GetToken(tt.args.input, tt.args.valTy, tt.args.pos); got != tt.want {
				t.Errorf("GetToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetToke2(t *testing.T) {
	var val ValType
	pos := 0
	if got := GetToken(" >= ", &val, &pos); got != GE {
		t.Errorf("GetToken() = %v, want %v", got, GE)
	}
}

func TestGetToke3(t *testing.T) {
	var val ValType
	pos := 0
	if got := GetToken(" == ", &val, &pos); got != EQ {
		t.Errorf("GetToken() = %v, want %v", got, EQ)
	}
}

func TestGetToke4(t *testing.T) {
	var val ValType
	pos := 0
	if got := GetToken("1 != 1", &val, &pos); got != NUM {
		t.Errorf("GetToken() = %v, want %v", got, NUM)
	}
	if got := GetToken("1 != 1", &val, &pos); got != NE {
		t.Errorf("GetToken() = %v, want %v", got, NE)
	}
	if got := GetToken("1 != 1", &val, &pos); got != NUM {
		t.Errorf("GetToken() = %v, want %v", got, NUM)
	}
}
