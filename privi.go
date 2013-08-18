package main

import (
	"fmt"
	"github.com/travissimon/go-mvc"
	"github.com/travissimon/privi/controllers"
	_ "html/template"
	"net/http"
	_ "net/url"
)

func main() {
	url := "localhost:8000"
	fmt.Printf("Listening on http://%s\n", url)

	handler := mvc.NewMvcHandler()

	// Set go html-templates
	// handler.SetTemplates(templates)
	// handler.AddRoute("Article", "/Article", mvc.GET, ArticleController)

	// Routes

	// Homepage
	handler.AddRoute("Home", "/", mvc.GET, controllers.HomeController)

	// News
	handler.AddRoute("News", "/news", mvc.GET, controllers.NewsController)
	handler.AddRoute("New Feed Post", "/news/newfeed", mvc.POST, controllers.AddFeedJSONController)
	handler.AddRoute("Entries", "/news/{FeedId}", mvc.GET, controllers.NewsEntriesJSONController)

	// Login/New User
	handler.AddRoute("Login", "/login", mvc.GET, controllers.LoginController)
	handler.AddRoute("Login Post", "/login", mvc.POST, controllers.LoginPostController)
	handler.AddRoute("New User", "/newuser", mvc.GET, controllers.NewUserController)
	handler.AddRoute("New User Post", "/newuser", mvc.POST, controllers.NewUserPostController)
	handler.AddRoute("Logout", "/logout", mvc.GET, controllers.LogoutController)

	// Static Files
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static/"))))
	http.Handle("/", handler)
	http.ListenAndServe(url, nil)
}
