package routers

import (
	"github.com/physpeach/oauth2-demo/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	ns := beego.NewNamespace("/google",
		beego.NSRouter("/", &controllers.GoogleController{}, "get:New"),
		beego.NSRouter("callback", &controllers.GoogleController{}, "get:Create"),
	)
	beego.AddNamespace(ns)
}
