package rescode

// @Description  公共返回对象
type BaseReply struct {
	ResCode int32  `json:"resCode"` // 返回码
	ResDesc string `json:"resDesc"` // 返回描述
}

type PageData struct {
	TotalNum int32 `json:"totalNum"` // 总条数
	PageNo   int32 `json:"pageNo"`   // 当前页
	PageSize int32 `json:"pageSize"` // 每页条数
}

func (s *PageData) InitData(pageNo, pageSize, totalNum int32) {
	s.PageNo = pageNo
	s.PageSize = pageSize
	s.TotalNum = totalNum
}

func (s *BaseReply) Failed(resCode int32, resDesc string) {
	s.ResCode = resCode
	s.ResDesc = resDesc
}

func (s *BaseReply) Success() {
	s.ResCode = SUCCESS
	s.ResDesc = DEFAULT_SUCCESS_DESC
}

func (s *BaseReply) SetSuccess() {
	s.ResCode = DEFAULT_SUCCESS
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
