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
