package main

import "fmt"

func numA(a, b int) int {
	return a + b
}

func numB() (v int) {
	return 20
}

func numC(a, b, c int, s string) (v1 string, v2, v3 int) {
	return s + "哈哈哈", a + b, b + c
}

func testFun() {
	a := numA(1, 2)
	b := numB()
	str, _, v3 := numC(1, 6, 8, "哈哈")
	fmt.Print(a)
	fmt.Print(b)
	fmt.Printf("str::%v,,v2::%v,,v3::%v", str, "测试下划线", v3)
}
