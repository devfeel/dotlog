package config

import (
	"encoding/xml"
	"errors"
	"github.com/devfeel/dotlog/const"
	"github.com/devfeel/dotlog/util/file"
	"io/ioutil"
)

const (
	ConfigType_Xml = "xml"
)

var (
	GlobalAppConfig *AppConfig
)

func init() {
	GlobalAppConfig = &AppConfig{}
}

type (
	AppConfig struct {
		XMLName   xml.Name          `xml:"config" json:"-"`
		Global    *GlobalConfig     `xml:"global"`
		Targets   *TargetList       `xml:"targets"`
		Variables []*VariableConfig `xml:"variable>var"`
		Loggers   []*LoggerConfig   `xml:"loggers>logger"`
	}

	GlobalConfig struct {
		IsLog          bool   `xml:"islog,attr"`
		ChanSize       int    `xml:"chansize,attr"` //日志队列长度，默认为DefaultChanSize = 1000
		InnerLogPath   string `xml:"innerlogpath,attr"`
		InnerLogEncode string `xml:"innerlogencode,attr"`
	}

	VariableConfig struct {
		Name  string `xml:"name,attr"`
		Value string `xml:"value,attr"`
	}

	LoggerConfig struct {
		Name       string               `xml:"name,attr"`
		IsLog      bool                 `xml:"islog,attr"`
		Layout     string               `xml:"layout,attr"`
		ConfigMode string               `xml:"configmode,attr"`
		Levels     []*LoggerLevelConfig `xml:"level"`
	}

	LoggerLevelConfig struct {
		Level   string `xml:"level,attr"`
		Targets string `xml:"targets,attr"`
		IsLog   bool   `xml:"islog,attr"`
	}

	TargetList struct {
		FileTargets  []*FileTargetConfig  `xml:"file>target"`
		UdpTargets   []*UdpTargetConfig   `xml:"udp>target"`
		HttpTargets  []*HttpTargetConfig  `xml:"http>target"`
		EMailTargets []*EMailTargetConfig `xml:"email>target"`
		FmtTargets []*FmtTargetConfig `xml:"fmt>target"`
	}

	FileTargetConfig struct {
		Name        string `xml:"name,attr"`
		IsLog       bool   `xml:"islog,attr"`
		Layout      string `xml:"layout,attr"`
		Encode      string `xml:"encode,attr"`
		FileMaxSize int64  `xml:"filemaxsize,attr"` //日志文件最大容量，单位为KB
		FileName    string `xml:"filename,attr"`
	}

	UdpTargetConfig struct {
		Name     string `xml:"name,attr"`
		IsLog    bool   `xml:"islog,attr"`
		Layout   string `xml:"layout,attr"`
		Encode   string `xml:"encode,attr"`
		RemoteIP string `xml:"remoteip,attr"`
	}

	HttpTargetConfig struct {
		Name    string `xml:"name,attr"`
		IsLog   bool   `xml:"islog,attr"`
		Layout  string `xml:"layout,attr"`
		Encode  string `xml:"encode,attr"`
		HttpUrl string `xml:"httpurl,attr"`
	}

	FmtTargetConfig struct {
		Name    string `xml:"name,attr"`
		IsLog   bool   `xml:"islog,attr"`
		Layout  string `xml:"layout,attr"`
		Encode  string `xml:"encode,attr"`
	}

	EMailTargetConfig struct {
		Name         string `xml:"name,attr"`
		IsLog        bool   `xml:"islog,attr"`
		Layout       string `xml:"layout,attr"`
		Encode       string `xml:"encode,attr"`
		MailServer   string `xml:"mailserver,attr"`
		MailAccount  string `xml:"mailaccount,attr"`
		MailNickName string `xml:"mailnickname,attr"`
		MailPassword string `xml:"mailpassword,attr"`
		ToMail       string `xml:"tomail,attr"`
		Subject      string `xml:"subject,attr"`
	}
)

func (c *GlobalConfig) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type innerType GlobalConfig // new type to prevent recursion
	item := innerType{
		IsLog:          true,
		InnerLogEncode: _const.DefaultEncode,
		InnerLogPath:   _file.GetCurrentDirectory(),
	}
	if err := d.DecodeElement(&item, &start); err != nil {
		return err
	}
	*c = (GlobalConfig)(item)
	return nil
}

func (c *LoggerLevelConfig) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type innerType LoggerLevelConfig // new type to prevent recursion
	item := innerType{
		IsLog: true,
	}
	if err := d.DecodeElement(&item, &start); err != nil {
		return err
	}
	*c = (LoggerLevelConfig)(item)
	return nil
}

func (c *FileTargetConfig) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type innerType FileTargetConfig // new type to prevent recursion
	item := innerType{
		IsLog: true,
	}
	if err := d.DecodeElement(&item, &start); err != nil {
		return err
	}
	*c = (FileTargetConfig)(item)
	return nil
}
func (c *UdpTargetConfig) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type innerType UdpTargetConfig // new type to prevent recursion
	item := innerType{
		IsLog: true,
	}
	if err := d.DecodeElement(&item, &start); err != nil {
		return err
	}
	*c = (UdpTargetConfig)(item)
	return nil
}

func (c *HttpTargetConfig) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type innerType HttpTargetConfig // new type to prevent recursion
	item := innerType{
		IsLog: true,
	}
	if err := d.DecodeElement(&item, &start); err != nil {
		return err
	}
	*c = (HttpTargetConfig)(item)
	return nil
}

func (c *EMailTargetConfig) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type innerType EMailTargetConfig // new type to prevent recursion
	item := innerType{
		IsLog: true,
	}
	if err := d.DecodeElement(&item, &start); err != nil {
		return err
	}
	*c = (EMailTargetConfig)(item)
	return nil
}

func (c *FmtTargetConfig) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type innerType FmtTargetConfig // new type to prevent recursion
	item := innerType{
		IsLog: true,
	}
	if err := d.DecodeElement(&item, &start); err != nil {
		return err
	}
	*c = (FmtTargetConfig)(item)
	return nil
}

func (c *LoggerConfig) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type innerType LoggerConfig // new type to prevent recursion
	item := innerType{
		IsLog: true,
	}
	if err := d.DecodeElement(&item, &start); err != nil {
		return err
	}
	*c = (LoggerConfig)(item)
	return nil
}

//初始化配置文件
//如果发生异常，返回异常
func InitConfig(configFile string, confType ...interface{}) (config *AppConfig, err error) {

	//检查配置文件有效性
	//1、按绝对路径检查
	//2、尝试在当前进程根目录下寻找
	//3、尝试在当前进程根目录/config/ 下寻找
	//fixed for issue #15 读取配置文件路径
	realFile := configFile
	if !_file.Exist(realFile) {
		realFile = _file.GetCurrentDirectory() + "/" + configFile
		if !_file.Exist(realFile) {
			realFile = _file.GetCurrentDirectory() + "/config/" + configFile
			if !_file.Exist(realFile) {
				return nil, errors.New("no exists config file => " + configFile)
			}
		}
	}

	cType := ConfigType_Xml

	config, err = initConfig(realFile, cType, fromXml)

	if err != nil {
		return config, err
	}

	return config, nil
}

func initConfig(configFile string, ctType string, f func([]byte, interface{}) error) (*AppConfig, error) {
	content, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, errors.New("GoLog:Config:initConfig 当前cType:" + ctType + " 配置文件[" + configFile + "]读取失败 - " + err.Error())
	}

	var config *AppConfig
	err = f(content, &config)
	if err != nil {
		return nil, errors.New("GoLog:Config:initConfig 当前cType:" + ctType + " 配置文件[" + configFile + "]解析失败 - " + err.Error())
	}
	return config, nil
}
