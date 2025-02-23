package ui

import (
	"fmt"

	"github.com/mikestefanello/pagoda/pkg/pager"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type (
	Posts struct {
		Posts []Post
		Pager pager.Pager
	}

	Post struct {
		Title, Body string
	}

	SearchResult struct {
		Title string
		URL   string
	}

	File struct {
		Name     string
		Size     int64
		Modified string
	}
)

func (p *Posts) render(path string) Node {
	g := make(Group, 0, len(p.Posts))
	for _, post := range p.Posts {
		g = append(g, post.render())
	}

	return Div(
		ID("posts"),
		g,
		Div(
			Class("field is-grouped is-grouped-centered"),
			P(
				Class("control"),
				Button(
					Class("button is-primary"),
					Attr("hx-swap", "outerHTML"),
					Attr("hx-get", fmt.Sprintf("%s?%s=%d", path, pager.QueryKey, p.Pager.Page-1)),
					Attr("hx-target", "#posts"),
					Text("Previous page"),
				),
			),
			P(
				Class("control"),
				Button(
					Class("button is-primary"),
					Attr("hx-swap", "outerHTML"),
					Attr("hx-get", fmt.Sprintf("%s?%s=%d", path, pager.QueryKey, p.Pager.Page+1)),
					Attr("hx-target", "#posts"),
					Text("Next page"),
				),
			),
		),
	)
}

func (p *Post) render() Node {
	return Article(
		Class("media"),
		Figure(
			Class("media-left"),
			P(
				Class("image is-64x64"),
				Img(
					Src(file("gopher.png")),
					Alt("Gopher"),
				),
			),
		),
		Div(
			Class("media-content"),
			Div(
				Class("content"),
				P(
					Strong(
						Text(p.Title),
					),
					Br(),
					Text(p.Body),
				),
			),
		),
	)
}

func (s *SearchResult) render() Node {
	return A(
		Class("panel-block"),
		Href(s.URL),
		Text(s.Title),
	)
}

func (f *File) render() Node {
	return Tr(
		Td(Text(f.Name)),
		Td(Text(fmt.Sprint(f.Size))),
		Td(Text(f.Modified)),
	)
}
