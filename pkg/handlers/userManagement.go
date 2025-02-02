package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/ent/user"
	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/page"
	"github.com/mikestefanello/pagoda/templates"
	"net/http"

	"github.com/mikestefanello/pagoda/pkg/middleware"
	"github.com/mikestefanello/pagoda/pkg/services"
)

const (
	routeNameUserDetails = "userDetails"
	routeNameUpdateUser  = "updateUser"
)

type (
	UserManagement struct {
		mail *services.MailClient
		orm  *ent.Client
		*services.TemplateRenderer
		auth *services.AuthClient
	}

	UserRequest struct {
		Users []string `json:"users"`
	}

	NotAuthenticatedError struct{}
)

func init() {
	Register(new(UserManagement))
}

func (u *UserManagement) Init(c *services.Container) error {
	u.TemplateRenderer = c.TemplateRenderer
	u.mail = c.Mail
	u.orm = c.ORM
	return nil
}

func (u *UserManagement) Routes(g *echo.Group) {
	UMGroup := g.Group("/userManagement", middleware.RequireAuthentication())
	UMGroup.GET("", u.UserDetails).Name = routeNameUserDetails
	UMGroup.POST("/updateUserState", u.UpdateUserState).Name = routeNameUpdateUser
}

func (u *UserManagement) UserDetails(ctx echo.Context) error {
	currentUser := ctx.Get(context.AuthenticatedUserKey).(*ent.User)
	p := page.New(ctx)
	p.Layout = templates.LayoutMain

	if currentUser.Role != "admin" {
		p.Name = templates.PageError
		p.StatusCode = http.StatusForbidden
		return u.RenderPage(ctx, p)
	}

	allUsers, err := u.orm.User.Query().All(ctx.Request().Context())
	if err != nil {
		return err
	}

	p.Name = templates.PageUserManagement
	p.Title = "User Management"
	p.Data = allUsers

	return u.RenderPage(ctx, p)
}

func (u *UserManagement) UpdateUserState(ctx echo.Context) error {
	usr := ctx.Get(context.AuthenticatedUserKey).(*ent.User)
	p := page.New(ctx)
	p.Layout = templates.LayoutMain
	if usr.Role != "admin" {
		p.Name = templates.PageError
		p.StatusCode = http.StatusForbidden
		return u.RenderPage(ctx, p)
	}

	state := ctx.Request().Form.Get("state")
	users := ctx.Request().Form["users"]
	if len(users) == 0 {
		p.Name = templates.PageError
		p.StatusCode = http.StatusBadRequest
		return u.RenderPage(ctx, p)
	}

	// Determine new state (true = disable, false = enable)
	disable := state == "disable"
	newState := "enable"
	newClass := "button is-warning is-rounded"
	buttonText := "Enable"

	if !disable {
		newState = "disable"
		buttonText = "Disable"
		newClass = "button is-danger is-rounded"
	}

	// Update all users in a single query
	updatedUserState, err := u.orm.User.Update().
		Where(user.EmailIn(users...), user.RoleNEQ("admin")).
		SetDisabled(disable).
		Save(ctx.Request().Context())

	if err != nil || updatedUserState == 0 {
		return fail(err, "failed to update user state")
	}

	// Generate button HTML response
	buttonHTML := fmt.Sprintf(
		`<button hx-post="/userManagement/updateUserState" hx-vals='{"state":"%s","users":["%s"]}' hx-target="this" hx-swap="outerHTML" class="%s">%s</button>`,
		newState, users[0], newClass, buttonText,
	)

	return ctx.HTML(http.StatusOK, buttonHTML)
}
