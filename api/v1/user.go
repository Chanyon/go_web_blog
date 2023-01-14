package v1

import (
	"go_blog/model"
	"go_blog/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
	controller
*/
// add user
func AddUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	if len(data.Password) < 6 || len(data.Password) > 16 {
		c.JSON(http.StatusOK, gin.H{
			"status":  401,
			"message": "密码长度需要6-16位",
		})
		return
	}
	if len(data.Username) < 6 || len(data.Username) > 8 {
		c.JSON(http.StatusOK, gin.H{
			"status":  401,
			"message": "用户名长度需要6-8位",
		})
		return
	}
	code := model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		model.CreateUserService(&data)
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": "创建用户成功",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrorMsg(code),
		})
	}
}

// getUserinfo
func GetUserInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var maps = make(map[string]interface{})
	data, code := model.GetUserInfoService(id)
	maps["username"] = data.Username
	maps["role"] = data.Role
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    maps,
		"total":   1,
		"message": errmsg.GetErrorMsg(code),
	})
}

// GetUser list
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	username := c.Query("username")

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	if pageNum == 0 {
		pageNum = 1
	}

	data, total := model.GetUsersService(username, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  errmsg.SUCCESS,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrorMsg(errmsg.SUCCESS),
	})
}

// edit user
func EditUser(c *gin.Context) {
	var data model.User
	var code int
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	validate_code := model.CheckUpUser(id, data.Username)
	if validate_code == 200 {
		code = model.EditUserService(id, &data)
	} else {
		code = validate_code
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrorMsg(code),
	})
}

// change user password
func ChangeUserPassword(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	// 密码长度判断
	if len(data.Password) < 6 || len(data.Password) > 16 {
		c.JSON(http.StatusOK, gin.H{
			"status":  401,
			"message": "密码长度需要6-16位",
		})
	} else {
		// validate_code := model.CheckUpUser(id, data.Username)
		// if validate_code == 200 {
		code := model.ChangePasswordService(id, &data)
		if code == 200 {
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrorMsg(code),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrorMsg(code),
			})
		}
		// } else {
		// 	c.JSON(http.StatusOK, gin.H{
		// 		"status":  validate_code,
		// 		"message": errmsg.GetErrorMsg(validate_code),
		// 	})
		// }
	}
}

//delete user
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteUserService(id)
	if code == 200 {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrorMsg(code),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrorMsg(code),
		})
	}
}
