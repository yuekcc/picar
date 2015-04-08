package web

import (
	"net/http"

	"github.com/labstack/echo"
)

type File struct {
	MIME    string
	Content string
}

// css、javascript 一类的资源文件
var ASSETS = make(map[string]File)

// 模板
var TEMPLATES = make(map[string]string)

func init() {
	ASSETS["style.css"] = File{
		MIME:    "text/css",
		Content: assets_style_css,
	}

	ASSETS["normalize.css"] = File{
		MIME:    "text/css",
		Content: assets_normalize_css,
	}

	ASSETS["app.js"] = File{
		MIME:    "text/javascript",
		Content: assets_app_js,
	}

	ASSETS["jquery2.min.js"] = File{
		MIME:    "text/javascript",
		Content: assets_jquery2_min_js,
	}
	/*
		ASSETS["jbone.min.js"] = File{
			MIME:    "text/javascript",
			Content: assets_jbone_mini_js,
		}

		ASSETS["reqwest.min.js"] = File{
			MIME:    "text/javascript",
			Content: assets_reqwest_min_js,
		}
	*/

	TEMPLATES["view_main"] = view_main
}

// 通过 Http 获取内置的 Assets 资源
//
func assetsHandler(c *echo.Context) {
	resName := c.Param("name")

	res, ok := ASSETS[resName]
	if !ok {
		c.String(http.StatusNotFound, "Not Found")
	}
	c.Response.ResponseWriter.Header().Set("Content-Type", res.MIME)
	c.Response.ResponseWriter.Write([]byte(res.Content))
}
