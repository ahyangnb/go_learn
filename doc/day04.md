# 类型
自定义类型打印出来的是自己定义的类型名，类型别名打印出来的还是原来的类型。
```go
// type 后面跟的是类型
type myInt int         // 自定义类型
type myString = string // 类型别名
```

# 结构体

### 结构体的定义
使用type和struct关键字来定义结构体：
```go
type 类型名 struct {
	字段名 字段类型
	字段名 字段类型
	s string
}
```
