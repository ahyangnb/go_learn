package demo_defer

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func demoDefer() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

//1.a := 1
//2.b := 2
//3.defer calc("1", a, calc("10", a, b))
//4.calc("10", 1, 2) // 10 1 2 3
//5.a = 0
//6.defer calc("2", a, calc("20", a, b))
//7.calc("20", 0, 2)  // 20 0 2 2
//8.b = 1
//9.calc("2", 0, 2)  // 2 0 2 2
//10.calc("1", 1, 3) // 1 1 3 4
