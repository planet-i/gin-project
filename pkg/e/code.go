package e

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400 // 请求参数错误

	ERROR_EXIST_TAG       = 10001 // 已存在该标签名称
	ERROR_EXIST_TAG_FAIL  = 10002 // 获取已存在标签失败
	ERROR_NOT_EXIST_TAG   = 10003 // 该标签不存在
	ERROR_GET_TAGS_FAIL   = 10004 // 获取所有标签失败
	ERROR_COUNT_TAG_FAIL  = 10005 // 统计标签失败
	ERROR_ADD_TAG_FAIL    = 10006 // 新增标签失败
	ERROR_EDIT_TAG_FAIL   = 10007 // 修改标签失败
	ERROR_DELETE_TAG_FAIL = 10008 // 删除标签失败
	ERROR_EXPORT_TAG_FAIL = 10009 // 导出标签失败
	ERROR_IMPORT_TAG_FAIL = 10010 // 导入标签失败

	ERROR_NOT_EXIST_ARTICLE        = 10011 // 该文章不存在
	ERROR_CHECK_EXIST_ARTICLE_FAIL = 10012 // 新增文章失败
	ERROR_ADD_ARTICLE_FAIL         = 10013 // 删除文章失败
	ERROR_DELETE_ARTICLE_FAIL      = 10014 // 检查文章是否存在失败
	ERROR_EDIT_ARTICLE_FAIL        = 10015 // 修改文章失败
	ERROR_COUNT_ARTICLE_FAIL       = 10016 // 统计文章失败
	ERROR_GET_ARTICLES_FAIL        = 10017 // 获取多个文章失败
	ERROR_GET_ARTICLE_FAIL         = 10018 // 获取单个文章失败
	ERROR_GEN_ARTICLE_POSTER_FAIL  = 10019 // 生成文章海报失败

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001 // Token鉴权失败
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002 // Token已超时
	ERROR_AUTH_TOKEN               = 20003 // Token生成失败
	ERROR_AUTH                     = 20004 // Token错误

	ERROR_UPLOAD_SAVE_IMAGE_FAIL    = 30001 // 保存图片失败
	ERROR_UPLOAD_CHECK_IMAGE_FAIL   = 30002 // 检查图片失败
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT = 30003 // 校验图片错误，图片格式或大小有问题
)
