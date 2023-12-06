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
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"time"
)

var requestWriter io.Writer
var sqlWriter io.Writer
var appWriter io.Writer
var appLogrus *log.Logger

type MyWriter struct {
	mlog *log.Logger
}

func (m *MyWriter) Printf(format string, v ...interface{}) {
	logstr := fmt.Sprintf(format, v...)
	//利用loggus记录日志
	m.mlog.Info(logstr)
}

// Request 请求日志输出到文件
func Request(app *fiber.App) {
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} ${locals:requestid} - ${latency} ${method} ${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Asia/Shanghai",
		Output:     requestWriter,
	}))
}

func Logg() *MyWriter {
	logg := log.New()
	if isProd() {
		appLogrus.SetFormatter(&log.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			PrettyPrint:     false,
		})
		appLogrus.SetOutput(appWriter)
	} else {
		logg.SetFormatter(&log.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			ForceColors:     true,
			FullTimestamp:   true,
		})
	}
	logg.SetLevel(log.TraceLevel)
	logg.SetReportCaller(false)
	logg.SetOutput(sqlWriter)
	return &MyWriter{mlog: logg}
}

func init() {
	// 初始化Request日志文件
	reqLogFilePath := "log/"
	reqLogFileName := "server.log"
	reqFileName := path.Join(reqLogFilePath, reqLogFileName)
	reqWriter, _ := rotatelogs.New(
		reqLogFilePath+"server_%Y%m%d.log",
		rotatelogs.WithLinkName(reqFileName),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	requestWriter = reqWriter
}

func init() {
	// 初始化SQL日志文件
	reqLogFilePath := "log/"
	reqLogFileName := "sql.log"
	reqFileName := path.Join(reqLogFilePath, reqLogFileName)
	reqWriter, _ := rotatelogs.New(
		reqLogFilePath+"sql_%Y%m%d.log",
		rotatelogs.WithLinkName(reqFileName),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	sqlWriter = reqWriter
}

func init() {
	// 初始化SLog日志
	logrus := log.New()
	logrus.SetLevel(log.TraceLevel)
	logrus.SetReportCaller(false)
	logrus.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		ForceColors:     true,
		FullTimestamp:   true,
	})
	logrus.SetOutput(os.Stdout)
	global.SLog = logrus
}

func init() {
	// 初始化业务日志
	logFilePath := "log/"
	logFileName := "app.log"
	fileName := path.Join(logFilePath, logFileName)
	_ = os.Mkdir(logFilePath, 0755)
	global.LogFile, _ = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
	logrus := log.New()
	logrus.Out = global.LogFile
	logrus.SetLevel(log.TraceLevel)
	logrus.SetReportCaller(false)
	appWriter, _ = rotatelogs.New(
		logFilePath+"app_%Y%m%d.log",
		rotatelogs.WithLinkName(fileName),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	appLogrus = logrus
}

func isProd() bool {
	return config.Config("APP_ENV") == "prod" || config.Config("APP_ENV") == "gray"
}

func New() fiber.Handler {
	return func(c *fiber.Ctx) error {
		reqId := c.Locals(requestid.ConfigDefault.ContextKey)
		if isProd() {
			appLogrus.SetFormatter(&log.JSONFormatter{
				TimestampFormat:  "2006-01-02 15:04:05",
				PrettyPrint:      false,
				DisableTimestamp: false,
			})
			appLogrus.SetOutput(appWriter)
		} else {
			appLogrus.SetFormatter(&log.TextFormatter{
				TimestampFormat: "2006-01-02 15:04:05",
				ForceColors:     true,
				FullTimestamp:   true,
			})
			appLogrus.SetOutput(io.MultiWriter(os.Stdout, appWriter))
		}
		global.BLog = appLogrus.WithFields(log.Fields{
			"requestId": reqId,
			"ip":        c.IP(),
		})
		return c.Next()
	}
}
