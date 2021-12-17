package middleware

import (
	"net/http"
	"strconv"

	"goweb/context"
	"goweb/ent"
	"goweb/ent/user"

	"github.com/labstack/echo/v4"
)

func LoadUser(orm *ent.Client) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userID, err := strconv.Atoi(c.Param("user"))
			if err != nil {
				return echo.NewHTTPError(http.StatusNotFound, "Not found")
			}

			u, err := orm.User.
				Query().
				Where(user.ID(userID)).
				Only(c.Request().Context())

			switch err.(type) {
			case nil:
			case *ent.NotFoundError:
				return echo.NewHTTPError(http.StatusNotFound, "Not found")
			default:
				c.Logger().Error(err)
				return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
			}

			c.Set(context.UserKey, u)

			return next(c)
		}
	}
}
