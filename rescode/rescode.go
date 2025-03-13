package rescode

const (
	SUCCESS   int32 = 1  //成功
	FAILED    int32 = -1 //失败
	PARAMNULL int32 = -2 //参数缺失
	PARAMERR  int32 = -3 //参数错误
	SERVERERR int32 = -4 //服务器异常
	DBERROR   int32 = -5 //数据异常

	TOKENERROR      int32 = -6  //登录TOKEN错误
	DATAEXIST       int32 = -7  //数据已存在
	DATANOTEXIST    int32 = -8  //数据不存在
	NOTENOUGHDEVICE int32 = -9  //没有足够的设备
	NEEDADDEQU      int32 = -10 //需要添加虚拟设备
	NOTOWNER        int32 = -11 //不是拥有者
	FORMATERROR     int32 = -12 //格式错误
	TASKERR         int32 = -13 //任务冲突
	TIMEOUT         int32 = -14 //超时
	NO_PERMISSION   int32 = -15 //没有权限
	STATUSERROR     int32 = -16 //状态异常
	DUPLICATE_DATA  int32 = -17
	PWDERROR        int32 = -19 //密码错误
	NEED_LOGIN      int32 = -20 //需要登录
	NEED_CODE       int32 = -22 //需要发送验证码

	NEEDSHOWERR int32 = -33 //需要显示返回
	CHANGEPHONE int32 = -34 //更换手机

	MOREACTION int32 = -88 //需要下一步操作
	NEEDAUTH   int32 = -99 //需要鉴权

	NEED_REDIRECT int32 = -301

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
