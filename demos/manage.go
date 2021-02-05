package main

import (
	"fmt"
	"os"
)

var (
	allStudent map[int64]*student
)

type student struct {
	id   int64
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
			viewAllStudent()
		case 2:
			addStudent()
		case 3:
			delStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("没")
		}
	}

}

func viewAllStudent() {
	for k, v := range allStudent {
		fmt.Printf("学号::%v,名字::%v\n", k, v.name)
	}
}

func addStudent() {
	var (
		id   int64
		name string
	)

	fmt.Print("清输入学号")
	fmt.Scanln(&id)
	fmt.Print("清输入名字")
	fmt.Scanln(&name)

	newStudent := &student{
		id:   id,
		name: name,
	}

	allStudent[id] = newStudent
}

func delStudent() {
	var (
		id int64
	)
	fmt.Print("清输入学号")
	fmt.Scanln(&id)
	delete(allStudent, id)
}
