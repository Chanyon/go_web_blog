package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	// user error code 1000~
	ERROR_USERNAME_USED       = 1001
	ERROR_USER_NOT_FOUND      = 1002
	ERROR_USER_PASSWORD_WRONG = 1003

	// article error code 2000~
	ERROR_ART_NOT_EXIST = 2001

	// category
	ERROR_CATE_NAME_USED = 3001

	// upload
	ERROR_FILE_MAX_SIZE = 4001
	ERROR_FILE_FORMAT   = 4002
)

var codeMsg map[int]string = map[int]string{
	SUCCESS:                   "ok",
	ERROR:                     "服务器发生错误",
	ERROR_USERNAME_USED:       "用户名已存在,请更换",
	ERROR_ART_NOT_EXIST:       "获取文章信息错误",
	ERROR_CATE_NAME_USED:      "分类名已被使用",
	ERROR_FILE_MAX_SIZE:       "文件内存需要小于8M",
	ERROR_FILE_FORMAT:         "文件类型错误",
	ERROR_USER_NOT_FOUND:      "用户不存在",
	ERROR_USER_PASSWORD_WRONG: "用户密码错误",
}

func GetErrorMsg(code int) string {
	return codeMsg[code]
}
