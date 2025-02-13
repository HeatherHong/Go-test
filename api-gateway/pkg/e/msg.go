package e

var MsgFlags = map[uint]string{
	SUCCESS: "ok",
	ERROR:   "fail",

	InvalidParams:              "Request parameter error",
	HaveSignUp:                 "Already signed up",
	ErrorActivityTimeout:       "The event has expired",
	ErrorAuthCheckTokenFail:    "Token authentication failure",
	ErrorAuthCheckTokenTimeout: "Token timed out",
	ErrorAuthToken:             "Token generation failure",
	ErrorAuth:                  "Token Error",
	ErrorNotCompare:            "Mismatch",
	ErrorDatabase:              "Database operation error, please try again",
	ErrorAuthNotFound:          "Token cannot be empty",

	ErrorServiceUnavailable: "Overload protection, service temporarily unavailable",
	ErrorDeadline:           "Service call timeout",
}

// GetMsg 获取状态码对应信息
func GetMsg(code uint) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
