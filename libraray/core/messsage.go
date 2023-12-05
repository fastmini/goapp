package core

import (
	log "github.com/sirupsen/logrus"
)

func StartupMessage(addr string) {
	log.Infof("HTTP Server is listen on 0.0.0.0%v", addr)
}
