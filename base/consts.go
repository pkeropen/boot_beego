package base

const ADMIN_SESSION_KEY string = "adminmodel"
const USER_SESSION_KEY string = "usermodel"
/*const WIDTH int = 5
var area int
const a, b, c = 1, false, "str" //多重赋值*/


type JsonResultCode int

const (
	JRCodeSucc JsonResultCode = iota
	JRCodeFailed
	JRCode302 = 302 //跳转至地址
	JRCode401 = 401 //未授权访问
)

const (
	Deleted = iota - 1
	Disabled
	Enabled
)
