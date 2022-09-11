package main

import "fmt"

/*
	切片又叫slice
	内部有capacity和length()
	capacity指总数组长度
	length指代已经使用的长度
*/
func main() {

	// nullCut()

	// copyCut()

	// mapCut()

	// appendCut()

	// splitArr()

	// changeCut()

	copyfun()
}

func fullList() {
	var s []int = []int{1, 2, 3, 4, 5} // 满容的
	fmt.Println(s, len(s), cap(s))
}

// 在创建切片时，还有两个非常特殊的情况需要考虑，那就是容量和长度都是零的切片，叫着「空切片」
func nullCut() {
	var s1 []int           // nil 切片 只是声明了切片,自动赋予了0值
	var s2 []int = []int{} // 空切片
	var s3 []int = make([]int, 0)
	fmt.Println(s1, s2, s3)
	fmt.Println(len(s1), len(s2), len(s3))
	fmt.Println(cap(s1), cap(s2), cap(s3))
}

// 切片赋值本质上是一次浅拷贝
func copyCut() {
	var arr = make([]int, 8)

	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1
	}

	var arr1 = arr

	arr1[2] = 10

	// 修改值后发现值都改变了
	fmt.Println(arr, arr1)
}

// 切片的便利

func mapCut() {
	// 满容切片
	var s = []int{1, 2, 3, 4, 5}

	for i, v := range s {
		fmt.Println(i, v)
	}

	for j := 0; j < len(s); j++ {
		fmt.Println(j, s[j])
	}
}

func appendCut() {
	s1 := []int{1, 1}

	_ = append(s1, 5)
	// s3 := append(s2, 5)
	// s4 := append(s3, 3)

	fmt.Println(s1)
	// fmt.Println(s1, s2, s3, s4)
	// fmt.Println(cap(s1), cap(s2), cap(s3))

}

/*
	切割数组
	arr[start: end] => [start: end)

	arr[:] => 所有
	arr[: length] => 从最开始返回长度为
	arr[start:] => 从strat到最后
	1. 切割数组
	开始和结束位置都是可选的
	不写开始位置 默认是从母切片的开始位置
	不写结束位置 默认是母切片
*/
func splitArr() {
	arr1 := []string{"s", "s1", "s2", "s3", "s2", "123"}

	var arr2 = arr1[1:3]

	fmt.Println(arr2, cap(arr2), len(arr2))
	fmt.Println(arr1, cap(arr1), len(arr1))
}

/*
	共享底层数组
	只是浅拷贝
*/
func changeCut() {
	var a = [5]int{1, 2, 3, 4, 5}
	var b = a[1:4]

	b[0] = 10

	fmt.Println(a, b)
}

/*
	copy函数, 用来进行切片的深拷贝

	注意:
	如果数组里面装的是指针，比如 []*int 类型，那么指针指向的内容还是共享的。
*/

// copy 只能用于slice
func copyfun() {
	var a = []int{1, 2, 3, 4, 5}
	var b = make([]int, len(a))
	var d = copy(b, a)

	fmt.Println(d, b)

}
