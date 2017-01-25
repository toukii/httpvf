package httpvf

import (
	"github.com/astaxie/beego/httplib"
	"github.com/toukii/goutils"

	"fmt"
	"time"
)

func Verify(req *Req) *Msg {
	msg := newMsg(req)

	beegoReq := httplib.NewBeegoRequest(req.URL, req.Method)

	//  start
	start := time.Now()
	resp, err := beegoReq.DoRequest()
	duration := time.Now().Sub(start)
	// end

	if goutils.CheckErr(err) {
		msg.Append(FATAL, err.Error())
	}
	// fmt.Printf("Request[%s] cost:%d ms\n", req.URL, duration.Nanoseconds()/1e6)
	cost := int(duration.Nanoseconds() / 1e6)
	if cost > req.Resp.Cost {
		msg.Append(ERROR, fmt.Sprintf("time cost: %d ms more than %d ms;", cost, req.Resp.Cost))
	} else if cost > req.Resp.Cost*3/4 {
		msg.Append(WARN, fmt.Sprintf("time cost: %d ms near by %d ms;", cost, req.Resp.Cost))
	} else {
		msg.Append(INFO, fmt.Sprintf("time cost: %d ms / %d ms;", cost, req.Resp.Cost))
	}
	if req.Resp.Code != resp.StatusCode {
		msg.ErrList = append(msg.ErrList, fmt.Sprintf("error code::%d gotten, %d wanted", resp.StatusCode, req.Resp.Code))
	}
	if len(msg.ErrList) <= 0 {
		return nil
	}
	return msg
}
