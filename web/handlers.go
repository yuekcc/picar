package web

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"picar/core"
	"picar/web/assets"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

// 通过 http 获取内置的 Assets 资源
//
func assetsHandler(c *echo.Context) {
	resName := c.Param("name")

	res := assets.Open(resName)
	c.Response().Header().Set("Content-Type", res.MIME())
	c.Response().Write([]byte(res.ReadAll()))
}

// 显示首页
//
func hello(c *echo.Context) {
	c.HTML(http.StatusOK, assets.ReadFile("main.html"))
}

// 建立 websocket 连接
// 通用 websocket 进行 RPC
//
func wsConn(c *echo.Context) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	// 将普通 HTTP-Handler 升级为 webscoket
	conn, _ := upgrader.Upgrade(c.Response().Writer(), c.Request(), nil)
	defer conn.Close()
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			continue
		}

		fmt.Printf("WS GET -> %v, %s\n", messageType, p)

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
