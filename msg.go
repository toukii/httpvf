package httpvf

import "fmt"

const (
	INFO  = "INFO"
	WARN  = "WARN"
	ERROR = "ERROR"
	FATAL = "FATAL"
)

var (
	MsgLevel = WARN // msg 的级别
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
	case "INFO":
		m.InfoLog = append(m.InfoLog, newLog(level, out))
		break
	case "WARN":
		m.WarnLog = append(m.WarnLog, newLog(level, out))
		break
	case "ERROR":
		m.ErrorLog = append(m.ErrorLog, newLog(level, out))
		break
	case "FATAL":
		m.FatalLog = append(m.FatalLog, newLog(level, out))
		break
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
	var logs []*Log
	switch MsgLevel {
	case "INFO":
		logs = m.InfoLog
		break
	case "WARN":
		logs = m.WarnLog
		break
	case "ERROR":
		logs = m.ErrorLog
		break
	case "FATAL":
		logs = m.FatalLog
		break
	}
	if len(logs) <= 0 {
		return ""
	}
	ret := fmt.Sprintf("# request [%s] error:\n", m.req.URL)
	for i, it := range logs {
		ret += fmt.Sprintf("%d. %s\n", i+1, it.Out)
	}
	return ret
}
