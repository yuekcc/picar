package web

import (
	"fmt"
	"net/http"
	"picar/core"
	"picar/web/assets"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

type msg struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func UI(port string) {

	// here should parsing some commandline args

	// starting db

	// starting httpd
	e := echo.New()
	e.Use(mw.Logger)
	// parse buildin assets
	e.Get("/assets/:name", assetsHandler)

	// do picar
	e.Post("/call/", call)

	e.Get("/", hello)
	// run
	e.Run(port)
}

func hello(c *echo.Context) {
	c.HTMLString(http.StatusOK, assets.ReadFile("main.html"))
}

func call(c *echo.Context) {

	path := checkEmptyStr(c.Request.FormValue("path"), "./")
	prefix := checkEmptyStr(c.Request.FormValue("prefix"), "")
	noarchiving := checkBool(c.Request.FormValue("noarchiving"))
	debug := checkBool(c.Request.FormValue("debug"))

	picar := core.NewPicar(path, prefix, noarchiving, debug)
	err := picar.Parse()

	if err != nil {
		c.JSON(http.StatusInternalServerError, &msg{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("出错，请尝试调试模式\n%s", err),
		})
	}

	c.JSON(http.StatusOK, &msg{
		Code:    200,
		Message: "完成",
	})
}

func checkEmptyStr(in string, d string) string {
	if in == "" {
		return d
	} else {
		return in
	}
}

func checkBool(in string) bool {
	if in == "true" {
		return true
	} else {
		return false
	}
}
