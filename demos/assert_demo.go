package main

import "fmt"

func assert(value interface{}) {
	// 判断是否为string
	str, strOk := value.(string)
	if strOk {
		fmt.Printf("猜对了，值是：%v\n", str)
	} else {
		fmt.Printf("猜错了\n")
	}
}

func main() {
	assert("哈哈哈") //对的
	assert(1)     //错的
}
