package targets

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/devfeel/dotlog/config"
	"github.com/devfeel/dotlog/const"
)

// JSONTarget outputs logs in JSON format
type JSONTarget struct {
	BaseTarget
	FileName    string
	MaxSize    int64 // in KB
	Encode      string
	prettyPrint bool
}

// NewJSONTarget creates a new JSON target
func NewJSONTarget(conf *config.JSONTargetConfig) *JSONTarget {
	t := &JSONTarget{}
	t.TargetType = _const.TargetType_JSON
	t.Name = conf.Name
	t.IsLog = conf.IsLog
	t.Encode = conf.Encode
	t.FileName = conf.FileName
	t.MaxSize = conf.FileMaxSize
	if conf.PrettyPrint != nil {
		t.prettyPrint = *conf.PrettyPrint
	}
	return t
}

// WriteLog writes log in JSON format
func (t *JSONTarget) WriteLog(message string, useLayout string, level string) {
	if !t.IsLog {
		return
	}

	entry := map[string]interface{}{
		"timestamp": time.Now().Format(time.RFC3339),
		"level":     level,
		"message":   message,
		"logger":    t.Name,
	}

	var output []byte
	var err error

	if t.prettyPrint {
		output, err = json.MarshalIndent(entry, "", "  ")
	} else {
		output, err = json.Marshal(entry)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "dotlog: failed to marshal JSON: %v\n", err)
		return
	}

	t.writeTarget(string(output), level)
}

func (t *JSONTarget) writeTarget(log string, level string) {
	if t.FileName != "" {
		// Write to file
		t.writeToFile(log)
	} else {
		// Write to stdout/stderr
		if level == _const.LogLevel_Error {
			os.Stderr.WriteString(log + "\n")
		} else {
			fmt.Println(log)
		}
	}
}

func (t *JSONTarget) writeToFile(log string) {
	// Simple file write - in production should handle rotation
	f, err := os.OpenFile(t.FileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dotlog: failed to open file: %v\n", err)
		return
	}
	defer f.Close()

	_, err = f.WriteString(log + "\n")
	if err != nil {
		fmt.Fprintf(os.Stderr, "dotlog: failed to write to file: %v\n", err)
	}
}
