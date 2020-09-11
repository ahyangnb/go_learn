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
	fmt.Printf("字符::%v %v %v\n", c1, c2, c3)

	// 数组初始化
	var b1 [3]bool
	// 数组赋值
	b1 = [3]bool{false, true, true}
	// 打印类型[数组长度是类型的一部分]
	fmt.Printf("%T\n", b1)
}
