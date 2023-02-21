package main

import (
	"fmt"
)

func main() {
	fmt.Printf("\"main\": %v\n", "main")
}

func init() {
	fmt.Printf("\"110\": %v\n", "110")
}

var i int = initvar()

func initvar() int {
	fmt.Printf("\"100\": %v\n", "100")
	return 100
}

/*
	golang有一个特殊的函数 init 函数，先于 main 函数执行，实现包級別的一些初始化操作。
	init 函数特点
	1.优先于main函数;自动执行;
	2.没有参数和返回值
	3.每个包可以有多个init函数
	4.不同包的int函数按照包导入的依赖关系决定执行顺序。
*/

/*
	golang初始化顺序
	变量初始化 -> init -> main
*/
