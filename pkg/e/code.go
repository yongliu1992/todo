package e

const (
	/* 状态码定义 */
	Success       = 200
	Error         = 500
	InvalidParams = 400

	ErrorNotExist              = 10003
	ErrorParamError            = 10011
	ErrorExistRecord           = 10013
	ENeedParam                 = 10014
	ErrorRecordNotFound        = 10018
	ErrorAuthCheckTokenFail    = 20001
	ErrorAuthCheckTokenTimeout = 20002
	ErrorAuthToken             = 20003
	ErrorAuth                  = 20004
)
