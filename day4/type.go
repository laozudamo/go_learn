package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	type f1 func(int, int) int

	var ff f1
	ff = sum
	v := ff(1, 2)

	fmt.Printf("v: %v\n", v)

}

/*
可以使用type 关键字来定义一个函数类型，语法格式如下：
type fun func(int, int) int
*/
