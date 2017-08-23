package internal

import (
	"fmt"
	"github.com/devfeel/dotlog/const"
	"github.com/devfeel/dotlog/layout"
	"github.com/devfeel/dotlog/util/exception"
	"github.com/devfeel/dotlog/util/file"
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

const innerLayout = "{DateTime} {LogLevel} {message}"

var GlobalInnerLogger *InnerLogger

type InnerLogger struct {
	TraceFileName, DebugFileName, InfoFileName, WarnFileName, ErrorFileName string

	FilePath string
	Encode   string
	Layout   string
}

func InitInnerLogger(logPath, logEncode string) {
	l := &InnerLogger{}
	l.FilePath = logPath
	l.TraceFileName = logPath + "devfeel.dotlog.trace.log"
	l.DebugFileName = logPath + "devfeel.dotlog.debug.log"
	l.InfoFileName = logPath + "devfeel.dotlog.info.log"
	l.WarnFileName = logPath + "devfeel.dotlog.warn.log"
	l.ErrorFileName = logPath + "devfeel.dotlog.error.log"
	l.Encode = logEncode
	l.Layout = innerLayout
	GlobalInnerLogger = l
}

func (log *InnerLogger) Trace(content ...interface{}) *InnerLogger {
	return log.writeLog(nil, fmt.Sprint(content), _const.LogLevel_Trace, log.TraceFileName)
}

func (log *InnerLogger) Debug(content ...interface{}) *InnerLogger {
	return log.writeLog(nil, fmt.Sprint(content), _const.LogLevel_Debug, log.DebugFileName)
}

func (log *InnerLogger) Info(content ...interface{}) *InnerLogger {
	return log.writeLog(nil, fmt.Sprint(content), _const.LogLevel_Info, log.InfoFileName)
}

func (log *InnerLogger) Warn(content ...interface{}) *InnerLogger {
	return log.writeLog(nil, fmt.Sprint(content), _const.LogLevel_Warn, log.WarnFileName)
}

func (log *InnerLogger) Error(err error, content ...interface{}) *InnerLogger {
	return log.writeLog(err, fmt.Sprint(content), _const.LogLevel_Error, log.ErrorFileName)
}

func (log *InnerLogger) writeLog(err error, content string, level string, fileName string) *InnerLogger {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	if err != nil && strings.ToUpper(level) == _const.LogLevel_Error {
		content = exception.ConvertError(err) + "\r\n" + content
	}
	logContent := layout.CompileLayout(log.Layout)
	logContent = layout.ReplaceLogLevelLayout(logContent, level)
	if log.Layout != "" {
		logContent = strings.Replace(logContent, "{message}", content, -1)
		log.writeFile(fileName, logContent)
	}
	return log
}

func (log *InnerLogger) writeFile(fileName, content string) {
	pathDir := filepath.Dir(fileName)
	pathExists := _file.Exist(pathDir)
	if pathExists == false {
		//create path
		err := os.MkdirAll(pathDir, 0777)
		if err != nil {

			return
		}
	}

	var mode os.FileMode
	flag := syscall.O_RDWR | syscall.O_APPEND | syscall.O_CREAT
	mode = 0666
	logstr := content + "\r\n"
	file, err := os.OpenFile(fileName, flag, mode)
	defer file.Close()
	if err != nil {

		return
	}
	file.WriteString(logstr)
}
