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
		field.String("id"),
		field.String("parent_id"),
		field.String("user_id"),
		field.String("name"),
		field.Bool("deleted"),
		field.Time("create_time"),
		field.Time("mod_time"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
