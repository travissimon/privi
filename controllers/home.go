package controllers

import (
	"github.com/travissimon/go-mvc"
	_ "github.com/travissimon/privi/models"
	"github.com/travissimon/privi/views"
	"net/url"
)

// HomeController handles GET requests for home page
func HomeController(ctx *mvc.WebContext, params url.Values) mvc.ControllerResult {
	wr := views.NewHomeWriter()
	return mvc.Haml(wr, "", ctx)
}
