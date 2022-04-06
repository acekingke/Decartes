package cartesian

import (
	"fmt"
	"testing"
)

func TestCartN(t *testing.T) {
	a := []string{"1", "2", "3"}
	b := []string{"4", "5", "6"}
	c := []string{"7", "8", "9"}
	all := [][]string{a, b, c}
	result := CartN(all...)
	if len(result) != 27 {
		t.Error("not correct result")
	}
	for _, v := range result {
		fmt.Println(v)
	}
}
