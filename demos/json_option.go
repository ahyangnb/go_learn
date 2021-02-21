package main

import (
	"encoding/json"
	"fmt"
)

type testJson struct {
	// json:"id"表示在json里key的名字，这里是写成小写开头的
	Id   int64  `json:"id" db:"name" ini:"name"`
	Name string `json:"name"`
}

func main() {
	demoJson := newTestJson(11, "哈哈哈")

	// 1.json序列化
	v, err := json.Marshal(demoJson)
	if err != nil {
		fmt.Printf("error::%v", err)
	}
	// 以string方式打印json
	fmt.Printf("json序列化:%v\n", string(v))

	// 2.json反序列化
	str := `{"id":11,"name":"哈哈哈"}`
	var testJson1 testJson
	err1 := json.Unmarshal([]byte(str), &testJson1)
	if err1 != nil {
		fmt.Printf("err1::%v", err1)
	}
	// 以string方式打印反序列化json
	fmt.Printf("json反序列化:%v\n", testJson1)
}

func newTestJson(Id int64, Name string) testJson {
	return testJson{
		Id:   Id,
		Name: Name,
	}
}
