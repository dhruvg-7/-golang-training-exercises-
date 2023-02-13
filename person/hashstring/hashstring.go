package hashstring

import (
	"crypto/md5"
	"fmt"
)

func Md5(s string) string {
	res := md5.Sum([]byte(s))
	return fmt.Sprintf("%x", res)
}
