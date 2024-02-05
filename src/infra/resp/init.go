package resp

import "github.com/gin-gonic/gin"

// gin.H封装了生成json数据的工具
var (
	ResolveParamsError         = gin.H{"errorCode": 10000, "msg": "无法解析参数"}
	UsernameNotExistError      = gin.H{"errorCode": 10001, "msg": "用户名不存在"}
	PasswordError              = gin.H{"errorCode": 10002, "msg": "密码错误"}
	UsernameOrPasswordError    = gin.H{"errorCode": 10003, "msg": "用户名或密码错误"}
	NoOperationPermissionError = gin.H{"errorCode": 10004, "msg": "无操作权限"}
	NoLoginError               = gin.H{"errorCode": 10005, "msg": "未登录"}
	SystemInnerError           = gin.H{"errorCode": 10006, "msg": "系统内部错误"}
)

func Fail(errorCode int, data interface{}) gin.H {
	return gin.H{"errorCode": errorCode, "data": data}
}

func Success(data interface{}) gin.H {
	return gin.H{"errorCode": 200, "data": data}
}
