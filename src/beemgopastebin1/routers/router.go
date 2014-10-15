package routers

import (
	"beemgopastebin1/conf"
	"beemgopastebin1/controllers"
	"github.com/astaxie/beego"
	//	"model"
	"github.com/ozonesurfer/pastemgomodel/src/model"
)

func init() {
	model.Init(conf.DATABASE)
	beego.Router("/", &controllers.MainController{})
	beego.Router("/new", &controllers.CreateController{})
	beego.Router("/paste/:id", &controllers.ShowController{})
}
