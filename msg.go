package httpvf

import "fmt"

type Msg struct {
	ErrList []string
}

func NewMsg() *Msg {
	return &Msg{
		ErrList: make([]string, 0, 3),
	}
}

func (m Msg) String() string {
	return fmt.Sprintf("%#v\n", m.ErrList)
}
