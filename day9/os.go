package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// criteFile()
	// mkdir()
	// mkdirAll()
	// remove()
	// readFile()
	// chdir()
	// makeTmp()
	getPWd()
}

/*
	增加目录
	创建文件
	删除目录
	修改
	获得目录
	创建临时目录
	重命名
	读文件
	写文件
*/

func criteFile() {
	f, err := os.Create("heelo.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("f.Name(): %v\n", f.Name())
}

func mkdir() {
	// 创建文件件
	err := os.Mkdir("mikeDir", 0660)
	if err != nil {
		log.Fatal(err)
	}
	// 创建文件
	// err1 := os.WriteFile("mikeDir/hello.txt", []byte("helllo jack, i'm lucy"), 0660)
	// if err1 != nil {
	// 	log.Fatal(err1)
	// }

}

func mkdirAll() {
	err := os.MkdirAll("test/subdir", 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
	err = os.WriteFile("test/subdir/testfile.txt", []byte("Hello, Gophers!"), 0660)
	if err != nil {
		log.Fatal(err)
	}
}

func remove() {
	err := os.RemoveAll("jack")
	if err != nil {
		// return
		log.Fatal(err)
	}
}

func readFile() {
	b, err := os.ReadFile("test/subdir/testfile.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	os.Stdout.Write(b)
}

func chdir() {
	err := os.Chdir("test/subdir")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
}

/*
	创建一个临时文件
*/
func makeTmp() {
	s, err := os.MkdirTemp("hello", "jack")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Printf("s: %v\n", s)
}

/*
	获取当前路径
*/

func getPWd() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Printf("dir: %v\n", dir)
}

/*
	是否存在
*/

/*
	是否不存在
*/

/*
	是否具有权限
*/
func isPermission() {
}
