package main

import (
	"encoding/json"
	"fmt"
)

type testJson struct {
	// json:"id"表示在json里key的名字，这里是写成小写开头的
	Id   int64 `json:"id"`
	Name string
}

func main() {
	demoJson := newTestJson(11, "哈哈哈")

	// 1.json序列化
	v, err := json.Marshal(demoJson)
	if err != nil {
		fmt.Printf("error::%v", err)
	}
	// 以string方式打印json
	fmt.Printf("json序列化:%v", string(v))

}

func newTestJson(Id int64, Name string) testJson {
	return testJson{
		Id:   Id,
		Name: Name,
	}
}
