package perm

import (
	"testing"
)

func TestPerm(t *testing.T) {
	type args struct {
		a []string
		f func([]string)
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"test1", args{[]string{"a", "b", "c"}, func(a []string) {
			t.Logf("%v\n", a)
		}}},
		{"test2", args{[]string{"1", "2", "3"}, func(a []string) {
			t.Logf("%v\n", a)
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Perm(tt.args.a, tt.args.f)
		})
	}
}
