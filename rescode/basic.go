package rescode

// @Description  公共返回对象
type BasicReply struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

// // @Description  通用返回对象
// type GeneralReply struct {
// 	BasicReply
// 	Data interface{} `json:"data"`
// }

type PageInfo struct {
	PageNo   int32 `json:"pageNo"`
	PageSize int32 `json:"pageSize"`
	Total    int32 `json:"total"`
}

func (s *BasicReply) Failed(code int32, msg string) {
	s.Code = code
	s.Msg = msg
}

func (s *BasicReply) Success() {
	s.Code = SUCCESS
	s.Msg = DEFAULT_SUCCESS_DESC
}

func (s *BasicReply) SetSuccess() {
	s.Code = DEFAULT_SUCCESS
	s.Msg = DEFAULT_SUCCESS_DESC
}

// @Description  字符串数组返回对象
type StrListReply struct {
	BasicReply
	DataList []string `json:"dataList"`
}

// @Description  字符串返回对象
type StrReply struct {
	BasicReply
	Data string `json:"data"`
}

// @Description 数字返回对象
type IntBasicReply struct {
	BasicReply
	Data int `json:"data"`
}
