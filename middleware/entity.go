package middleware

import (
	"net/http"
	"strconv"

	"github.com/mikestefanello/pagoda/context"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/ent/user"

	"github.com/labstack/echo/v4"
)

// LoadUser loads the user based on the ID provided as a path parameter
func LoadUser(orm *ent.Client) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userID, err := strconv.Atoi(c.Param("user"))
			if err != nil {
				return echo.NewHTTPError(http.StatusNotFound)
			}

			u, err := orm.User.
				Query().
				Where(user.ID(userID)).
				Only(c.Request().Context())

			switch err.(type) {
			case nil:
				c.Set(context.UserKey, u)
				return next(c)
			case *ent.NotFoundError:
				return echo.NewHTTPError(http.StatusNotFound)
			default:
				if context.IsCanceledError(err) {
					return nil
				}
				c.Logger().Error(err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
		}
	}
}
