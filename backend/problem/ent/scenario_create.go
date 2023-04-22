// Code generated by ent, DO NOT EDIT.

package ent

import (
	"code-connect/problem/ent/scenario"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ScenarioCreate is the builder for creating a Scenario entity.
type ScenarioCreate struct {
	config
	mutation *ScenarioMutation
	hooks    []Hook
}

// SetUUID sets the "uuid" field.
func (sc *ScenarioCreate) SetUUID(s string) *ScenarioCreate {
	sc.mutation.SetUUID(s)
	return sc
}

// SetTitle sets the "title" field.
func (sc *ScenarioCreate) SetTitle(s string) *ScenarioCreate {
	sc.mutation.SetTitle(s)
	return sc
}

// SetContent sets the "content" field.
func (sc *ScenarioCreate) SetContent(s string) *ScenarioCreate {
	sc.mutation.SetContent(s)
	return sc
}

// SetRequestID sets the "request_id" field.
func (sc *ScenarioCreate) SetRequestID(s string) *ScenarioCreate {
	sc.mutation.SetRequestID(s)
	return sc
}

// Mutation returns the ScenarioMutation object of the builder.
func (sc *ScenarioCreate) Mutation() *ScenarioMutation {
	return sc.mutation
}

// Save creates the Scenario in the database.
func (sc *ScenarioCreate) Save(ctx context.Context) (*Scenario, error) {
	return withHooks[*Scenario, ScenarioMutation](ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *ScenarioCreate) SaveX(ctx context.Context) *Scenario {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *ScenarioCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *ScenarioCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *ScenarioCreate) check() error {
	if _, ok := sc.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New(`ent: missing required field "Scenario.uuid"`)}
	}
	if _, ok := sc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Scenario.title"`)}
	}
	if _, ok := sc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "Scenario.content"`)}
	}
	if _, ok := sc.mutation.RequestID(); !ok {
		return &ValidationError{Name: "request_id", err: errors.New(`ent: missing required field "Scenario.request_id"`)}
	}
	return nil
}

func (sc *ScenarioCreate) sqlSave(ctx context.Context) (*Scenario, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *ScenarioCreate) createSpec() (*Scenario, *sqlgraph.CreateSpec) {
	var (
		_node = &Scenario{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(scenario.Table, sqlgraph.NewFieldSpec(scenario.FieldID, field.TypeInt))
	)
	if value, ok := sc.mutation.UUID(); ok {
		_spec.SetField(scenario.FieldUUID, field.TypeString, value)
		_node.UUID = value
	}
	if value, ok := sc.mutation.Title(); ok {
		_spec.SetField(scenario.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := sc.mutation.Content(); ok {
		_spec.SetField(scenario.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if value, ok := sc.mutation.RequestID(); ok {
		_spec.SetField(scenario.FieldRequestID, field.TypeString, value)
		_node.RequestID = value
	}
	return _node, _spec
}

// ScenarioCreateBulk is the builder for creating many Scenario entities in bulk.
type ScenarioCreateBulk struct {
	config
	builders []*ScenarioCreate
}

// Save creates the Scenario entities in the database.
func (scb *ScenarioCreateBulk) Save(ctx context.Context) ([]*Scenario, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Scenario, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ScenarioMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *ScenarioCreateBulk) SaveX(ctx context.Context) []*Scenario {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *ScenarioCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *ScenarioCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
