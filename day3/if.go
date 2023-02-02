package main

import "fmt"

// mian主函数
func main() {
	// ifFun()

	judgeNum()
}

func ifFun() {
	// if表达式里面不需要括号
	flag := true
	if flag {
		fmt.Printf("flag: %v\n", flag)
	}

	a := 1
	b := 2
	if a > b {
		fmt.Printf("\"a > b\": %v\n", "a > b")
		return
	}
	fmt.Printf("\"a < b\": %v\n", "a < b")
}

// 不能使用0或非0表示真假

func judgeNum() {
	var num int
	fmt.Println("请输入一个数字")
	fmt.Scan(&num)

	if num%2 == 0 {
		fmt.Printf("num: 偶数%v\n", num)
	} else {
		fmt.Printf("num: 鸡数%v\n", num)
	}
}
