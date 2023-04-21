// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"fsm_client/pkg/ent/dir"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Dir is the model entity for the Dir schema.
type Dir struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// SyncID holds the value of the "sync_id" field.
	SyncID string `json:"sync_id,omitempty" db:"sync_id" `
	// Dir holds the value of the "dir" field.
	Dir string `json:"dir,omitempty"`
	// Level holds the value of the "level" field.
	Level int `json:"level,omitempty"`
	// Deleted holds the value of the "deleted" field.
	Deleted bool `json:"deleted,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime int64 `json:"create_time,omitempty" db:"create_time" `
	// ModTime holds the value of the "mod_time" field.
	ModTime int64 `json:"mod_time,omitempty" db:"mod_time" `
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Dir) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case dir.FieldDeleted:
			values[i] = new(sql.NullBool)
		case dir.FieldLevel, dir.FieldCreateTime, dir.FieldModTime:
			values[i] = new(sql.NullInt64)
		case dir.FieldID, dir.FieldSyncID, dir.FieldDir:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Dir", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Dir fields.
func (d *Dir) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case dir.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				d.ID = value.String
			}
		case dir.FieldSyncID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sync_id", values[i])
			} else if value.Valid {
				d.SyncID = value.String
			}
		case dir.FieldDir:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field dir", values[i])
			} else if value.Valid {
				d.Dir = value.String
			}
		case dir.FieldLevel:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field level", values[i])
			} else if value.Valid {
				d.Level = int(value.Int64)
			}
		case dir.FieldDeleted:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field deleted", values[i])
			} else if value.Valid {
				d.Deleted = value.Bool
			}
		case dir.FieldCreateTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				d.CreateTime = value.Int64
			}
		case dir.FieldModTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field mod_time", values[i])
			} else if value.Valid {
				d.ModTime = value.Int64
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Dir.
// Note that you need to call Dir.Unwrap() before calling this method if this Dir
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Dir) Update() *DirUpdateOne {
	return NewDirClient(d.config).UpdateOne(d)
}

// Unwrap unwraps the Dir entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Dir) Unwrap() *Dir {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Dir is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Dir) String() string {
	var builder strings.Builder
	builder.WriteString("Dir(")
	builder.WriteString(fmt.Sprintf("id=%v, ", d.ID))
	builder.WriteString("sync_id=")
	builder.WriteString(d.SyncID)
	builder.WriteString(", ")
	builder.WriteString("dir=")
	builder.WriteString(d.Dir)
	builder.WriteString(", ")
	builder.WriteString("level=")
	builder.WriteString(fmt.Sprintf("%v", d.Level))
	builder.WriteString(", ")
	builder.WriteString("deleted=")
	builder.WriteString(fmt.Sprintf("%v", d.Deleted))
	builder.WriteString(", ")
	builder.WriteString("create_time=")
	builder.WriteString(fmt.Sprintf("%v", d.CreateTime))
	builder.WriteString(", ")
	builder.WriteString("mod_time=")
	builder.WriteString(fmt.Sprintf("%v", d.ModTime))
	builder.WriteByte(')')
	return builder.String()
}

// Dirs is a parsable slice of Dir.
type Dirs []*Dir
