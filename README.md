# DotLog
Simple and easy go log micro framework

## 1. Install

```
go get -u github.com/devfeel/dotlog
```

## 2. Getting Started
```go
func main() {
	dotlog.StartLogService("/home/log.conf")
	log1 = dotlog.GetLogger("FileLogger")
	log1.Info("example-normal test main")
}
```

## 3. Features
* 简单易用，100%配置化
* 支持文件、UDP、Http、EMail四种日志目标
* 支持自定义变量
* 更多待完善