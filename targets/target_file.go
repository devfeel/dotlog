package targets

import (
	"os"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/devfeel/dotlog/config"
	"github.com/devfeel/dotlog/const"
	"github.com/devfeel/dotlog/internal"
	"github.com/devfeel/dotlog/layout"
	"github.com/devfeel/dotlog/util/file"
)

type FileTarget struct {
	BaseTarget

	FileName     string
	FileMaxSize  int64 //日志文件最大字节数
	MaxBackups   int   // 保留备份文件数量
	RotateInterval time.Duration // 时间轮转间隔，0 表示不启用
	RealFileName string

	logChan chan string
	mu      sync.Mutex
	rotator Rotator
}

func NewFileTarget(conf *config.FileTargetConfig) *FileTarget {
	t := &FileTarget{
		logChan: make(chan string, GetChanSize()),
		mu:      sync.Mutex{},
	}
	t.TargetType = _const.TargetType_File
	t.Name = conf.Name
	t.IsLog = conf.IsLog
	t.Encode = conf.Encode
	t.Layout = conf.Layout
	t.FileName = conf.FileName
	t.FileMaxSize = conf.FileMaxSize * 1024
	t.MaxBackups = conf.MaxBackups

	// 初始化轮转器
	if t.FileMaxSize > 0 {
		t.rotator = NewSizeRotator(t.FileMaxSize/1024, t.MaxBackups)
	}

	//启动异步写文件
	go t.handleLog()
	return t
}

// GetRealFileName get real filename with compile layout
func (t *FileTarget) getRealFileName() string {
	t.RealFileName = filepath.Clean(layout.CompileLayout(t.FileName))
	return t.RealFileName
}

// 处理日志内部函数
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
	fileName := t.getRealFileName()

	// 先确保目录存在
	pathDir := filepath.Dir(fileName)
	pathExists := _file.Exist(pathDir)
	if !pathExists {
		err := os.MkdirAll(pathDir, 0777)
		if err != nil {
			internal.GlobalInnerLogger.Error(err, "golog.writeFile create path error")
			return
		}
	}

	// 检查并执行轮转
	if t.rotator != nil {
		should, err := t.rotator.ShouldRotate(fileName)
		if err == nil && should {
			t.rotator.Rotate(fileName)
		}
	}

	// 写入文件
	t.writeFile(fileName, log)
}

func (t *FileTarget) writeFile(fileName, content string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	var mode os.FileMode
	flag := syscall.O_RDWR | syscall.O_APPEND | syscall.O_CREAT
	mode = 0666
	logstr := content + "\r\n"

	file, err := os.OpenFile(fileName, flag, mode)
	if err != nil {
		internal.GlobalInnerLogger.Error(err, "golog.writeFile OpenFile error")
		return
	}
	defer file.Close()

	_, err = file.WriteString(logstr)
	if err != nil {
		internal.GlobalInnerLogger.Error(err, "golog.writeFile WriteString error")
	}
}
