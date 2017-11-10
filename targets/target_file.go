package targets

import (
	"github.com/devfeel/dotlog/config"
	"github.com/devfeel/dotlog/const"
	"github.com/devfeel/dotlog/internal"
	"github.com/devfeel/dotlog/layout"
	"github.com/devfeel/dotlog/util/file"
	"os"
	"path"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

type FileTarget struct {
	BaseTarget

	FileName     string
	FileMaxSize  int64 //日志文件最大字节数
	RealFileName string

	logChan chan string
}

func NewFileTarget(conf *config.FileTargetConfig) *FileTarget {
	t := &FileTarget{logChan: make(chan string, GetChanSize())}
	t.TargetType = _const.TargetType_File
	t.Name = conf.Name
	t.IsLog = conf.IsLog
	t.Encode = conf.Encode
	t.Layout = conf.Layout
	t.FileName = conf.FileName
	t.FileMaxSize = conf.FileMaxSize * 1024
	//启动异步写文件
	go t.handleLog()
	return t
}

//GetRealFileName get real filename with compile layout
func (t *FileTarget) getRealFileName() string {
	t.RealFileName = path.Clean(layout.CompileLayout(t.FileName))
	return t.RealFileName
}

//处理日志内部函数
func (t *FileTarget) handleLog() {
	for {
		log := <-t.logChan
		t.writeTarget(log)
	}
}

func (t *FileTarget) WriteLog(log string, useLayout string, level string) {
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

func (t *FileTarget) writeTarget(log string) {
	//TODO:如何规避每次都需要CompileLayout?
	fileName := t.getRealFileName()

	if fileInfo, err := os.Stat(fileName); err != nil {
		//ignore stat error, fixed for #1 bug: golog.writeTarget os.Stat error
		//internal.GlobalInnerLogger.Error(err, "golog.writeTarget os.Stat error")
	} else {
		if t.FileMaxSize > 0 {
			//如果设置了FileMaxSize，则进行判断
			if fileInfo.Size() >= t.FileMaxSize {
				//modify old filename
				modifyFileName := fileName + "." + time.Now().Format(_const.DefaultNoSeparatorTimeLayout) + ".logbak"
				err := os.Rename(fileName, modifyFileName)
				if err != nil {
					internal.GlobalInnerLogger.Error(err, "golog.writeTarget os.Rename(", fileName, ", ", modifyFileName, ") error")
				}
			}
		}
	}

	pathDir := filepath.Dir(fileName)
	pathExists := _file.Exist(pathDir)
	if pathExists == false {
		//create path
		err := os.MkdirAll(pathDir, 0777)
		if err != nil {
			internal.GlobalInnerLogger.Error(err, "golog.writeFile create path error")
			return
		}
	}

	var mode os.FileMode
	flag := syscall.O_RDWR | syscall.O_APPEND | syscall.O_CREAT
	mode = 0666
	logstr := log + "\r\n"
	file, err := os.OpenFile(fileName, flag, mode)
	defer file.Close()
	if err != nil {
		internal.GlobalInnerLogger.Error(err, "golog.writeFile OpenFile error")
		return
	}
	file.WriteString(logstr)
}
