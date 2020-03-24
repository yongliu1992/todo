package util

import (
	tC "github.com/yongliu1992/todo/config"
)

// Setup Initialize the util
func Setup() {
	jwtSecret = []byte(tC.JwtSecret)
}
