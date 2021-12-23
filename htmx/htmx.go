package htmx

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

//  Headers (https://.org/docs/#requests)
const (
	HeaderRequest            = "HX-Request"
	HeaderBoosted            = "HX-Boosted"
	HeaderTrigger            = "HX-Trigger"
	HeaderTriggerName        = "HX-Trigger-Name"
	HeaderTriggerAfterSwap   = "HX-Trigger-After-Swap"
	HeaderTriggerAfterSettle = "HX-Trigger-After-Settle"
	HeaderTarget             = "HX-Target"
	HeaderPrompt             = "HX-Prompt"
	HeaderPush               = "HX-Push"
	HeaderRedirect           = "HX-Redirect"
	HeaderRefresh            = "HX-Refresh"
)

type (
	Request struct {
		Enabled     bool
		Boosted     bool
		Trigger     string
		TriggerName string
		Target      string
		Prompt      string
	}

	Response struct {
		Push               string
		Redirect           string
		Refresh            bool
		Trigger            string
		TriggerAfterSwap   string
		TriggerAfterSettle string
		NoContent          bool
	}
)

func GetRequest(ctx echo.Context) Request {
	return Request{
		Enabled:     ctx.Request().Header.Get(HeaderRequest) == "true",
		Boosted:     ctx.Request().Header.Get(HeaderBoosted) == "true",
		Trigger:     ctx.Request().Header.Get(HeaderTrigger),
		TriggerName: ctx.Request().Header.Get(HeaderTriggerName),
		Target:      ctx.Request().Header.Get(HeaderTarget),
		Prompt:      ctx.Request().Header.Get(HeaderPrompt),
	}
}

func (r *Response) Apply(ctx echo.Context) {
	if r.Push != "" {
		ctx.Response().Header().Set(HeaderPush, r.Push)
	}
	if r.Redirect != "" {
		ctx.Response().Header().Set(HeaderRedirect, r.Redirect)
	}
	if r.Refresh {
		ctx.Response().Header().Set(HeaderRefresh, "true")
	}
	if r.Trigger != "" {
		ctx.Response().Header().Set(HeaderTrigger, r.Trigger)
	}
	if r.TriggerAfterSwap != "" {
		ctx.Response().Header().Set(HeaderTriggerAfterSwap, r.TriggerAfterSwap)
	}
	if r.TriggerAfterSettle != "" {
		ctx.Response().Header().Set(HeaderTriggerAfterSettle, r.TriggerAfterSettle)
	}
	if r.NoContent {
		ctx.Response().Status = http.StatusNoContent
	}
}
