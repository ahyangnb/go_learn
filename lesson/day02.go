package lesson

import "fmt"

func foo() (int, string) {
	return 10, "ha"
}

func main() {
	x, _ := foo()
	_, y := foo()
	fmt.Print("x=", x)
	fmt.Print("y=", y)
}