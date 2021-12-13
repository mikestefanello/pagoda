package middleware

import (
	"net/http"

	"goweb/auth"
	"goweb/context"
	"goweb/ent"
	"goweb/ent/user"

	"github.com/labstack/echo/v4"
)

func LoadAuthenticatedUser(orm *ent.Client) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if userID, err := auth.GetUserID(c); err == nil {
				u, err := orm.User.Query().
					Where(user.ID(userID)).
					First(c.Request().Context())

				switch err.(type) {
				case *ent.NotFoundError:
					c.Logger().Debug("auth user not found: %d", userID)
				case nil:
					c.Set(context.AuthenticatedUserKey, u)
					c.Logger().Info("auth user loaded in to context: %d", userID)
				default:
					c.Logger().Errorf("error querying for authenticated user: %v", err)
				}
			}

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
