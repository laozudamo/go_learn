package main

import "fmt"

func main() {

	test()
}

type Phone struct {
	price   string
	num     int
	version int
}

func test() {

	// newPhone := Phone{
	// 	"1888",
	// 	100,
	// 	11,
	// }

	// var p_phone *Phone

	/*
		用new关键字 创结构体指针
	*/
	p_phone := new(Phone)

	fmt.Printf("p_phone: %p\n", p_phone)

}
