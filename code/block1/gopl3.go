package main

import (
	"strings"
	"fmt"
	"os"
)

func main() {
	//打印被执行命令本身的名称
	fmt.Println(os.Args[0])
	//切片输出格式 : [1 2 3 4 5]
	fmt.Println(os.Args[1:])
	fmt.Println(strings.Join(os.Args[1:], ","))
	//输出每个参数的索引与值
	for i, arg := range os.Args[1:] {
		fmt.Println(i)
		fmt.Println(arg)
	}
}