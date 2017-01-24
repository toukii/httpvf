package httpvf


import (
	yaml "gopkg.in/yaml.v2"
	"github.com/toukii/goutils"
	"fmt"
)

type Resp struct{
	Code int
	Cost int
	Body string
}

const(
	GET = "GET"
	POST = "POST"
)

type Req struct {
	URL string
	Method string
	Body string
	Resp Resp
}

func Reqs(filename string) (reqs []Req,err error) {
	in:=goutils.ReadFile(filename)

	err=yaml.Unmarshal(in, &reqs)
	if goutils.CheckErr(err) {
		return nil,err
	}
	return reqs,nil
}

func Req1(filename string)(*Req, error) {
	bs:=goutils.ReadFile(filename)

	var req Req
	err := yaml.Unmarshal(bs,&req)
	if goutils.CheckErr(err){
		return nil, err
	}

	req.Body = "hello"
	req.Resp.Body = "world"
	req.Method = GET
	reqs:=[]Req{req,req}
	bs2,err:=yaml.Marshal(reqs)
	goutils.CheckErr(err)
	fmt.Print(goutils.ToString(bs2))

	return &req,nil
}
