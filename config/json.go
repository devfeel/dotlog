package config

// JSONTargetConfig JSON 输出目标配置
type JSONTargetConfig struct {
	Name        string `yaml:"name"`
	IsLog       bool   `yaml:"isLog"`
	Layout      string `yaml:"layout"`
	Encode      string `yaml:"encode"`
	FileName    string `yaml:"fileName"`
	FileMaxSize int64  `yaml:"fileMaxSize"`
	PrettyPrint *bool  `yaml:"prettyPrint"`
}
