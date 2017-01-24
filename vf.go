package httpvf

import (
	"github.com/astaxie/beego/httplib"

	"time"
	"fmt"
)

func Verify(req *Req) *Msg  {
	beegoReq := httplib.NewBeegoRequest(req.URL, req.Method)

	var msg *Msg
	start:=time.Now()
	//resp := beegoReq.DoRequest()
	respstr,err:=beegoReq.String()
	duration := time.Now().Sub(start)
	fmt.Printf("cost:%d ms\n",duration.Nanoseconds()/1e6)
	cost:=int(duration.Nanoseconds()/1e6)
	if cost > req.Resp.Cost{
		msg=NewMsg()
		msg.ErrList = append(msg.ErrList,fmt.Sprintf("Request[%s] time out(%d ms> %d ms)",req.URL,cost,req.Resp.Cost))
	}
	fmt.Println("resp error:",err)
	_=respstr

	return msg
}
