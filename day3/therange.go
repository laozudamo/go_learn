package main

import "fmt"

/*
Go语言中可以使用for range遍历数組、切片、字符串、map 及通道(channel)）。通过 for range 遍历的返回
值有以下规律：
1，数组、切片、宇符串返回索引值。
2. map返回键和值。
3.通道（channel）只返回通道内的值。
*/
func main() {
	rangeFun()
	mapObj()
	rangeSlice()
	theMap()
}

func rangeFun() {
	arr := [5]string{"h", "e", "l", "l", "o"}
	for i, v := range arr {
		fmt.Print(i, " ", v, "\n")
	}

}

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func mapObj() {
	obj := Books{"忘了吃饭", "12", "21", 44}
	fmt.Printf("obj: %v\n", obj)
	// for i, v := range obj {
	// 	fmt.Print(i, " ", v, "\n")
	// }
}

func rangeSlice() {
	var s = []int{1, 2, 33}
	for _, v := range s {
		fmt.Printf("v: %v\n", v)
	}
}

func theMap() {
	m := map[string]int{"height": 188, "age": 20}
	for key, v := range m {
		fmt.Printf("key: %v\n", key)
		fmt.Printf("v: %v\n", v)
	}
}
