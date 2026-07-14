package rescode

// @Title 常量文档
// @Description 生成包含常量的文档
// @Version 1.0
// @Contact.name Money
// @Contact.email qianlaijian@gmail.com
// @BasePath /api/v1
// @Schemes http https
// @Consumes json
// @Produces json
const (
	// @Description 成功
	// @swagger:constant 1
	SUCCESS int32 = 1 //成功
	// @Description 成功
	// @swagger:constant 1
	DEFAULT_SUCCESS int32 = 200 //成功
	// @Description 失败
	// @swagger:constant -1
	FAILED int32 = -1 //失败
	// @Description 参数缺失
	// @swagger:constant -2
	PARAMNULL int32 = -2 //参数缺失
	// @Description 参数错误
	// @swagger:constant -3
	PARAMERR int32 = -3 //参数错误
	// @Description 服务器异常
	// @Swagger:constant -4
	SERVERERR int32 = -4 //服务器异常
	// @Description 数据异常
	// @Swagger:constant -5
	DBERROR int32 = -5 //数据异常
	// @Description 登录TOKEN错误
	// @Swagger:constant -6
	TOKENERROR int32 = -6 //登录TOKEN错误
	// @Description 数据已存在
	// @Swagger:constant -7
	DATAEXIST int32 = -7 //数据已存在
	// @Description 数据不存在
	// @Swagger:constant -8
	DATANOTEXIST int32 = -8 //数据不存在

	// @Description 不是拥有者
	// @Swagger:constant -11
	NOTOWNER int32 = -11 //不是拥有者
	// @Description 格式错误
	// @Swagger:constant -12
	FORMATERROR int32 = -12 //格式错误

	// @Description 超时
	// @Swagger:constant -14
	TIMEOUT int32 = -14 //超时
	// @Description 没有权限
	// @Swagger:constant -15
	NO_PERMISSION int32 = -15 //没有权限
	// @Description 状态异常
	// @Swagger:constant -16
	STATUSERROR int32 = -16 //状态异常
	// @Description 争议数据
	// @Swagger:constant -17
	DUPLICATE_DATA int32 = -17 //争议数据
	// @Description 密码错误
	// @Swagger:constant -19
	PWDERROR int32 = -19 //密码错误
	// @Description 需要登录
	// @Swagger:constant -20
	NEED_LOGIN int32 = -20 //需要登录
	// @Description 需要发送验证码
	// @Swagger:constant -22
	NEED_CODE int32 = -22 //需要发送验证码
	// @Description 需要显示返回
	// @Swagger:constant -33
	NEEDSHOWERR int32 = -33 //需要显示返回
	// @Description 更换手机
	// @Swagger:constant -34
	CHANGEPHONE int32 = -34 //更换手机
	// @Description 需要下一步操作
	// @Swagger:constant -88
	MOREACTION int32 = -88 //需要下一步操作
	// @Description 需要鉴权
	// @Swagger:constant -99
	NEEDAUTH int32 = -99 //需要鉴权
	// @Description 需要重定向
	// @Swagger:constant -301
	NEED_REDIRECT int32 = -301

	// @Description 需要完成信息填写
	// @Swagger:constant -111
	NEED_USER_INFO int32 = -111

	NEED_LOGIN_1K   = 1001 //需要登录
	PARAMNULL_1K    = 1002 //参数为空
	FORMATERROR_1K  = 1003 //格式错误
	DBERROR_1K      = 1004 //数据库错误
	DATANOTEXIST_1K = 1005 //数据不存在
	NOPERMISSION_1K = 1006 //无权限
	BALANCE_LESS    = 1009 //余额不足
	PRICE_ERROR     = 1010 //价格错误

	DEFAULT_SUCCESS_DESC    string = "success"
	DATA_NOT_EXIST_DESC     string = "the data is not exist"
	DATA_ALREADY_EXIST_DESC string = "the data is already exist"
	NOTENOUGHDEVICE_DESC    string = "not enough device"
	NEEDADDEQU_DESC         string = "need add more equ"
	NOTOWNER_DESC           string = "you are not the owner"
	FORMATERROR_DESC        string = "format error"
	TASKERR_DESC            string = "some task is not over"
	PARAMERR_DESC           string = "some param is error"
	DBERROR_DESC            string = "db error"
	PARAMNULL_DESC          string = "some params is null"
	DEFAULT_FAILED_DESC     string = "failed"
	CHANGEPHONE_DESC        string = "please use another phonenum"
	NO_PERMISSION_DESC      string = "no permission"
	NEED_LOGIN_DESC         string = "need login"
	STATUSERROR_DESC        string = "status error"
	DUPLICATE_DATA_DESC     string = "Contains some duplicate data."
	PWDERROR_DESC           string = "password error"
)

var ErrorCodeDescriptions1 = map[int32]string{
	SUCCESS:      "成功",
	FAILED:       "失败",
	PARAMNULL:    "参数缺失",
	PARAMERR:     "参数错误",
	SERVERERR:    "服务器异常",
	DBERROR:      "数据异常",
	TOKENERROR:   "登录TOKEN错误",
	DATAEXIST:    "数据已存在",
	DATANOTEXIST: "数据不存在",
	// NOTENOUGHDEVICE: "没有足够的设备",
	// NEEDADDEQU:      "需要添加虚拟设备",
	NOTOWNER:    "不是拥有者",
	FORMATERROR: "格式错误",
	// TASKERR:         "任务冲突",
	TIMEOUT:        "超时",
	NO_PERMISSION:  "没有权限",
	STATUSERROR:    "状态异常",
	DUPLICATE_DATA: "争议数据",
	PWDERROR:       "密码错误",
	NEED_LOGIN:     "需要登录",
	NEED_CODE:      "需要发送验证码",
	NEEDSHOWERR:    "需要显示返回",
	CHANGEPHONE:    "更换手机",
	MOREACTION:     "需要下一步操作",
	NEEDAUTH:       "需要鉴权",
	NEED_REDIRECT:  "需要重定向",
}

// 校验并初始化分页参数
func InitPageParams(pageNo, pageSize int32) (int32, int32) {
	if pageNo == 0 {
		pageNo = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	pageFrom := (pageNo - 1) * pageSize

	return pageSize, pageFrom
}
