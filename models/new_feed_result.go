package models

import (
	"github.com/travissimon/rss"
)

type NewEntryResponse struct {
	FeedId       int64
	FeedName     string
	Entries      []*rss.Entry
	ErrorMessage string
}
