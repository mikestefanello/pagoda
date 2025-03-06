package emails

import (
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/ui"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func ConfirmEmailAddress(ctx echo.Context, username, token string) Node {
	url := ui.NewRequest(ctx).
		Url(routenames.VerifyEmail, token)

	return Group{
		Strong(Textf("Hello %s,", username)),
		Br(),
		P(Text("Please click on the following link to confirm your email address:")),
		Br(),
		A(Href(url), Text(url)),
	}
}
