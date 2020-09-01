package controllers

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"github.com/najeira/randstr"
)
func init() {
	goth.UseProviders(
		google.New(
			beego.AppConfig.String("googleClientId"),
			beego.AppConfig.String("googleClientSecret"),
			beego.AppConfig.String("googleRedirectUri"),
			"openid",
			"email",
			"profile",
		),
	)
}

// GoogleController operations for Google
type GoogleController struct {
	beego.Controller
}

// URLMapping ...
func (c *GoogleController) URLMapping() {
	c.Mapping("New", c.New)
	c.Mapping("Create", c.Create)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

func (c *GoogleController) New() {
	p, err := goth.GetProvider("google")
	if err != nil{
		panic(err)
	}
	state := randstr.CryptoString(16)
	sess, err := p.BeginAuth(state)
	if err != nil{
		panic(err)
	}
	c.SetSession("goth", sess)
	c.SetSession("state", state)
	url, err := sess.GetAuthURL()
	if err != nil{
		panic(err)
	}
	fmt.Printf("Visit the URL for the auth dialog: %v", url)
	c.Redirect(url, 302)
}

func (c *GoogleController) Create() {
	p, err := goth.GetProvider("google")
	if err != nil{
		panic(err)
	}
	sess := c.GetSession("goth").(goth.Session)
	params := c.Ctx.Request.URL.Query()
	if params.Get("state") != c.GetSession("state"){
		panic(errors.New("wrong state"))
	}
	sess.Authorize(p, params)
	user, err := p.FetchUser(sess)
	if err != nil{
		panic(err)
	}
	c.Data["Name"] = user.Name
	c.Data["Email"] = user.Email
	c.TplName = "google/callback.tpl"
}

func (c *GoogleController) Put() {

}

func (c *GoogleController) Delete() {

}
