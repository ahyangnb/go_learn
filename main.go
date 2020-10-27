package main

import (
	"encoding/json"
	"fmt"
	"go_learn/src/server_text"
	"go_learn/src/type_def"
	"io"
	"net/http"
)

func main() {
	server_text.TemplateText()
	//demo_digui.DemoDiGui()
	//demo_error.ErrorHandle()
	//demo_type.TimeTest()
	//demo_file.BufioReadFile()
	//server_text.TemplateText()
	//////第一个参数是接口名，第二个参数 http handle demo_func
	//http.HandleFunc("/h1", h1)
	//http.HandleFunc("/test", testHandle)
	//////服务器要监听的主机地址和端口号
	//fmt.Println("===================")
	//fmt.Println("welcome to my serve")
	//fmt.Println("===================")
	//host := "0.0.0.0:3389"
	////host := "127.0.0.1:3389"
	////host := "81.71.139.64:3389"
	//err := http.ListenAndServe(host, nil)
	//if err != nil {
	//	fmt.Printf("%v", err)
	//}
	//fmt.Printf("listen tcp %v 成功\n", host)

}

func h1(rw http.ResponseWriter, req *http.Request) {
	data := type_def.Data{Name: "why", Age: 19}

	ret := new(type_def.Ret)
	id := req.FormValue("id")

	ret.Code = 0
	ret.Param = id
	ret.Msg = "success"
	ret.Data = append(ret.Data, data)
	ret.Data = append(ret.Data, data)
	ret.Data = append(ret.Data, data)

	retJson, _ := json.Marshal(ret)
	rw.Header().Set("Content-Type", "application/json")
	_, e := io.WriteString(rw, string(retJson))
	if e != nil {
		fmt.Println(e)
	}
}

func testHandle(rw http.ResponseWriter, req *http.Request) {
	data := type_def.TestData{Num: "testNum", Status: "成功", Count: 10, Price: 19.22}

	ret := new(type_def.TestRet)
	id := req.FormValue("id")

	ret.Code = 0
	ret.Param = id
	ret.Msg = "success"
	ret.TestData = append(ret.TestData, data)

	retJson, _ := json.Marshal(ret)
	rw.Header().Set("Content-Type", "application/json")
	_, e := io.WriteString(rw, string(retJson))
	if e != nil {
		fmt.Println(e)
	}
}
