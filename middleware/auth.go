package middleware

import (
	"net/http"

	"goweb/context"
	"goweb/ent"
	"goweb/msg"
	"goweb/services"

	"github.com/labstack/echo/v4"
)

func LoadAuthenticatedUser(authClient *services.AuthClient) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			u, err := authClient.GetAuthenticatedUser(c)
			switch err.(type) {
			case *ent.NotFoundError:
				c.Logger().Debug("auth user not found")
			case services.NotAuthenticatedError:
			case nil:
				c.Set(context.AuthenticatedUserKey, u)
				c.Logger().Info("auth user loaded in to context: %d", u.ID)
			default:
				c.Logger().Errorf("error querying for authenticated user: %v", err)
			}

			return next(c)
		}
	}
}

func LoadValidPasswordToken(authClient *services.AuthClient) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var usr *ent.User

			if c.Get(context.UserKey) == nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
			}
			usr = c.Get(context.UserKey).(*ent.User)

			tokenParam := c.Param("password_token")
			token, err := authClient.GetValidPasswordToken(c, tokenParam, usr.ID)

			switch err.(type) {
			case nil:
			case services.InvalidPasswordTokenError:
				msg.Warning(c, "The link is either invalid or has expired. Please request a new one.")
				return c.Redirect(http.StatusFound, c.Echo().Reverse("forgot_password"))
			default:
				c.Logger().Error(err)
				return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
			}

			c.Set(context.PasswordTokenKey, token)

			return next(c)
		}
	}
}

func RequireAuthentication() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if u := c.Get(context.AuthenticatedUserKey); u == nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
			}

			return next(c)
		}
	}
}

func RequireNoAuthentication() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if u := c.Get(context.AuthenticatedUserKey); u != nil {
				return echo.NewHTTPError(http.StatusForbidden, "Forbidden")
			}

			return next(c)
		}
	}
}
