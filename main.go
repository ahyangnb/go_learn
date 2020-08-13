package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	//第一个参数是接口名，第二个参数 http handle func
	http.HandleFunc("/h1", h1)
	http.HandleFunc("/test", testHandle)
	//服务器要监听的主机地址和端口号
	http.ListenAndServe("127.0.0.1:8081", nil)
}

type Data struct {
	Name string
	Age  int
}

type Ret struct {
	Code  int
	Param string
	Msg   string
	Data  []Data
}

func h1(rw http.ResponseWriter, req *http.Request) {
	data := Data{Name: "why", Age: 19}

	ret := new(Ret)
	id := req.FormValue("id")
	//id := req.PostFormValue('id')

	ret.Code = 0
	ret.Param = id
	ret.Msg = "success"
	ret.Data = append(ret.Data, data)
	ret.Data = append(ret.Data, data)
	ret.Data = append(ret.Data, data)
	retJson, _ := json.Marshal(ret)
	rw.Header().Set("Content-Type", "application/json")
	io.WriteString(rw, string(retJson))
}

func testHandle(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprint(rw, "进入到d 测试的接口了")
}
