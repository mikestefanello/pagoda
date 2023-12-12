package templates

import (
	"embed"
)

//go:embed *
var Templates embed.FS

type (
	Layout string
	Page   string
)

const (
	LayoutMain Layout = "main"
	LayoutAuth Layout = "auth"
)

const (
	PageAbout          Page = "about"
	PageContact        Page = "contact"
	PageError          Page = "error"
	PageForgotPassword Page = "forgot-password"
	PageHome           Page = "home"
	PageLogin          Page = "login"
	PageRegister       Page = "register"
	PageResetPassword  Page = "reset-password"
	PageSearch         Page = "search"
)
