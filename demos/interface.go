package main

import "fmt"

type dog struct {
}

type cat struct {
}

func (dog) speak() {
	fmt.Print("::旺旺\n")
}

func (cat) speak() {
	fmt.Print("::喵呜\n")
}

func da(a animal) {
	//接受一个参数，传进来什么就调用什么的speak方法
	a.speak()
}

// 定义一个接口
type animal interface {
	speak()
}

type returnValueType struct {
}

// 特殊的接口
type returnValueInterFace interface {
	// 有返回值的接口
	returnValue() (age int, name string)
	// 有参数的接口
	paramValue(age int, name string)
}

func (returnValueType) returnValue() (age int, name string) {
	fmt.Print("returnValue\n")
	return 1, "哈哈"
}

func (returnValueType) paramValue(age int, name string) {
	fmt.Printf("paramValue,,age:%v,name:%v\n", age, name)
}

func test(r returnValueInterFace) {
	r.paramValue(1, "a")
	r.returnValue()
}

type testA struct {
	name string
	age  int
}

// 指针接收者示例
type pointerType struct {
	name string
}

// 接收的是指针值
func (in *pointerType) test() {
	fmt.Printf("调用指针接收者方式接口::%v\n", in.name)
}

type pointerInterFace interface {
	test()
}

func main() {
	var d dog
	da(d)

	var c cat
	da(c)

	// 特殊
	var r returnValueType
	test(r)

	// 另一个构造的示例
	a := testA{"哈哈", 1}
	fmt.Printf("name:%v,age:%v\n", a.name, a.age)

	// 指针接受者调用接口示例
	var pTest pointerInterFace
	pType := pointerType{"one"}
	// 把指针值给需要接受值的【没问题】
	// 把值给需要接受指针值的【有问题】
	pTest = &pType // 传的是指针值
	pTest.test()
}
