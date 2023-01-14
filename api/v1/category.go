package v1

import (
	"go_blog/model"
	"go_blog/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	var category model.Category
	_ = c.ShouldBindJSON(&category)
	code := model.CheckCategory(category.Name)
	if code == errmsg.SUCCESS {
		create_code := model.CreateCategoryService(&category)
		c.JSON(http.StatusOK, gin.H{
			"status":  create_code,
			"message": errmsg.GetErrorMsg(create_code),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrorMsg(code),
		})
	}
}

func GetCategory(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageNum = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}
	data, total := model.GetCategoryService(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  errmsg.SUCCESS,
		"message": errmsg.GetErrorMsg(errmsg.SUCCESS),
		"data":    data,
		"total":   total,
	})
}

func GetCategoryInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetCategoryInfoService(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrorMsg(code),
	})
}

func EditCategory(c *gin.Context) {
	var data model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code := model.CheckCategory(data.Name)
	if code == errmsg.SUCCESS {
		editCode := model.EditCategoryService(id, &data)
		c.JSON(http.StatusOK, gin.H{
			"status":  editCode,
			"message": errmsg.GetErrorMsg(editCode),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrorMsg(code),
		})
	}
	if code == errmsg.ERROR_CATE_NAME_USED {
		c.Abort()
	}
}

func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteCategoryService(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrorMsg(code),
	})
}
