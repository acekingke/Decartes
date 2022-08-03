%{
package expr
import "fmt"

%}


	%union {
		val float64
	}
	
	%type	<val>	E
	%token '+' '-' '*'   '(' ')' '/' '>' '<' 
	%nonassoc '>'  '<' GE LE EQ	NE
	%left '+'  '-'
	%left '*'   '/'
	%token	<val>	NUM
	%token GE LE EQ	NE
	%token NUM 100
	%start E
%%

E:
	E '+' E {
		$$	=	$1 + $3
	}	
	|E '-' E {
		$$	=	$1 - $3
	}
	| E '*' E {
		$$	=	$1 * $3
	}
	| E '/' E {
		$$	=	$1 / $3
	}
	| E '>' E {
		if $1 > $3 {
			$$	=	1
		} else {
			$$	=	0
		}
	}
	| E '<' E {
		if $1 < $3 {
			$$	=	1
		} else {
			$$	=	0
		}
	}
	| E GE E {
		if $1 >= $3 {
			$$	=	1
		} else {
			$$	=	0
		}
	}
	| E LE E {
		if $1 <= $3 {
			$$	=	1
		} else {
			$$	=	0
		}
	}
	| '(' E ')' {
		$$	=	$2
	}
	| NUM {
		$$	=	$1
	}
	
%%
