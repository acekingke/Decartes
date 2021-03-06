package parser

import (
	"fmt"
	"testing"
)

func TestScan(t *testing.T) {
	_, tok, err := Scan("1ahga")

	if err != nil {
		t.Error(err)
	} else if tok.TokenType != IDENTIFIER {
		t.Error("not correct token")
	}

	_, tok, err = Scan("(")
	if err != nil {
		t.Error(err)
	} else if tok.TokenType != LPAREN {
		t.Error("not correct token")
	}
}

func TestScan2(t *testing.T) {
	str := `cart (1 2 3) (4 5 6) (7 8 9) each (a b c) {
		shell {xxx}
	}
	`
	for rest, tok, _ := Scan(str); tok != nil && tok.TokenType != -1; rest, tok, _ = Scan(rest) {
		fmt.Printf("tokenType:%d,%s\n", tok.TokenType, tok.TokenString)
	}
}

func TestScan3(t *testing.T) {
	str := `puts hello $x
	`

	for rest, tok, _ := Scan(str); tok != nil && tok.TokenType != -1; rest, tok, _ = Scan(rest) {
		fmt.Printf("tokenType:%d,%s\n", tok.TokenType, tok.TokenString)
	}
}

func TestScan4(t *testing.T) {
	str := `{
		shell {echo  hello $a $b $c}
}
	`
	for rest, tok, _ := Scan(str); tok != nil && tok.TokenType != -1; rest, tok, _ = Scan(rest) {
		fmt.Printf("tokenType:%d,%s\n", tok.TokenType, tok.TokenString)
	}
}
func TestComment(t *testing.T) {
	str := `
	#Type:          Single, Span, Stripe, Mirror, RAID-5
	#Size:          10, 100, 500, 1000, 5000, 10000, 40000
	#Format method: Quick, Slow
	#File system:   FAT, FAT32, NTFS
	#Compression:   On, Off
	
	puts Type   size  format   "filesystem"   compression
`
	for rest, tok, _ := Scan(str); tok != nil && tok.TokenType != -1; rest, tok, _ = Scan(rest) {
		fmt.Printf("tokenType:%d,%s\n", tok.TokenType, tok.TokenString)
	}
	//_ = Parser(str)

}
