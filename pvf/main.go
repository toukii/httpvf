package main

import (
	"plugin"

	"github.com/toukii/httpvf"
)

func main() {
	pg, err := plugin.Open("vf.so")
	println(err)
	verify, _ := pg.Lookup("Verify")
	httpvf.MsgLevel = httpvf.INFO

	verify.(func(string))("../vf/vf.yml")

}
