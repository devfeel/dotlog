package targets

import (
	"github.com/devfeel/dotlog/config"
	"github.com/devfeel/dotlog/const"
	"github.com/devfeel/dotlog/internal"
	"github.com/devfeel/dotlog/layout"
	"net"
	"strings"
)

type UdpTarget struct {
	BaseTarget
	RemoteIP string
}

func NewUdpTarget(conf *config.UdpTargetConfig) *UdpTarget {
	t := &UdpTarget{}
	t.TargetType = _const.TargetType_Udp
	t.Name = conf.Name
	t.IsLog = conf.IsLog
	t.Encode = conf.Encode
	t.Layout = conf.Layout
	t.RemoteIP = conf.RemoteIP
	return t
}

func (t *UdpTarget) WriteLog(log string, useLayout string, level string) {
	if t.IsLog {
		if t.Layout != "" {
			useLayout = t.Layout
		}
		logContent := layout.CompileLayout(useLayout)
		logContent = layout.ReplaceLogLevelLayout(logContent, level)

		if useLayout != "" {
			logContent = strings.Replace(logContent, "{message}", log, -1)
			udpAddr, err := net.ResolveUDPAddr("udp4", t.RemoteIP)
			if err != nil {
				internal.GlobalInnerLogger.Error(err, "UdpTarget:WriteLog:ResolveUDPAddr error", log, t.RemoteIP)
			}
			conn, err := net.DialUDP("udp", nil, udpAddr)
			defer conn.Close()
			if err != nil {
				internal.GlobalInnerLogger.Error(err, "UdpTarget:WriteLog:DialUDP error", log, t.RemoteIP)
			}
			_, err = conn.Write([]byte(logContent))
			if err != nil {
				internal.GlobalInnerLogger.Error(err, "UdpTarget:WriteLog:DialUDP error", log, t.RemoteIP)
			}
		}
	}
}
