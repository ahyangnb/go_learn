package demo_type

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"time"
)

func foo() (int, string) {
	return 10, "ha"
}

func day02() {
	x, _ := foo()
	_, y := foo()
	fmt.Print("x=", x)
	fmt.Print("y=", y)
}

func formatTest() {
	var n = 100
	// 查看类型 T
	fmt.Printf("类型::%T\n", n)
	// 查看值 v
	fmt.Printf("值::%v\n", n)

	// 字符串
	s1 := "我是测试文字"
	fmt.Printf("打印::%v\n", s1)
	// 加了类型描述符后的
	fmt.Printf("打印::%#v", s1)

	// 字符
	c1 := '1'
	c2 := 'h'
	c3 := '哈'
	fmt.Printf("字符::%v %v %v\n", c1, c2, c3)

	// 数组初始化1
	var b1 [3]bool
	// 数组赋值
	b1 = [3]bool{false, true, true}
	fmt.Printf("%T\n", b1)

	// 数组初始化2
	b2 := [2]bool{false, true}
	fmt.Printf("%v\n", b2)

	// 数组初始化3：根据初始值自动推断数组长度
	b3 := [...]int{0, 1, 2, 3}
	fmt.Printf("值::%v，类型::%T\n", b3, b3)

	// 数组初始化4：不给的值使用默认值
	b4 := [5]int{0, 1}
	fmt.Printf("值::%v，类型::%T\n", b4, b4)

	// 数组初始化5：根据索引来初始化
	b5 := [5]int{0: 0, 3: 1}
	fmt.Printf("值::%v，类型::%T\n", b5, b5)

	// 数组遍历
	d1 := [...]int{3, 1, 2, 9}
	for i := 0; i < len(d1); i += 1 {
		fmt.Printf("当前索引::%v 值::%v", i, d1[i])
	}

	// Range遍历
	for i, v := range d1 {
		fmt.Printf("Range遍历当前索引::%v 值::%v", i, v)
	}

	// 多维数组
	d3 := [3][2]int{{0, 2}, {9, 2}, {7, 3}}
	for i, v := range d3 {
		fmt.Printf("多维遍历当前索引::%v 值::%v", i, v)
	}

	fmt.Println("================")

	// 练习：找出和为5的两个下标{2,1,4,5,3}
	d4 := [...]int{2, 1, 4, 5, 3}
	for i := 0; i < len(d4); i++ {
		for j := i + 1; j < len(d4); j++ {
			if d4[i]+d4[j] == 5 {
				fmt.Printf("(i::%v,,j::%v)\n", i, j)
			}
		}
	}
}

func sliceDemo() {
	a1 := [...]int{2, 4, 5, 9, 2}
	s1 := a1[1:3] //从1切到3但不包含3
	s2 := a1[1:]  //从1切到最后
	s3 := a1[:3]  //从最开始切到3
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	fmt.Println("=============")

	// 删除下标为2的值
	a2 := []int{6, 1, 2}
	fmt.Println(a1)
	a2 = append(a2[:1], a2[2:]...)
	fmt.Println(a1)
}

// 类型别名和自定义类型
type myInt int         // 自定义类型
type myString = string // 类型别名

func MyType() {
	var i myInt
	i = 1
	fmt.Println(i)

	var s myString
	s = "test"
	fmt.Println(s)

	// rune就是官方帮我们写好的类型别名
	var r rune
	r = 2
	fmt.Println(r)
}

// 结构体
type Person struct {
	age   int
	name  string
	hobby []string
}

func StructDemo() {
	var p Person
	p.name = "q1"
	p.age = 19
	p.hobby = []string{"篮球", "乒乓球"}

	fmt.Println(p.name)
	fmt.Printf("%T", p)

	// 匿名结构体
	var s struct {
		x string
		y int
	}
	s.x = "嘿嘿黑"
	s.y = 1
	fmt.Printf("hahah::%T", s)

	// 结构体初始化1
	var p1 = new(Person)
	p1.hobby = []string{"2", "嘿嘿嘿"}
	p1.name = "q2"
	p1.age = 18
	fmt.Println(p1)

	// 结构体初始化2(Key-value形式初始化)
	var p2 = Person{
		name:  "q3",
		age:   10,
		hobby: []string{"啦啦啦"},
	}
	fmt.Println(p2)

}

type s struct {
	data map[string]interface{}
}

// https://www.liwenzhou.com/posts/Go/gob_msgpack/
func GobTest() {
	var v1 = s{
		data: make(map[string]interface{}, 8),
	}
	v1.data["a"] = 1

	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(v1.data)
	if err != nil {
		fmt.Println("gob encode failed, err:", err)
		return
	}
	b := buf.Bytes()
	fmt.Println(b)
	var s2 = s{
		data: make(map[string]interface{}, 8),
	}
	// decode
	dec := gob.NewDecoder(bytes.NewBuffer(b))
	err = dec.Decode(&s2.data)
	if err != nil {
		fmt.Println("gob decode failed, err", err)
		return
	}
	fmt.Println("====s2.data===")
	fmt.Println(s2.data)
	for _, v := range s2.data {
		fmt.Printf("value:%v, type:%T\n", v, v)
	}

	//json
	ret, err := json.Marshal(s2.data)
	if err != nil {
		fmt.Println("marshal failed", err)
	}
	fmt.Println("====json===")
	fmt.Printf("%v\n", string(ret))
	fmt.Println("====json(encode)===")
	fmt.Printf("%#v\n", string(ret))
}

// 时间 https://www.liwenzhou.com/posts/Go/go_time/
func TimeTest() bool {
	now := time.Now()
	fmt.Printf("current time:%v\n", now)

	minute := now.Minute()
	fmt.Printf("%v \n", minute)

	// add
	later := now.Add(time.Hour)
	fmt.Printf("%v", later.Minute())

	return true
}
