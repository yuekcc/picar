package web

import (
	"net/http"
	"picar/core"

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
	c.HTMLString(http.StatusOK, TEMPLATES["view_main"])
}

func call(c *echo.Context) {
	path := checkStr(c.Request.FormValue("path"), "./")
	prefix := checkStr(c.Request.FormValue("prefix"), "")
	noarchiving := checkBool(c.Request.FormValue("noarchiving"))
	debug := checkBool(c.Request.FormValue("debug"))

	picar := core.NewPicar(path, prefix, noarchiving, debug)
	err := picar.Parse()

	if err != nil {
		c.JSON(http.StatusOK, &msg{
			Code:    500,
			Message: "Got a error. please retry with debug mode.",
		})
	}

	c.JSON(http.StatusOK, &msg{
		Code:    200,
		Message: "DONE.",
	})
	/*
		// for debug -start
		type r struct {
			Path        string
			Prefix      string
			Noarchiving bool
			Debug       bool
		}
		c.JSON(http.StatusOK, &r{
			Path:        path,
			Prefix:      prefix,
			Noarchiving: noarchiving,
			Debug:       debug,
		})
		// for debug -end
	*/
}

func checkStr(in string, d string) string {
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
