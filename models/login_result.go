package models

import (
	"github.com/travissimon/go-mvc"
)

type LoginSource int

const (
	Session LoginSource = iota
	Form
)

type LoginResult struct {
	IsLoggedIn  bool
	LoginSource LoginSource
	User        *mvc.User
}
