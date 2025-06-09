package pages

import (
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/ui"
	"github.com/mikestefanello/pagoda/pkg/ui/cache"
	. "github.com/mikestefanello/pagoda/pkg/ui/components"
	"github.com/mikestefanello/pagoda/pkg/ui/layouts"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func About(ctx echo.Context) error {
	r := ui.NewRequest(ctx)
	r.Title = "About"
	r.Metatags.Description = "Learn a little about what's included in Pagoda."

	// The tabs are static, so we can render and cache them.
	tabs := cache.SetIfNotExists("pages.about.Tabs", func() Node {
		return Group{
			H2(Text("Frontend")),
			P(Text("The following incredible projects make developing advanced, modern frontends possible and simple without having to write a single line of JS or CSS. You can go extremely far without leaving the comfort of Go with server-side rendered HTML.")),
			Tabs(
				[]Tab{
					{
						Title: "HTMX",
						Body:  "Completes HTML as a hypertext by providing attributes to AJAXify anything and much more. Visit <a href=\"https://htmx.org/\">htmx.org</a> to learn more.",
					},
					{
						Title: "Alpine.js",
						Body:  "Drop-in, Vue-like functionality written directly in your markup. Visit <a href=\"https://alpinejs.dev/\">alpinejs.dev</a> to learn more.",
					},
					{
						Title: "DaisyUI",
						Body:  "DaisyUI is the Tailwind CSS plugin you will love! It provides useful component class names to help you write less code and build faster. No JavaScript requirements. Visit <a href=\"https://daisyui.com/\">daisyui.com</a> to learn more.",
					},
				},
			),
			H2(Text("Backend")),
			P(Text("The following incredible projects provide the foundation of the Go backend. See the repository for a complete list of included projects.")),
			Tabs(
				[]Tab{
					{
						Title: "Echo",
						Body:  "High performance, extensible, minimalist Go web framework. Visit <a href=\"https://echo.labstack.com/\">echo.labstack.com</a> to learn more.",
					},
					{
						Title: "Ent",
						Body:  "Simple, yet powerful ORM for modeling and querying data. Visit <a href=\"https://entgo.io/\">entgo.io</a> to learn more.",
					},
					{
						Title: "Gomponents",
						Body:  "HTML components written in pure Go. They render to HTML 5, and make it easy for you to build reusable components. Visit <a href=\"https://gomponents.com/\">gomponents.com</a> to learn more.",
					},
				},
			),
		}
	})

	return r.Render(layouts.Primary, tabs)
}
