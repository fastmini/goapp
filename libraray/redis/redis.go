// Package redis
// @Description:
// @Author AN 2023-12-06 23:20:06
package redis

import (
	"fiber/config"
	"fiber/global"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"strconv"
	"time"
)

func InitRedis() {
	// 默认连接池
	if config.ConfigWithBool("REDIS_ENABLE") == false {
		return
	}
	global.Redis = defaultPool().Get()
}

func defaultPool() *redis.Pool {
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			db, _ := strconv.Atoi(config.Config("REDIS_DB"))
			conn, err := redis.Dial(
				"tcp",
				fmt.Sprintf("%s:%s", config.Config("REDIS_HOST"), config.Config("REDIS_PORT")),
				redis.DialReadTimeout(time.Second),
				redis.DialWriteTimeout(time.Second*2),
				redis.DialConnectTimeout(10*time.Second),
				redis.DialDatabase(db),
				redis.DialPassword(config.Config("REDIS_PASSWORD")),
			)
			if err != nil {
				global.SLog.Errorf("连接redis失败, ERR: %v", err)
			}
			global.SLog.Infof("连接redis成功, 地址: %v, DB: %v", config.Config("REDIS_HOST"), config.Config("REDIS_DB"))
			return conn, err
		},
		MaxIdle:     10,
		MaxActive:   1000,
		IdleTimeout: 60 * time.Second,
	}
}
