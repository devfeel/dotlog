package main

import (
	"github.com/devfeel/dotlog"
)

var log1 dotlog.Logger

func main() {
	dotlog.StartLogService("./log.conf")
	log1 = dotlog.GetLogger("log1")
	for i := 0; i < 100; i++ {
		log1.Info("example-normal test main")
	}
	//log1.Error(errors.New("example error"), "normal error")
	for true {

	}
}
