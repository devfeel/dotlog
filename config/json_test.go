package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadYamlConfigWithJSON(t *testing.T) {
	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "test.json.yaml")

	yamlContent := `
global:
  isLog: true
  chanSize: 1000
  innerLogPath: "./"
  innerLogEncode: "utf-8"

targets:
  json:
    - name: JSONLogger
      isLog: true
      layout: "{datetime} - {message}"
      encode: "utf-8"
      fileName: "./logs/app.json"
      fileMaxSize: 10240
      prettyPrint: true

loggers:
  - name: JSONLogger
    isLog: true
    layout: "{datetime} - {message}"
    configMode: "json"
    levels:
      - level: info
        targets: "JSONLogger"
        isLog: true
`
	err := os.WriteFile(configFile, []byte(yamlContent), 0644)
	if err != nil {
		t.Fatalf("failed to create test config: %v", err)
	}

	config, err := LoadYamlConfig(configFile)
	if err != nil {
		t.Fatalf("failed to load YAML config: %v", err)
	}

	// Verify JSON targets
	if len(config.Targets.JSONTargets) != 1 {
		t.Errorf("expected 1 JSON target, got %d", len(config.Targets.JSONTargets))
	}

	jsonTarget := config.Targets.JSONTargets[0]
	if jsonTarget.Name != "JSONLogger" {
		t.Errorf("expected target name JSONLogger, got %s", jsonTarget.Name)
	}
	if jsonTarget.FileName != "./logs/app.json" {
		t.Errorf("expected fileName ./logs/app.json, got %s", jsonTarget.FileName)
	}
	if jsonTarget.FileMaxSize != 10240 {
		t.Errorf("expected fileMaxSize 10240, got %d", jsonTarget.FileMaxSize)
	}
	if jsonTarget.PrettyPrint == nil || !*jsonTarget.PrettyPrint {
		t.Error("expected prettyPrint to be true")
	}
}
