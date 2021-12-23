package controller

import (
	"github.com/labstack/echo/v4"
)

// HTMX headers (https://htmx.org/docs/#requests)
const (
	HTMXHeaderRequest            = "HX-Request"
	HTMXHeaderTrigger            = "HX-Trigger"
	HTMXHeaderTriggerName        = "HX-Trigger-Name"
	HTMXHeaderTriggerAfterSwap   = "HX-Trigger-After-Swap"
	HTMXHeaderTriggerAfterSettle = "HX-Trigger-After-Settle"
	HTMXHeaderTarget             = "HX-Target"
	HTMXHeaderPrompt             = "HX-Prompt"
	HTMXHeaderPush               = "HX-Push"
	HTMXHeaderRedirect           = "HX-Redirect"
	HTMXHeaderRefresh            = "HX-Refresh"
)

type (
	HTMXRequest struct {
		Enabled     bool
		Trigger     string
		TriggerName string
		Target      string
		Prompt      string
	}

	HTMXResponse struct {
		Push               string
		Redirect           string
		Refresh            bool
		Trigger            string
		TriggerAfterSwap   string
		TriggerAfterSettle string
		// TODO: No content 204 response?
	}
)

func GetHTMXRequest(ctx echo.Context) HTMXRequest {
	return HTMXRequest{
		Enabled:     ctx.Request().Header.Get(HTMXHeaderRequest) == "true",
		Trigger:     ctx.Request().Header.Get(HTMXHeaderTrigger),
		TriggerName: ctx.Request().Header.Get(HTMXHeaderTriggerName),
		Target:      ctx.Request().Header.Get(HTMXHeaderTarget),
		Prompt:      ctx.Request().Header.Get(HTMXHeaderPrompt),
	}
}
