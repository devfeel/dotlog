package main

import (
	"github.com/devfeel/dotlog"
	"github.com/pkg/errors"
)

func main() {
	dotlog.StartLogService("./log.conf")
	log1 := dotlog.GetLogger("FileLogger")
	log1.Trace("example-normal trace main")
	log1.Debug("example-normal debug main")
	log1.Info("example-normal info main")
	log1.Warn("example-normal warn main")
	log1.Error(errors.New("example-normal error main"), "example-normal error main")
	for true {

	}
}
