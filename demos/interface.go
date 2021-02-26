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

func main() {
	var d dog
	da(d)

	var c cat
	da(c)

	// 特殊
	var r returnValueType
	test(r)

	a := testA{"哈哈", 1}
	fmt.Printf("name:%v,age:%v\n", a.name, a.age)
}
