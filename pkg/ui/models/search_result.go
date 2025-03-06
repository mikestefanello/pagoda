package models

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type SearchResult struct {
	Title string
	URL   string
}

func (s *SearchResult) Render() Node {
	return A(
		Class("panel-block"),
		Href(s.URL),
		Text(s.Title),
	)
}
