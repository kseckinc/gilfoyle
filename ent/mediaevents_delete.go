// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/dreamvo/gilfoyle/ent/mediaevents"
	"github.com/dreamvo/gilfoyle/ent/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// MediaEventsDelete is the builder for deleting a MediaEvents entity.
type MediaEventsDelete struct {
	config
	hooks    []Hook
	mutation *MediaEventsMutation
}

// Where adds a new predicate to the delete builder.
func (med *MediaEventsDelete) Where(ps ...predicate.MediaEvents) *MediaEventsDelete {
	med.mutation.predicates = append(med.mutation.predicates, ps...)
	return med
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (med *MediaEventsDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(med.hooks) == 0 {
		affected, err = med.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MediaEventsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			med.mutation = mutation
			affected, err = med.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(med.hooks) - 1; i >= 0; i-- {
			mut = med.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, med.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (med *MediaEventsDelete) ExecX(ctx context.Context) int {
	n, err := med.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (med *MediaEventsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: mediaevents.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: mediaevents.FieldID,
			},
		},
	}
	if ps := med.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, med.driver, _spec)
}

// MediaEventsDeleteOne is the builder for deleting a single MediaEvents entity.
type MediaEventsDeleteOne struct {
	med *MediaEventsDelete
}

// Exec executes the deletion query.
func (medo *MediaEventsDeleteOne) Exec(ctx context.Context) error {
	n, err := medo.med.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{mediaevents.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (medo *MediaEventsDeleteOne) ExecX(ctx context.Context) {
	medo.med.ExecX(ctx)
}