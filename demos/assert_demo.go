package main

import "fmt"

func assert0(value interface{}) {
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

func assert1(value interface{}) {
	str, isOk := value.(string)
	if isOk {
		fmt.Printf("isOk，string，值是：%v\n", str)
	} else {
		fmt.Printf("isOk，猜错了\n")
	}
}

func main() {
	assert0("哈哈哈") //对的[string]
	assert0(1)     //对的[int]
	assert0(1.1)   //错的

	//第二种判断方式
	assert1(0)
}
