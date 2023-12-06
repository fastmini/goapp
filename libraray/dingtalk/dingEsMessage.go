// Package dingtalk
// @Description:
// @Author AN 2023-12-06 23:19:13
package dingtalk

import (
	"fiber/config"
	"fiber/global"
	"fmt"
	"github.com/CodyGuo/dingtalk"
	"io/ioutil"
	"time"
)

func connectDingTalk() *dingtalk.DingTalk {
	webHook := "https://oapi.dingtalk.com/robot/send?access_token=" + config.Config("DINGDING_TOKEN")
	secret := config.Config("DINGDING_SECRET")
	if len(secret) == 0 {
		global.SLog.Error("缺少钉钉密钥，请检查配置")
	}
	return dingtalk.New(webHook, dingtalk.WithSecret(secret))
}

func SendEsError(message string, title string) {
	// 发送钉钉消息
	dingtalkInstance := connectDingTalk()
	date := time.Now().Format("2006-01-02")
	if len(title) == 0 {
		title = "ES集群每日数据检查异常"
	}
	markdownTitle := title + "⚠️"
	markdownText := fmt.Sprintf("# **<font color=#FF0000 face='黑体'>%s</font>**\n"+
		"+ **触发时间：** %v\n"+
		"%s", title, date, message)
	if err := dingtalkInstance.RobotSendMarkdown(markdownTitle, markdownText); err != nil {
		global.SLog.Error("发送钉钉失败: %v", err)
	}
	printResult(dingtalkInstance)
}

func SendEsSuccess(message string, title string) {
	// 发送钉钉消息
	dingtalkInstance := connectDingTalk()
	date := time.Now().Format("2006-01-02")
	if len(title) == 0 {
		title = "ES集群每日数据检查成功"
	}
	markdownTitle := title + "😄"
	markdownText := fmt.Sprintf("# **<font color=#53B809 face='黑体'>%s</font>**\n"+
		"+ **触发时间：** %v\n"+
		"+ **检查参数：** %v\n"+
		"+ **检查完成，数据一切正常。😄**", title, date, message)
	if err := dingtalkInstance.RobotSendMarkdown(markdownTitle, markdownText); err != nil {
		global.SLog.Infof("发送钉钉失败: %v", err)
	}
	printResult(dingtalkInstance)
}

func printResult(dt *dingtalk.DingTalk) {
	response, err := dt.GetResponse()
	if err != nil {
		global.SLog.Infof("发送钉钉失败->printResult: %v", err)
	}
	reqBody, err := response.Request.GetBody()
	if err != nil {
		global.SLog.Infof("发送钉钉失败->printResult: %v", err)
	}
	reqData, err := ioutil.ReadAll(reqBody)
	if err != nil {
		global.SLog.Infof("发送钉钉失败->printResult: %v", err)
	}
	global.SLog.Infof("发送消息成功->printResult, core: %s", reqData)
}
