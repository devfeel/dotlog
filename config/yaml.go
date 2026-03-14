package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// yamlConfig YAML 配置结构
type yamlConfig struct {
	Global    globalConfig     `yaml:"global"`
	Variables []variableConfig `yaml:"variables"`
	Targets   targetList       `yaml:"targets"`
	Loggers   []loggerConfig  `yaml:"loggers"`
}

type globalConfig struct {
	IsLog          bool   `yaml:"isLog"`
	ChanSize       int    `yaml:"chanSize"`
	InnerLogPath   string `yaml:"innerLogPath"`
	InnerLogEncode string `yaml:"innerLogEncode"`
}

type variableConfig struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

type targetList struct {
	File  []fileTargetConfig  `yaml:"file"`
	Udp   []udpTargetConfig   `yaml:"udp"`
	Http  []httpTargetConfig  `yaml:"http"`
	EMail []emailTargetConfig `yaml:"email"`
	Fmt   []fmtTargetConfig   `yaml:"fmt"`
}

type fileTargetConfig struct {
	Name        string `yaml:"name"`
	IsLog       bool   `yaml:"isLog"`
	Layout      string `yaml:"layout"`
	Encode      string `yaml:"encode"`
	FileMaxSize int64  `yaml:"fileMaxSize"`
	FileName    string `yaml:"fileName"`
}

type udpTargetConfig struct {
	Name     string `yaml:"name"`
	IsLog    bool   `yaml:"isLog"`
	Layout   string `yaml:"layout"`
	Encode   string `yaml:"encode"`
	RemoteIP string `yaml:"remoteIP"`
}

type httpTargetConfig struct {
	Name    string `yaml:"name"`
	IsLog   bool   `yaml:"isLog"`
	Layout  string `yaml:"layout"`
	Encode  string `yaml:"encode"`
	HttpUrl string `yaml:"httpUrl"`
}

type emailTargetConfig struct {
	Name         string `yaml:"name"`
	IsLog        bool   `yaml:"isLog"`
	Layout       string `yaml:"layout"`
	Encode       string `yaml:"encode"`
	MailServer   string `yaml:"mailServer"`
	MailAccount  string `yaml:"mailAccount"`
	MailNickName string `yaml:"mailNickName"`
	MailPassword string `yaml:"mailPassword"`
	ToMail       string `yaml:"toMail"`
	Subject      string `yaml:"subject"`
}

type fmtTargetConfig struct {
	Name   string `yaml:"name"`
	IsLog  bool   `yaml:"isLog"`
	Layout string `yaml:"layout"`
	Encode string `yaml:"encode"`
}

type loggerConfig struct {
	Name       string              `yaml:"name"`
	IsLog      bool                `yaml:"isLog"`
	Layout     string              `yaml:"layout"`
	ConfigMode string              `yaml:"configMode"`
	Levels     []loggerLevelConfig `yaml:"levels"`
}

type loggerLevelConfig struct {
	Level   string `yaml:"level"`
	Targets string `yaml:"targets"`
	IsLog   bool   `yaml:"isLog"`
}

// LoadYamlConfig loads configuration from YAML file
func LoadYamlConfig(configFile string) (*AppConfig, error) {
	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var yc yamlConfig
	if err := yaml.Unmarshal(data, &yc); err != nil {
		return nil, fmt.Errorf("failed to parse YAML config: %w", err)
	}

	// Convert YAML config to AppConfig
	appConfig := &AppConfig{
		Global: &GlobalConfig{
			IsLog:          yc.Global.IsLog,
			ChanSize:       yc.Global.ChanSize,
			InnerLogPath:   yc.Global.InnerLogPath,
			InnerLogEncode: yc.Global.InnerLogEncode,
		},
		Variables: make([]*VariableConfig, len(yc.Variables)),
		Loggers:   make([]*LoggerConfig, len(yc.Loggers)),
		Targets: &TargetList{
			FileTargets:  make([]*FileTargetConfig, len(yc.Targets.File)),
			UdpTargets:   make([]*UdpTargetConfig, len(yc.Targets.Udp)),
			HttpTargets:  make([]*HttpTargetConfig, len(yc.Targets.Http)),
			EMailTargets: make([]*EMailTargetConfig, len(yc.Targets.EMail)),
			FmtTargets:   make([]*FmtTargetConfig, len(yc.Targets.Fmt)),
		},
	}

	// Convert variables
	for i, v := range yc.Variables {
		appConfig.Variables[i] = &VariableConfig{Name: v.Name, Value: v.Value}
	}

	// Convert loggers
	for i, l := range yc.Loggers {
		levels := make([]*LoggerLevelConfig, len(l.Levels))
		for j, level := range l.Levels {
			levels[j] = &LoggerLevelConfig{
				Level:   level.Level,
				Targets: level.Targets,
				IsLog:   level.IsLog,
			}
		}
		appConfig.Loggers[i] = &LoggerConfig{
			Name:       l.Name,
			IsLog:      l.IsLog,
			Layout:     l.Layout,
			ConfigMode: l.ConfigMode,
			Levels:     levels,
		}
	}

	// Convert file targets
	for i, t := range yc.Targets.File {
		appConfig.Targets.FileTargets[i] = &FileTargetConfig{
			Name:        t.Name,
			IsLog:       t.IsLog,
			Layout:      t.Layout,
			Encode:      t.Encode,
			FileMaxSize: t.FileMaxSize,
			FileName:    t.FileName,
		}
	}

	// Convert UDP targets
	for i, t := range yc.Targets.Udp {
		appConfig.Targets.UdpTargets[i] = &UdpTargetConfig{
			Name:     t.Name,
			IsLog:    t.IsLog,
			Layout:   t.Layout,
			Encode:   t.Encode,
			RemoteIP: t.RemoteIP,
		}
	}

	// Convert HTTP targets
	for i, t := range yc.Targets.Http {
		appConfig.Targets.HttpTargets[i] = &HttpTargetConfig{
			Name:    t.Name,
			IsLog:   t.IsLog,
			Layout:  t.Layout,
			Encode:  t.Encode,
			HttpUrl: t.HttpUrl,
		}
	}

	// Convert Email targets
	for i, t := range yc.Targets.EMail {
		appConfig.Targets.EMailTargets[i] = &EMailTargetConfig{
			Name:         t.Name,
			IsLog:        t.IsLog,
			Layout:       t.Layout,
			Encode:       t.Encode,
			MailServer:   t.MailServer,
			MailAccount:  t.MailAccount,
			MailNickName: t.MailNickName,
			MailPassword: t.MailPassword,
			ToMail:       t.ToMail,
			Subject:      t.Subject,
		}
	}

	// Convert Fmt targets
	for i, t := range yc.Targets.Fmt {
		appConfig.Targets.FmtTargets[i] = &FmtTargetConfig{
			Name:   t.Name,
			IsLog:  t.IsLog,
			Layout: t.Layout,
			Encode: t.Encode,
		}
	}

	return appConfig, nil
}
