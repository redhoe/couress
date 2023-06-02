package tips

// 公共返回错误字典
var (
	Success             = NewAppError(0, "Success")             // 业务成功 0 业务错误
	PasswordEditSuccess = NewAppError(0, "PasswordEditSuccess") // 业务成功 0 业务错误
	SysTemError         = NewAppError(300, "SysTemError")
	TokenError          = NewAppError(301, "TokenExpire")
	PermissionError     = NewAppError(401, "PermissionError")
)

func init() {

	// 普通消息提示
	cn["Success"] = "成功"
	en["Success"] = "success"
	kr["Success"] = "성공"

	cn["SysTemError"] = "系统错误 请联系管理员解决"
	en["SysTemError"] = "system error"

	cn["TokenExpire"] = "令牌过期"
	en["TokenExpire"] = "token expire"

	cn["PermissionError"] = "权限不足"
	en["PermissionError"] = "permission error"

}
