package main

// CRUD
import (
	"fmt"
)

func main() {
	var a, b, c = 1, false, "str"
	a = 2
	fmt.Println(a, b, c)

	x, y, z := 2, 3, "hahhah"
	fmt.Println(x, y, z)

	const PI float64 = 3.14
	const PI2 = 3.1415

	const (
		weight = 32
		height = 32
	)

	const i, j = 1, 2

	fmt.Printf("i: %v\n", i)
	fmt.Printf("z: %v\n", z)
	fmt.Printf("weight: %v\n", weight)

	testIota()

}

// iota 可被修改的常量 默认开始值是0 每次被调用加一 遇到const 关键字被重置为0

// 使用 _ 跳过某些值

// iota 中间 插入值
func testIota() {
	const (
		a1 = iota
		// _
		a2 = 10
		a3 = iota
	)
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)
	fmt.Printf("a3: %v\n", a3)
	const a4 = iota
	fmt.Printf("a4: %v\n", a4)

}
