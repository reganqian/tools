package rescode

// @Description  公共返回对象
type BaseReply struct {
	ResCode int32  `json:"resCode"`
	ResDesc string `json:"resDesc"`
}

func (s *BaseReply) Failed(resCode int32, resDesc string) {
	s.ResCode = resCode
	s.ResDesc = resDesc
}

func (s *BaseReply) Success() {
	s.ResCode = SUCCESS
	s.ResDesc = DEFAULT_SUCCESS_DESC
}

// @Description  字符串数组返回对象
type StringListReply struct {
	BaseReply
	DataList []string `json:"dataList"`
}

// @Description  字符串返回对象
type StringReply struct {
	BaseReply
	Data string `json:"data"`
}

// @Description 数字返回对象
type IntReply struct {
	BaseReply
	Data int `json:"data"`
}
