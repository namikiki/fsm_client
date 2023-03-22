package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// File holds the schema definition for the File entity.
type File struct {
	ent.Schema
}

// Fields of the File.
func (File) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("sync_id"),

		field.String("name"),
		field.String("parent_dir_id"),
		field.Uint64("level"),
		field.String("hash"),
		field.Int64("size"),

		field.Bool("deleted"),
		field.Int64("create_time"),
		field.Int64("mod_time"),
	}
}

// Edges of the File.
func (File) Edges() []ent.Edge {
	return nil
}
