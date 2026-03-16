package main

import (
	"fmt"
	"time"

	"github.com/devfeel/dotlog"
)

// 本示例展示日志轮转功能
// 运行方式: go run main.go
// 配置文件: rotation.yaml
func main() {
	// 启动日志服务（使用 YAML 配置）
	err := dotlog.StartLogService("rotation.yaml")
	if err != nil {
		fmt.Printf("Failed to start log service: %v\n", err)
		return
	}

	// 获取带轮转的 FileLogger
	fileLog := dotlog.GetLogger("FileLogger")

	// 获取 JSON Logger
	jsonLog := dotlog.GetLogger("JSONLogger")

	fmt.Println("开始写入日志（测试轮转功能）...")
	fmt.Println("每条日志后会打印当前文件大小，达到 FileMaxSize 后会自动轮转")
	fmt.Println("按 Ctrl+C 停止")

	// 写入大量日志来触发轮转
	count := 0
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		count++
		msg := fmt.Sprintf("Test log message #%d - %s", count, time.Now().Format("15:04:05.000"))

		// 写入文件日志
		fileLog.Info(msg)

		// 写入 JSON 日志
		jsonLog.Info(msg)

		// 每 100 条打印一次
		if count%100 == 0 {
			fmt.Printf("已写入 %d 条日志\n", count)
		}
	}
}
