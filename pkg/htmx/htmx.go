package htmx

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Request headers: https://htmx.org/docs/#request-headers
const (
	HeaderBoosted               = "HX-Boosted"
	HeaderHistoryRestoreRequest = "HX-History-Restore-Request"
	HeaderPrompt                = "HX-Prompt"
	HeaderRequest               = "HX-Request"
	HeaderTarget                = "HX-Target"
	HeaderTrigger               = "HX-Trigger"
	HeaderTriggerName           = "HX-Trigger-Name"
)

// Response headers: https://htmx.org/docs/#response-headers
const (
	HeaderPushURL            = "HX-Push-Url"
	HeaderRedirect           = "HX-Redirect"
	HeaderReplaceURL         = "HX-Replace-Url"
	HeaderRefresh            = "HX-Refresh"
	HeaderTriggerAfterSettle = "HX-Trigger-After-Settle"
	HeaderTriggerAfterSwap   = "HX-Trigger-After-Swap"
)

type (
	// Request contains data that HTMX provides during requests
	Request struct {
		Enabled        bool
		Boosted        bool
		HistoryRestore bool
		Trigger        string
		TriggerName    string
		Target         string
		Prompt         string
	}

	// Response contain data that the server can communicate back to HTMX
	Response struct {
		PushURL            string
		Redirect           string
		Refresh            bool
		ReplaceURL         string
		Trigger            string
		TriggerAfterSwap   string
		TriggerAfterSettle string
		NoContent          bool
	}
)

// GetRequest extracts HTMX data from the request
func GetRequest(ctx echo.Context) Request {
	return Request{
		Enabled:        ctx.Request().Header.Get(HeaderRequest) == "true",
		Boosted:        ctx.Request().Header.Get(HeaderBoosted) == "true",
		Trigger:        ctx.Request().Header.Get(HeaderTrigger),
		TriggerName:    ctx.Request().Header.Get(HeaderTriggerName),
		Target:         ctx.Request().Header.Get(HeaderTarget),
		Prompt:         ctx.Request().Header.Get(HeaderPrompt),
		HistoryRestore: ctx.Request().Header.Get(HeaderHistoryRestoreRequest) == "true",
	}
}

// Apply applies data from a Response to a server response
func (r Response) Apply(ctx echo.Context) {
	if r.PushURL != "" {
		ctx.Response().Header().Set(HeaderPushURL, r.PushURL)
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
	if r.ReplaceURL != "" {
		ctx.Response().Header().Set(HeaderReplaceURL, r.ReplaceURL)
	}
	if r.NoContent {
		ctx.Response().Status = http.StatusNoContent
	}
}
