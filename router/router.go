/**
 * @Author: AF
 * @Date: 2021/8/9 14:53
 */

package router

import (
	taskApi "fiber/app/api/task"
	userApi "fiber/app/api/user"
	businessError "fiber/error"
	"fiber/global"
	"fiber/middleware"
	"fiber/resultVo"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"time"
)

func AppRouter(app *fiber.App) {
	app.Get("/metrics", monitor.New(monitor.Config{Title: "GoApp Monitor", Refresh: 2 * time.Second}))
	app.Get("/_startup", func(ctx *fiber.Ctx) error {
		global.SLog.Infof("hello")
		return ctx.JSON(resultVo.Success("ok", ctx), fiber.MIMEApplicationJSONCharsetUTF8)
	})
	app.Get("/_healthz", func(ctx *fiber.Ctx) error {
		return ctx.JSON(resultVo.Success("success", ctx), fiber.MIMEApplicationJSONCharsetUTF8)
	})
	app.Post("/task/testCheck", taskApi.TestCheck)
	app.Get("/user/count", userApi.GetUser)
	// 需要登录鉴权的路由
	apiRoute := app.Group("", middleware.AuthMiddleware())
	apiRoute.Get("/userInfo", func(ctx *fiber.Ctx) error {
		return ctx.JSON(resultVo.Success(nil, ctx), fiber.MIMEApplicationJSONCharsetUTF8)
	})
	// 其他
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(resultVo.Success(nil, ctx), fiber.MIMEApplicationJSONCharsetUTF8)
	})
	// 404返回
	app.Use(func(c *fiber.Ctx) error {
		return c.JSON(resultVo.Fail(businessError.New(businessError.NOT_FOUND), c), fiber.MIMEApplicationJSONCharsetUTF8)
	})
	// 这个后面不要写
}
