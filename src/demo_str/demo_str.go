package demo_str

import "fmt"

func wordTest() {
	v1 := "上山打老虎"

	r := make([]rune, 0, len(v1))
	for _, c := range v1 {
		r = append(r, c)
		fmt.Printf("value::%v\n", r)
	}

	len1 := len(v1)
	fmt.Print(len1)
}
