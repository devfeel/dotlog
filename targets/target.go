package targets

type Target interface {
	WriteLog(content string, useLayout string, level string)
	GetName() string
	GetLayout() string
}
