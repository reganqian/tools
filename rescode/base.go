package rescode

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

type StringListReply struct {
	BaseReply
	DataList []string `json:"dataList"`
}

type StringReply struct {
	BaseReply
	Data string `json:"data"`
}

type IntReply struct {
	BaseReply
	Data int `json:"data"`
}
