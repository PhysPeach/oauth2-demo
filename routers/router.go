package routers

import (
	"github.com/physpeach/oauth2-demo/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	googleNs := beego.NewNamespace("/google",
		beego.NSRouter("/", &controllers.GoogleController{}, "get:New"),
		beego.NSRouter("callback", &controllers.GoogleController{}, "get:Create"),
	)
	beego.AddNamespace(googleNs)
	discordNs := beego.NewNamespace("/discord",
		beego.NSRouter("/", &controllers.DiscordController{}, "get:New"),
		beego.NSRouter("callback", &controllers.DiscordController{}, "get:Create"),
	)
	beego.AddNamespace(discordNs)
}
