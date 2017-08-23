package main

import (
	"errors"
	"github.com/devfeel/dotlog"
)

var log1 dotlog.Logger

func main() {
	dotlog.StartLogService("d:/gotmp/golog/log.conf")
	log1 = dotlog.GetLogger("FileLogger")
	log1.Info("example-normal test main")
	log1.Error(errors.New("example error"), "normal error")
	for true {
	}
}
