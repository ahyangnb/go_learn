package main

import "fmt"

// 指针
func pointer() {
	// 1. &：取地址
	n := 2
	p := &n
	fmt.Println(p)
	fmt.Printf("类型为：%T", p)

	// 2. *：根据地址取值
	v := *p
	fmt.Println(v)
	fmt.Printf("类型为：%T", v)
}
