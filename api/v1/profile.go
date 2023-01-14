package v1

import (
	"go_blog/model"
	"go_blog/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetProfileService(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrorMsg(code),
	})
}

func UpdateProfile(c *gin.Context) {
	var data model.Profile
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code := model.UpdateProfileService(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrorMsg(code),
	})
}

func CreatProfile(c *gin.Context) {
	var profile model.Profile
	_ = c.ShouldBindJSON(&profile)
	code := model.CreatProfileService(&profile)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrorMsg(code),
	})
}
