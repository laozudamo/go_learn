package main

import (
	"fmt"
)

func main() {
	// sliceFun()

	// test2()

	// copyFun()
	// reverseFun()
	// capsFun()
	// appendFun()
	deleteFun()
}

/*
切片长度是可变的
前面我们学习了数组，数组是固定长度，可以容纳相同数据类型的元素的集合。
当长度固定时，使用还是带来一些限制，比如：我们申请的长度太大浪费内存，太小又不够用。
鉴于上述原因，我们有了g0语言的切片，可以把切片理解为，可变长度的数组，其实它底层就是使用数组实现的，增加了自动扩容功能。切片(Slice）是一个拥有相同类型元素的可变长度的序列。
*/

// 增加了自动扩容功能
func sliceFun() {
	// var names []string
	// ages := []int{1, 2, 3}
}

// nil 空切片

func test2() {
	var names []string
	names = append(names, "jack")

	// 分配了内存
	var s2 = make([]string, 3)
	s2 = append(s2, "jack", "kacs", "wqwq")

	// 在原来slice的后面添加

}

func copyFun() {
	s := []string{"jack", "luck", "hammy"}

	sCopy := make([]string, len(s))

	copy(sCopy, s)

	fmt.Printf("s: %v\n", s)

}

func reverseFun() {
	s := []string{"jack", "luck", "hammy"}

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	fmt.Printf("s: %v\n", s)
}

// caps获取容量
func capsFun() {
	s := []string{"jack", "luck", "hammy"}
	fmt.Printf("s: %v\n", cap(s))
}

func appendFun() {
	s := []int{}
	s = append(s, 1, 2, 3)
	fmt.Printf("s: %v\n", s)
}

// 删除
func deleteFun() {
	s := []string{"jack", "lucy", "dany"}

	// lucy
	var s2 = []string{}
	s2 = append(s2, s[:1]...)
	s2 = append(s2, s[2:]...)
	fmt.Printf("s2: %v\n", s2)
}

// 更新 替换就可

// query
