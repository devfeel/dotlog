package dotlog

import (
	"fmt"
	"testing"
)

func Test_StartLogService(t *testing.T) {
	err := StartLogService("d:/gotmp/golog/log.conf")
	fmt.Println(err)
}
