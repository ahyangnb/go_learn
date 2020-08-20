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











