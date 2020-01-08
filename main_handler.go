package main

import (
	"github.com/astaxie/beego"
)



type IndexController struct {
	beego.Controller
}



func (this *IndexController) Get() {
	this.Ctx.Output.Body([]byte("Up and kicking"))
}


