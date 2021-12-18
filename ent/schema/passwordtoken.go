package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
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
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the PasswordToken.
func (PasswordToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Required().
			Unique(),
	}
}
