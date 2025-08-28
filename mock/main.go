package main

// import (
// 	"fmt"
// 	"html/template"
// 	"io"
// 	"os"

// 	"gopkg.in/macaron.v1"
// )

// type Msg struct {
// 	Map     map[string]interface{}
// 	Message string
// 	Cost    float64
// }

// func newMsg(msg string, cost float64) *Msg {
// 	return &Msg{
// 		Message: msg,
// 		Cost:    cost,
// 		Map:     make(map[string]interface{}),
// 	}
// }

// func main() {
// 	m := macaron.Classic()
// 	m.Use(macaron.Renderer(macaron.RenderOptions{
// 		Funcs: []template.FuncMap{
// 			{},
// 		},
// 	}))
// 	m.Group("/", func() {
// 		m.Get("", func(ctx *macaron.Context) {
// 			bd, er := ctx.Req.Body().String()
// 			fmt.Println(bd, er)
// 			ctx.JSON(200, newMsg("This is toukii,root", 3.14))
// 		})
// 		m.Get("r1", func(ctx *macaron.Context) {
// 			bd, er := ctx.Req.Body().String()
// 			fmt.Println(bd, er)
// 			msg := newMsg("This is toukii,r1", 0.315)
// 			msg.Map["1"] = "hello"
// 			ctx.JSON(201, []*Msg{msg})
// 		})
// 		m.Combo("r2").Post(r2)
// 		m.Post("upload", func(ctx *macaron.Context) {
// 			r := ctx.Req
// 			r.ParseMultipartForm(32 << 20)
// 			file, handler, err := r.FormFile("verify")
// 			if err != nil {
// 				fmt.Println(err)
// 				return
// 			}
// 			defer file.Close()
// 			f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
// 			if err != nil {
// 				fmt.Println(err)
// 				return
// 			}
// 			defer f.Close()
// 			io.Copy(f, file)
// 			w := ctx.Resp
// 			fmt.Fprint(w, "upload ok!")
// 		})
// 	})
// 	macaron.Env = macaron.PROD
// 	//macaron.Env = macaron.DEV
// 	m.Run()
// }

// func r2(ctx *macaron.Context) {
// 	ctx.JSON(200, newMsg("This is toukii,r2", 0.101))
// }
