package task

import (
	"context"
	"fmt"
	"github.com/xxl-job/xxl-job-executor-go"
	"time"
)

func TestJobHandler(cxt context.Context, param *xxl.RunReq) (msg string) {
	t := time.Now()
	endTime := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, t.Location())
	startTime := endTime.AddDate(0, 0, -1)
	fmt.Println(startTime, endTime)
	fmt.Println("params:" + param.ExecutorParams)
	return "success"
}
