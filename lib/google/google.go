package google

import (
	"golang.org/x/oauth2"
	"github.com/astaxie/beego"
)

func GetConnect() *oauth2.Config {
	conf := &oauth2.Config{
		ClientID: beego.AppConfig.String("googleClientId"),
		ClientSecret: beego.AppConfig.String("googleClientSecret"),
		Endpoint: oauth2.Endpoint{
			AuthURL:   beego.AppConfig.String("googleAuthUri"),
			TokenURL:  beego.AppConfig.String("googleTokenUri"),
		},
		Scopes: []string{"openid", "email", "profile"},
		RedirectURL: beego.AppConfig.String("googleRedirectUri"),
	}

	return conf
}