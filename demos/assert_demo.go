package main

import "fmt"

func assert(value interface{}) {

	switch value.(type) {
	case string:
		str := value.(string)
		fmt.Printf("string，值是：%v\n", str)
	case int:
		intValue := value.(int)
		fmt.Printf("int，值是：%v\n", intValue)
	default:
		fmt.Printf("猜错了\n")
	}
}

func main() {
	assert("哈哈哈") //对的[string]
	assert(1)     //对的[int]
	assert(1.1)   //错的
}
