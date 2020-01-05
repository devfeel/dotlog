package targets

import (
	"fmt"
	"github.com/devfeel/dotlog/config"
	"github.com/devfeel/dotlog/const"
	"github.com/devfeel/dotlog/layout"
	"os"
	"strings"
)

type FmtTarget struct {
	BaseTarget
}

func NewFmtTarget(conf *config.FmtTargetConfig) *FmtTarget {
	t := &FmtTarget{}
	t.TargetType = _const.TargetType_Fmt
	t.Name = conf.Name
	t.IsLog = conf.IsLog
	t.Encode = conf.Encode
	t.Layout = conf.Layout
	return t
}

func (t *FmtTarget) WriteLog(log string, useLayout string, level string) {
	if t.IsLog {
		if t.Layout != "" {
			useLayout = t.Layout
		}
		logContent := layout.CompileLayout(useLayout)
		logContent = layout.ReplaceLogLevelLayout(logContent, level)
		if useLayout != "" {
			logContent = strings.Replace(logContent, "{message}", log, -1)
			t.writeTarget(logContent, level)
		}
	}
}

func (t *FmtTarget) writeTarget(log string, level string) {
	if level == _const.LogLevel_Error {
		os.Stderr.WriteString(log)
	} else {
		fmt.Println(log)
	}
}
