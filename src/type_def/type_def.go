package type_def

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

type TestData struct {
	Num    string
	Status string
	Count  int
	Price  float32
}

type TestRet struct {
	Code     int
	Param    string
	Msg      string
	TestData []TestData
}
