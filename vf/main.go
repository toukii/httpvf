package main

import (
	"github.com/toukii/httpvf"

	"fmt"
)

func main() {
	//t1()
	t2()
}

func t2()  {
	reqs,_:=httpvf.Reqs("hts.yaml")
	for _,it := range reqs{
		msg := httpvf.Verify(&it)
		if nil != msg {
			fmt.Println(msg)
		}
	}
}

func t1()  {
	req,_:=httpvf.Req1("ht.yaml")
	fmt.Println(req)
}
