/**
 * @Author: AF
 * @Date: 2021/8/10 12:03
 */

package logger

import (
	"fiber/config"
	"fiber/global"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

var requestWriter io.Writer
var sqlLogger *log.Logger
var appLogger *log.Logger

type MyWriter struct {
	mlog *log.Logger
}

func (m *MyWriter) Printf(format string, v ...interface{}) {
	logstr := fmt.Sprintf(format, v...)
	// 利用loggus记录日志
	m.mlog.Info(logstr)
}

// Request 请求日志输出到文件
func Request(app *fiber.App) {
	app.Use(logger.New(logger.Config{
		Format:        "INFO[${time}] ${status} ${locals:requestid} - ${latency} ${method} ${path}\n",
		TimeFormat:    "2006-01-02 15:04:05",
		TimeZone:      "Asia/Shanghai",
		DisableColors: false,
		Output: &lumberjack.Logger{
			Filename:   "./log/server.log",
			MaxSize:    500, // megabytes
			MaxBackups: 1,
			MaxAge:     30,   // days
			Compress:   true, // disabled by default
		},
	}))
}

func SqlLogg() *MyWriter {
	if isProd() {
		sqlLogger.SetFormatter(&log.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			PrettyPrint:     false,
		})
		sqlLogger.SetOutput(&lumberjack.Logger{
			Filename:   "./log/sql.log",
			MaxSize:    500, // megabytes
			MaxBackups: 1,
			MaxAge:     30,   // days
			Compress:   true, // disabled by default
		})
	} else {
		sqlLogger.SetFormatter(&log.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			ForceColors:     true,
			FullTimestamp:   true,
		})
	}
	return &MyWriter{mlog: sqlLogger}
}

func init() {
}

func init() {
	// 初始化SQL日志文件
	sqlLogger = log.New()
	sqlLogger.SetLevel(log.TraceLevel)
	sqlLogger.SetReportCaller(false)
}

func init() {
	// 初始化SLog日志
	global.SLog = log.New()
	global.SLog.SetLevel(log.TraceLevel)
	global.SLog.SetReportCaller(false)
	global.SLog.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		ForceColors:     true,
		FullTimestamp:   true,
	})
	global.SLog.SetOutput(os.Stdout)
}

func init() {
	// 初始化业务日志
	appLogger = log.New()
	appLogger.SetLevel(log.TraceLevel)
	appLogger.SetReportCaller(false)
	appLogger.SetOutput(&lumberjack.Logger{
		Filename:   "./log/app.log",
		MaxSize:    500, // megabytes
		MaxBackups: 1,
		MaxAge:     30,   // days
		Compress:   true, // disabled by default
	})
}

func isProd() bool {
	return config.Config("APP_ENV") == "prod" || config.Config("APP_ENV") == "gray"
}

func New() fiber.Handler {
	return func(c *fiber.Ctx) error {
		reqId := c.Locals(requestid.ConfigDefault.ContextKey)
		if isProd() {
			// 生产环境为了高性能，终端不输出日志
			appLogger.SetFormatter(&log.JSONFormatter{
				TimestampFormat:  "2006-01-02 15:04:05",
				PrettyPrint:      false,
				DisableTimestamp: false,
			})
		} else {
			// 非线上环境无所谓，慢点就慢点吧
			appLogger.SetFormatter(&log.TextFormatter{
				TimestampFormat: "2006-01-02 15:04:05",
				ForceColors:     true,
				FullTimestamp:   true,
			})
			appLogger.SetOutput(io.MultiWriter(os.Stdout, &lumberjack.Logger{
				Filename:   "./log/app.log",
				MaxSize:    500, // megabytes
				MaxBackups: 1,
				MaxAge:     30,   // days
				Compress:   true, // disabled by default
			}))
		}
		global.BLog = appLogger.WithFields(log.Fields{
			"requestId": reqId,
			"ip":        c.IP(),
		})
		return c.Next()
	}
}
