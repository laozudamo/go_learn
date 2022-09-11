package main

import (
	"fmt"
)

/*
	字符串通常有两种设计，一种是「字符」串，一种是「字节」串。
	「字符」串中的每个字都是定长的，而「字节」串中每个字是不定长的。
	type rune int32

	使用「字符」串来表示字符串势必会浪费空间，因为所有的英文字符本来只需要 1 个字节来表示，
	用 rune 字符来表示的话那么剩余的 3 个字节都是零。但是「字符」串有一个好处，那就是可以快速定位
*/
func main() {
	// mapStr()

	splitStr()
}

func mapStr() {
	var s = "嘻哈china"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func splitStr() {
	var str = "hello world"
	var str2 = str[1:3]

	var b = []byte(str)
	var b2 = string(b)
	fmt.Println(b, b2)
	fmt.Println(str2)
}

// 字节码可以修改 而字符串不可修改
