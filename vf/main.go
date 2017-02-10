package main

import (
	"github.com/toukii/httpvf"

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
	 //test()
	// t1()
	//t2()
	httpvf.Verify(vf)
}

func verify(vf string)  {
	httpvf.Verify(vf)
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
