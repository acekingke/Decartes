package parser

import (
	"fmt"

	"github.com/acekingke/lexergo"
)

type Token struct {
	TokenType   int
	TokenString string
}

//const EOF = -1

// keywords
var keyword = map[string]int{
	"break": BREAK,
	"cart":  CARTESIAN,
	"each":  EACH,
	"perm":  PERMUTATION,
	"step":  STEP,
	"if":    IF,
	"else":  ELSE,
}

// alphadigit_fn , alpha or number
var alphadigit_fn lexergo.Bool_string_string_fn = lexergo.Repeat_fn(lexergo.Select_fn(
	lexergo.Match_alpha_fn, lexergo.Match_digit_fn, func(x string) (bool, string, string) {
		return lexergo.Match_fn(x, func(in string) bool {
			return in == "$"
		})
	}), 1, lexergo.MaxInt)

// Left parentheses.
var lp_fn lexergo.Bool_string_string_fn = func(x string) (bool, string, string) {
	return lexergo.Match_fn(x, func(in string) bool {
		return in == "("
	})
}

// Right parentheses
var rp_fn lexergo.Bool_string_string_fn = func(x string) (bool, string, string) {
	return lexergo.Match_fn(x, func(in string) bool {
		return in == ")"
	})
}

// {xxx}
var brace_string_fn lexergo.Bool_string_string_fn = func(x string) (bool, string, string) {
	count := 0
	braceLeft := func(x string) (bool, string, string) {
		return lexergo.Match_fn(x,
			func(in string) bool {
				count += 1
				return in == "{" && count == 1
			})
	}
	nonBrace := func(x string) (bool, string, string) {
		return lexergo.Match_fn(x, func(in string) bool {
			if in == "{" {
				count++
				return true
			} else if in == "}" {
				count--
				return count == 1
			}
			return in != "}"
		})
	}
	braceRight := func(x string) (bool, string, string) {
		return lexergo.Match_fn(x, func(in string) bool {
			return in == "}"
		})
	}

	return lexergo.Concat_fn(braceLeft,
		lexergo.Repeat_fn(nonBrace, 0, lexergo.MaxInt),
		braceRight)(x)
}

// double quote string
var double_quote_string_fn lexergo.Bool_string_string_fn = lexergo.Concat_fn(func(x string) (bool, string, string) {
	return lexergo.Match_fn(x, func(in string) bool {
		return in == "\""
	})
}, // double quote,
	lexergo.Repeat_fn(func(x string) (bool, string, string) {
		return lexergo.Match_fn(x, func(in string) bool {
			return in != "\"" // none double quote
		})
	}, 0, lexergo.MaxInt),
	func(x string) (bool, string, string) {
		return lexergo.Match_fn(x, func(in string) bool {
			return in == "\""
		})
	})
var single_quote_string_fn lexergo.Bool_string_string_fn = lexergo.Concat_fn(func(x string) (bool, string, string) {
	return lexergo.Match_fn(x, func(in string) bool {
		return in == "'"
	})
}, // single quote
	lexergo.Repeat_fn(func(x string) (bool, string, string) {
		return lexergo.Match_fn(x, func(in string) bool {
			return in != "'"
		})
	},
		0, lexergo.MaxInt),
	func(x string) (bool, string, string) {
		return lexergo.Match_fn(x, func(in string) bool {
			return in == "'"
		})
	})

var square_quote_string_fn lexergo.Bool_string_string_fn = lexergo.Concat_fn(func(x string) (bool, string, string) {
	return lexergo.Match_fn(x, func(in string) bool {
		return in == "["
	})
}, // double quote,
	lexergo.Repeat_fn(func(x string) (bool, string, string) {
		return lexergo.Match_fn(x, func(in string) bool {
			return in != "]" // none double quote
		})
	}, 0, lexergo.MaxInt),
	func(x string) (bool, string, string) {
		return lexergo.Match_fn(x, func(in string) bool {
			return in == "]"
		})
	})
var blank_fn lexergo.Bool_string_string_fn = lexergo.Repeat_fn(lexergo.Select_fn(
	lexergo.MatchSpace_fn,
	lexergo.MatchTab_fn,
	/*func(in string) (bool, string, string) {
		return lexergo.Match_fn(in, func(in string) bool {
			return in == "\r" || in == "\n"
		}),
	}*/), 1, lexergo.MaxInt)

// Comment start with #
var comment_fn lexergo.Bool_string_string_fn = lexergo.Concat_fn(func(x string) (bool, string, string) {
	return lexergo.Match_fn(x, func(in string) bool {
		return in == "#"
	})
},
	lexergo.Repeat_fn(func(x string) (bool, string, string) {
		return lexergo.Match_fn(x, func(in string) bool {
			return in != "\r" && in != "\n"
		})
	}, 1, lexergo.MaxInt),
	func(x string) (bool, string, string) {
		return lexergo.Match_fn(x, func(in string) bool {
			return in == "\r" || in == "\n"
		})
	},
)
var newline_fn lexergo.Bool_string_string_fn = func(x string) (bool, string, string) {
	return lexergo.Match_fn(x, func(in string) bool {
		return in == "\r" || in == "\n"
	})
}
var Scan = func(src string) (string, *Token, error) {
	_, _, rest := blank_fn(src)
	for len(rest) != 0 && rest[0] == '#' {
		_, _, rest = comment_fn(rest) // skip comment
		_, _, rest = blank_fn(rest)
	}
	var flag bool
	var read string
	if flag, read, rest = alphadigit_fn(rest); flag {
		var token *Token
		if keyword[read] != 0 {
			token = &Token{
				TokenType:   keyword[read],
				TokenString: read,
			}
		} else if CommandMap[read] != nil {
			token = &Token{
				TokenType:   NORMALCMD,
				TokenString: read,
			}
		} else {
			token = &Token{
				TokenType:   IDENTIFIER,
				TokenString: read,
			}
		}
		return rest, token, nil
	}
	if flag, read, rest = newline_fn(rest); flag {
		token := &Token{
			TokenType:   NEWLINE,
			TokenString: read,
		}
		return rest, token, nil
	}
	//LP
	if flag, read, rest = lp_fn(rest); flag {
		token := &Token{
			TokenType:   LPAREN,
			TokenString: read,
		}
		return rest, token, nil
	}
	// RP
	if flag, read, rest = rp_fn(rest); flag {
		token := &Token{
			TokenType:   RPAREN,
			TokenString: read,
		}
		return rest, token, nil
	}
	// {xxxx} string can replace
	if flag, read, rest = brace_string_fn(rest); flag {
		token := &Token{
			TokenType:   BRACE_STRING,
			TokenString: read,
		}
		return rest, token, nil
	}
	// "xxx" string can replace
	if flag, read, rest = double_quote_string_fn(rest); flag {
		token := &Token{
			TokenType:   DQUOTE_STRING,
			TokenString: read,
		}
		return rest, token, nil
	}
	// 'xxx' string  cannot replace
	if flag, read, rest = single_quote_string_fn(rest); flag {
		token := &Token{
			TokenType:   SINGLEQUOTE_STRING,
			TokenString: read,
		}
		return rest, token, nil
	}
	// [xxx] string  cmd replace
	if flag, read, rest = square_quote_string_fn(rest); flag {
		token := &Token{
			TokenType:   SQUAREQUOTE_STRING,
			TokenString: read,
		}
		return rest, token, nil
	}
	return rest, nil, fmt.Errorf("not correct token")
}
