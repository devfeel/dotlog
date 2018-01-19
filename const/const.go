package _const

const(
	Version = "0.9.4"
)

const (
	DefaultEncode   = "gb2312"
	DefaultChanSize = 1000
)

const (
	DefaultDateFormatForFileName = "2006_01_02"
	DefaultDateLayout            = "2006-01-02"
	DefaultFullTimeLayout        = "2006-01-02 15:04:05.999999"
	DefaultTimeLayout            = "2006-01-02 15:04:05"
	DefaultNoSeparatorTimeLayout = "20060102150405.999999"
)

const (
	ConfigMode_Classics = "classics"
	ConfigMode_File     = "file"
	ConfigMode_Fmt     = "fmt"
	ConfigMode_FileFmt     = "filefmt"
)

const (
	TargetType_File  = "File"
	TargetType_Udp   = "Udp"
	TargetType_Http  = "Http"
	TargetType_EMail = "EMail"
	TargetType_Fmt  = "Fmt"
)

const (
	LogLevel_Trace = "TRACE"
	LogLevel_Debug = "DEBUG"
	LogLevel_Info  = "INFO"
	LogLevel_Warn  = "WARN"
	LogLevel_Error = "ERROR"
)
