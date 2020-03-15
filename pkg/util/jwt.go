package util

import (
	"fmt"
	jwt  "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/yongliu1992/todo/pkg/e"
	"strings"
	"time"
)

var jwtSecret []byte

type Claims struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	RealName  string `json:"realName"`
	jwt.StandardClaims
}

// GenerateToken generate tokens used for auth
func GenerateToken(id int, username, password string, organId, chatUid int, realName string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(12 * time.Hour)
	claims := Claims{
		id,
		username,
		EncodeMD5(password),
		realName,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-customer",
		},
	}

	fmt.Println("jwtSecret", jwtSecret)
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken parsing token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}


// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS

		Authorization := c.GetHeader("Authorization")
		token := strings.Split(Authorization, " ")

		if Authorization == "" {
			code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
		} else {
			userInfo, err := ParseToken(token[1])
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}
			c.Set("userInfo", userInfo)
		}

		if code != e.SUCCESS {
			c.JSON(200, gin.H{
				"code":    code,
				"error":   e.GetMsg(code),
				"data":    data,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}