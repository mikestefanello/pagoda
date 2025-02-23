package ui

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Home(ctx echo.Context, posts Posts) error {
	r := newRequest(ctx)
	r.Metatags.Description = "This is the home page."
	r.Metatags.Keywords = []string{"Software", "Coding", "Go"}

	g := make(Group, 0)

	if r.Htmx.Target != "posts" {
		var hello string
		if r.IsAuth {
			hello = fmt.Sprintf("Hello, %s", r.AuthUser.Name)
		} else {
			hello = "Hello"
		}

		g = append(g,
			Section(
				Class("hero is-info welcome is-small mb-5"),
				Div(
					Class("hero-body"),
					Div(
						Class("container"),
						H1(
							Class("title"),
							Text(hello),
						),
						H2(
							Class("subtitle"),
							If(!r.IsAuth, Text("Please login in to your account.")),
							If(r.IsAuth, Text("Welcome back!")),
						),
					),
				),
			),
			H2(Class("title"), Text("Recent posts")),
			H3(Class("subtitle"), Text("Below is an example of both paging and AJAX fetching using HTMX")),
		)
	}

	g = append(g, posts.render(r.path(routenames.Home)))

	if r.Htmx.Target != "posts" {
		g = append(g, message(
			"is-small is-warning mt-5",
			"Serving files",
			Group{
				Text("In the example posts above, check how the file URL contains a cache-buster query parameter which changes only when the app is restarted."),
				Text("Static files also contain cache-control headers which are configured via middleware."),
			},
		))
	}

	return r.render(layoutPrimary, g)
}

func ContactUs(ctx echo.Context, form *ContactForm) error {
	r := newRequest(ctx)
	r.Title = "Contact us"
	r.Metatags.Description = "Get in touch with us."

	g := make(Group, 0)

	if r.Htmx.Target != "contact" {
		g = append(g, message(
			"is-link",
			"",
			Group{
				P(Text("This is an example of a form with inline, server-side validation and HTMX-powered AJAX submissions without writing a single line of JavaScript.")),
				P(Text("Only the form below will update async upon submission.")),
			}))
	}

	if form.IsDone() {
		g = append(g, message(
			"is-large is-success",
			"Thank you!",
			Text("No email was actually sent but this entire operation was handled server-side and degrades without JavaScript enabled."),
		))
	} else {
		g = append(g, form.render(r))
	}

	return r.render(layoutPrimary, g)
}

func Login(ctx echo.Context, form *LoginForm) error {
	r := newRequest(ctx)
	r.Title = "Login"

	return r.render(layoutAuth, form.render(r))
}

func Register(ctx echo.Context, form *RegisterForm) error {
	r := newRequest(ctx)
	r.Title = "Register"

	return r.render(layoutAuth, form.render(r))
}

func ForgotPassword(ctx echo.Context, form *ForgotPasswordForm) error {
	r := newRequest(ctx)
	r.Title = "Forgot password"

	g := Group{
		Div(
			Class("content"),
			P(Text("Enter your email address and we'll email you a link that allows you to reset your password.")),
		),
		form.render(r),
	}

	return r.render(layoutAuth, g)
}

func ResetPassword(ctx echo.Context, form *ResetPasswordForm) error {
	r := newRequest(ctx)
	r.Title = "Reset your password"

	return r.render(layoutAuth, form.render(r))
}

func SearchResults(ctx echo.Context, results []*SearchResult) error {
	r := newRequest(ctx)

	g := make(Group, 0, len(results))
	for _, result := range results {
		g = append(g, result.render())
	}

	return r.render(layoutPrimary, g)
}

func AddTask(ctx echo.Context, form *TaskForm) error {
	r := newRequest(ctx)
	r.Title = "Create a task"
	r.Metatags.Description = "Test creating a task to see how it works."

	g := make(Group, 0)

	if r.Htmx.Target != "task" {
		g = append(g, message(
			"is-link",
			"",
			Group{
				P(Raw("Submitting this form will create an <i>ExampleTask</i> in the task queue. After the specified delay, the message will be logged by the queue processor.")),
				P(Text("See pkg/tasks and the README for more information.")),
			}))
	}

	g = append(g, form.render(r))

	return r.render(layoutPrimary, g)
}

func UpdateCache(ctx echo.Context, form *CacheForm) error {
	r := newRequest(ctx)
	r.Title = "Set a cache entry"

	return r.render(layoutPrimary, form.render(r))
}

func Error(ctx echo.Context, code int) error {
	r := newRequest(ctx)
	r.Title = http.StatusText(code)
	var body Node

	switch code {
	case http.StatusInternalServerError:
		body = Text("Please try again.")
	case http.StatusForbidden, http.StatusUnauthorized:
		body = Text("You are not authorized to view the requested page.")
	case http.StatusNotFound:
		body = Group{
			Text("Click "),
			A(
				Href(r.path(routenames.Home)),
				Text("here"),
			),
			Text(" to go return home."),
		}
	default:
		body = Text("Something went wrong.")
	}

	return r.render(layoutPrimary, P(body))
}

func UploadFile(ctx echo.Context, files []*File) error {
	r := newRequest(ctx)
	r.Title = "Upload a file"

	fileList := make(Group, 0, len(files))
	for _, file := range files {
		fileList = append(fileList, file.render())
	}

	n := Group{
		message(
			"is-link",
			"",
			P(Text("This is a very basic example of how to handle file uploads. Files uploaded will be saved to the directory specified in your configuration.")),
		),
		Hr(),
		Form(
			ID("files"),
			Method(http.MethodPost),
			Action(r.path(routenames.FilesSubmit)),
			EncType("multipart/form-data"),
			formFile("file", "Choose a file.. "),
			formControlGroup(
				button("is-link", "Upload"),
			),
			csrf(r),
		),
		Hr(),
		H3(
			Class("title"),
			Text("Uploaded files"),
		),
		message("is-warning", "", P(Text("Below are all files in the configured upload directory."))),
		Table(
			Class("table"),
			THead(
				Tr(
					Th(Text("Filename")),
					Th(Text("Size")),
					Th(Text("Modified on")),
				),
			),
			TBody(
				fileList,
			),
		),
	}

	return r.render(layoutPrimary, n)
}

func About(ctx echo.Context) error {
	r := newRequest(ctx)
	r.Title = "About"
	r.Metatags.Description = "Learn a little about what's included in Pagoda."

	return r.render(layoutPrimary, Group{
		tabs(
			"Frontend",
			"The following incredible projects make developing advanced, modern frontends possible and simple without having to write a single line of JS or CSS. You can go extremely far without leaving the comfort of Go with server-side rendered HTML.",
			[]tab{
				{
					title: "HTMX",
					body:  "Completes HTML as a hypertext by providing attributes to AJAXify anything and much more. Visit <a href=\"https://htmx.org/\">htmx.org</a> to learn more.",
				},
				{
					title: "Alpine.js",
					body:  "Drop-in, Vue-like functionality written directly in your markup. Visit <a href=\"https://alpinejs.dev/\">alpinejs.dev</a> to learn more.",
				},
				{
					title: "Bulma",
					body:  "Ready-to-use frontend components that you can easily combine to build responsive web interfaces with no JavaScript requirements. Visit <a href=\"https://bulma.io/\">bulma.io</a> to learn more.",
				},
			},
		),
		Div(Class("mb-4")),
		tabs(
			"Backend",
			"The following incredible projects provide the foundation of the Go backend. See the repository for a complete list of included projects.",
			[]tab{
				{
					title: "Echo",
					body:  "High performance, extensible, minimalist Go web framework. Visit <a href=\"https://echo.labstack.com/\">echo.labstack.com</a> to learn more.",
				},
				{
					title: "Ent",
					body:  "Simple, yet powerful ORM for modeling and querying data. Visit <a href=\"https://entgo.io/\">entgo.io</a> to learn more.",
				},
			},
		),
	})
}
