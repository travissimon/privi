package controllers

import (
	"github.com/travissimon/go-mvc"
	"github.com/travissimon/privi/models"
	"github.com/travissimon/privi/views"
	"net/url"
)

// LoginController Handles 'GET' Requests to login page
func LoginController(ctx *mvc.WebContext, params url.Values) mvc.ControllerResult {
	login := new(models.LoginResult)
	login.LoginSource = models.Session
	login.User = ctx.User
	login.IsLoggedIn = (ctx.IsUserLoggedIn())

	wr := views.NewLoginWriter()
	return mvc.Haml(wr, login, ctx)
}

// Handles 'POST's from login page
func LoginPostController(ctx *mvc.WebContext, params url.Values) mvc.ControllerResult {
	redirectUrl := params.Get("RedirectUrl")
	username := params.Get("txtUsername")
	password := params.Get("txtPassword")
	user, err := ctx.Login(username, password)

	login := new(models.LoginResult)
	login.LoginSource = models.Form
	login.User = user
	login.IsLoggedIn = err == nil

	if login.IsLoggedIn {
		newUrl := redirectUrl
		if len(newUrl) == 0 {
			newUrl = "/"
		}
		return mvc.Redirect(newUrl, ctx)
	}

	wr := views.NewLoginWriter()
	return mvc.Haml(wr, login, ctx)
}

// Logs a user out and redirects to the homepage
func LogoutController(ctx *mvc.WebContext, params url.Values) mvc.ControllerResult {
	ctx.Logout()
	return mvc.Redirect("/", ctx)
}

// Handles GETs to New User controller
func NewUserController(ctx *mvc.WebContext, params url.Values) mvc.ControllerResult {
	login := new(models.LoginResult)
	login.LoginSource = models.Session
	login.User = ctx.User
	login.IsLoggedIn = (ctx.IsUserLoggedIn())

	wr := views.NewNewUserWriter()
	return mvc.Haml(wr, login, ctx)

}

// Handles POSTs from New User controller
func NewUserPostController(ctx *mvc.WebContext, params url.Values) mvc.ControllerResult {
	username := params.Get("txtUsername")
	password := params.Get("txtPassword")
	passwordConfirm := params.Get("txtPasswordConfirm")
	email := params.Get("txtRecoveryEmail")

	if password != passwordConfirm {
		login := new(models.LoginResult)
		login.LoginSource = models.Form
		login.IsLoggedIn = false
		login.Error = "Password and confirmation don't match"
	}

	user, err := ctx.CreateUser(username, password, email)

	login := new(models.LoginResult)
	login.LoginSource = models.Form
	login.User = user
	login.IsLoggedIn = err == nil

	wr := views.NewLoginWriter()
	return mvc.Haml(wr, login, ctx)
}
