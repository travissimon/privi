package models

import (
	_ "github.com/travissimon/go-mvc"
	"github.com/travissimon/rss"
)

type NewsPage struct {
	Username string
	Feeds    []*rss.Feed
}
