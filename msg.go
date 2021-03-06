package httpvf

import "fmt"

const (
	INFO       = "INFO"
	WARN       = "WARN"
	ERROR      = "ERROR"
	FATAL      = "FATAL"
	CONCLUSION = "CONCLUSION"
)

var (
	MsgLevel = INFO // msg 的级别
)

type Msg struct {
	Req           *Req
	InfoLog       []*Log
	WarnLog       []*Log
	ErrorLog      []*Log
	FatalLog      []*Log
	ConclusionLog []*Log
}

func (m *Msg) AppendLogs(logs []*Log) {
	for _, log := range logs {
		m.AppendLog(log)
	}
}

func (m *Msg) AppendLog(log *Log) {
	switch log.Level {
	case "CONCLUSION":
		m.ConclusionLog = append(m.ConclusionLog, log)
	case "FATAL":
		m.FatalLog = append(m.FatalLog, log)
	case "ERROR":
		m.ErrorLog = append(m.ErrorLog, log)
	case "WARN":
		m.WarnLog = append(m.WarnLog, log)
	case "INFO":
		m.InfoLog = append(m.InfoLog, log)
	}
}

func (m *Msg) Append(level, out string) {
	switch level {
	case "CONCLUSION":
		m.ConclusionLog = append(m.ConclusionLog, newLog(level, out))
	case "FATAL":
		m.FatalLog = append(m.FatalLog, newLog(level, out))
	case "ERROR":
		m.ErrorLog = append(m.ErrorLog, newLog(level, out))
	case "WARN":
		m.WarnLog = append(m.WarnLog, newLog(level, out))
	case "INFO":
		m.InfoLog = append(m.InfoLog, newLog(level, out))
	}
}

func newMsg(req *Req) *Msg {
	return &Msg{
		Req:           req,
		InfoLog:       make([]*Log, 0, 3),
		WarnLog:       make([]*Log, 0, 3),
		ErrorLog:      make([]*Log, 0, 3),
		FatalLog:      make([]*Log, 0, 3),
		ConclusionLog: make([]*Log, 0, 3),
	}
}

type Log struct {
	Level string
	Out   string
}

func newLog(level, out string) *Log {
	return &Log{
		Level: level,
		Out:   out,
	}
}

func (m Msg) Logs() []*Log {
	logs := make([]*Log, 0, 10)
	logs = append(logs, m.InfoLog...)
	logs = append(logs, m.WarnLog...)
	logs = append(logs, m.ErrorLog...)
	logs = append(logs, m.FatalLog...)
	logs = append(logs, m.ConclusionLog...)
	return logs
}

func (m Msg) String() string {
	// var logs []*Log
	logs := make([]*Log, 0, 10)
	switch MsgLevel {
	case "INFO":
		logs = append(logs, m.InfoLog...)
		logs = append(logs, m.WarnLog...)
		logs = append(logs, m.ErrorLog...)
		logs = append(logs, m.FatalLog...)
	case "WARN":
		logs = append(logs, m.WarnLog...)
		logs = append(logs, m.ErrorLog...)
		logs = append(logs, m.FatalLog...)
	case "ERROR":
		logs = append(logs, m.ErrorLog...)
		logs = append(logs, m.FatalLog...)
	case "FATAL":
		logs = append(logs, m.FatalLog...)
	}
	logs = append(logs, m.ConclusionLog...)
	if len(logs) <= 0 {
		return ""
	}
	ret := fmt.Sprintf("%s\n", m.Req.URL)
	for _, it := range logs {
		ret += fmt.Sprintf("  [%s] %s\n", it.Level, it.Out)
	}
	return ret
}
