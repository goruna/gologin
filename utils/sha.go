package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"crypto/md5"
	"fmt"
)

func SHA1(s string) string  {
	o := sha1.New()
	o.Write([]byte(s))
	return  hex.EncodeToString(o.Sum(nil))
}

func md5V1(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}
