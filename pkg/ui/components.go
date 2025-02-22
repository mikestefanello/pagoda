package ui

import (
	"strings"

	"github.com/mikestefanello/pagoda/pkg/msg"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func head(r *request) Node {
	return Head(
		Meta(Charset("utf-8")),
		Meta(Name("viewport"), Content("width=device-width, initial-scale=1")),
		Link(Rel("icon"), Href(file("favicon.png"))),
		TitleEl(Text(r.Title)),
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
		Attr("hx-boost", "true"),
		P(
			Class("menu-label"),
			Text("General"),
		),
		Ul(
			Class("menu-list"),
			menuLink(r, "Dashboard", routenames.Home),
			menuLink(r, "About", "about"),
			menuLink(r, "Contact", routenames.Contact),
			menuLink(r, "Cache", "cache"),
			menuLink(r, "Task", "task"),
			menuLink(r, "Files", "files"),
		),
		P(
			Class("menu-label"),
			Text("Account"),
		),
		Ul(
			Class("menu-list"),
			If(r.IsAuth, menuLink(r, "Logout", "logout")),
			If(!r.IsAuth, menuLink(r, "Login", "login")),
			If(!r.IsAuth, menuLink(r, "Register", "register")),
			If(!r.IsAuth, menuLink(r, "Forgot password", "forgot_password")),
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

func navBar(r *request) Node {
	return Nav(
		Class("navbar is-dark"),
		Div(
			Class("container"),
			Div(
				Class("navbar-brand"),
				Attr("hx-boost", "true"),
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
							Attr("hx-get", r.path("search")),
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
