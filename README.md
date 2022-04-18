# Decartes
English | [简体中文](README_CN.md) 

The essence of testing is exhaustive.  how to exhaust all possibilities to find bugs, it is a greate challenge. You shouldn‘t do it manually, You should do it with a machine. Language **Decartes** make it. which core is the **Cartesian product** and **permutation** algorithm.

# Philosophy

inspired from tcl,  has philosopy

1. every grammar is `command` 
2. every data is string, 
4. interpreted language
4. small enough
5. easy to extend your own grammar, 
# Install and usage
install : 

`go get https://github.com/acekingke/Decartes/Decartes`

run: 

`Decartes examples/print.dkr`

# Tool
[grammar generator](https://github.com/acekingke/yaccgo)

[lexer package](https://github.com/acekingke/lexergo)
# Grammar
## 1 string
All data is string!. It has three form of string:
* string without quote
follow the command, all things is string, and they are separated by space.
for example:
```
puts hello world
```
`puts` is a command, and `hello` `world` are strings.
* string with single quote
single quote string, it is a string, and it will not use variable replacement.
for example:
```
set ax 1
puts 'hello $ax'
puts "hello $ax"
```
the output is:
```
hello $ax 
hello 1 
```

* string with double quote
if you want to use a string contain space in it, you should use `"` to wrap it.
for example:
```
set a "hello world"
puts $a
```
`set a xxx` command ,means set variables `a` with value `xxx`, and `xxx` is a string.
see [variable assignment](#variable) for more detail.
* string with curly bracket
It also a string, but it use '{' and '}' to wrap it. it can be nested.
single quote or double quote can not be nested.
for example:
```
set a {hello world}
puts $a
set a { {hello world} }
puts $a
```
the output is:
```
hello world 
 {hello world}  
```
`set nice day` is setting variable `nice` with `day` string, the command result is `day`.
so the `puts [set nice  day]` it means `puts day`.
## 2 normal command
cmd arg1 arg2 ...
`set`, `puts`, `shell` are normal command.
for example:

```
set x val
puts $x
```
## 3 variable assignment
use `set` cmd to assign a value to a variable.
the form is:
`set variable  value`
for example:
```
set x 1
```
When you use the variable, you can use $variable to get the value. the variable can use in string, quoted string, and curly bracket string.
for example:
```
set x 1
puts $x
puts hello $x
puts "hello $x"
puts {hello $x}
```
## 4 Command result replecement
square bracket wrap a command, and it is represented by the command result.
for example:

```
puts [set nice  day]
```
## 5 cartesian product
the form is:
```
cart (x1 x2 x3...) ... (y1 y2 y3...) each (a1 ... an) {

}
```
is command use (x1 x2 x3...)...(y1 y2 y3...) to generate all cartesian product. and every cartesian product values set to (a1 ... an), for each value of (a1 ... an), it will execute the command block, which is wrapped by `{}`.
for exampls:

```
puts now test the cartesian product for test
cart (1 2 3) (4 5 6) (7 8 9) each (a b c) {
		shell {echo  now $a $b $c}
}
```
the output is:
```
now test the cartesian product for test 
now 1 4 7
now 1 4 8
now 1 4 9
now 1 5 7
now 1 5 8
now 1 5 9
now 1 6 7
now 1 6 8
now 1 6 9
now 2 4 7
now 2 4 8
now 2 4 9
now 2 5 7
now 2 5 8
now 2 5 9
now 2 6 7
now 2 6 8
now 2 6 9
now 3 4 7
now 3 4 8
now 3 4 9
now 3 5 7
now 3 5 8
now 3 5 9
now 3 6 7
now 3 6 8
now 3 6 9

```
## 6 step 
`step` command tell the machine the command string has a alias name, and the alias name is used in permutation command.
the form is:
`step alias CommandString`
for example:

```
step a {puts do STEP 1}
step b {puts "do STEP 2"}
step c {puts "do STEP 3"}
```

## 7 perm
`perm` is a command use to generate all permutation of steps, and run each steps sequence commands.
for example:

```
puts permutation the steps

step a {puts do STEP 1}
step b {puts "do STEP 2"}
step c {puts "do STEP 3"}
perm (a b c)
```
it will run 3!=6 sequences of steps, the output is:
```
permutation the steps 
do STEP 1 
do STEP 2 
do STEP 3 
do STEP 1 
do STEP 3 
do STEP 2 
do STEP 2 
do STEP 1 
do STEP 3 
do STEP 2 
do STEP 3 
do STEP 1 
do STEP 3 
do STEP 2 
do STEP 1 
do STEP 3 
do STEP 1 
do STEP 2 
```



# COMMAND

step  s1 {xxxx}

set x val

puts $x

cart (x1 x2 ... xk) ... (z1 z2 ... zk) each (a1 ... an) {

}

shell {xxxx}

perm (s1 s2 ... sn)

# WIP

* template code generator


# roadmap

- [x]  set 
- [x]  puts
- [x]  step
- [x]  cart  each command
- [ ]  permutation
- [ ]  if
- [ ]  while

# License

[MIT](LICENSE)
