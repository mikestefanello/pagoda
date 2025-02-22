package ui

import (
	"github.com/labstack/echo/v4"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Home(ctx echo.Context, posts Posts) error {
	r := newRequest(ctx)
	r.Title = "Home"
	r.Metatags.Description = "This is the home page."
	r.Metatags.Keywords = []string{"Software", "Coding", "Go"}

	g := make(Group, 0, len(posts.Posts)+1)
	g = append(g, Div(Class("dashboard"), Text("hello")))

	for _, post := range posts.Posts {
		g = append(g, Div(
			Class("post"),
			H3(Text(post.Title)),
			Span(Text(post.Body)),
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
