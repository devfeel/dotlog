package targets

import (
	"github.com/devfeel/dotlog/config"
	"github.com/devfeel/dotlog/const"
	"github.com/devfeel/dotlog/internal"
	"github.com/devfeel/dotlog/layout"
	"github.com/devfeel/dotlog/util/email"
	"strings"
)

type EMailTarget struct {
	BaseTarget

	MailServer   string
	MailNickName string
	MailAccount  string
	MailPassword string
	ToMail       string
	Subject      string
	logChan      chan string
}

func NewEMailTarget(conf *config.EMailTargetConfig) *EMailTarget {
	t := &EMailTarget{logChan: make(chan string, GetChanSize())}
	t.TargetType = _const.TargetType_EMail
	t.Name = conf.Name
	t.IsLog = conf.IsLog
	t.Encode = conf.Encode
	t.Layout = conf.Layout
	t.MailServer = layout.CompileLayout(conf.MailServer)
	t.MailNickName = layout.CompileLayout(conf.MailNickName)
	t.MailAccount = layout.CompileLayout(conf.MailAccount)
	t.MailPassword = layout.CompileLayout(conf.MailPassword)
	t.ToMail = layout.CompileLayout(conf.ToMail)
	t.Subject = layout.CompileLayout(conf.Subject)

	if t.MailNickName == ""{
		t.MailNickName = t.MailAccount
	}

	//启动异步写文件
	go t.handleLog()
	return t
}

//处理日志内部函数
func (t *EMailTarget) handleLog() {
	for {
		log := <-t.logChan
		t.writeTarget(log)
	}
}

func (t *EMailTarget) WriteLog(log string, useLayout string, level string) {
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

func (t *EMailTarget) writeTarget(log string) {
	mail := new(_email.MailConfig)
	mail.Host = t.MailServer
	mail.FromNickName = t.MailNickName
	mail.FromAccount = t.MailAccount
	mail.FromPassword = t.MailPassword
	mail.ToMail = t.ToMail
	mail.BodyType = "text"
	mail.Subject = t.Subject

	mail.Body = log
	if err := _email.SendEMail(mail); err != nil {
		internal.GlobalInnerLogger.Error(err, "EMailTarget:writeTarget error", log, t.MailServer, t.MailAccount)
	}
}
