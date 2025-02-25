package ui

import (
	"fmt"
	"strings"

	"github.com/mikestefanello/pagoda/pkg/msg"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type tab struct {
	title, body string
}

func head(r *request) Node {
	return Head(
		Meta(Charset("utf-8")),
		Meta(Name("viewport"), Content("width=device-width, initial-scale=1")),
		Link(Rel("icon"), Href(file("favicon.png"))),
		TitleEl(Text(appName), If(r.Title != "", Text(" | "+r.Title))),
		If(r.Metatags.Description != "", Meta(Name("description"), Content(r.Metatags.Description))),
		If(len(r.Metatags.Keywords) > 0, Meta(Name("keywords"), Content(strings.Join(r.Metatags.Keywords, ", ")))),
		Link(Href("https://cdn.jsdelivr.net/npm/bulma@1.0.2/css/bulma.min.css"), Rel("stylesheet")),
		Script(Src("https://unpkg.com/htmx.org@2.0.0/dist/htmx.min.js")),
		Script(Src("https://unpkg.com/alpinejs@3.x.x/dist/cdn.min.js"), Defer()),
	)
}

func flashMessages(r *request) Node {
	var g Group
	for _, typ := range []msg.Type{
		msg.TypeSuccess,
		msg.TypeInfo,
		msg.TypeWarning,
		msg.TypeDanger,
	} {
		for _, str := range msg.Get(r.Context, typ) {
			g = append(g, notification(typ, str))
		}
	}

	return g
}

func notification(typ msg.Type, text string) Node {
	var class string

	switch typ {
	case msg.TypeSuccess:
		class = "success"
	case msg.TypeInfo:
		class = "info"
	case msg.TypeWarning:
		class = "warning"
	case msg.TypeDanger:
		class = "danger"
	}

	return Div(
		Class("notification is-"+class),
		Attr("x-data", "{show: true}"),
		Attr("x-show", "show"),
		Button(
			Class("delete"),
			Attr("@click", "show = false"),
		),
		Text(text),
	)
}

func message(class, header string, body Node) Node {
	return Article(
		Class("message "+class),
		If(header != "", Div(
			Class("message-header"),
			P(Text(header)),
		)),
		Div(
			Class("message-body"),
			body,
		),
	)
}

func sidebarMenu(r *request) Node {
	return Aside(
		Class("menu"),
		hxBoost(),
		P(
			Class("menu-label"),
			Text("General"),
		),
		Ul(
			Class("menu-list"),
			menuLink(r, "Dashboard", routenames.Home),
			menuLink(r, "About", routenames.About),
			menuLink(r, "Contact", routenames.Contact),
			menuLink(r, "Cache", routenames.Cache),
			menuLink(r, "Task", routenames.Task),
			menuLink(r, "Files", routenames.Files),
		),
		P(
			Class("menu-label"),
			Text("Account"),
		),
		Ul(
			Class("menu-list"),
			If(r.IsAuth, menuLink(r, "Logout", routenames.Logout)),
			If(!r.IsAuth, menuLink(r, "Login", routenames.Login)),
			If(!r.IsAuth, menuLink(r, "Register", routenames.Register)),
			If(!r.IsAuth, menuLink(r, "Forgot password", routenames.ForgotPasswordSubmit)),
		),
	)
}

func menuLink(r *request, title, routeName string, routeParams ...string) Node {
	href := r.path(routeName, routeParams...)

	return Li(
		A(
			Href(href),
			Text(title),
			If(href == r.Path, Class("is-active")),
		),
	)
}

func authNavBar(r *request) Node {
	return Nav(
		Class("navbar"),
		Div(
			Class("navbar-menu"),
			Div(
				Class("navbar-start"),
				A(Class("navbar-item"), Href(r.path(routenames.Login)), Text("Login")),
				A(Class("navbar-item"), Href(r.path(routenames.Register)), Text("Create an account")),
				A(Class("navbar-item"), Href(r.path(routenames.ForgotPassword)), Text("Forgot password")),
			),
		),
	)
}

func navBar(r *request) Node {
	return Nav(
		Class("navbar is-dark"),
		Div(
			Class("container"),
			Div(
				Class("navbar-brand"),
				hxBoost(),
				A(
					Href(r.path(routenames.Home)),
					Class("navbar-item"),
					Text("Pagoda"),
				),
			),
			Div(
				ID("navbarMenu"),
				Class("navbar-menu"),
				Div(
					Class("navbar-end"),
					search(r),
				),
			),
		),
	)
}

func search(r *request) Node {
	return Div(
		Class("search mr-2 mt-1"),
		Attr("x-data", "{modal:false}"),
		Input(
			Class("input"),
			Type("search"),
			Placeholder("Search..."),
			Attr("@click", "modal = true; $nextTick(() => $refs.input.focus());"),
		),
		Div(
			Class("modal"),
			Attr(":class", "modal ? 'is-active' : ''"),
			Attr("x-show", "modal == true"),
			Div(
				Class("modal-background"),
			),
			Div(
				Class("modal-content"),
				Attr("@click.outside", "modal = false;"),
				Div(
					Class("box"),
					H2(
						Class("subtitle"),
						Text("Search"),
					),
					P(
						Class("control"),
						Input(
							Attr("hx-get", r.path(routenames.Search)),
							Attr("hx-trigger", "keyup changed delay:500ms"),
							Attr("hx-target", "#results"),
							Name("query"),
							Class("input"),
							Type("search"),
							Placeholder("Search..."),
							Attr("x-ref", "input"),
						),
					),
					Div(
						Class("block"),
					),
					Div(
						ID("results"),
					),
				),
			),
			Button(
				Class("modal-close is-large"),
				Aria("label", "close"),
			),
		),
	)
}

func footer(r *request) Node {
	return Group{
		If(len(r.CSRF) > 0, Script(
			Raw(`
  			document.body.addEventListener('htmx:configRequest', function(evt)  {
                if (evt.detail.verb !== "get") {
                    evt.detail.parameters['csrf'] = '`+r.CSRF+`';
                }
            })
			`),
		)),
		Script(Raw(`
		document.body.addEventListener('htmx:beforeSwap', function(evt) {
            if (evt.detail.xhr.status >= 400){
                evt.detail.shouldSwap = true;
                evt.detail.target = htmx.find("body");
            }
        });
		`)),
	}
}

func button(class, label string) Node {
	return Button(
		Class("button "+class),
		Text(label),
	)
}

func buttonLink(href, class, label string) Node {
	return A(
		Href(href),
		Class("button "+class),
		Text(label),
	)
}

func hxBoost() Node {
	return Attr("hx-boost", "true")
}

func tabs(heading, description string, items []tab) Node {
	renderTitles := func() Node {
		g := make(Group, 0, len(items))
		for i, item := range items {
			g = append(g, Li(
				Attr(":class", fmt.Sprintf("{'is-active': tab === %d}", i)),
				Attr("@click", fmt.Sprintf("tab = %d", i)),
				A(Text(item.title)),
			))
		}
		return g
	}

	renderBodies := func() Node {
		g := make(Group, 0, len(items))
		for i, item := range items {
			g = append(g, Div(
				Attr("x-show", fmt.Sprintf("tab == %d", i)),
				P(Raw(" "+item.body)),
			))
		}
		return g
	}

	return Div(
		P(
			Class("subtitle mt-5"),
			Text(heading),
		),
		P(
			Class("mb-4"),
			Text(description),
		),
		Div(
			Attr("x-data", "{tab: 0}"),
			Div(
				Class("tabs"),
				Ul(renderTitles()),
			),
			renderBodies(),
		),
	)
}
