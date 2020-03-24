package e

var MsgFlags = map[int]string{
	SUCCESS:                        "成功",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "请求参数错误",
	ERROR_EXIST_RECORD:             "记录已存在,请勿重复添加",
	ERROR_NOT_EXIST:                "记录不存在",
	ERROR_NEED_PARAM:               "缺少参数",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
	ERROR_PARAM_QUESTION_ERROR:     "问题状态不对",
	ERROR_PARAM_ERROR:              "参数错误",
	ErrorRecordNotFound:            "记录找不到",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
