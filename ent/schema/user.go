package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("email").
			MaxLen(255).
			Optional().
			Nillable().
			Unique(),
		field.String("password").
			MaxLen(255).
			Default(""),
		field.String("nickname").
			MaxLen(50),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
