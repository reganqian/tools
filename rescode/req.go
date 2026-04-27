package rescode

type PageReq struct {
	PageNo   int32 `json:"pageNo"`
	PageSize int32 `json:"pageSize"`
}

func (p *PageReq) GetPageFrom() int32 {
	if p.PageNo <= 0 {
		p.PageNo = 1
	}
	if p.PageSize <= 0 {
		p.PageSize = 10
	}
	pageFrom := p.PageSize * (p.PageNo - 1)
	return pageFrom
}

// @Description  初始化分页请求
func InitPageReq(pageNo, pageSize int32) (int32, int32) {
	if pageNo <= 0 {
		pageNo = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	pageFrom := (pageNo - 1) * pageSize
	return pageSize, pageFrom
}
