package main

import (
	"fmt"
	"os"
)

// 文件处理
func main() {
	// 相对路径，需要在当前目录运行
	file, err := os.Open("./file_demo.go")
	if err != nil {
		fmt.Printf("发生异常::%v\n", err)
	}

	// 要记得关闭文件
	defer file.Close()

	fmt.Printf("读取文件名::%v\n", file.Name())

	var s [128]byte
	// 读取并切片
	length, err := file.Read(s[:])
	if err != nil {
		fmt.Printf("Read file error::%v\n", err)
	}

	fmt.Printf("文件长度::%v\n", length)
	// 打印内容（切片转string）
	fmt.Printf("文件内容::%v\n", string(s[:length]))

}
