package httpvf

import "fmt"

const (
	INFO  = "INFO"
	WARN  = "WARN"
	ERROR = "ERROR"
	FATAL = "FATAL"
)

var (
	MsgLevel = INFO // msg 的级别
)

type Msg struct {
	req      *Req
	ErrList  []string
	InfoLog  []*Log
	WarnLog  []*Log
	ErrorLog []*Log
	FatalLog []*Log
}

func (m *Msg) Append(level, out string) {
	switch level {
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
		req:      req,
		ErrList:  make([]string, 0, 3),
		InfoLog:  make([]*Log, 0, 3),
		WarnLog:  make([]*Log, 0, 3),
		ErrorLog: make([]*Log, 0, 3),
		FatalLog: make([]*Log, 0, 3),
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

func (m Msg) String() string {
	// var logs []*Log
	logs := make([]*Log, 0, 10)
	switch MsgLevel {
	case "INFO":
		logs = append(logs, m.FatalLog...)
		logs = append(logs, m.ErrorLog...)
		logs = append(logs, m.WarnLog...)
		logs = append(logs, m.InfoLog...)
	case "WARN":
		logs = append(logs, m.FatalLog...)
		logs = append(logs, m.ErrorLog...)
		logs = append(logs, m.WarnLog...)
	case "ERROR":
		logs = append(logs, m.FatalLog...)
		logs = append(logs, m.ErrorLog...)
	case "FATAL":
		logs = append(logs, m.FatalLog...)
	}
	if len(logs) <= 0 {
		return ""
	}
	ret := fmt.Sprintf("# %s\n", m.req.URL)
	for i, it := range logs {
		ret += fmt.Sprintf("%d. [%s] %s\n", i+1, it.Level, it.Out)
	}
	return ret
}
