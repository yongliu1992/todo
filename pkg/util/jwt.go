package util

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/yongliu1992/todo/pkg/e"
	"strings"
	"time"
)

var jwtSecret []byte

// Claims data
type Claims struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	RealName string `json:"realName"`
	jwt.StandardClaims
}

// GenerateToken generate tokens used for auth
func GenerateToken(id int, username, password string, organID, chatUID int, realName string) (string, error) {
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

		code = e.Success

		Authorization := c.GetHeader("Authorization")
		token := strings.Split(Authorization, " ")

		if Authorization == "" {
			code = e.ErrorAuthCheckTokenFail
		} else {
			userInfo, err := ParseToken(token[1])
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.ErrorAuthCheckTokenTimeout
				default:
					code = e.ErrorAuthCheckTokenFail
				}
			}
			c.Set("userInfo", userInfo)
		}

		if code != e.Success {
			c.JSON(200, gin.H{
				"code":  code,
				"error": e.GetMsg(code),
				"data":  data,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
