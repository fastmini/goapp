package userApi

import (
	"fiber/global"
	"fiber/resultVo"
	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	var count int64
	global.DB.Table("system_tenant").Limit(1).Count(&count)
	global.BLog.Infof("count数量：%d", count)
	return c.JSON(resultVo.Success("success", c))
}
