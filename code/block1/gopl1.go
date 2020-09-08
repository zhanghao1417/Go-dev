package main

import (
    "fmt"
    "os"
)

func main() {
	//var 定义两个string类型的变量,变量会在声明的时候直接初始化.如果变量没有被显示初始化,则会被隐式的赋予其类型的零值.数值类型为0,字符串类型为""
    var s, sep string
    //i++ 是自增语句,等价于i += 1或i = i + 1,++与--是语句,而不是表达式,所以只能放在变量名之后,因此--i是非法的
    for i := 1; i < len(os.Args); i++ {
        //对于数值类型,go提供常规的数值与逻辑运算符.对于string类型,+运算符起连接字符串的作用
    	//+ 运算符标识连接字符串,下面的语句等价于 s = s + sep + os.Args[i]
        s += sep + os.Args[i]
        sep = ","
    }
    fmt.Println(s)
}