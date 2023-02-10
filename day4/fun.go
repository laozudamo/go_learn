package main

import "fmt"

func main() {
	// v := sum(2, 5)
	// fmt.Printf("v: %v\n", v)

	// max := getMax(2, 19)
	// fmt.Printf("max: %v\n", max)

	// v1 := moreRet("jack", 20)
	// fmt.Printf("v1: %v\n", v1)

	theSlice := []int{1, 2, 4}

	// 传递得到是地址
	theFun(theSlice)
	fmt.Printf("theSlice: %v\n", theSlice)

	v := getSum(1, 3, 4)
	fmt.Printf("v: %v\n", v)
}

/*
函数的go语言中的一级公民，
我们把所有的功能单元都定义在函数中，可以重复使用。
函数包含函数的名称参数列表和返回类型，这些构成了函数的签名 (signature）
	1.go语言函数里面只能嵌套匿名函数

	func function_name( [parameter list] ) [return_types]{
		函数体
	}
*/

func sum(a int, b int) int {
	return a + b
}

func getMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

type Person struct {
	name string
	age  int
}

func moreRet(name string, age int) Person {
	var person1 Person
	person1.name = name
	person1.age = age
	return person1
}

/*
	go语言的函数是通过传值方式传递参数的;传递的是拷贝后的副本
	0语言可以使用变长参数，有时候并不能确定爹数的个数，可以使用变长参数，可以在函数定义语句的参数部分使用 ARGS...FYPE 的方式。
*/

func theFun(s []int) {
	s[0] = 33
}

// ...变长参数
func getSum(args ...int) int {
	var sum int
	for _, v := range args {
		sum += v
	}
	return sum
}
