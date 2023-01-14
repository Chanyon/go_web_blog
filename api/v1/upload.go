package v1

import (
	"go_blog/model"
	"go_blog/utils/errmsg"
	"net/http"
	"path"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// 简单处理文件上传
func UploadFile(c *gin.Context) {
	url := "/upload/"
	file, _ := c.FormFile("file")
	extString := path.Ext(file.Filename)     //扩展名
	fileNameInt := time.Now().Unix() + 99990 //时间戳
	fileNameStr := strconv.FormatInt(fileNameInt, 10)
	newFileName := fileNameStr + extString //新文件名
	dst := "./upload/" + newFileName
	fileSize := file.Size
	code := model.UploadFileService(extString, fileSize)
	if code == 200 {
		c.SaveUploadedFile(file, dst)
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrorMsg(code),
			"url":     url + newFileName,
		})
	} else if code != 200 {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrorMsg(code),
			"url":     "",
		})
	}
}
