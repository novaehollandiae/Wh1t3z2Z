package main

import (
	"github.com/gin-gonic/gin"
)

var (
	app *gin.Engine
)

// CoreRequest is Receive t.cn/t1213121/apiv1/
type CoreRequest struct {
	Secret string `json:"secret"`
	Method string `json:"method"`
	Params string `json:"params"`

	Version string `json:"version"`
	Channel string `json:"channel"`
	/*
		32位的key    认证用
		操作id           前后端返回id，方便记录日志
		请求函数：
			// 系统管理类
			system.init
			system.nginx
			// 区服操作类
			server.create
			server.update
			server.hotupdate
			server.clear
			// 信息查询获取类
			get.channel
			get.node
			get.server
		渠道：
			操作的是哪一个渠道的服务器
		版本：

	*/
}

// CoreRequestHandler is xxx
func CoreRequestHandler() {
	var crq CoreRequest

}

func main() {
	app = gin.Default()
	app.POST("/t1213121/apiv1/")
	app.Run(":9001")
}
