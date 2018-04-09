package main

import (
	"errors"
	"os"
	"sort"

	"github.com/kataras/iris"

	"github.com/kataras/iris/sessions"

	"github.com/gorilla/securecookie" //encoder/decoder

	"github.com/markbates/goth"
	"github.com/markbates/facebook"
	"github.com/markbates/twitter"
	"github.com/markbates/github"
	"github.com/markbates/gitlab"
)

var SessionsManager *sessions.Sessions

func init() {

	cookieName := "mycustomsessionid"
	hashkey := []byte("hgsoft-hashkey-custom")
	blockkey := []byte("hgsoft-blockkey-custom")
	securecookie := securecookie.New(hashkey,blockkey)

	sessionsManager = sessions.New(sessions.Config{
		Cookie: cokkieName,
		Encode: secureCookie.Encode,
		Decore: secureCookie.Decode,
	})

}

var GetProviderName = func(ctx iris.Context) (string, error) {
	//try to get it from the url param "provider"
	if p := ctx.URLParam("provider"); p != "" {
		return p, nil
	}

	//try to get it from the url PATH parameter ...
	if p := ctx.Params().Get("provider"); p != "" {
		return p, nil
	}

	if p := ctx.Values().GetString("provider"); p != "" {
		return p, nil
	}

	//if not found then return an empty string with the corresponding error
	return "", errors.New("Você precisa selecionar um provedor")

}

//Iniciar processo de autenticação
func BeginAuthHandler(ctx iris.Context) {
	url, err := GetAuthURL(ctx)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.Writef("%v",err)
		return
	}

	ctx.Redirect(url, iris.StatusTemporaryRedirect)
}

//Continua em : https://iris-go.com/v10/recipe#Oauth211
