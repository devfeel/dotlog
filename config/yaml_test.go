package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadYamlConfig(t *testing.T) {
	// Create temp YAML config file
	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "test.yaml")

	yamlContent := `
global:
  isLog: true
  chanSize: 1000
  innerLogPath: "./"
  innerLogEncode: "utf-8"

variables:
  - name: LogDir
    value: "./logs/"

targets:
  file:
    - name: FileLogger
      isLog: true
      layout: "{datetime} - {message}"
      encode: "utf-8"
      fileMaxSize: 10240
      fileName: "./logs/app.log"

  fmt:
    - name: StdoutLogger
      isLog: true
      layout: "[{level}] {datetime} - {message}"
      encode: "utf-8"

loggers:
  - name: FileLogger
    isLog: true
    layout: "{datetime} - {message}"
    configMode: "file"
    levels:
      - level: info
        targets: "FileLogger"
        isLog: true
`
	err := os.WriteFile(configFile, []byte(yamlContent), 0644)
	if err != nil {
		t.Fatalf("failed to create test config: %v", err)
	}

	// Test loading
	config, err := LoadYamlConfig(configFile)
	if err != nil {
		t.Fatalf("failed to load YAML config: %v", err)
	}

	// Verify global config
	if !config.Global.IsLog {
		t.Error("expected IsLog to be true")
	}
	if config.Global.ChanSize != 1000 {
		t.Errorf("expected ChanSize to be 1000, got %d", config.Global.ChanSize)
	}

	// Verify variables
	if len(config.Variables) != 1 {
		t.Errorf("expected 1 variable, got %d", len(config.Variables))
	}
	if config.Variables[0].Name != "LogDir" {
		t.Errorf("expected variable name LogDir, got %s", config.Variables[0].Name)
	}

	// Verify file targets
	if len(config.Targets.FileTargets) != 1 {
		t.Errorf("expected 1 file target, got %d", len(config.Targets.FileTargets))
	}
	if config.Targets.FileTargets[0].Name != "FileLogger" {
		t.Errorf("expected target name FileLogger, got %s", config.Targets.FileTargets[0].Name)
	}

	// Verify fmt targets
	if len(config.Targets.FmtTargets) != 1 {
		t.Errorf("expected 1 fmt target, got %d", len(config.Targets.FmtTargets))
	}

	// Verify loggers
	if len(config.Loggers) != 1 {
		t.Errorf("expected 1 logger, got %d", len(config.Loggers))
	}
	if config.Loggers[0].Name != "FileLogger" {
		t.Errorf("expected logger name FileLogger, got %s", config.Loggers[0].Name)
	}
}

func TestLoadYamlConfigFileNotFound(t *testing.T) {
	_, err := LoadYamlConfig("/nonexistent/path/config.yaml")
	if err == nil {
		t.Error("expected error for nonexistent file")
	}
}

func TestLoadYamlConfigInvalidYAML(t *testing.T) {
	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "invalid.yaml")

	// Invalid YAML content
	invalidYAML := `
global:
  isLog: true
  invalid: [unclosed
`
	err := os.WriteFile(configFile, []byte(invalidYAML), 0644)
	if err != nil {
		t.Fatalf("failed to create test config: %v", err)
	}

	_, err = LoadYamlConfig(configFile)
	if err == nil {
		t.Error("expected error for invalid YAML")
	}
}

func TestLoadYamlConfigEmpty(t *testing.T) {
	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "empty.yaml")

	// Empty YAML content
	err := os.WriteFile(configFile, []byte(""), 0644)
	if err != nil {
		t.Fatalf("failed to create test config: %v", err)
	}

	config, err := LoadYamlConfig(configFile)
	if err != nil {
		t.Fatalf("failed to load empty YAML config: %v", err)
	}

	// Verify default values
	if config.Global.IsLog {
		t.Error("expected IsLog to be false for empty config")
	}
	if len(config.Variables) != 0 {
		t.Errorf("expected 0 variables, got %d", len(config.Variables))
	}
}
