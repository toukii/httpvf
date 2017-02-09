package main

import (
	"plugin"
	"github.com/toukii/httpvf"
	"fmt"
)

func main() {
	pg,err:=plugin.Open("./vf.so")
	println(err)
	pg.Lookup()

	reqs, _ := httpvf.Reqs(vf)
	for _, it := range reqs {
		msg := httpvf.Verify(it)
		if nil != msg {
			fmt.Println(msg)
		}
	}
}