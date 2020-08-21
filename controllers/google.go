package controllers

import (
	"context"
	"errors"
	"fmt"
	"github.com/physpeach/oauth2-demo/lib"
	"github.com/astaxie/beego"
	"google.golang.org/api/oauth2/v2"
)

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
	conf := google.GetConnect()
	url := conf.AuthCodeURL("state")
	fmt.Printf("Visit the URL for the auth dialog: %v", url)
	c.Redirect(url, 302)
}

type CallbackRequest struct {
	Code  string `form:"code"`
	State string `form:"state"`
}
func (c *GoogleController) Create() {
	request := CallbackRequest{}
	if err := c.ParseForm(&request); err != nil {
		panic(err)
	}
	conf := google.GetConnect()
	ctx := context.Background()
	tok, err := conf.Exchange(ctx, request.Code)
	if err != nil {
		panic(err)
	}
	if tok.Valid() == false {
		panic(errors.New("vaild token"))
	}
	service, err := oauth2.New(conf.Client(ctx, tok))
	if err != nil {
		panic(err)
	}
	tokenInfo, err := service.Tokeninfo().AccessToken(tok.AccessToken).Context(ctx).Do()
	if err != nil {
		panic(err)
	}
	fmt.Println(tokenInfo)

	c.Data["ID"] = tokenInfo.UserId
	c.Data["Email"] = tokenInfo.Email
	c.TplName = "google/callback.tpl"
}

func (c *GoogleController) Put() {

}

func (c *GoogleController) Delete() {

}
