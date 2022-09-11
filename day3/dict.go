package main

import (
	"fmt"
	"unsafe"
)

func main() {

	// createMap()

	// handleDict()

	handKeys()

	// test1()
}

// 字典的 key 是因，字典的 value 是果

func createMap() {

	// 使用类型推导
	var m1 = map[int]string{
		90: "优秀",
		80: "良好",
		60: "及格",
	}
	fmt.Println("m1", m1)

	// 创建空字典使用make
	var m2 = make(map[string]string)
	fmt.Println(m2, len(m2))

	// 赋值

	var m3 map[int]string = map[int]string{
		90: "优秀",
		80: "良好",
		60: "及格",
	}

	fmt.Println("m3", m3)
}

/*
	字典读写
*/

// 读操作时，如果 key 不存在，也不会抛出异常。它会返回 value 类型对应的零值。
func handleDict() {
	var fruits = map[string]int{
		"apple":  2,
		"banana": 5,
		"orange": 8,
	}
	// 读取元素
	var score = fruits["banana"]
	fmt.Println(score)
	// 增加或修改元素
	fruits["pear"] = 3
	fmt.Println(fruits)
	// 删除元素
	delete(fruits, "pear")
	fmt.Println(fruits)

	var num, ok = fruits["durin"]
	if ok {
		fmt.Println("num", num)
	} else {
		fmt.Println("no", "不存在"+"durin")
	}

}

// 遍历使用range

//获取所有key值

func handKeys() {
	var dict = map[string]string{
		"age":    "19",
		"sex":    "男",
		"height": "188",
		"hight":  "188",
	}

	var keys = make([]string, 0, len(dict))

	// var values = make([]string, len(dict))
	for key := range dict {
		keys = append(keys, key)
		// 	values = append(values, value)
	}
	fmt.Println("keys", keys)
}

/*
	字典变量里存的只是一个地址指针，这个指针指向字典的头部对象。
	所以字典变量占用的空间是一个字，也就是一个指针的大小，64 位机器是 8 字节，32 位机器是 4 字节。

	下面dict变量中用的空间是一个指针的大小;也就是一个字的大小
*/

func test1() {
	var dict = map[string]string{
		"age":    "19",
		"sex":    "男",
		"height": "188",
	}

	var s = unsafe.Sizeof(dict)

	fmt.Println("s", s)
}
