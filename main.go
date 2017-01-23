package main

import (
	"github.com/ghodss/yaml"
	"github.com/toukii/goutils"
)

func main() {
	bs:=goutils.ReadFile("ht.yaml")

	var v interface{}
	err:=yaml.Unmarshal(bs,&v)
	goutils.CheckErr(err)
}

