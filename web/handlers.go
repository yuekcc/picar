package web

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"picar/core"
	"picar/web/assets"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

// 通过 http 获取内置的 Assets 资源
//
func assetsHandler(c *echo.Context) {
	resName := c.Param("name")

	res := assets.Open(resName)
	c.Response.ResponseWriter.Header().Set("Content-Type", res.MIME())
	c.Response.ResponseWriter.Write([]byte(res.ReadAll()))
}

// 显示首页
//
func hello(c *echo.Context) {
	c.HTMLString(http.StatusOK, assets.ReadFile("main.html"))
}

// 调用 picar
//
func call(c *echo.Context) {

	// 处理 picar 参数
	path := emptyStr(c.Request.FormValue("path"), "./")
	prefix := emptyStr(c.Request.FormValue("prefix"), "")
	noarchiving, _ := strconv.ParseBool(c.Request.FormValue("noarchiving"))
	debug, _ := strconv.ParseBool(c.Request.FormValue("debug"))

	// 调用 picar
	picar := core.NewPicar(path, prefix, noarchiving, debug)

	logfile, _ := os.Create("picar.log")
	defer logfile.Close()

	picar.SetOutput(logfile)
	err := picar.Parse()
	// 出错
	if err != nil {
		c.JSON(http.StatusInternalServerError, &msg{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("出错啦！请尝试「调试模式」\n%s", err),
		})
	}

	// 成功
	c.JSON(http.StatusOK, &msg{
		Code:    200,
		Message: "完成",
	})
}

func wsConn(c *echo.Context) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conn, _ := upgrader.Upgrade(c.Response.ResponseWriter, c.Request, nil)
	defer conn.Close()
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			continue
		}

		fmt.Printf("GOT -> %v, %s\n", messageType, p)

		// 将输出写入 websocket 中
		var buf bytes.Buffer
		tty := bufio.NewWriter(&buf)

		// 解析客户端发来的 JSON
		picarArgs := &Parms{}
		err = json.Unmarshal(p, &picarArgs)
		if err != nil {
			// 解析 JSON 出错，跳出这次循环
			conn.WriteMessage(messageType, []byte("参数解析出错。"))
			continue
		} else {
			picar := core.NewPicar(picarArgs.Path, picarArgs.Prefix, picarArgs.Noarchiving, picarArgs.Debug)

			// 将输出设置为 websocket
			picar.SetOutput(tty)

			err = picar.Parse()
			if err != nil {
				conn.WriteMessage(messageType, []byte(err.Error()))
				continue
			}
			tty.Flush()
		}

		err = conn.WriteMessage(messageType, []byte(buf.String()))
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
	}
}

type Parms struct {
	Path        string `json:"path"`
	Prefix      string `json:"prefix"`
	Noarchiving bool   `json:"noarchiving"`
	Debug       bool   `json:"debug"`
}
