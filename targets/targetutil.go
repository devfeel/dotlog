package targets

import (
	"github.com/devfeel/dotlog/config"
	"github.com/devfeel/dotlog/const"
	"strings"
)

// GetChanSize get chan size with config or const
func GetChanSize() int {
	if config.GlobalAppConfig.Global.ChanSize >= 0 {
		return config.GlobalAppConfig.Global.ChanSize
	}
	return _const.DefaultChanSize
}

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

func GetDefaultFmtTarget(name, level string) Target {
	conf := &config.FmtTargetConfig{
		Name:     name + "_fmt_" + level,
		IsLog:    true,
		Encode:   _const.DefaultEncode,
		Layout:   "{DateTime} {LogLevel} {message}",
	}
	return NewFmtTarget(conf)
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
