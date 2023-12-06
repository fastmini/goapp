// Package task
// @Description:
// @Author AN 2023-12-06 23:21:34
package task

import (
	"context"
	"encoding/json"
	"fiber/app/task"
	businessError "fiber/error"
	"fiber/resultVo"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/xxl-job/xxl-job-executor-go"
)

type CheckQo struct {
	ExecutorParams string
}

func TestCheck(c *fiber.Ctx) error {
	var a context.Context
	var qo CheckQo
	if err := c.BodyParser(&qo); err != nil {
		panic(businessError.New(businessError.BAD_REQUEST))
	}
	xxlParams := &xxl.RunReq{
		ExecutorParams: qo.ExecutorParams,
	}
	jsonStr, _ := json.Marshal(xxlParams)
	fmt.Println("xxlParams: " + string(jsonStr))
	task.TestJobHandler(a, xxlParams)
	return c.JSON(resultVo.Success("success", c))
}
