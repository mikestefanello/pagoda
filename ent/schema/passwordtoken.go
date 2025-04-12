package schema

import (
	"context"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	ge "github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/ent/hook"
	"golang.org/x/crypto/bcrypt"
)

// PasswordToken holds the schema definition for the PasswordToken entity.
type PasswordToken struct {
	ent.Schema
}

// Fields of the PasswordToken.
func (PasswordToken) Fields() []ent.Field {
	return []ent.Field{
		field.String("hash").
			Sensitive().
			NotEmpty(),
		field.Int("user_id"),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the PasswordToken.
func (PasswordToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Field("user_id").
			Required().
			Unique(),
	}
}

// Hooks of the PasswordToken.
func (PasswordToken) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.PasswordTokenFunc(func(ctx context.Context, m *ge.PasswordTokenMutation) (ent.Value, error) {
					if v, exists := m.Hash(); exists {
						hash, err := bcrypt.GenerateFromPassword([]byte(v), bcrypt.DefaultCost)
						if err != nil {
							return "", err
						}
						m.SetHash(string(hash))
					}
					return next.Mutate(ctx, m)
				})
			},
			// Limit the hook only for these operations.
			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne,
		),
	}
}
