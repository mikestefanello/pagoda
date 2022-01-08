package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/context"
	"github.com/mikestefanello/pagoda/controller"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/ent/user"
	"github.com/mikestefanello/pagoda/msg"
)

type VerifyEmail struct {
	controller.Controller
}

func (c *VerifyEmail) Get(ctx echo.Context) error {
	c.verifyToken(ctx)

	return c.Redirect(ctx, "home")
}

func (c *VerifyEmail) verifyToken(ctx echo.Context) {
	var usr *ent.User

	// Validate the token
	token := ctx.Param("token")
	email, err := c.Container.Auth.ValidateEmailVerificationToken(token)
	if err != nil {
		msg.Warning(ctx, "The link is either invalid or has expired.")
		return
	}

	// Check if it matches the authenticated user
	if u := ctx.Get(context.AuthenticatedUserKey); u != nil {
		authUser := u.(*ent.User)

		if authUser.Email == email {
			usr = authUser
		}
	}

	// Query to find a matching user, if needed
	if usr == nil {
		usr, err = c.Container.ORM.User.
			Query().
			Where(user.Email(email)).
			Only(ctx.Request().Context())

		if err != nil {
			ctx.Logger().Errorf("error querying user during email verification: %v", err)
			msg.Danger(ctx, "An error occurred. Please try again.")
			return
		}
	}

	// Verify the user
	err = c.Container.ORM.User.
		Update().
		SetVerified(true).
		Where(user.ID(usr.ID)).
		Exec(ctx.Request().Context())

	if err != nil {
		ctx.Logger().Errorf("error setting user as verified: %v", err)
		msg.Danger(ctx, "An error occurred. Please try again.")
		return
	}

	msg.Success(ctx, "You email has been successfully verified.")
}
