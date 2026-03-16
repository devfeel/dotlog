package targets

import (
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/devfeel/dotlog/const"
	"github.com/devfeel/dotlog/internal"
)

// Rotator 日志轮转接口
type Rotator interface {
	ShouldRotate(fileName string) (bool, error)
	Rotate(fileName string) error
	Clean() error
}

// SizeRotator 按大小轮转
type SizeRotator struct {
	MaxSize    int64
	MaxBackups int // 保留备份文件数量，0 表示不清理
	mu         sync.Mutex
}

// NewSizeRotator 创建大小轮转器
func NewSizeRotator(maxSizeKB int64, maxBackups int) *SizeRotator {
	return &SizeRotator{
		MaxSize:    maxSizeKB * 1024,
		MaxBackups: maxBackups,
	}
}

// ShouldRotate 检查是否需要轮转
func (r *SizeRotator) ShouldRotate(fileName string) (bool, error) {
	if r.MaxSize <= 0 {
		return false, nil
	}

	info, err := os.Stat(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	return info.Size() >= r.MaxSize, nil
}

// Rotate 执行轮转
func (r *SizeRotator) Rotate(fileName string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// 再次检查（防止并发）
	should, err := r.ShouldRotate(fileName)
	if err != nil || !should {
		return err
	}

	// 生成备份文件名
	backupName := fileName + "." + time.Now().Format(_const.DefaultNoSeparatorTimeLayout) + ".logbak"

	// 重命名当前文件
	err = os.Rename(fileName, backupName)
	if err != nil {
		internal.GlobalInnerLogger.Error(err, "SizeRotator.Rename error: ", fileName, " -> ", backupName)
		return err
	}

	// 清理旧备份
	if r.MaxBackups > 0 {
		r.cleanBackups(fileName)
	}

	return nil
}

// cleanBackups 清理超过保留数量的备份文件
func (r *SizeRotator) cleanBackups(fileName string) {
	dir := filepath.Dir(fileName)
	baseName := filepath.Base(fileName)

	entries, err := os.ReadDir(dir)
	if err != nil {
		return
	}

	var backups []os.FileInfo
	prefix := baseName + "."

	for _, entry := range entries {
		if !entry.IsDir() && strings.HasPrefix(entry.Name(), prefix) {
			if info, err := entry.Info(); err == nil {
				backups = append(backups, info)
			}
		}
	}

	// 按修改时间排序，删除旧的
	if len(backups) > r.MaxBackups {
		for i := 0; i < len(backups)-r.MaxBackups; i++ {
			os.Remove(filepath.Join(dir, backups[i].Name()))
		}
	}
}

// Clean 清理备份文件
func (r *SizeRotator) Clean() error {
	return nil
}

// TimeRotator 按时间轮转
type TimeRotator struct {
	Interval   time.Duration // 轮转间隔
	MaxBackups int           // 保留备份文件数量
	format     string       // 文件名时间格式
	mu         sync.Mutex
	lastRotate time.Time
}

// NewTimeRotator 创建时间轮转器
func NewTimeRotator(interval time.Duration, maxBackups int) *TimeRotator {
	return &TimeRotator{
		Interval:   interval,
		MaxBackups: maxBackups,
		format:     "2006-01-02",
		lastRotate: time.Now(),
	}
}

// ShouldRotate 检查是否需要轮转
func (r *TimeRotator) ShouldRotate(fileName string) (bool, error) {
	// 检查文件是否存在以及是否需要按时间轮转
	info, err := os.Stat(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			r.lastRotate = time.Now()
			return false, nil
		}
		return false, err
	}

	// 检查修改时间
	modTime := info.ModTime()
	now := time.Now()

	// 如果距离上次轮转超过间隔，且文件有内容
	if now.Sub(r.lastRotate) >= r.Interval && modTime.After(r.lastRotate) {
		return true, nil
	}

	return false, nil
}

// Rotate 执行轮转
func (r *TimeRotator) Rotate(fileName string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// 再次检查
	should, err := r.ShouldRotate(fileName)
	if err != nil || !should {
		return err
	}

	// 备份当前文件
	backupName := fileName + "." + r.lastRotate.Format(r.format) + ".logbak"
	err = os.Rename(fileName, backupName)
	if err != nil {
		internal.GlobalInnerLogger.Error(err, "TimeRotator.Rename error: ", fileName, " -> ", backupName)
		return err
	}

	r.lastRotate = time.Now()

	// 清理旧备份
	if r.MaxBackups > 0 {
		r.cleanBackups(fileName)
	}

	return nil
}

// cleanBackups 清理超过保留数量的备份文件
func (r *TimeRotator) cleanBackups(fileName string) {
	dir := filepath.Dir(fileName)
	baseName := filepath.Base(fileName)

	entries, err := os.ReadDir(dir)
	if err != nil {
		return
	}

	var backups []os.FileInfo
	prefix := baseName + "."

	for _, entry := range entries {
		if !entry.IsDir() && strings.HasPrefix(entry.Name(), prefix) {
			if info, err := entry.Info(); err == nil {
				backups = append(backups, info)
			}
		}
	}

	if len(backups) > r.MaxBackups {
		for i := 0; i < len(backups)-r.MaxBackups; i++ {
			os.Remove(filepath.Join(dir, backups[i].Name()))
		}
	}
}

// Clean 清理备份文件
func (r *TimeRotator) Clean() error {
	return nil
}
