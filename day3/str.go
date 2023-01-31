package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	var html string = `
	<html>
		<head>  HELLO WORLD </head>
	</ html>
`
	fmt.Printf("html: %v\n", html)

	// strJoin()

	// getBuffer()

	// sliceStr()

	// splitStr()

	trimStr()
}

// strings.join()

func strJoin() {
	str1 := "hello"
	str2 := "world"
	s := strings.Join([]string{str1, str2}, "-")
	c := str1 + "_" + str2
	fmt.Printf("s: %v\n", s)
	fmt.Printf("c: %v\n", c)
}

// buffer.writeString 进行写入
// 得到值 buffer.String()
func getBuffer() {
	var buffer bytes.Buffer
	buffer.WriteString("hello")
	buffer.WriteString("__")
	buffer.WriteString("jack")
	fmt.Printf("buffer: %v\n", buffer.String())
}

// 转义字符
// \t ,\,  \\

// 字符串切片
func sliceStr() {
	str := "hello world"
	n := 3
	m := 5
	fmt.Println(str[n:m]) //llo
	fmt.Println(str[n])   // 获取取字符串素引位置为n的原始字节
	fmt.Println(str[n:])  // n后面所有
	fmt.Println(str[:m])  //到m
}

// 字符串函数 len(str)
func getStrLen() {
	str := "helllo jackkk"
	length := len(str)
	fmt.Printf("length: %v\n", length)
}

// split
func splitStr() {
	str := "helllo,ackkk"
	v := strings.Split(str, ",")

	fmt.Printf("str: %v\n", str)
	fmt.Printf("v: %v\n", v)
}

//trim

func trimStr() {
	str := " heckkk"
	// strings.Trim(str)
	fmt.Printf("str: %v\n", str)
}

//replace

//repeat

//index

//fileds
