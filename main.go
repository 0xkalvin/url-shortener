package main

import (
	"github.com/astaxie/beego"
)

func main() {
	beego.Router("/", &IndexController{})
	beego.Router("/short", &ShortController{})
	beego.Run()
}