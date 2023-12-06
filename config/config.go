// Package config
// @Description:
// @Author AN 2023-12-06 23:21:25
package config

import (
	"os"
	"strconv"
)

func Config(key string) string {
	return os.Getenv(key)
}

func ConfigWithBool(key string) bool {
	b, _ := strconv.ParseBool(os.Getenv(key))
	return b
}
