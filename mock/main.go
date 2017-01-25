package main

import (
	"fmt"
	"github.com/go-macaron/macaron"
	"html/template"
)

type Msg struct {
	Message string
	Cost    float64
}

func newMsg(msg string, cost float64) *Msg {
	return &Msg{
		Message: msg,
		Cost:    cost,
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
			fmt.Println("This is toukii,root")
			ctx.JSON(200, newMsg("This is toukii,root", 3.14))
		})
		m.Get("r1", func(ctx *macaron.Context) {
			fmt.Println("This is toukii,r1")
			ctx.JSON(201, newMsg("This is toukii,r1", 0.315))
		})
		m.Combo("r2").Post(r2)
	})
	// macaron.Env = macaron.PROD
	macaron.Env = macaron.DEV
	m.Run()
}

func r2(ctx *macaron.Context) {
	fmt.Println("This is toukii,r2")
	ctx.JSON(200, newMsg("This is toukii,r2", 0.101))
}
