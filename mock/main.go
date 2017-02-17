package main

import (
	"fmt"
	"github.com/go-macaron/macaron"
	"github.com/toukii/goutils"
	"html/template"
	"io"
	"os"
)

type Msg struct {
	Map map[string]interface{}
	Message string
	Cost    float64
}

func newMsg(msg string, cost float64) *Msg {
	return &Msg{
		Message: msg,
		Cost:    cost,
		Map: make(map[string]interface{}),
	}
}

func main() {
	m := macaron.Classic()
	m.Use(macaron.Renderer(macaron.RenderOptions{
		Funcs: []template.FuncMap{
			{},
		},
	}))
	m.Group("/", func() {
		m.Get("", func(ctx *macaron.Context) {
			bd,er:=ctx.Req.Body().String()
			fmt.Println("req-body:",bd,er)
			fmt.Println("This is toukii,root")
			ctx.JSON(200, newMsg("This is toukii,root", 3.14))
		})
		m.Get("r1", func(ctx *macaron.Context) {
			bd,er:=ctx.Req.Body().String()
			fmt.Println("req-body:",bd,er)
			fmt.Println("This is toukii,r1")
			msg:=newMsg("This is toukii,r1", 0.315)
			msg.Map["1"]="hello"
			ctx.JSON(201, []*Msg{msg})
		})
		m.Combo("r2").Post(r2)
		m.Post("upload", func(ctx *macaron.Context) {
			bd,er:=ctx.Req.Body().String()
			fmt.Println("req-body:",bd,er)
			mul, header, err := ctx.GetFile("filename")
			if goutils.CheckErr(err) {
				ctx.JSON(403, err)
				return
			}
			fmt.Println("upload a file:", header.Filename)

			/*goutils.CheckErr(err)
			err2:=ctx.SaveToFile(header.Filename,"./")
			goutils.CheckErr(err2)*/
			fmt.Println(err)

			file, _ := os.OpenFile(header.Filename, os.O_CREATE|os.O_RDWR, 0644)
			defer file.Close()
			_, err2 := io.Copy(file, mul)
			goutils.CheckErr(err2)
			defer mul.Close()

		})
	})
	// macaron.Env = macaron.PROD
	macaron.Env = macaron.DEV
	m.Run()
}

func r2(ctx *macaron.Context) {
	fmt.Println("This is toukii,r2")
	ctx.JSON(200, newMsg("This is toukii,r2", 0.101))
}
