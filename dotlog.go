package dotlog

import (
	"fmt"
	"strings"
)

func GetLogger(name string) Logger {
	if logger, exists := GlobalLoggerMap[name]; !exists {
		return EmptyLogger()
	} else {
		return logger
	}
}

// SprintSpacing formats using the default formats for its operands and returns the resulting string.
// Spaces are always added between operands.
func SprintSpacing(a ...interface{}) string {
	return strings.TrimSuffix(fmt.Sprintln(a...), "\n")
}
