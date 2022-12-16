package main

import "fmt"

func main() {
	a := 5

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
