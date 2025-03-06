package models

import (
	"fmt"

	"github.com/mikestefanello/pagoda/pkg/pager"
	"github.com/mikestefanello/pagoda/pkg/ui"
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
)

func (p *Posts) Render(path string) Node {
	g := make(Group, len(p.Posts))
	for i, post := range p.Posts {
		g[i] = post.Render()
	}

	return Div(
		ID("posts"),
		g,
		Div(
			Class("field is-grouped is-grouped-centered"),
			If(!p.Pager.IsBeginning(), P(
				Class("control"),
				Button(
					Class("button is-primary"),
					Attr("hx-swap", "outerHTML"),
					Attr("hx-get", fmt.Sprintf("%s?%s=%d", path, pager.QueryKey, p.Pager.Page-1)),
					Attr("hx-target", "#posts"),
					Text("Previous page"),
				),
			)),
			If(!p.Pager.IsEnd(), P(
				Class("control"),
				Button(
					Class("button is-primary"),
					Attr("hx-swap", "outerHTML"),
					Attr("hx-get", fmt.Sprintf("%s?%s=%d", path, pager.QueryKey, p.Pager.Page+1)),
					Attr("hx-target", "#posts"),
					Text("Next page"),
				),
			)),
		),
	)
}

func (p *Post) Render() Node {
	return Article(
		Class("media"),
		Figure(
			Class("media-left"),
			P(
				Class("image is-64x64"),
				Img(
					Src(ui.File("gopher.png")),
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
