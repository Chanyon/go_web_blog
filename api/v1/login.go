package v1

import (
	"go_blog/middleware"
	"go_blog/model"
	"go_blog/utils/errmsg"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// 后台登录
func Login(c *gin.Context) {
	var formData model.User
	_ = c.ShouldBindJSON(&formData)
	var code int
	formData, code = model.CheckLogin(formData.Username, formData.Password)
	if code == errmsg.SUCCESS {
		setToken(c, formData)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    formData.Username,
			"id":      formData.ID,
			"message": errmsg.GetErrorMsg(code),
			"token":   "",
		})
	}
}

// front user login
func LoginFront(c *gin.Context) {
	var formData model.User
	_ = c.ShouldBindJSON(&formData)
	formData, errCode := model.CheckLoginFront(formData.Username, formData.Password)
	if errCode == errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  errCode,
			"data":    formData.Username,
			"id":      formData.ID,
			"message": errmsg.GetErrorMsg(errCode),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  errCode,
			"data":    formData.Username,
			"id":      -1,
			"message": errmsg.GetErrorMsg(errCode),
		})
	}
}

func setToken(c *gin.Context, user model.User) {
	j := middleware.NewJWT()
	claims := middleware.MyClaims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 100,
			ExpiresAt: time.Now().Unix() + 604800,
			Issuer:    "GinBlog",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  errmsg.ERROR,
			"message": errmsg.GetErrorMsg(errmsg.ERROR),
			"token":   token,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  errmsg.SUCCESS,
			"data":    user.Username,
			"id":      user.ID,
			"message": errmsg.GetErrorMsg(errmsg.SUCCESS),
			"token":   token,
		})
	}
}
