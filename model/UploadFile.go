package model

import (
	"go_blog/utils/errmsg"
	"strings"
)

func UploadFileService(extStr string, fileSize int64) int {
	// l := len(fileName) - 4 //简单截取扩展名
	name := []string{".png", ".jpg", ".jpeg", ".gif"}
	// fileFormat := string([]rune(fileName)[l:]) //
	if fileSize > 8*1024*1024 {
		return errmsg.ERROR_FILE_MAX_SIZE
	}
	if stringContains(name, strings.ToLower(extStr)) {
		return errmsg.SUCCESS
	} else {
		return errmsg.ERROR_FILE_FORMAT
	}
}

// 字符串截取
func Str(str string, length int) string {
	var i, n int
	for i = range str {
		if length == n {
			break
		}
		n++
	}
	return str[:i]
}

func stringContains(array []string, val string) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			return true
		}
	}
	return false
}
