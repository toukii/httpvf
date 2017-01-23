package main

import (
	//"github.com/ghodss/yaml"
	yaml "gopkg.in/yaml.v2"
	"github.com/toukii/goutils"
)

type Req struct{
	Code int
	Cost int
}


func main() {
	bs:=goutils.ReadFile("ht.yaml")

	var v Req
	err:=yaml.Unmarshal(bs,&v)
	goutils.CheckErr(err)
}

