// Package elastic
// @Description:
// @Author AN 2023-12-06 23:19:27
package elastic

import (
	"fiber/config"
	"fiber/global"
	"github.com/elastic/go-elasticsearch/v7"
)

func ConnectES() {
	if config.ConfigWithBool("ELASTIC_ENABLE") == false {
		return
	}
	var err error
	cfg := elasticsearch.Config{
		Addresses: []string{
			config.Config("ELASTIC_HOST"),
		},
		Username: config.Config("ELASTIC_USER"),
		Password: config.Config("ELASTIC_PASSWORD"),
	}
	global.ES, err = elasticsearch.NewClient(cfg)
	if err != nil {
		global.SLog.Errorf("Error creating the client: %s", err)
	}
	global.SLog.Infof("连接ES数据源成功, 地址: %v, 账号：%v", config.Config("ELASTIC_HOST"), config.Config("ELASTIC_USER"))
}
