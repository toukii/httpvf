package main

import (
	"github.com/toukii/httpvf"

	"fmt"
	"flag"
)

var(
	vf string
)

func init() {
	flag.StringVar(&vf,"v","vf.yml","vf -v vf.yml")
}

func main() {
	flag.Parse()
	httpvf.MsgLevel = httpvf.INFO
	// test()
	// t1()
	//t2()
	verify(vf)
}

func verify(vf string)  {
	reqs, _ := httpvf.Reqs(vf)
	for _, it := range reqs {
		msg := httpvf.Verify(it)
		if nil != msg {
			fmt.Println(msg)
		}
	}
}

func t2() {
	verify("vfs.yml")
}

func t1() {
	verify("vf.yml")
}

func test() {
	httpvf.Test()
}
