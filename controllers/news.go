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

func FeedStubsForUserController(ctx *mvc.WebContext, params url.Values) mvc.ControllerResult {
	if ctx == nil {
		return mvc.Error("Context is nil, unable to proceed", ctx)
	}
	if !ctx.IsUserLoggedIn() {
		return mvc.Redirect("/login?RedirectUrl=/news", ctx)
	}

	userId := ctx.User.Id
	feeds, err := rssEngine.GetFeedStubsForUser(userId)
	if err != nil {
		fmt.Println(err)
	}

	return mvc.Json(feeds, ctx)
}

// NewsEntriesJSONController handles JSON requests for entries for a given feed
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

	fmt.Printf("Adding URL: %s\n", feedUrl)

	feed, entries, err := rssEngine.AddFeedForUser(ctx.User.Id, feedUrl)

	if feed == nil {
		fmt.Printf("Feed is nil\n")
	} else {
		fmt.Printf("AddFeed: feed is: %d - %s\n", feed.Id, feed.Title)
	}

	if entries == nil {
		fmt.Printf("Entries are nil\n")
	} else {
		fmt.Printf("Entries parsed: %d\n", len(entries))
	}

	if err != nil {
		fmt.Println(err)
	}

	newEntry := &models.NewEntryResponse{}
	if feed == nil {
		newEntry.ErrorMessage = "Returned nil feed"
	} else if err == nil {
		newEntry.FeedId = feed.Id
		newEntry.FeedName = feed.Title
		newEntry.Entries = entries
	} else {
		newEntry.ErrorMessage = err.Error()
	}

	fmt.Printf("Returning: %v\n", newEntry)

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
