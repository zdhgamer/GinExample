package e

const (
	//成功
	SUCCESS = 200
	//错误
	ERROR = 500
	//非法参数
	INVALID_PARAMS = 400

	//已存在tag
	ERROR_EXIST_TAG = 10001
	//不存在tag
	ERROR_NOT_EXIST_TAG = 10002
	//不存在文章
	ERROR_NOT_EXIST_ARTICLE = 10003


	//token验证失败
	ERROR_AUTH_CHECK_TOKEN_FAIL = 20001
	//token超时
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	//token生成失败
	ERROR_AUTH_TOKEN = 20003
	//token错误
	ERROR_AUTH = 20004
)