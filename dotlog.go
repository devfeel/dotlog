package dotlog

func GetLogger(name string) Logger {
	if logger, exists := GlobalLoggerMap[name]; !exists {
		return EmptyLogger()
	} else {
		return logger
	}
}
