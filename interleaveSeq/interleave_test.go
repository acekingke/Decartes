package interleave

import (
	"fmt"
	"testing"
)

func TestInterleave(t *testing.T) {
	type args struct {
		left  []string
		right []string
		f     func([]string)
	}
	tests := []struct {
		name string
		args args
	}{

		{"testOne", args{left: []string{"a", "b", "c"}, right: []string{"x", "y", "z"},
			f: func(a []string) {
				for _, x := range a {
					fmt.Print(x)
				}
				fmt.Println()
			}}},
		{"test2", args{left: []string{"a", "b"}, right: []string{"x", "y", "z"},
			f: func(a []string) {
				for _, x := range a {
					fmt.Print(x)
				}
				fmt.Println()
			}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Interleave(tt.args.left, tt.args.right, tt.args.f)
		})
	}
}
