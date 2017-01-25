package httpvf

import (
	"github.com/astaxie/beego/httplib"
	"github.com/toukii/goutils"

	"fmt"
	"time"
)

func Verify(req *Req) *Msg {
	msg := NewMsg()

	beegoReq := httplib.NewBeegoRequest(req.URL, req.Method)

	//  start
	start := time.Now()
	resp, err := beegoReq.DoRequest()
	duration := time.Now().Sub(start)
	// end

	if goutils.CheckErr(err) {
		msg.ErrList = append(msg.ErrList, err.Error())
	}
	// fmt.Printf("Request[%s] cost:%d ms\n", req.URL, duration.Nanoseconds()/1e6)
	cost := int(duration.Nanoseconds() / 1e6)
	if cost > req.Resp.Cost {
		msg.ErrList = append(msg.ErrList, fmt.Sprintf("%s :time last longer(%d ms> %d ms)", req.URL, cost, req.Resp.Cost))
	}
	if req.Resp.Code != resp.StatusCode {
		msg.ErrList = append(msg.ErrList, fmt.Sprintf("%s :error code::%d, %d wanted", req.URL, resp.StatusCode, req.Resp.Code))
	}
	if len(msg.ErrList) <= 0 {
		return nil
	}
	return msg
}
