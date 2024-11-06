package codeutil

import (
	"crypto/md5"
	"fmt"
)

func Md5Str(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
