package ui

import (
	"github.com/mikestefanello/pagoda/pkg/page"
)

type (
	Posts struct {
		Posts []Post
		Pager page.Pager
	}
	Post struct {
		Title, Body string
	}
)
