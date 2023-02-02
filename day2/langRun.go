package main

import "fmt"

func main() {
	ifFun()
	ifElseFun()
	funSwitch()
}

func ifFun() {
	num := 20

	if num > 10 {
		fmt.Println("我的值大于10", num)
	}
}

// if else语句
func ifElseFun() {
	a := "我是个字符串"

	flag := len(a)

	if flag > 3 {
		fmt.Println("字符串长度大于3")
	} else {
		fmt.Println("字符串长度小于等于是3")
	}
}

// switch case 语句
func funSwitch() {
	var num int = 10
	switch num {
	case 10:
		fmt.Println("num等于10")
	case 20:
		fmt.Println("num等于20")
	default:
		fmt.Println("num不等于10或20")
	}

	var grade string = "A"

	switch {
	case grade == "A":
		fmt.Println("优秀")
	case grade == "B":
		fmt.Println("有进步空间")
	case grade == "c":
		fmt.Println("该加油了")
	default:
		fmt.Println("你该不会是老板吧")
	}
}

// select 语句：
// select {
// case <- channel1:
// 	// 执行的代码
// case value := <- channel2:
// 	// 执行的代码
// case channel3 <- value:
// 	// 执行的代码

// 	// 你可以定义任意数量的 case

// default:
// 	// 所有通道都没有准备好，执行的代码
// }
