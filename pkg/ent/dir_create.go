// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"fsm_client/pkg/ent/dir"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DirCreate is the builder for creating a Dir entity.
type DirCreate struct {
	config
	mutation *DirMutation
	hooks    []Hook
}

// SetSyncID sets the "sync_id" field.
func (dc *DirCreate) SetSyncID(s string) *DirCreate {
	dc.mutation.SetSyncID(s)
	return dc
}

// SetDir sets the "dir" field.
func (dc *DirCreate) SetDir(s string) *DirCreate {
	dc.mutation.SetDir(s)
	return dc
}

// SetLevel sets the "level" field.
func (dc *DirCreate) SetLevel(u uint64) *DirCreate {
	dc.mutation.SetLevel(u)
	return dc
}

// SetDeleted sets the "deleted" field.
func (dc *DirCreate) SetDeleted(b bool) *DirCreate {
	dc.mutation.SetDeleted(b)
	return dc
}

// SetCreateTime sets the "create_time" field.
func (dc *DirCreate) SetCreateTime(t time.Time) *DirCreate {
	dc.mutation.SetCreateTime(t)
	return dc
}

// SetModTime sets the "mod_time" field.
func (dc *DirCreate) SetModTime(t time.Time) *DirCreate {
	dc.mutation.SetModTime(t)
	return dc
}

// SetID sets the "id" field.
func (dc *DirCreate) SetID(s string) *DirCreate {
	dc.mutation.SetID(s)
	return dc
}

// Mutation returns the DirMutation object of the builder.
func (dc *DirCreate) Mutation() *DirMutation {
	return dc.mutation
}

// Save creates the Dir in the database.
func (dc *DirCreate) Save(ctx context.Context) (*Dir, error) {
	return withHooks[*Dir, DirMutation](ctx, dc.sqlSave, dc.mutation, dc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DirCreate) SaveX(ctx context.Context) *Dir {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DirCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DirCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DirCreate) check() error {
	if _, ok := dc.mutation.SyncID(); !ok {
		return &ValidationError{Name: "sync_id", err: errors.New(`ent: missing required field "Dir.sync_id"`)}
	}
	if _, ok := dc.mutation.Dir(); !ok {
		return &ValidationError{Name: "dir", err: errors.New(`ent: missing required field "Dir.dir"`)}
	}
	if _, ok := dc.mutation.Level(); !ok {
		return &ValidationError{Name: "level", err: errors.New(`ent: missing required field "Dir.level"`)}
	}
	if _, ok := dc.mutation.Deleted(); !ok {
		return &ValidationError{Name: "deleted", err: errors.New(`ent: missing required field "Dir.deleted"`)}
	}
	if _, ok := dc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Dir.create_time"`)}
	}
	if _, ok := dc.mutation.ModTime(); !ok {
		return &ValidationError{Name: "mod_time", err: errors.New(`ent: missing required field "Dir.mod_time"`)}
	}
	return nil
}

func (dc *DirCreate) sqlSave(ctx context.Context) (*Dir, error) {
	if err := dc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Dir.ID type: %T", _spec.ID.Value)
		}
	}
	dc.mutation.id = &_node.ID
	dc.mutation.done = true
	return _node, nil
}

func (dc *DirCreate) createSpec() (*Dir, *sqlgraph.CreateSpec) {
	var (
		_node = &Dir{config: dc.config}
		_spec = sqlgraph.NewCreateSpec(dir.Table, sqlgraph.NewFieldSpec(dir.FieldID, field.TypeString))
	)
	if id, ok := dc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := dc.mutation.SyncID(); ok {
		_spec.SetField(dir.FieldSyncID, field.TypeString, value)
		_node.SyncID = value
	}
	if value, ok := dc.mutation.Dir(); ok {
		_spec.SetField(dir.FieldDir, field.TypeString, value)
		_node.Dir = value
	}
	if value, ok := dc.mutation.Level(); ok {
		_spec.SetField(dir.FieldLevel, field.TypeUint64, value)
		_node.Level = value
	}
	if value, ok := dc.mutation.Deleted(); ok {
		_spec.SetField(dir.FieldDeleted, field.TypeBool, value)
		_node.Deleted = value
	}
	if value, ok := dc.mutation.CreateTime(); ok {
		_spec.SetField(dir.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := dc.mutation.ModTime(); ok {
		_spec.SetField(dir.FieldModTime, field.TypeTime, value)
		_node.ModTime = value
	}
	return _node, _spec
}

// DirCreateBulk is the builder for creating many Dir entities in bulk.
type DirCreateBulk struct {
	config
	builders []*DirCreate
}

// Save creates the Dir entities in the database.
func (dcb *DirCreateBulk) Save(ctx context.Context) ([]*Dir, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Dir, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DirMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DirCreateBulk) SaveX(ctx context.Context) []*Dir {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DirCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DirCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}