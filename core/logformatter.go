package core

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
)

type logFormatter struct{}

func (l *logFormatter) Format(entry *log.Entry) ([]byte, error) {
	output := fmt.Sprintf("%v\n", entry.Message)
	return []byte(output), nil
}
