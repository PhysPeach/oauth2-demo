package controllers

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/discord"
	"github.com/najeira/randstr"
)
func init() {
	goth.UseProviders(
		discord.New(
			beego.AppConfig.String("discordClientId"),
			beego.AppConfig.String("discordClientSecret"),
			beego.AppConfig.String("discordRedirectUri"),
			"email", 
			"identify",
		),
	)
}

// DiscordController operations for Google
type DiscordController struct {
	beego.Controller
}

// URLMapping ...
func (c *DiscordController) URLMapping() {
	c.Mapping("New", c.New)
	c.Mapping("Create", c.Create)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

func (c *DiscordController) New() {
	p, err := goth.GetProvider("discord")
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

func (c *DiscordController) Create() {
	p, err := goth.GetProvider("discord")
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
	c.TplName = "discord/callback.tpl"
}

func (c *DiscordController) Put() {

}

func (c *DiscordController) Delete() {

}
