package global

import (
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/gomodule/redigo/redis"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"os"
)

// Redis 默认redis连接池
var Redis redis.Conn

// DB 数据库
var DB *gorm.DB

// SLog 系统日志
var SLog *log.Logger

// BLog 系统日志
var BLog *log.Entry

// LogFile 日志文件
var LogFile *os.File

// ES ES客户端
var ES *elasticsearch.Client

type AuthUserPayload struct {
	UserId string
}

var AuthUser *AuthUserPayload

//func SetAuthUser(user *model.UserModel) {
//	AuthUser = &AuthUserPayload{UserId: user.Id}
//}
