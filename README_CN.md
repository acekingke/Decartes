# 笛卡尔语言

测试本质是穷举, 如何穷举所有情况,来发现bug,是一个巨大的挑战. 您不应当手动去穷举, 而应当将这个工作交给机器去做.**笛卡尔**语言就是用来做这个的.

# 设计哲学
受TCL语言启发,  有如下哲学

1. 每一个语法单位都是 `command` 
2. 每一种数据类型都是字符串
4. 解释型语言
4. 足够小
5. 容易拓展自定义的语法 

# 工具
[grammar generator](https://github.com/acekingke/yaccgo)

[lexer package](https://github.com/acekingke/lexergo)
# 语法

    WIP....

# 命令

step  {xxxx}

set x val

cart (),... () each (a1 ... an) {

}

shell {xxxx}

permutation s1 s2 ... sn

# WIP

* step 


# 路线图

- [x]  set 
- [x]  puts
- [x]  cart  each command
- [ ]  permutation
- [ ]  if
- [ ]  while
