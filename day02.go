package main

import "fmt"

func foo() (int, string) {
	return 10, "ha"
}

func day02() {
	x, _ := foo()
	_, y := foo()
	fmt.Print("x=", x)
	fmt.Print("y=", y)
}

func fmtTest() {
	var n = 100
	// 查看类型 T
	fmt.Printf("类型::%T\n", n)
	// 查看值 v
	fmt.Printf("值::%v\n", n)

	// 字符串
	s1 := "我是测试文字"
	fmt.Printf("打印::%v\n", s1)
	// 加了类型描述符后的
	fmt.Printf("打印::%#v", s1)

	// 字符
	c1 := '1'
	c2 := 'h'
	c3 := '哈'
	fmt.Printf("字符::%v %v %v", c1, c2, c3)
}
