package utility

import (
	"crypto/md5"
	"fmt"
)

func Createmd5(s string) string {
	has := md5.Sum([]byte(s))
	md5pwd := fmt.Sprintf("%x", has) //将[]byte转成16进制
	fmt.Println(len(md5pwd))
	return md5pwd
}
