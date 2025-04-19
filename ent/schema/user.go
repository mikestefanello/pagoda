package schema

import (
	"context"
	"net/mail"
	"strings"
	"time"

	ge "github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/ent/hook"
	"golang.org/x/crypto/bcrypt"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
		field.String("email").
			NotEmpty().
			Unique().
			Validate(func(s string) error {
				_, err := mail.ParseAddress(s)
				return err
			}),
		field.String("password").
			Sensitive().
			NotEmpty(),
		field.Bool("verified").
			Default(false),
		field.Bool("admin").
			Default(false),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", PasswordToken.Type).
			Ref("user"),
	}
}

// Hooks of the User.
func (User) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.UserFunc(func(ctx context.Context, m *ge.UserMutation) (ent.Value, error) {
					if v, exists := m.Email(); exists {
						m.SetEmail(strings.ToLower(v))
					}

					if v, exists := m.Password(); exists {
						hash, err := bcrypt.GenerateFromPassword([]byte(v), bcrypt.DefaultCost)
						if err != nil {
							return "", err
						}
						m.SetPassword(string(hash))
					}
					return next.Mutate(ctx, m)
				})
			},
			// Limit the hook only for these operations.
			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne,
		),
	}
}
