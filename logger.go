package dotlog

import (
	"fmt"
	"github.com/devfeel/dotlog/config"
	"github.com/devfeel/dotlog/const"
	"github.com/devfeel/dotlog/internal"
	"github.com/devfeel/dotlog/targets"
	"github.com/devfeel/dotlog/util/exception"
	"github.com/devfeel/dotlog/util/string"
	"strings"
)

type Logger interface {
	LoggerName() string
	IsLog() bool

	Trace(content interface{})
	TraceS(content ...interface{})
	TraceF(format string, content ...interface{})
	Debug(content interface{})
	DebugS(content ...interface{})
	DebugF(format string, content ...interface{})
	Info(content interface{})
	InfoS(content ...interface{})
	InfoF(format string, content ...interface{})
	Warn(content interface{})
	WarnS(content ...interface{})
	WarnF(format string, content ...interface{})
	Error(err error, content interface{})
	ErrorS(err error, content ...interface{})
	ErrorF(err error, format string, content ...interface{})
}

type (
	LoggerLevel struct {
		Level       string
		IsLog       bool
		Targets     []string
		TargetArray []targets.Target
	}

	logger struct {
		isTraceEnabled, isDebugEnabled, isInfoEnabled, isWarnEnabled, isErrorEnabled bool

		loggerName string
		isLog      bool
		layout     string
		configMode string

		loggerLevelMap map[string]*LoggerLevel
	}
)

//NewLogger create new *LoggerLevel with LoggerLevelConfig
func NewLoggerLevel(conf *config.LoggerLevelConfig) *LoggerLevel {
	l := &LoggerLevel{}
	l.Level = conf.Level
	l.IsLog = conf.IsLog
	l.Targets = strings.Split(conf.Targets, ",")

	//load Target interface array
	for _, tName := range l.Targets {
		if t, exists := GlobalTargetMap[tName]; exists {
			l.TargetArray = append(l.TargetArray, t)
		}
	}
	return l
}

//NewLogger create Empty *LoggerLevel with level
func EmptyLoggerLevel(level string) *LoggerLevel {
	l := &LoggerLevel{}
	l.Level = level
	l.IsLog = true
	return l
}

//AddTarget add target to loglevel
func (l *LoggerLevel) AddTarget(t targets.Target) {
	l.Targets = append(l.Targets, t.GetName())
	l.TargetArray = append(l.TargetArray, t)
}

//NewLogger create new *logger with LoggerConfig
func NewLogger(conf *config.LoggerConfig) *logger {
	log := &logger{
		isTraceEnabled: true,
		isDebugEnabled: true,
		isInfoEnabled:  true,
		isWarnEnabled:  true,
		isErrorEnabled: true,

		loggerName: conf.Name,
		isLog:      conf.IsLog,
		layout:     conf.Layout,
		configMode: conf.ConfigMode,

		loggerLevelMap: make(map[string]*LoggerLevel),
	}

	//init loglevel
	for _, l := range conf.Levels {
		log.loggerLevelMap[strings.ToUpper(l.Level)] = NewLoggerLevel(l)
	}

	//parse config-mode
	if conf.ConfigMode == _const.ConfigMode_Classics {
		log = updateClassicsLogger(log)
	}

	if conf.ConfigMode == _const.ConfigMode_File {
		log = updateFileLogger(log)
	}

	if conf.ConfigMode == _const.ConfigMode_Fmt {
		log = updateFmtLogger(log)
	}

	if conf.ConfigMode == _const.ConfigMode_FileFmt {
		log = updateFileFmtLogger(log)
	}

	return log
}

func updateClassicsLogger(logger *logger) *logger {
	targetName := logger.LoggerName()
	if strings.LastIndex(strings.ToLower(targetName), "logger") == (len(targetName) - 6) {
		targetName = _string.Substr(targetName, 0, len(targetName)-6)
	}

	logger.addLevelTarget(_const.LogLevel_Trace, targets.GetDefaultFileTarget(targetName, _const.LogLevel_Trace, logger.layout))
	logger.addLevelTarget(_const.LogLevel_Debug, targets.GetDefaultFileTarget(targetName, _const.LogLevel_Debug, logger.layout))
	logger.addLevelTarget(_const.LogLevel_Info, targets.GetDefaultFileTarget(targetName, _const.LogLevel_Info, logger.layout))
	logger.addLevelTarget(_const.LogLevel_Warn, targets.GetDefaultFileTarget(targetName, _const.LogLevel_Warn, logger.layout))
	logger.addLevelTarget(_const.LogLevel_Warn, targets.GetDefaultEMailTarget(targetName, _const.LogLevel_Warn, logger.layout))
	logger.addLevelTarget(_const.LogLevel_Error, targets.GetDefaultFileTarget(targetName, _const.LogLevel_Error, logger.layout))
	logger.addLevelTarget(_const.LogLevel_Error, targets.GetDefaultEMailTarget(targetName, _const.LogLevel_Error, logger.layout))
	return logger
}

func updateFileLogger(logger *logger) *logger {
	targetName := logger.LoggerName()
	if strings.LastIndex(strings.ToLower(targetName), "logger") == (len(targetName) - 6) {
		targetName = _string.Substr(targetName, 0, len(targetName)-6)
	}

	logger.addLevelTarget(_const.LogLevel_Trace, targets.GetDefaultFileTarget(targetName, _const.LogLevel_Trace, logger.layout))
	logger.addLevelTarget(_const.LogLevel_Debug, targets.GetDefaultFileTarget(targetName, _const.LogLevel_Debug, logger.layout))
	logger.addLevelTarget(_const.LogLevel_Info, targets.GetDefaultFileTarget(targetName, _const.LogLevel_Info, logger.layout))
	logger.addLevelTarget(_const.LogLevel_Warn, targets.GetDefaultFileTarget(targetName, _const.LogLevel_Warn, logger.layout))
	logger.addLevelTarget(_const.LogLevel_Error, targets.GetDefaultFileTarget(targetName, _const.LogLevel_Error, logger.layout))
	return logger
}

func updateFmtLogger(logger *logger) *logger {
	targetName := logger.LoggerName()
	if strings.LastIndex(strings.ToLower(targetName), "logger") == (len(targetName) - 6) {
		targetName = _string.Substr(targetName, 0, len(targetName)-6)
	}

	logger.addLevelTarget(_const.LogLevel_Trace, targets.GetDefaultFmtTarget(targetName, _const.LogLevel_Trace, logger.layout))
	logger.addLevelTarget(_const.LogLevel_Debug, targets.GetDefaultFmtTarget(targetName, _const.LogLevel_Debug, logger.layout))
	logger.addLevelTarget(_const.LogLevel_Info, targets.GetDefaultFmtTarget(targetName, _const.LogLevel_Info, logger.layout))
	logger.addLevelTarget(_const.LogLevel_Warn, targets.GetDefaultFmtTarget(targetName, _const.LogLevel_Warn, logger.layout))
	logger.addLevelTarget(_const.LogLevel_Error, targets.GetDefaultFmtTarget(targetName, _const.LogLevel_Error, logger.layout))
	return logger
}

func updateFileFmtLogger(logger *logger) *logger {
	targetName := logger.LoggerName()
	if strings.LastIndex(strings.ToLower(targetName), "logger") == (len(targetName) - 6) {
		targetName = _string.Substr(targetName, 0, len(targetName)-6)
	}
	logger.addLevelTarget(_const.LogLevel_Trace, targets.GetDefaultFileTarget(targetName, _const.LogLevel_Trace, logger.layout))
	logger.addLevelTarget(_const.LogLevel_Debug, targets.GetDefaultFileTarget(targetName, _const.LogLevel_Debug, logger.layout))
	logger.addLevelTarget(_const.LogLevel_Info, targets.GetDefaultFileTarget(targetName, _const.LogLevel_Info, logger.layout))
	logger.addLevelTarget(_const.LogLevel_Warn, targets.GetDefaultFileTarget(targetName, _const.LogLevel_Warn, logger.layout))
	logger.addLevelTarget(_const.LogLevel_Error, targets.GetDefaultFileTarget(targetName, _const.LogLevel_Error, logger.layout))

	logger.addLevelTarget(_const.LogLevel_Trace, targets.GetDefaultFmtTarget(targetName, _const.LogLevel_Trace, logger.layout))
	logger.addLevelTarget(_const.LogLevel_Debug, targets.GetDefaultFmtTarget(targetName, _const.LogLevel_Debug, logger.layout))
	logger.addLevelTarget(_const.LogLevel_Info, targets.GetDefaultFmtTarget(targetName, _const.LogLevel_Info, logger.layout))
	logger.addLevelTarget(_const.LogLevel_Warn, targets.GetDefaultFmtTarget(targetName, _const.LogLevel_Warn, logger.layout))
	logger.addLevelTarget(_const.LogLevel_Error, targets.GetDefaultFmtTarget(targetName, _const.LogLevel_Error, logger.layout))
	return logger
}

func (log *logger) addLevelTarget(level string, target targets.Target) {
	logLevel, exists := log.loggerLevelMap[level]
	if !exists {
		logLevel = EmptyLoggerLevel(level)
		log.loggerLevelMap[level] = logLevel
	}
	logLevel.AddTarget(target)
}

//EmptyLogger create new empty *logger
func EmptyLogger() *logger {
	return new(logger)
}

func (log *logger) getLoggerLevel(logLevel string) *LoggerLevel {
	level, _ := log.loggerLevelMap[logLevel]
	return level
}

//LoggerName get logger name
func (log *logger) LoggerName() string {
	return log.loggerName
}

//LoggerName get logger's is start log
func (log *logger) IsLog() bool {
	return log.isLog
}

func (log *logger) Trace(content interface{}) {
	log.writeLog(nil, fmt.Sprint(content), log.getLoggerLevel(_const.LogLevel_Trace))
}

func (log *logger) TraceS(content ...interface{}) {
	log.writeLog(nil, SprintSpacing(content...), log.getLoggerLevel(_const.LogLevel_Trace))
}

func (log *logger) TraceF(format string, content ...interface{}) {
	log.writeLog(nil, fmt.Sprintf(format, content...), log.getLoggerLevel(_const.LogLevel_Trace))
}

func (log *logger) Debug(content interface{}) {
	log.writeLog(nil, fmt.Sprint(content), log.getLoggerLevel(_const.LogLevel_Debug))
}

func (log *logger) DebugS(content ...interface{}) {
	log.writeLog(nil, SprintSpacing(content...), log.getLoggerLevel(_const.LogLevel_Debug))
}

func (log *logger) DebugF(format string, content ...interface{}) {
	log.writeLog(nil, fmt.Sprintf(format, content...), log.getLoggerLevel(_const.LogLevel_Debug))
}

func (log *logger) Info(content interface{}) {
	log.writeLog(nil, fmt.Sprint(content), log.getLoggerLevel(_const.LogLevel_Info))
}

func (log *logger) InfoS(content ...interface{}) {
	log.writeLog(nil, SprintSpacing(content...), log.getLoggerLevel(_const.LogLevel_Info))
}

func (log *logger) InfoF(format string, content ...interface{}) {
	log.writeLog(nil, fmt.Sprintf(format, content...), log.getLoggerLevel(_const.LogLevel_Info))
}

func (log *logger) Warn(content interface{}) {
	log.writeLog(nil, fmt.Sprint(content), log.getLoggerLevel(_const.LogLevel_Warn))
}

func (log *logger) WarnS(content ...interface{}) {
	log.writeLog(nil, SprintSpacing(content...), log.getLoggerLevel(_const.LogLevel_Warn))
}

func (log *logger) WarnF(format string, content ...interface{}) {
	log.writeLog(nil, fmt.Sprintf(format, content...), log.getLoggerLevel(_const.LogLevel_Warn))
}

func (log *logger) Error(err error, content interface{}) {
	log.writeLog(err, fmt.Sprint(content), log.getLoggerLevel(_const.LogLevel_Error))
}

func (log *logger) ErrorS(err error, content ...interface{}) {
	log.writeLog(err, SprintSpacing(content...), log.getLoggerLevel(_const.LogLevel_Error))
}

func (log *logger) ErrorF(err error, format string, content ...interface{}) {
	log.writeLog(err, fmt.Sprintf(format, content...), log.getLoggerLevel(_const.LogLevel_Error))
}

func (log *logger) writeLog(err error, content string, level *LoggerLevel) Logger {
	defer func() {
		if err := recover(); err != nil {
			internal.GlobalInnerLogger.Error(fmt.Errorf("%v", err), "Logger:writeLog error", log.LoggerName())
		}
	}()
	if log.isLog && level.IsLog {
		layout := log.layout
		if err != nil && strings.ToUpper(level.Level) == _const.LogLevel_Error {
			content = exception.ConvertError(err) + "\r\n" + content
		}
		for _, t := range level.TargetArray {
			t.WriteLog(content, layout, level.Level)
		}

	}
	return log
}
