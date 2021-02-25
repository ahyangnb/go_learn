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

func main() {
	var d dog
	da(d)

	var c cat
	da(c)
}
