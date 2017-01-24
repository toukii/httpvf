package main

import (
	//"github.com/ghodss/yaml"
	yaml "gopkg.in/yaml.v2"
	"github.com/toukii/goutils"

	"fmt"
)

type Resp struct{
	Code int
	Cost int
	Body string
}

type Req struct {
	URL string
	Body string
	Resp Resp
}

func main() {
	//t1()
	t2()
}

func t2()  {
	bs:=goutils.ReadFile("hts.yaml")

	var v []Req
	err:=yaml.Unmarshal(bs,&v)
	goutils.CheckErr(err)
	goutils.Print(v)
}

func t1() {
	bs:=goutils.ReadFile("ht.yaml")

	var v Req
	err:=yaml.Unmarshal(bs,&v)
	goutils.CheckErr(err)
	goutils.Log(v)

	v.Body = "hello"
	v.Resp.Body = "world"
	reqs:=[]Req{v,v}
	bs2,err:=yaml.Marshal(reqs)
	goutils.CheckErr(err)
	fmt.Print(goutils.ToString(bs2))

}

