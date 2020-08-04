package main

import (
	"github.com/gkzy/gow"
	"github.com/gkzy/gow/lib/config"
)

func main() {
	r := gow.Default()
	// gow.GetAppConfig 会读取当前所使用配置文件中的基础配置
	// 再通过r.SetAppConfig设置
	// 如果 env 的 APP_RUN_MODE==prod 时，会读取 prod.app.conf文件
	// 可以使用：  env APP_RUN_MODE="prod" go run main.go 来实现(mac/linux)
	r.SetAppConfig(gow.GetAppConfig())
	r.GET("/", GetRedisConfig)

	r.Run()
}

// GetRedisConfig get current config file
//	curl -i  http://127.0.0.1:8080
// HTTP/1.1 200 OK
// Content-Type: application/json; charset=utf-8
// Date: Tue, 04 Aug 2020 10:12:26 GMT
// Content-Length: 138
//
// {
//  "code": 0,
//  "data": {
//    "db": 0,
//    "host": "192.168.0.197",
//    "password": "123456",
//    "port": 6379
//  },
//  "msg": "success"
// }
func GetRedisConfig(c *gow.Context) {
	host := config.GetString("redis::host")
	port := config.DefaultInt("redis::port", 6379)
	db := config.DefaultInt("redis::db", 0)
	password := config.GetString("redis::password")

	c.JSON(gow.H{
		"code": 0,
		"msg":  "success",
		"data": gow.H{
			"host":     host,
			"port":     port,
			"db":       db,
			"password": password,
		},
	})
}
