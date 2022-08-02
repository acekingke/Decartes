%{
package parser
import "fmt"
import "os"
%}
%token NEWLINE
%token BREAK
%token IF ELSE
%token WHILE
%token 	CARTESIAN
%token  PERMUTATION 
%token 	STEP
%token	NORMALCMD
%token 	EACH
%token 	IDENTIFIER
%token 	LPAREN
%token 	RPAREN
//DOLLOR
%token 	BRACE_STRING
%token 	DQUOTE_STRING
%token 	SINGLEQUOTE_STRING
%token 	SQUAREQUOTE_STRING
%start COMMANDS
%type <Array>  WORDS
%type <Value>  WORD
%type <Value>  IDENTIFIER
%type <Value>  STRING
%type <Array>  ARRAY
%type <Arrays>  ARRAYS
%type <Value>  BRACE_STRING
%type <Value>  DQUOTE_STRING
%type <Value>  SINGLEQUOTE_STRING
%type <Value>  SQUAREQUOTE_STRING
%type <Value>  NORMALCMD
%left ELSE
%union {
    WorkValue
}
	
%%
COMMANDS : /*empty*/
    | COMMANDS COMMAND
COMMAND : END
    | CMD END
CMD: NORMALCMD WORDS {
       err := ExecuteNormalCmd($1, $2);
       if err != nil {
            ParserError(err)
       }
    }
    |CARTESIAN ARRAYS EACH ARRAY BRACE_STRING {
       w := NewWorkCart($2, $4, $5[1:len($5)-1]);
       err := w.Execute();
       if err != nil {
          ParserError(err)
       }
    }
    | PERMUTATION ARRAY  {
       p := &PermType{
            StepNames: $2,
       }
       err := p.Execute();
         if err != nil {
             ParserError(err)
         }
    }
    | STEP IDENTIFIER STRING {
       err := NewStep($2, $3)
       if err != nil {
          ParserError(err)
       }
    }
    | IFCMD {
       
    }
    | BREAK {
        globalEnv.IsBreak = true;
    }
    | WHILE STRING STRING {
        w := NewWhile($2, $3)
        w.Execute()
        globalEnv.IsBreak = false;
    }
IFCMD: IF STRING  STRING  {
         ifcmd := NewIfCmd($2, $3, "")
          ifcmd.Execute()
        }
       | IF STRING  STRING ELSE STRING {
                ifcmd := NewIfCmd($2, $3, $5)
          ifcmd.Execute()
       }

WORDS : /*empty*/ {
    $$= make([]string,0)
}
    |WORDS WORD {
        $$ = append($1, $2)
    }
WORD : IDENTIFIER {$$ = $1 }
    | STRING { $$ = $1 }
STRING : BRACE_STRING {$$ = $1 }
        | DQUOTE_STRING {$$ = $1 }
        | SINGLEQUOTE_STRING {$$ = $1 }
        | SQUAREQUOTE_STRING {$$ = $1 }
ARRAYS :/*empty*/ {$$ =make([][]string ,0)} 
    |ARRAYS ARRAY {
     $$ = append($1, $2) 
     }
ARRAY : LPAREN WORDS RPAREN {
    $$ = $2
}
END :NEWLINE
%%
func ParserError(err error) {
    fmt.Print("ParserError: ")
    fmt.Println(err)
    os.Exit(1)
}
