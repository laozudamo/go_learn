package main

import "fmt"

func main() {
	a := 5

	foFun()

theItem:
	if a == 9 {
		fmt.Println("这里当a等于9的时候跳转到这里了")
	}

	for a < 10 {
		a += 1
		fmt.Println("a", a)
		goto theItem
	}

	fmt.Println("结束", a)
}

func foFun() {
	for i := 0; i < 5; i++ {
		fmt.Println("i", i)
	}
}
