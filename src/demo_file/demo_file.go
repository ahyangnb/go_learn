package demo_file

import (
	"fmt"
	"io"
	"os"
)

func FileDemo() {
	fmt.Println("===FileDemo===")

	// open
	a, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("error::%v", err)
		return
	}
	fmt.Printf("打开文件成功%v\n", a.Name())
	defer a.Close()

	// read
	var r1 = make([]byte, 128)
	b, err1 := a.Read(r1)
	if err1 == io.EOF {
		fmt.Println("文件读完了")
		return
	}
	if err1 != nil {
		fmt.Printf("error_red::%v", err1)
		return
	}

	fmt.Printf("read success::%v\n", b)
	fmt.Printf("读取了%d字节数据\n", b)
	fmt.Println(string(r1[:b]))
}
