package main

import "fmt"

func main() {

	phones := [3]string{"小米", "iphone", "华为"}

	fmt.Println("phones", phones)

	item := phones[2]

	fmt.Printf(item)

	// var 二维数组
	// list := [3][2]int{
	// 	{1, 2},
	// 	{4, 5},
	// 	{4, 5},
	// }
	// fmt.Println(list)

	list := [3][2][5]int{
		{
			{1, 4, 3},
			{1, 3, 3},
		},
		{
			{1, 3, 3},
			{1, 3, 3},
		},
		{
			{1, 3, 3},
			{1, 3, 3},
		},
	}
	fmt.Println(list)
}
