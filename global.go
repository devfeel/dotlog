package dotlog

import (
	"fmt"
	"github.com/devfeel/dotlog/config"
	"github.com/devfeel/dotlog/internal"
	"github.com/devfeel/dotlog/layout"
	"github.com/devfeel/dotlog/targets"
	"github.com/devfeel/dotlog/const"
)

var (
	GlobalTargetMap map[string]targets.Target
	GlobalVariable  *layout.Variable
	GlobalLoggerMap map[string]Logger
)

func init() {
	GlobalTargetMap = make(map[string]targets.Target)
	GlobalVariable = layout.GetVariable()
	GlobalLoggerMap = make(map[string]Logger)
}

func StartLogService(configFile string) error {
	conf, err := config.InitConfig(configFile)
	if err != nil {
		fmt.Println(err)
		return err
	}

	config.GlobalAppConfig = conf

	//init innerlogger
	internal.InitInnerLogger(conf.Global.InnerLogPath, conf.Global.InnerLogEncode)

	internal.GlobalInnerLogger.Debug("*******************New Begin***********************")
	internal.GlobalInnerLogger.Debug("devfeel.dotlog ["+ _const.Version+"] InitConfig success")

	//init variable
	for _, v := range conf.Variables {
		GlobalVariable.RegisterUserVar(v.Name, v.Value)
	}
	internal.GlobalInnerLogger.Debug("RegisterUserVar success - total:", len(GlobalVariable.UserVar))

	//init file target
	var count int
	for _, v := range conf.Targets.FileTargets {
		GlobalTargetMap[v.Name] = targets.NewFileTarget(v)
		count++
	}
	internal.GlobalInnerLogger.Debug("InitFileTargets success - total:", count)

	//init udp target
	count = 0
	for _, v := range conf.Targets.UdpTargets {
		GlobalTargetMap[v.Name] = targets.NewUdpTarget(v)
		count++
	}
	internal.GlobalInnerLogger.Debug("InitUdpTargets success - total:", count)

	//init http target
	count = 0
	for _, v := range conf.Targets.HttpTargets {
		GlobalTargetMap[v.Name] = targets.NewHttpTarget(v)
		count++
	}
	internal.GlobalInnerLogger.Debug("InitHttpTargets success - total:", count)

	//init email target
	count = 0
	for _, v := range conf.Targets.EMailTargets {
		GlobalTargetMap[v.Name] = targets.NewEMailTarget(v)
		count++
	}
	internal.GlobalInnerLogger.Debug("InitEMailTargets success - total:", count)

	//init fmt target
	count = 0
	for _, v := range conf.Targets.FmtTargets {
		GlobalTargetMap[v.Name] = targets.NewFmtTarget(v)
		count++
	}
	internal.GlobalInnerLogger.Debug("InitFmtTargets success - total:", count)


	//init logger
	for _, v := range conf.Loggers {
		GlobalLoggerMap[v.Name] = NewLogger(v)
	}
	internal.GlobalInnerLogger.Debug("InitLogger success - total:", len(GlobalLoggerMap))

	return nil
}
