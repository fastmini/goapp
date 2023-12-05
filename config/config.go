/**
 * @Author: AF
 * @Date: 2021/8/9 14:22
 */

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
