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







