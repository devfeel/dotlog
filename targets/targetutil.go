package targets

import (
	"github.com/devfeel/dotlog/config"
	"github.com/devfeel/dotlog/const"
	"strings"
)

func GetDefaultFileTarget(name, level string) Target {
	conf := &config.FileTargetConfig{
		Name:     name + "_file_" + level,
		FileName: "{LogDateDir}/" + name + "_" + strings.ToLower(level) + ".log",
		IsLog:    true,
		Encode:   _const.DefaultEncode,
		Layout:   "{DateTime} {LogLevel} {message}",
	}
	return NewFileTarget(conf)
}

func GetDefaultEMailTarget(name, level string) Target {
	conf := &config.EMailTargetConfig{
		Name:         name + "_mail_" + level,
		IsLog:        true,
		Encode:       _const.DefaultEncode,
		Layout:       "{DateTime} {LogLevel} {message}",
		MailServer:   "{MailServer}",
		MailAccount:  "{MailAccount}",
		MailPassword: "{MailPassword}",
		ToMail:       "{ToMail}",
		Subject:      "{SysName} " + strings.ToLower(level) + "_Mail",
	}
	return NewEMailTarget(conf)
}
