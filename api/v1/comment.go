package v1

import (
	"go_blog/model"
	"go_blog/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	var data model.Comment
	_ = c.ShouldBindJSON(&data)

	code := model.CreateCommentService(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrorMsg(code),
	})
}

func GetComment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetCommentService(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrorMsg(code),
	})
}

func DeleteComment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteCommentService(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrorMsg(code),
	})
}

func GetCommentCount(c *gin.Context) {
	article_id, _ := strconv.Atoi(c.Param("id"))
	total := model.GetCommentCountService(article_id)
	c.JSON(http.StatusOK, gin.H{
		"total": total,
	})
}

func CheckComment(c *gin.Context) {
	var data model.Comment
	_ = c.ShouldBindJSON(&data)
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.CheckCommentService(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrorMsg(code),
	})
}

func UnCheckComment(c *gin.Context) {
	var data model.Comment
	_ = c.ShouldBindJSON(&data)
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.CheckCommentService(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrorMsg(code),
	})
}

func GetCommentList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Param("pagesize"))
	pageNum, _ := strconv.Atoi(c.Param("pagenum"))
	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	if pageNum == 0 {
		pageNum = 1
	}
	data, total, code := model.GetCommentListService(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrorMsg(code),
	})
}

func GetCommentListFront(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	if pageNum == 0 {
		pageNum = 1
	}
	data, total, code := model.GetCommentListFrontService(id, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrorMsg(code),
	})
}
