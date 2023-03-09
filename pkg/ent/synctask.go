// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"fsm_client/pkg/ent/synctask"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// SyncTask is the model entity for the SyncTask schema.
type SyncTask struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID string `json:"user_id,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// RootDir holds the value of the "root_dir" field.
	RootDir string `json:"root_dir,omitempty"`
	// Deleted holds the value of the "deleted" field.
	Deleted bool `json:"deleted,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SyncTask) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case synctask.FieldDeleted:
			values[i] = new(sql.NullBool)
		case synctask.FieldID, synctask.FieldUserID, synctask.FieldType, synctask.FieldName, synctask.FieldRootDir:
			values[i] = new(sql.NullString)
		case synctask.FieldCreateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SyncTask", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SyncTask fields.
func (st *SyncTask) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case synctask.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				st.ID = value.String
			}
		case synctask.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				st.UserID = value.String
			}
		case synctask.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				st.Type = value.String
			}
		case synctask.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				st.Name = value.String
			}
		case synctask.FieldRootDir:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field root_dir", values[i])
			} else if value.Valid {
				st.RootDir = value.String
			}
		case synctask.FieldDeleted:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field deleted", values[i])
			} else if value.Valid {
				st.Deleted = value.Bool
			}
		case synctask.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				st.CreateTime = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this SyncTask.
// Note that you need to call SyncTask.Unwrap() before calling this method if this SyncTask
// was returned from a transaction, and the transaction was committed or rolled back.
func (st *SyncTask) Update() *SyncTaskUpdateOne {
	return NewSyncTaskClient(st.config).UpdateOne(st)
}

// Unwrap unwraps the SyncTask entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (st *SyncTask) Unwrap() *SyncTask {
	_tx, ok := st.config.driver.(*txDriver)
	if !ok {
		panic("ent: SyncTask is not a transactional entity")
	}
	st.config.driver = _tx.drv
	return st
}

// String implements the fmt.Stringer.
func (st *SyncTask) String() string {
	var builder strings.Builder
	builder.WriteString("SyncTask(")
	builder.WriteString(fmt.Sprintf("id=%v, ", st.ID))
	builder.WriteString("user_id=")
	builder.WriteString(st.UserID)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(st.Type)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(st.Name)
	builder.WriteString(", ")
	builder.WriteString("root_dir=")
	builder.WriteString(st.RootDir)
	builder.WriteString(", ")
	builder.WriteString("deleted=")
	builder.WriteString(fmt.Sprintf("%v", st.Deleted))
	builder.WriteString(", ")
	builder.WriteString("create_time=")
	builder.WriteString(st.CreateTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// SyncTasks is a parsable slice of SyncTask.
type SyncTasks []*SyncTask
