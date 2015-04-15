package web

import (
	"picar/web/assets"

	"github.com/labstack/echo"
)

// 通过 Http 获取内置的 Assets 资源
//
func assetsHandler(c *echo.Context) {
	resName := c.Param("name")

	res := assets.Open(resName)
	c.Response.ResponseWriter.Header().Set("Content-Type", res.MIME())
	c.Response.ResponseWriter.Write([]byte(res.ReadAll()))
}
