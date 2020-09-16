> 声明非全局变量必须使用，否则无法编译通过。
>
> 同一个作用域不能重复声明变量。

# 变量声明

```go
var name string
var age int
var isOk bool
```



# 批量声明

```go
var (
  name string // default=""
  age int // default=0
  isOk bool // default=false
)
```

# 格式化代码

```sh
go fmt main.go
```

main.go就是要格式化的文件。

# 类型推导

根据值判断是什么类型。

```go
var s1 = ""
```

# 简短变量声明

就是类型推导的简写，只能在函数里面使用。

```go
s1 := "值"
```



# 匿名变量

多重赋值时需要忽略的变量使用下划线“_”来接收，如：

```go
func foo() (int, string) {
	return 10, "ha"
}

func main() {
	x, _ := foo()
	_, y := foo()
	fmt.Print("x=", x)
	fmt.Print("y=", y)
}
```

# iota

计数器，声明常量的时候设置值为iota则第一个的值为0，第二个为1，以此类推。

```go
const (
	a1 = iota // 0
	a2 = iota // 1
  a3	//2
  a4	//3
)
```



# 浮点数

Go语言支持两种浮点型数：`float32`和`float64`。这两种浮点型数据格式遵循`IEEE 754`标准： `float32` 的浮点数的最大范围约为 `3.4e38`，可以使用常量定义：`math.MaxFloat32`。 `float64` 的浮点数的最大范围约为 `1.8e308`，可以使用一个常量定义：`math.MaxFloat64`。

打印浮点数时，可以使用`fmt`包配合动词`%f`，代码如下：

```go
package main
import (
        "fmt"
        "math"
)
func main() {
        fmt.Printf("%f\n", math.Pi)
        fmt.Printf("%.2f\n", math.Pi)
}
```

float32的值不能赋值给float64，因为两个是完全不同的类型。

# fmt占位符

```go
func fmtTest() {
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
}
```

GO语言字符串必须用双引号包裹，单引号包裹的是字符。 

# 字符

单独的字母、汉字、符号表示一个字符。

```go
c1 := '1'
c2 := 'h'
c3 := '哈'
fmt.Printf("字符::%v %v %v", c1, c2, c3)
```

# 位运算

针对的是二进制

| 运算符 |                             描述                             |
| :----: | :----------------------------------------------------------: |
|   &    |    参与运算的两数各对应的二进位相与。 （两位均为1才为1）     |
|   \|   |  参与运算的两数各对应的二进位相或。 （两位有一个为1就为1）   |
|   ^    | 参与运算的两数各对应的二进位相异或，当两对应的二进位相异时，结果为1。 （两位不一样则为1） |
|   <<   | 左移n位就是乘以2的n次方。 “a<<b”是把a的各二进位全部左移b位，高位丢弃，低位补0。 |
|   >>   | 右移n位就是除以2的n次方。 “a>>b”是把a的各二进位全部右移b位。 |



# 数组初始化

数组长度是类型的一部分

```go
// 数组初始化1
var b1 [3]bool
// 数组赋值
b1 = [3]bool{false, true, true}
fmt.Printf("%T\n", b1)

// 数组初始化2
b2 := [2]bool {false,true}
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
```

# 数组遍历

```go
// 数组遍历
d1 := [...]int{0, 1, 2, 9}
for i := 0; i < len(d1); i += 1 {
fmt.Printf("当前索引::%v 值::%v", i, d1[i])
}
```

Range遍历

```go
for i, v := range d1 {
	fmt.Printf("Range遍历当前索引::%v 值::%v", i, v)
}
```

多维数组

```go
d3 := [3][2]int{{0, 2}, {9, 2}, {7, 3}}
for i, v := range d3 {
	fmt.Printf("多维遍历当前索引::%v 值::%v", i, v)
}
```

练习：找出和为5的两个下标{2,1,4,5,3}

```go
d4 := [...]int{2, 1, 4, 5, 3}
for i := 0; i < len(d4); i++ {
	for j := i + 1; j < len(d4); j++ {
		if d4[i]+d4[j] == 5 {
			fmt.Printf("(i::%v,,j::%v)\n", i, j)
		}
	}
}
```

# 切片

完整：https://www.liwenzhou.com/posts/Go/06_slice/

```go
a1 := [...]int{2, 4, 5, 9, 2}
s1 := a1[1:3] //从1切到3但不包含3
s2 := a1[1:]  //从1切到最后
s3 := a1[:3]  //从最开始切到3
fmt.Println(s1)
fmt.Println(s2)
fmt.Println(s3)
```

切片的容量为切片开始位置到数组的长度，

##### 本质

切片就是一个框，框住了一块连续内存。

切片属于引用类型，真正的数据其实都是保存在底层数组里的。

##### 判断

要判断一个切片是否为空需要用len(s) == 0来判断长度是否等于空，不应该使用s == nil来判断。

### 切片内部逻辑

![image-20200911115418624](/Users/q1/Library/Application Support/typora-user-images/image-20200911115418624.png)

### 切片中删除数值

删除下标为2的值

```go
a2 := []int{6, 1, 2}
fmt.Println(a1)
a2 = append(a2[:1], a2[2:]...)
fmt.Println(a1)
```

打印

```
[6 1 2]
[6 2]
```

# 指针

```go
// 1. &：取地址
n := 2
p := &n
fmt.Println(p)
fmt.Printf("类型为：%T",p)

// 2. *：根据地址取值
v := *p
fmt.Println(v)
fmt.Printf("类型为：%T",v)
```

# 异常处理
程序异常抛出: `panic`
处理异常: `recover`
```go
	// 刚刚打开数据库连接
	defer func() {
		e := recover()
		fmt.Printf("发现错误::%v\n", e)
		fmt.Println("释放数据库连接...")
	}()
	panic("出现严重的错误")
```
### 注意事项
1. recover()必须搭配defer使用；

2. defer一定要在可能引发panic的语句之前定义；






