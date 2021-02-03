package main

import (
	"fmt"
	"os"
)

var (
	allStudent map[int64]*student
)

type student struct {
	id   int
	name string
}

func main() {
	allStudent = make(map[int64]*student, 48)
	for {
		fmt.Println("欢迎进入学生管理系统")
		fmt.Println(`
		"1.查看所有学生；" +
		"2.新增学生；" +
		"3.删除学生" +
		"4.退出"`)
		fmt.Print("请输入指令:")

		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择了:%d\n", choice)

		switch choice {
		case 1:
			fmt.Println("查看所有学生")
			break
		case 2:
			fmt.Println("新增学生")
		case 3:
			fmt.Println("删除学生")
		case 4:
			os.Exit(1)
		default:
			fmt.Println("没")

		}
	}

}
