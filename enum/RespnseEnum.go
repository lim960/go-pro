package enum

var (
	ParamError        = Enum{"Param error", "参数错误", "參數錯誤"}
	ServiceError      = Enum{"Service error", "服务异常", "服務異常"}
	TokenInvalidation = Enum{"token invalidation", "token失效", "token失效"}

	AccountNotExist = Enum{"Account not exist", "账号不存在", "賬號不存在"}
	Registered      = Enum{"Account has been registered, please login", "账号已注册，请登录", "賬號已注冊，請登錄"}

	VerificationCodeError = Enum{"Verification code error", "验证码错误", "驗證碼錯誤"}
	PasswordError         = Enum{"Password error", "密码错误", "密碼錯誤"}

	InsufficientBalance = Enum{"Insufficient balance", "余额不足", "餘額不足"}
)

type Enum struct {
	En string
	Ch string
	Ft string
}
