// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fsm_client/pkg/ent/dir"
	"fsm_client/pkg/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DirDelete is the builder for deleting a Dir entity.
type DirDelete struct {
	config
	hooks    []Hook
	mutation *DirMutation
}

// Where appends a list predicates to the DirDelete builder.
func (dd *DirDelete) Where(ps ...predicate.Dir) *DirDelete {
	dd.mutation.Where(ps...)
	return dd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dd *DirDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, DirMutation](ctx, dd.sqlExec, dd.mutation, dd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (dd *DirDelete) ExecX(ctx context.Context) int {
	n, err := dd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dd *DirDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(dir.Table, sqlgraph.NewFieldSpec(dir.FieldID, field.TypeString))
	if ps := dd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, dd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	dd.mutation.done = true
	return affected, err
}

// DirDeleteOne is the builder for deleting a single Dir entity.
type DirDeleteOne struct {
	dd *DirDelete
}

// Where appends a list predicates to the DirDelete builder.
func (ddo *DirDeleteOne) Where(ps ...predicate.Dir) *DirDeleteOne {
	ddo.dd.mutation.Where(ps...)
	return ddo
}

// Exec executes the deletion query.
func (ddo *DirDeleteOne) Exec(ctx context.Context) error {
	n, err := ddo.dd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{dir.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ddo *DirDeleteOne) ExecX(ctx context.Context) {
	if err := ddo.Exec(ctx); err != nil {
		panic(err)
	}
}