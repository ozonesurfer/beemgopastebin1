package controllers

import (
	"github.com/astaxie/beego"
	"github.com/ozonesurfer/pastemgomodel/src/model"
	"html/template"
	"log"
	//	"model"
)

type MainController struct {
	beego.Controller
}
type CreateController MainController
type ShowController MainController

func (this *MainController) Get() {
	/*	this.Data["Website"] = "beego.me"
		this.Data["Email"] = "astaxie@gmail.com"
	*/
	this.Data["Pastes"] = model.GetAll()
	this.Data["Languages"] = model.Languages
	this.TplNames = "index.tpl"
}

func (this *CreateController) Post() {
	paste := model.Paste{
		Title:   this.GetString("title"),
		Content: this.GetString("content"),
		// LanguageId: this.GetInt("language"),
	}
	paste.LanguageId, _ = this.GetInt("language")
	_, _, now := paste.Add()
	paste.CreatedOn = now
	content, err := paste.HighlightKeywords()
	if err != nil {
		log.Fatalln("Display error:", err)
	}
	this.Data["Paste"] = paste
	this.Data["Content"] = template.HTML(content)
	this.Data["Language"] = model.Languages[paste.LanguageId].Name
	this.TplNames = "create.tpl"
}

func (this *ShowController) Get() {
	rawId := this.Ctx.Input.Param(":id")
	id := model.ToObjectId(rawId)
	paste := model.GetPaste(id)
	content, err := paste.HighlightKeywords()
	if err != nil {
		log.Fatalln("Display error:", err)
	}
	this.Data["Paste"] = paste
	this.Data["Content"] = template.HTML(content)
	this.Data["Language"] = model.Languages[paste.LanguageId].Name
	this.TplNames = "create.tpl"
}
