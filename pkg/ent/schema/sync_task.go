package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// SyncTask  holds the schema definition for the File entity.
type SyncTask struct {
	ent.Schema
}

// Fields of the SyncTask.
func (SyncTask) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("user_id"),

		field.String("type"),
		field.String("name"),
		field.String("root_dir"),

		field.Bool("deleted"),
		field.Time("create_time"),
	}
}
