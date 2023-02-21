package main

import "fmt"

func main() {
	test()
	arrPorinter()
}

/*
类型指针不能进行偏移和运算。
& 取地址
* 取值
*/
func test() {

	// i := map[string]int{
	// 	"age":    19,
	// 	"height": 188,
	// }
	// v := &i
	// fmt.Printf("v: %v\n", v)

	num := 10
	address := &num
	fmt.Printf("num: %v\n", *address)
}

/*
	指向数组的指针
	指向map的指针
*/

// const MAX int = 5

func arrPorinter() {
	a := [5]int{1, 2, 34, 5, 6}
	var ptr [len(a)]*int
	for i, v := range a {
		ptr[i] = &v
	}
	fmt.Printf("ptr: %v\n", ptr)
}
