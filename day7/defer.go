package main

import "fmt"

func main() {
	test()
}

/*
	g0语言中的 defer 语向会将其后面跟随的语向进行延迟处理。
	在defer 归厲的西数即将返回时，将延迟处理的语句按defer 定义的逆序进行执行，也就是说，先被 defer 的语句最后被执行，最后被 defer的语句，最先被执
5. stack
*/

/*
	先进后出
*/

/*
用途:
1. 关闭文件句柄
2.锁资源释放
3.数据库连接释放
*/

func test() {
	fmt.Printf("\"start\": %v\n", "start")
	defer fmt.Println("step1")
	defer fmt.Println("step2")
	defer fmt.Println("step3")
	fmt.Printf("\"end\": %v\n", "end")
}
