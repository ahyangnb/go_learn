package main

import (
	"testing"
)

// Tips：
// 文件命名一定要`name_test`，否则无法识别为单元测试文件。
// 详情：https://stackoverflow.com/questions/28240489/golang-testing-no-test-files/28240537

// 单元测试，使用`go test demo_test.go`命令
// 即可执行～
func TestFib(t *testing.T) {
	var (
		in       = 7
		expected = 13
	)
	actual := Fib(in)
	if actual != expected {
		t.Errorf("Fib(%d) = %d; expected %d", in, actual, expected)
	}
}

// 单元测试调用的方法
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
