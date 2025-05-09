// Code generated by ent, DO NOT EDIT.
package admin

import "time"

type PasswordToken struct {
	Token     *string    `form:"token"`
	UserID    int        `form:"user_id"`
	CreatedAt *time.Time `form:"created_at"`
}

type User struct {
	Name      string     `form:"name"`
	Email     string     `form:"email"`
	Password  *string    `form:"password"`
	Verified  bool       `form:"verified"`
	Admin     bool       `form:"admin"`
	CreatedAt *time.Time `form:"created_at"`
}

type EntityList struct {
	Columns     []string
	Entities    []EntityValues
	Page        int
	HasNextPage bool
}

type EntityValues struct {
	ID     int
	Values []string
}

type HandlerConfig struct {
	ItemsPerPage int
	PageQueryKey string
	TimeFormat   string
}

func GetEntityTypeNames() []string {
	return []string{
		"PasswordToken",
		"User",
	}
}
