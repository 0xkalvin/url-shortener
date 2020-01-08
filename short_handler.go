package main

import (
	"github.com/astaxie/beego"
	"encoding/json"
)

type ShortController struct {
	beego.Controller
}


type ShortResponse struct {
	Longurl string
	Shorturl string
}


func generateUniqueIdentifier(longurl string) string {

	id := ""

	return id
}


func (this *ShortController) Post() {

	var result ShortResponse
	
	json.Unmarshal(this.Ctx.Input.RequestBody, &result)
	
	beego.Info(result)

	result.Shorturl = generateUniqueIdentifier(result.Longurl)

	this.Data["json"] = &result
	this.ServeJSON()
}
