package main

import "fmt"

func main() {
	//可以使用单行注释
	/*
	也可以使用多行注释
	*/
	//运算符不可以当作标识符
	//go语言中使用+进行字符串的拼接
	fmt.Println("Hello World")
	fmt.Println("Hello" + " " +"World" + "!")

	var a string = "Test String"
	fmt.Println(a)

	//可以一次声明多个变量
	var b, c int = 1, 2
	fmt.Println(b, c)

	//变量如果没有初始化,则变量的默认为零值
	var d int
	fmt.Println(d)

	var e bool
	fmt.Println(e)
}

/*
go语言数据类型

1.基础类型
	1.数字类型
	2.布尔类型
	3.字符串类型
2.复合类型
	1.数组类型
	2.结构体类型
3.引用类型
	1.指针类型
	2.切片类型
	3.字典类型
	4.函数类型
	5.通道类型
4.接口类型
*/