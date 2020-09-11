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

	// 数组初始化1
	var b1 [3]bool
	// 数组赋值
	b1 = [3]bool{false, true, true}
	fmt.Printf("%T\n", b1)

	// 数组初始化2
	b2 := [2]bool{false, true}
	fmt.Printf("%v\n", b2)

	// 数组初始化3：根据初始值自动推断数组长度
	b3 := [...]int{0, 1, 2, 3}
	fmt.Printf("值::%v，类型::%T\n", b3, b3)

	// 数组初始化4：不给的值使用默认值
	b4 := [5]int{0, 1}
	fmt.Printf("值::%v，类型::%T\n", b4, b4)

	// 数组初始化5：根据索引来初始化
	b5 := [5]int{0: 0, 3: 1}
	fmt.Printf("值::%v，类型::%T\n", b5, b5)

	// 数组遍历
	d1 := [...]int{3, 1, 2, 9}
	for i := 0; i < len(d1); i += 1 {
		fmt.Printf("当前索引::%v 值::%v", i, d1[i])
	}

	// Range遍历
	for i, v := range d1 {
		fmt.Printf("Range遍历当前索引::%v 值::%v", i, v)
	}

	// 多维数组
	d3 := [3][2]int{{0, 2}, {9, 2}, {7, 3}}
	for i, v := range d3 {
		fmt.Printf("多维遍历当前索引::%v 值::%v", i, v)
	}

	fmt.Println("================")

	// 练习：找出和为5的两个下标{2,1,4,5,3}
	d4 := [...]int{2, 1, 4, 5, 3}
	for i := 0; i < len(d4); i++ {
		for j := i + 1; j < len(d4); j++ {
			if d4[i]+d4[j] == 5 {
				fmt.Printf("(i::%v,,j::%v)\n", i, j)
			}
		}
	}
}
