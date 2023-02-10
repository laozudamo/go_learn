package main

import "fmt"

func main() {

	// test1()
	// test2()

	// test3()
	test4()
}

// ...省略数组长度
func test1() {
	arr := [5]int{1, 2, 3, 4, 5}

	v := arr
	fmt.Printf("arr[0]: %v\n", arr[0])
	fmt.Printf("v: %v\n", v)
}

// 数组长度越界

func test2() {
	arr := [3]int{1, 2, 4}
	fmt.Printf("arr[3]: %v\n", arr[2])
}

// 遍历数组
func test3() {
	arr := [3]int{1, 2, 4}
	for _, v := range arr {
		fmt.Printf("v: %v\n", v)
	}
}

func test4() {
	var arrTest = [3]string{"jack", "ma", "lucy"}

	for _, v := range arrTest {
		fmt.Printf("v: %v\n", v)
	}

	arr := [3]int{1, 2, 4}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("arr[i]: %v\n", arr[i])
	}
}
