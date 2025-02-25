package pages

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/ui"
	"github.com/mikestefanello/pagoda/pkg/ui/layouts"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Error(ctx echo.Context, code int) error {
	r := ui.NewRequest(ctx)
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
				Href(r.Path(routenames.Home)),
				Text("here"),
			),
			Text(" to go return home."),
		}
	default:
		body = Text("Something went wrong.")
	}

	return r.Render(layouts.Primary, P(body))
}
