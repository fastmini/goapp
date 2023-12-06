// Package resultVo
// @Description:
// @Author AN 2023-12-06 23:18:23
package resultVo

import (
	businessError "fiber/error"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"time"
)

type ResultVo struct {
	*businessError.Err
	TimeStamp int         `json:"timeStamp"`
	RequestId interface{} `json:"requestId"`
	Data      interface{} `json:"data"`
}

func Success(data interface{}, c *fiber.Ctx) ResultVo {
	return ResultVo{
		Err:       businessError.New(businessError.SUCCESS),
		TimeStamp: time.Now().Nanosecond(),
		Data:      data,
		RequestId: c.Locals(requestid.ConfigDefault.ContextKey),
	}
}

func Fail(error *businessError.Err, c *fiber.Ctx) ResultVo {
	return ResultVo{
		Err:       error,
		TimeStamp: time.Now().Nanosecond(),
		Data:      nil,
		RequestId: c.Locals(requestid.ConfigDefault.ContextKey),
	}
}
