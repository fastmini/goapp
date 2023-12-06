// Package database
// @Description:
// @Author AN 2023-12-06 23:19:05
package database

import (
	"fiber/config"
	"fiber/global"
	libLog "fiber/logger"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func ConnectDB() {
	if config.ConfigWithBool("DB_ENABLE") == false {
		return
	}
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s",
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_HOST"),
		config.Config("DB_PORT"),
		config.Config("DB_NAME"),
		"10s",
	)
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true, // 关闭默认事务
		PrepareStmt:            true, // 开启缓存预编译，可以提高后续的调用速度
		Logger: logger.New(libLog.SqlLogg(), logger.Config{
			SlowThreshold:             2 * time.Second,
			Colorful:                  false,
			IgnoreRecordNotFoundError: true,
			LogLevel:                  logger.Info,
		}),
	})
	if err != nil {
		global.SLog.Errorf("连接DB数据源失败, 地址: %v, 账号：%v ERR: %v", config.Config("DB_HOST"), config.Config("DB_USER"), err)
	}
	// 一个坑，不设置这个参数，gorm会把表名转义后加个s，导致找不到数据库的表
	SqlDB, _ := global.DB.DB()
	// 设置连接池中最大的闲置连接数
	SqlDB.SetMaxIdleConns(3)
	// 设置数据库的最大连接数量
	SqlDB.SetMaxOpenConns(5)
	// 这是连接的最大可复用时间
	SqlDB.SetConnMaxLifetime(10 * time.Second)
	global.SLog.Infof("连接DB数据源成功, 地址: %v, 账号：%v", config.Config("DB_HOST"), config.Config("DB_USER"))
}
