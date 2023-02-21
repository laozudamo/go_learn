package main

import "fmt"

func main() {

	test()
}

/*

接口像是一个公司里面的领导，他会定义一些通用规范，只设计规范，而不实现规范。
go语言的接口，是一种新的类型定义，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口;。


type interface_name interface {
	method_namel [return_type]
	method_name2 [return_type]
	method_name3 [return_type]
	method_name4 [return_type]
}
*/
type USB interface {
	read()
	write()
}

type Mobile struct {
}

type Computer struct {
	name string
}

func (c Computer) read() {
	fmt.Printf("%v\n, read-----", c.name)
}

func (c Computer) write() {
	fmt.Printf("%v\n, write......", c.name)
}

func test() {
	c := Computer{
		name: "我的计算机",
	}
	c.read()
	c.write()
}
