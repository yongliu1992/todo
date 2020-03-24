package e
/* 存放信息映射*/
var MsgFlags = map[int]string{
	Success:                    "成功",
	Error:                      "fail",
	InvalidParams:              "请求参数错误",
	ErrorExistRecord:           "记录已存在,请勿重复添加",
	ErrorNotExist:              "记录不存在",
	ENeedParam:                 "缺少参数",
	ErrorAuthCheckTokenFail:    "Token鉴权失败",
	ErrorAuthCheckTokenTimeout: "Token已超时",
	ErrorAuthToken:             "Token生成失败",
	ErrorAuth:                  "Token错误",
	ErrorParamError:            "参数错误",
	ErrorRecordNotFound:        "记录找不到",
}
/* 获取返回值 */
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[Error]
}
