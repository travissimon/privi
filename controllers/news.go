package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/travissimon/go-mvc"
	"github.com/travissimon/privi/models"
	"github.com/travissimon/privi/views"
	"github.com/travissimon/rss"
	"net/http"
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
		return mvc.Error(err.Error(), ctx)
	}

	entries, err := rssEngine.GetEntriesForFeed(feedId)
	if err != nil {
		fmt.Println(err)
		return mvc.Error(err.Error(), ctx)
	}

	return mvc.Json(entries, ctx)
}

func AddFeedJSONController(ctx *mvc.WebContext, params url.Values) mvc.ControllerResult {
	if !ctx.IsUserLoggedIn() {
		return mvc.Redirect("/login?RedirectUrl=/news", ctx)
	}

	json, err := ParseJson(ctx.Request)
	if err != nil {
		return mvc.Error(err.Error(), ctx)
	}
	feedIdInterface := json["feedUrl"]
	if feedIdInterface == nil {
		return mvc.Error("parameter 'feedUrl' not found", ctx)
	}
	feedUrl := feedIdInterface.(string)

	_, err = rssEngine.AddFeedForUser(ctx.User.Id, feedUrl)

	if err != nil {
		fmt.Println(err)
	}

	newEntry := &models.NewEntryResponse{}
	if feedUrl == "" {
		newEntry.ErrorMessage = "Error, feedUrl is missing"
	} else {
		newEntry.ErrorMessage = "You posted: " + feedUrl
	}

	return mvc.Json(newEntry, ctx)
}

func ParseJson(r *http.Request) (data map[string]interface{}, err error) {
	var d interface{}
	e := json.NewDecoder(r.Body).Decode(&d)
	if e != nil {
		return nil, err
	}
	return d.(map[string]interface{}), nil
}
