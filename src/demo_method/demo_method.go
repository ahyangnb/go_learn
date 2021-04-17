package demo_method

import "fmt"

type Dog struct {
	call string
}

// MethodTest 方法（只有Dog的类型能调用）
func (d Dog) MethodTest() {
	fmt.Printf("%v:::旺", d.call)
}

// 构造
func newDog(call string) Dog {
	return Dog{call: call}
}

func DemoMethod() {
	//第一种方式
	d1 := newDog("1号")
	d1.MethodTest()

	//第二种方式
	d2 := Dog{call: "2号"}
	fmt.Printf("%v:::旺", d2.call)
}
