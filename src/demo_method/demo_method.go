package demo_method

import "fmt"

type Dog struct {
	call string
}

// 方法（只有Dog的类型能调用）
func (d Dog) MethodTest(a Dog) {
	fmt.Printf("%v:::旺", a.call)
}

// 构造
func newDog(call string) Dog {
	return Dog{call: call}
}

func DemoMethod() {
	d1 := newDog("1号")
	d1.MethodTest(d1)
}
