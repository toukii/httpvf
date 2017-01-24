package httpvf

import (
	"github.com/astaxie/beego/httplib"

	"time"
	"fmt"
)

func Verify(req Req)  {
	beegoReq := httplib.NewBeegoRequest(req.URL, req.Method)
	start:=time.Now()
	time.Sleep(1e9)
	//resp := beegoReq.DoRequest()
	respstr,err:=beegoReq.String()
	duration := time.Now().Sub(start)
	//fmt.Println("cost:",duration.Nanoseconds(),duration)
	fmt.Printf("cost:%d ms\n",duration.Nanoseconds()/1e5)
	fmt.Println("resp error:",err)
	_=respstr
}
