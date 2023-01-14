package middleware

import (
	"errors"
	"go_blog/utils"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JWT struct {
	JwtKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(utils.JwtKey),
	}
}

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func (j *JWT) CreateToken(claims MyClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.JwtKey)
}

// 解析token
var (
	ErrTokenMalformed   error = errors.New("token不正确,请重新登录")
	ErrTokenExpired     error = errors.New("token过期,请重新登录")
	ErrTokenNotValidYet error = errors.New("token无效,请重新登录")
	ErrTokenInvalid     error = errors.New("token不正确,请重新登录")
)

func (j *JWT) ParserToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		return j.JwtKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ErrTokenNotValidYet
			} else {
				return nil, ErrTokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*MyClaims); ok {
			return claims, nil
		}
		return nil, ErrTokenInvalid
	}
	return nil, ErrTokenInvalid
}

func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			code = 401
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": "token无效",
			})
			c.Abort()
			return
		}
		checkToken := strings.Split(token, " ")
		if len(checkToken) == 0 || len(checkToken) != 2 || checkToken[0] != "Bearer" {
			code = 401
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": "token无效",
			})
			c.Abort()
			return
		}
		j := NewJWT()
		claims, err := j.ParserToken(checkToken[1])
		if err != nil {
			if err == ErrTokenExpired {
				code = 401
				c.JSON(http.StatusOK, gin.H{
					"status":  code,
					"message": "token过期,请重新登录",
					"data":    "",
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": "token无效",
				"data":    "",
			})
			c.Abort()
			return
		}
		c.Set("username", claims)
		c.Next()
	}
}
