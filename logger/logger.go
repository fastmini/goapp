/**
 * @Author: AF
 * @Date: 2021/8/10 12:03
 */

package logger

import (
	"fiber/global"
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

// Request 请求日志输出到文件
func Request(app *fiber.App) {
	logFilePath := "log/"
	logFileName := "server.log"
	fileName := path.Join(logFilePath, logFileName)
	writer, _ := rotatelogs.New(
		logFilePath+"server_%Y%m%d.log",
		rotatelogs.WithLinkName(fileName),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} ${locals:requestid} - ${latency} ${method} ${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Asia/Shanghai",
		Output:     writer,
	}))
}
func init() {
	logFilePath := "log/"
	logFileName := "app.log"
	// 日志文件
	fileName := path.Join(logFilePath, logFileName)
	_ = os.Mkdir(logFilePath, 0755)
	global.LogFile, _ = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
	logrus := log.New()
	logrus.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		ForceColors:     true,
	})
	logrus.Out = global.LogFile
	logrus.SetLevel(log.TraceLevel)
	logrus.SetReportCaller(true)
	writer, _ := rotatelogs.New(
		logFilePath+"app_%Y%m%d.log",
		rotatelogs.WithLinkName(fileName),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	logrus.SetOutput(io.MultiWriter(os.Stdout, writer))
	global.SLog = logrus
}

func New() fiber.Handler {
	return func(c *fiber.Ctx) error {
		reqId := c.Locals(requestid.ConfigDefault.ContextKey)
		global.BLog = global.SLog.WithFields(log.Fields{
			"requestId": reqId,
		})
		return c.Next()
	}
}
