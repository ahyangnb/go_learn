package demo_error

import "fmt"

func funcA() {
	m1 := make(map[string]string)
	m1["a"] = "test"
	fmt.Println(m1)
}
func funcB() {
	// 刚刚打开数据库连接
	defer func() {
		e := recover()
		fmt.Printf("发现错误::%v\n", e)
		fmt.Println("释放数据库连接...")
	}()
	panic("出现严重的错误") // defer一定要写在我之前
	fmt.Println("我不会执行到，因为前面抛出异常导致崩溃退出了")
}
func funcC() {
	fmt.Println("c")
}

func ErrorHandle() {
	fmt.Println("welcome to error handle.")
	funcA()
	funcB()
	funcC()
}
