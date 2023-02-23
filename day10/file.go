package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// openFile()
	readFile()
}

/*
	1.打开文件
	2.读写文件
	3.
*/

func openFile() {
	f, err := os.OpenFile("hello.txt", os.O_RDWR|os.O_CREATE, 0750)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		_, err2 := f.Write([]byte("hello jack, 我是测试文件sasasa"))
		if err2 != nil {
			fmt.Printf("err2: %v\n", err2)
		}
		fmt.Printf("f.Name(): %v\n", f.Name())
		f.Close()
	}
}
func readFile() {
	f, _ := os.Open("hello.txt")

	for {
		buf := make([]byte, 10)
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		fmt.Printf("n: %v\n", n)
		fmt.Printf("string(buf): %v\n", string(buf))
	}
	f.Close()

}

/*
	ReadAt()从某个偏移量继续进行阅读
*/

/*
	目录下面进行遍历
*/
