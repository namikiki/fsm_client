// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// DirsColumns holds the columns for the "dirs" table.
	DirsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "sync_id", Type: field.TypeString},
		{Name: "dir", Type: field.TypeString},
		{Name: "level", Type: field.TypeUint64},
		{Name: "deleted", Type: field.TypeBool},
		{Name: "create_time", Type: field.TypeInt64},
		{Name: "mod_time", Type: field.TypeInt64},
	}
	// DirsTable holds the schema information for the "dirs" table.
	DirsTable = &schema.Table{
		Name:       "dirs",
		Columns:    DirsColumns,
		PrimaryKey: []*schema.Column{DirsColumns[0]},
	}
	// FilesColumns holds the columns for the "files" table.
	FilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "sync_id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "parent_dir_id", Type: field.TypeString},
		{Name: "level", Type: field.TypeUint64},
		{Name: "hash", Type: field.TypeString},
		{Name: "size", Type: field.TypeInt64},
		{Name: "deleted", Type: field.TypeBool},
		{Name: "create_time", Type: field.TypeInt64},
		{Name: "mod_time", Type: field.TypeInt64},
	}
	// FilesTable holds the schema information for the "files" table.
	FilesTable = &schema.Table{
		Name:       "files",
		Columns:    FilesColumns,
		PrimaryKey: []*schema.Column{FilesColumns[0]},
	}
	// SyncTasksColumns holds the columns for the "sync_tasks" table.
	SyncTasksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeString},
		{Name: "type", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "root_dir", Type: field.TypeString},
		{Name: "ignore", Type: field.TypeBool},
		{Name: "deleted", Type: field.TypeBool},
		{Name: "status", Type: field.TypeString},
		{Name: "create_time", Type: field.TypeInt64},
	}
	// SyncTasksTable holds the schema information for the "sync_tasks" table.
	SyncTasksTable = &schema.Table{
		Name:       "sync_tasks",
		Columns:    SyncTasksColumns,
		PrimaryKey: []*schema.Column{SyncTasksColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "parent_id", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "deleted", Type: field.TypeBool},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "mod_time", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DirsTable,
		FilesTable,
		SyncTasksTable,
		UsersTable,
	}
)

func init() {
}
