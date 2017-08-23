package layout

import (
	"github.com/devfeel/dotlog/const"
	"github.com/devfeel/dotlog/util/convert"
	"strconv"
	"strings"
	"time"
)

const (
	Key_LogLevel     = "{LogLevel}"
	Key_FullDateTime = "{fulldatetime}"
	Key_DateTime     = "{datetime}"
	Key_Date         = "{date}"
	Key_Year         = "{year}"
	Key_Month        = "{month}"
	Key_Day          = "{day}"
	Key_Hour         = "{hour}"
	Key_Minute       = "{minute}"
	Key_Second       = "{second}"
)

var CurrentVariable *Variable

type ParseSysVarHandle func() string

type Variable struct {
	SysVarFunc map[string]ParseSysVarHandle
	UserVar    map[string]string
}

func GetVariable() *Variable {
	if CurrentVariable == nil {
		v := new(Variable)
		v.SysVarFunc = make(map[string]ParseSysVarHandle)
		v.RegisterSysVar(Key_FullDateTime, GetFullDateTime)
		v.RegisterSysVar(Key_DateTime, GetDateTime)
		v.RegisterSysVar(Key_Date, GetDate)
		v.RegisterSysVar(Key_Year, GetYear)
		v.RegisterSysVar(Key_Month, GetMonth)
		v.RegisterSysVar(Key_Day, GetDay)
		v.RegisterSysVar(Key_Hour, GetHour)
		v.RegisterSysVar(Key_Minute, GetMinute)
		v.RegisterSysVar(Key_Second, GetSecond)

		v.UserVar = make(map[string]string)
		CurrentVariable = v
	}
	return CurrentVariable
}

//RegisterUserVar register user var with key/value
func (v *Variable) RegisterUserVar(key, value string) {
	v.UserVar["{"+key+"}"] = value
}

//RegisterSysVar register sys var with key/ParseSysVarHandle
func (v *Variable) RegisterSysVar(key string, handler ParseSysVarHandle) {
	v.SysVarFunc[key] = handler
}

//ConvertVariable convert layout var to value
func (v *Variable) ConvertVariable(variable string) string {
	value, isSysVar := v.convertSysVariable(variable)
	if !isSysVar {
		uval, exists := v.UserVar[variable]
		if exists {
			return CompileLayout(uval)
		}
	}
	return value
}

// convertSysVariable convert sys var to string
func (v *Variable) convertSysVariable(variable string) (value string, isSysVar bool) {
	isSysVar = false
	if variable == "" {
		return "", isSysVar
	}

	t_variable := strings.ToLower(variable)
	f, exists := v.SysVarFunc[t_variable]
	if exists {
		isSysVar = true
		return f(), isSysVar
	}
	return variable, isSysVar
}

func GetFullDateTime() string {
	return time.Now().Format(_const.DefaultFullTimeLayout)
}

func GetDateTime() string {
	return time.Now().Format(_const.DefaultTimeLayout)
}

func GetDate() string {
	return time.Now().Format(_const.DefaultDateLayout)
}

func GetYear() string {
	return convert.Int2String(time.Now().Year())
}
func GetMonth() string {
	return strconv.Itoa((int)(time.Now().Month()))
}
func GetDay() string {
	return convert.Int2String(time.Now().Day())
}
func GetHour() string {
	return convert.Int2String(time.Now().Hour())
}
func GetMinute() string {
	return convert.Int2String(time.Now().Minute())
}
func GetSecond() string {
	return convert.Int2String(time.Now().Second())
}
