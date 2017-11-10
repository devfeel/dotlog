package targets

import (
	"github.com/devfeel/dotlog/config"
	"github.com/devfeel/dotlog/const"
	"github.com/devfeel/dotlog/internal"
	"github.com/devfeel/dotlog/layout"
	"github.com/devfeel/dotlog/util/http"
	"strings"
)

type HttpTarget struct {
	BaseTarget

	HttpUrl string
	logChan chan string
}

func NewHttpTarget(conf *config.HttpTargetConfig) *HttpTarget {
	t := &HttpTarget{logChan: make(chan string, GetChanSize())}
	t.TargetType = _const.TargetType_Http
	t.Name = conf.Name
	t.IsLog = conf.IsLog
	t.Encode = conf.Encode
	t.Layout = conf.Layout
	t.HttpUrl = conf.HttpUrl
	//启动异步写文件
	go t.handleLog()
	return t
}

//处理日志内部函数
func (t *HttpTarget) handleLog() {
	for {
		log := <-t.logChan
		t.writeTarget(log)
	}
}

func (t *HttpTarget) WriteLog(log string, useLayout string, level string) {
	if t.IsLog {
		if t.Layout != "" {
			useLayout = t.Layout
		}
		logContent := layout.CompileLayout(useLayout)
		logContent = layout.ReplaceLogLevelLayout(logContent, level)
		if useLayout != "" {
			logContent = strings.Replace(logContent, "{message}", log, -1)
			t.logChan <- logContent
		}
	}
}

func (t *HttpTarget) writeTarget(log string) {
	err := _http.HttpPost(t.HttpUrl, log, "")
	if err != nil {
		internal.GlobalInnerLogger.Error(err, "HttpTarget:WriteLog error", log, t.HttpUrl)
	}
}
