package main

import "fmt"

func main() {

	// mapFun()
	// mapDel()
	hasDst()
}

// map key value
func mapFun() {
	theMap := make(map[string]int)
	theMap["age"] = 18
	fmt.Printf("theMap: %T\n", theMap)

	m := map[string]string{
		"sex":   "男",
		"phone": "1568131923",
	}
	fmt.Printf("m: %v\n", m)
}

func mapDel() {
	m := map[string]int{
		"age":    19,
		"height": 199,
	}

	// for k := range m {
	// 	delete(m, k)
	// }
	delete(m, "age")
	fmt.Printf("m: %v\n", m)
	newMap := m
	fmt.Printf("newMap: %v\n", newMap)
	// for k, v := range m {
	// 	fmt.Printf("v: %v\n", v)
	// 	fmt.Printf("k: %v\n", k)
	// }
}

// 判断属性存在

func hasDst() {
	m := map[string]int{
		"age":    19,
		"height": 199,
	}
	k := "age"
	// 第一个是值, 第二个是布尔值
	v, ok := m[k]
	fmt.Printf("v: %v\n", v)
	fmt.Printf("ok: %v\n", ok)
}
