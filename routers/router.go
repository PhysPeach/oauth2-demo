package routers

import (
	"github.com/physpeach/oauth2-demo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
