package core

import (
	"fiber/global"
)

func Tips() {
}

func StartupMessage(addr string) {
	global.SLog.Infof("HTTP Server listening at 0.0.0.0%v", addr)
}
