package middleware

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/msg"
	"github.com/mikestefanello/pagoda/pkg/services"

	"github.com/labstack/echo/v4"
)

// LoadAuthenticatedUser loads the authenticated user, if one, and stores in context
func LoadAuthenticatedUser(authClient *services.AuthClient) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			u, err := authClient.GetAuthenticatedUser(c)
			switch err.(type) {
			case *ent.NotFoundError:
				c.Logger().Warn("auth user not found")
			case services.NotAuthenticatedError:
			case nil:
				c.Set(context.AuthenticatedUserKey, u)
				c.Logger().Infof("auth user loaded in to context: %d", u.ID)
			default:
				return echo.NewHTTPError(
					http.StatusInternalServerError,
					fmt.Sprintf("error querying for authenticated user: %v", err),
				)
			}

			return next(c)
		}
	}
}

// LoadValidPasswordToken loads a valid password token entity that matches the user and token
// provided in path parameters
// If the token is invalid, the user will be redirected to the forgot password route
// This requires that the user owning the token is loaded in to context
func LoadValidPasswordToken(authClient *services.AuthClient) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Extract the user parameter
			if c.Get(context.UserKey) == nil {
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			usr := c.Get(context.UserKey).(*ent.User)

			// Extract the token ID
			tokenID, err := strconv.Atoi(c.Param("password_token"))
			if err != nil {
				return echo.NewHTTPError(http.StatusNotFound)
			}

			// Attempt to load a valid password token
			token, err := authClient.GetValidPasswordToken(
				c,
				usr.ID,
				tokenID,
				c.Param("token"),
			)

			switch err.(type) {
			case nil:
				c.Set(context.PasswordTokenKey, token)
				return next(c)
			case services.InvalidPasswordTokenError:
				msg.Warning(c, "The link is either invalid or has expired. Please request a new one.")
				return c.Redirect(http.StatusFound, c.Echo().Reverse("forgot_password"))
			default:
				return echo.NewHTTPError(
					http.StatusInternalServerError,
					fmt.Sprintf("error loading password token: %v", err),
				)
			}
		}
	}
}

// RequireAuthentication requires that the user be authenticated in order to proceed
func RequireAuthentication() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if u := c.Get(context.AuthenticatedUserKey); u == nil {
				return echo.NewHTTPError(http.StatusUnauthorized)
			}

			return next(c)
		}
	}
}

// RequireNoAuthentication requires that the user not be authenticated in order to proceed
func RequireNoAuthentication() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if u := c.Get(context.AuthenticatedUserKey); u != nil {
				return echo.NewHTTPError(http.StatusForbidden)
			}

			return next(c)
		}
	}
}
