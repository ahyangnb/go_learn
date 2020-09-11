package main

import "fmt"

func numA(a, b int) int {
	return a + b
}

func numB() (v int) {
	v = 20
	return v
}

func numC(a, b, c int, s string) (v1 string, v2, v3 int) {
	v1 = s + "哈哈哈"
	v2 = a + b
	v3 = b + c
	return
}

func numD() (int, string) {
	return 100, "哈哈哈"
}

func testFun() {
	a := numA(1, 2)
	b := numB()
	str, _, v3 := numC(1, 6, 8, "哈哈")
	n1, n2 := numD()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("str::%v,,v2::%v,,v3::%v\n", str, "测试下划线", v3)
	fmt.Printf("n1::%v,,n2::%v\n", n1, n2)
}
