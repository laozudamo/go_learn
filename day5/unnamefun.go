package main

import "fmt"

/*
	go语言函数不能嵌套，但是在函数内部可以定义居名函数，实现一下简单功能调用。
所谓匿名函数就是，没有名称的函数。
语法格式如下：
fund（参数列表）（返回值）
T
当然可以既没有参数，可以没有返回值
*/
func main() {
	max := func(x int, y int) int {
		if x > y {
			return x
		}
		return y
	}
	max(2, 3)
	fmt.Printf("max: %v\n", max)
}
