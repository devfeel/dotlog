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

func Test_SprintSpacing(t *testing.T) {
	checkResult := "a 1 b true"
	toCheck := SprintSpacing("a", 1, "b", true)
	if checkResult == toCheck {
		t.Log("CHECK SUCCESS:", toCheck, checkResult)
	} else {
		t.Error("CHECK ERROR:", toCheck, checkResult)
	}
}
