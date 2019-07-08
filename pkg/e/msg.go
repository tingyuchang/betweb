package e

var MsgFlags = map[int]string{
	SUCCESS:            "OK",
	ERROR:              "Fail",
	INVALID_PARAMS:     "Request params invalid",
	ERROR_EXIST_MEMBER: "User already exist",

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token authenticates error",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token checks timeout",
	ERROR_AUTH_TOKEN:               "Token can't be generated",
	ERROR_AUTH:                     "Token error",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]

	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
