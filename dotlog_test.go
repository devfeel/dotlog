package dotlog

import (
	"fmt"
	"testing"
)

func Test_GetLogger(t *testing.T) {
	err := StartLogService("d:/gotmp/golog/log.conf")
	fmt.Println(err)

	log1 := GetLogger("log1")
	fmt.Println(log1.LoggerName())
}
