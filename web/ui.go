/*
模块：webui

实现 webui 相关的 httphandler，包括显示首页、调用 core

*/

package web

import (
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

type msg struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func UI(port string) {

	// starting httpd
	e := echo.New()
	e.Use(mw.Logger) // 加载 logger 中间件

	// 处理内置的资源
	e.Get("/assets/:name", assetsHandler)

	// 调用 picar
	e.Get("/ws", wsConn)

	// 首页
	e.Get("/", hello)

	// 运行服务器
	e.Run(port)
}
