// Package utils
// @Description:
// @Author AN 2023-12-06 23:17:20
package utils

import (
	"crypto/sha1"
	"encoding/hex"
)

func SHA1(s string) string {
	o := sha1.New()
	o.Write([]byte(s))
	return hex.EncodeToString(o.Sum(nil))
}
