package core

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
)

// 日志输出格式化
// 用于 websocket 返回日志格式化
//
type logFormatter struct{}

func (l *logFormatter) Format(entry *log.Entry) ([]byte, error) {
	output := fmt.Sprintf("%v\n", entry.Message)
	return []byte(output), nil
}
