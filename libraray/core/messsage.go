// Package core
// @Description:
// @Author AN 2023-12-06 23:18:59
package core

import (
	"fiber/global"
)

func Tips() {
}

func StartupMessage(addr string) {
	global.SLog.Infof("HTTP Server listening at 0.0.0.0%v", addr)
}
