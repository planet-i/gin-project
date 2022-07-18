package e

// 错误对应错误码
const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400 // 请求参数错误

	ERROR_EXIST_TAG         = 10001 // 已存在该标签名称
	ERROR_NOT_EXIST_TAG     = 10002 // 该标签不存在
	ERROR_NOT_EXIST_ARTICLE = 10003 // 该文章不存在

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001 // Token鉴权失败
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002 // Token已超时
	ERROR_AUTH_TOKEN               = 20003 // Token生成失败
	ERROR_AUTH                     = 20004 // Token错误

	ERROR_UPLOAD_SAVE_IMAGE_FAIL    = 30001 // 保存图片失败
	ERROR_UPLOAD_CHECK_IMAGE_FAIL   = 30002 // 检查图片失败
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT = 30003 // 校验图片错误，图片格式或大小有问题
)
