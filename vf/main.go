package main

import (
	"github.com/toukii/httpvf"

	"fmt"
	"flag"
	"sync"
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
	var wg sync.WaitGroup
	for _, it := range reqs {
		wg.Add(1)
		go func(it *httpvf.Req){
			i:=0
			var cost int
			for {
				msg := httpvf.Verify(it)
				if nil != msg {
					fmt.Println(msg)
					cost += msg.Req.Resp.RealCost
				}
				i++
				if i>= it.N {
					fmt.Println("avg cost: ",cost/i,"ms")
					fmt.Println("TPS:",1000.0*float32(i)/float32(cost))
					break
				}
			}
			wg.Done()
		}(it)
	}
	wg.Wait()
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
