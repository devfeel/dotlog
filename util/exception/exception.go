package exception

import (
	"fmt"
	"runtime/debug"
)

//ConvertError convert err to string
func ConvertError(err interface{}) string {
	errmsg := fmt.Sprint(err)
	stack := string(debug.Stack())
	return "Exception: " + errmsg + "\r\nStack: " + stack
}
