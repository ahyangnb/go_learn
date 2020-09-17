package demo_digui

import "fmt"

// 递归：函数自己调用自己
func DemoDiGui() {
	i := f(4)
	fmt.Println(i)
}

// 计算的阶层
func f(i uint64) uint64 {
	if i <= 1 {
		return 1
	}
	return i * f(i-1)
}
