package targets

import (
	"github.com/devfeel/dotlog/config"
	"github.com/devfeel/dotlog/const"
	"strings"
)

const defaultLayout = "{DateTime} {LogLevel} {message}"

// GetChanSize get chan size with config or const
func GetChanSize() int {
	if config.GlobalAppConfig.Global.ChanSize >= 0 {
		return config.GlobalAppConfig.Global.ChanSize
	}
	return _const.DefaultChanSize
}

func GetDefaultFileTarget(name, level, layout string) Target {
	if layout == ""{
		layout = defaultLayout
	}
	conf := &config.FileTargetConfig{
		Name:     name + "_file_" + level,
		FileName: "{LogDateDir}/" + name + "_" + strings.ToLower(level) + ".log",
		IsLog:    true,
		Encode:   _const.DefaultEncode,
		Layout:   layout,
	}
	return NewFileTarget(conf)
}

func GetDefaultFmtTarget(name, level, layout string) Target {
	if layout == ""{
		layout = defaultLayout
	}
	conf := &config.FmtTargetConfig{
		Name:     name + "_fmt_" + level,
		IsLog:    true,
		Encode:   _const.DefaultEncode,
		Layout:   layout,
	}
	return NewFmtTarget(conf)
}


func GetDefaultEMailTarget(name, level, layout string) Target {
	if layout == ""{
		layout = defaultLayout
	}
	conf := &config.EMailTargetConfig{
		Name:         name + "_mail_" + level,
		IsLog:        true,
		Encode:       _const.DefaultEncode,
		Layout:       layout,
		MailServer:   "{MailServer}",
		MailNickName: "{MailNickName}",
		MailAccount:  "{MailAccount}",
		MailPassword: "{MailPassword}",
		ToMail:       "{ToMail}",
		Subject:      "{SysName} " + strings.ToLower(level) + "_Mail",
	}
	return NewEMailTarget(conf)
}
