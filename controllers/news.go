package controllers

import (
	"fmt"
	"github.com/travissimon/go-mvc"
	"github.com/travissimon/privi/models"
	"github.com/travissimon/privi/views"
	"github.com/travissimon/rss"
	"net/url"
	"strconv"
)

var rssEngine *rss.RssEngine = rss.NewRssEngine("rss", "rss", "rss")

// NewsController handles GET requests for /news
func NewsController(ctx *mvc.WebContext, params url.Values) mvc.ControllerResult {
	if ctx == nil {
		return mvc.Error("Context is nil, unable to proceed", ctx)
	}
	if !ctx.IsUserLoggedIn() {
		return mvc.Redirect("/login?RedirectUrl=/news", ctx)
	}

	userId := ctx.User.Id
	feeds, err := rssEngine.GetFeedsForUser(userId)
	if err != nil {
		fmt.Println(err)
	}

	newsPage := &models.NewsPage{
		Username: ctx.User.Username,
		Feeds:    feeds,
	}
	wr := views.NewNewsWriter()
	return mvc.Haml(wr, newsPage, ctx)
}

func NewsEntriesJSONController(ctx *mvc.WebContext, params url.Values) mvc.ControllerResult {
	if ctx == nil {
		return mvc.Error("Context is nil, unable to proceed", ctx)
	}
	if !ctx.IsUserLoggedIn() {
		return mvc.Redirect("/login?RedirectUrl=/news", ctx)
	}

	feedIdStr := params.Get("FeedId")
	if feedIdStr == "" {
		return mvc.Error("Parameter 'FeedId' not found", ctx)
	}

	feedId, err := strconv.ParseInt(feedIdStr, 10, 64)
	if err != nil {
		fmt.Println(err)
		return mvc.Error(err.Error(), ctx)
	}

	entries, err := rssEngine.GetEntriesForFeed(feedId)
	if err != nil {
		fmt.Println(err)
		return mvc.Error(err.Error(), ctx)
	}

	return mvc.Json(entries, ctx)
}
