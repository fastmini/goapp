// Package xxljob
// @Description:
// @Author AN 2023-12-06 23:20:19
package xxljob

import (
	"fiber/app/task"
	"fiber/config"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	xxl "github.com/xxl-job/xxl-job-executor-go"
	"log"
)

type logger struct{}

func (l *logger) Info(format string, a ...interface{}) {
	// log.Println(fmt.Sprintf("INFO日志 - "+format, a...))
}

func (l *logger) Error(format string, a ...interface{}) {
	log.Println(fmt.Sprintf("ERROR日志 - "+format, a...))
	// panic(businessError.New(businessError.SERVER_ERROR, fmt.Sprintf("ERROR日志 - "+format, a...)))
}

// ConnectXxlJob 连接xxlJob
func ConnectXxlJob(app *fiber.App) {
	if config.ConfigWithBool("XXL_ENABLE") == false {
		return
	}
	exec := xxl.NewExecutor(
		xxl.ServerAddr(fmt.Sprintf("http://%s/xxl-job-admin", config.Config("XXL_JOB_ADDRESS"))),
		xxl.AccessToken(""), // 请求令牌(默认为空)
		// xxl.ExecutorIp("127.0.0.1"), //可自动获取
		xxl.ExecutorPort(config.Config("PORT")),                                 // 默认9999（非必填）
		xxl.RegistryKey(config.Config("APP_ENV")+"-"+config.Config("APP_NAME")), // 执行器名称
		xxl.SetLogger(&logger{}),                                                // 自定义日志
	)
	exec.Init()
	defer exec.Stop()
	mux(app, exec)
	// 设置日志查看handler
	exec.LogHandler(func(req *xxl.LogReq) *xxl.LogRes {
		return &xxl.LogRes{Code: 200, Msg: "", Content: xxl.LogResContent{
			FromLineNum: req.FromLineNum,
			ToLineNum:   2,
			LogContent:  "这个是自定义日志handler",
			IsEnd:       true,
		}}
	})
	// ES检查任务handler
	exec.RegTask("testJobHandler", task.TestJobHandler)
}

func mux(app *fiber.App, exec xxl.Executor) {
	app.Post("/run", adaptor.HTTPHandlerFunc(exec.RunTask))
	app.Post("/kill", adaptor.HTTPHandlerFunc(exec.KillTask))
	app.Post("/log", adaptor.HTTPHandlerFunc(exec.TaskLog))
	app.Post("/beat", adaptor.HTTPHandlerFunc(exec.Beat))
	app.Post("/idleBeat", adaptor.HTTPHandlerFunc(exec.IdleBeat))
}
