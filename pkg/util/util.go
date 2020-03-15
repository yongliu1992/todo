package util

import (
	. "github.com/yongliu1992/todo/config"
)

// Setup Initialize the util
func Setup() {
	jwtSecret = []byte(JwtSecret)
}
