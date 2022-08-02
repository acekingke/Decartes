package expr

import (
	"strconv"
)

const EOF = -1

// The parser expects the lexer to return 0 on EOF.  Give it a name
// for clarity.
func GetToken(input string, valTy *ValType, pos *int) int {
	if *pos >= len(input) {
		return -1
	} else {
		*valTy = ValType{0}

		if *pos >= len(input) {
			return -1
		}
		c := input[*pos]
		for c == ' ' || c == '\t' || c == '\n' || c == '\r' {
			*pos++
			if *pos >= len(input) {
				return -1
			}
			c = input[*pos]
		}
		switch c {
		case '+':
			fallthrough
		case '-':
			fallthrough
		case '(':
			fallthrough
		case ')':
			fallthrough

		case '*':
			fallthrough
		case '/':
			*pos++
			return int(c)
		case '=':
			*pos++
			if *pos < len(input) && input[*pos] == '=' {
				*pos++
				return EQ
			} else {
				return -1
			}

		case '!':
			*pos++
			if *pos < len(input) && input[*pos] == '=' {
				*pos++
				return NE
			} else {
				return -1
			}
		case '<':
			*pos++
			if *pos < len(input) {
				if input[*pos] == '=' {
					*pos++
					return LE
				} else {
					return int(c)
				}
			}
		case '>':
			*pos++
			if *pos < len(input) {
				if input[*pos] == '=' {
					*pos++
					return GE
				} else {
					return int(c)
				}
			}

		default:
			if numstr := LexNumber(input, pos); numstr != "" { // is digit
				var err error
				if valTy.val, err = strconv.ParseFloat(numstr, 64); err != nil {
					panic(err)
				}
				return NUM
			}

		}
		return -1
	}
}

func LexNumber(input string, pos *int) string {
	if *pos >= len(input) || input[*pos] < '0' || input[*pos] > '9' {
		return ""
	}
	start := *pos
	for *pos < len(input) && input[*pos] >= '0' && input[*pos] <= '9' {
		*pos++
	}

	if *pos < len(input) && start <= *pos && input[*pos] == '.' { // has digits before '.
		*pos++
		decimalpoint := *pos

		for *pos < len(input) && input[*pos] >= '0' && input[*pos] <= '9' {
			*pos++
		}
		if *pos < len(input) && decimalpoint < *pos { //has digits after '.'
			if *pos < len(input) && (input[*pos] == 'e' || input[*pos] == 'E') {
				*pos++
				if *pos < len(input) && (input[*pos] == '+' || input[*pos] == '-') {
					*pos++
				}
				for *pos < len(input) && input[*pos] >= '0' && input[*pos] <= '9' {
					*pos++
				}
			}
		}
	}

	return input[start:*pos]
}

func RunParser(input string) float64 {
	PushContex()
	ParserInit()
	val := Parser(input).val
	PopContex()
	return val
}
