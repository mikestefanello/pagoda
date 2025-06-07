package models

import (
	"fmt"

	"github.com/mikestefanello/pagoda/pkg/pager"
	"github.com/mikestefanello/pagoda/pkg/ui"
	. "github.com/mikestefanello/pagoda/pkg/ui/components"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type (
	Posts struct {
		Posts []Post
		Pager pager.Pager
	}

	Post struct {
		ID          int
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
		Ul(
			Class("list bg-base-100 rounded-box shadow-md not-prose"),
			g,
		),
		Div(Class("mb-4")),
		Pager(p.Pager.Page, path, !p.Pager.IsEnd(), "#posts"),
	)
}

func (p *Post) Render() Node {
	return Li(
		Class("list-row"),
		Div(
			Class("text-4xl font-thin opacity-30 tabular-nums"),
			Text(fmt.Sprintf("%02d", p.ID)),
		),
		Div(
			Img(
				Class("size-10 rounded-box"),
				Src(ui.StaticFile("gopher.png")),
				Alt("Gopher"),
			),
		),
		Div(
			Class("list-col-grow"),
			Div(
				Text(p.Title),
			),
			Div(
				Class("text-xs font-semibold opacity-60"),
				Text(p.Body),
			),
		),
	)
}
