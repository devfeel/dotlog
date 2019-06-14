package main

import (
	"errors"
	"github.com/devfeel/dotlog"
	"time"
)

func main() {
	dotlog.StartLogService("log.conf")
	log1 := dotlog.GetLogger("ClassicsLogger")
	log1.Trace("example-normal trace main")
	log1.Debug("example-normal debug main")
	log1.Info("example-normal info main")
	log1.Warn("example-normal warn main")
	log1.Error(errors.New("example-normal error main"), "example-normal error main")

	log2 := dotlog.GetLogger("log1")
	log2.Trace("example-normal trace main - log1")
	for {
		time.Sleep(time.Hour)
	}
}
