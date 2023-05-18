// Code generated by ent, DO NOT EDIT.

package ent

import (
	"code-connect/problem/ent/predicate"
	"code-connect/problem/ent/problem"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProblemUpdate is the builder for updating Problem entities.
type ProblemUpdate struct {
	config
	hooks    []Hook
	mutation *ProblemMutation
}

// Where appends a list predicates to the ProblemUpdate builder.
func (pu *ProblemUpdate) Where(ps ...predicate.Problem) *ProblemUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetUUID sets the "uuid" field.
func (pu *ProblemUpdate) SetUUID(s string) *ProblemUpdate {
	pu.mutation.SetUUID(s)
	return pu
}

// SetTitle sets the "title" field.
func (pu *ProblemUpdate) SetTitle(s string) *ProblemUpdate {
	pu.mutation.SetTitle(s)
	return pu
}

// SetBackground sets the "background" field.
func (pu *ProblemUpdate) SetBackground(s string) *ProblemUpdate {
	pu.mutation.SetBackground(s)
	return pu
}

// SetCode sets the "code" field.
func (pu *ProblemUpdate) SetCode(s string) *ProblemUpdate {
	pu.mutation.SetCode(s)
	return pu
}

// SetTestCode sets the "test_code" field.
func (pu *ProblemUpdate) SetTestCode(s string) *ProblemUpdate {
	pu.mutation.SetTestCode(s)
	return pu
}

// SetNillableTestCode sets the "test_code" field if the given value is not nil.
func (pu *ProblemUpdate) SetNillableTestCode(s *string) *ProblemUpdate {
	if s != nil {
		pu.SetTestCode(*s)
	}
	return pu
}

// ClearTestCode clears the value of the "test_code" field.
func (pu *ProblemUpdate) ClearTestCode() *ProblemUpdate {
	pu.mutation.ClearTestCode()
	return pu
}

// SetEstimatedTime sets the "estimated_time" field.
func (pu *ProblemUpdate) SetEstimatedTime(i int) *ProblemUpdate {
	pu.mutation.ResetEstimatedTime()
	pu.mutation.SetEstimatedTime(i)
	return pu
}

// SetNillableEstimatedTime sets the "estimated_time" field if the given value is not nil.
func (pu *ProblemUpdate) SetNillableEstimatedTime(i *int) *ProblemUpdate {
	if i != nil {
		pu.SetEstimatedTime(*i)
	}
	return pu
}

// AddEstimatedTime adds i to the "estimated_time" field.
func (pu *ProblemUpdate) AddEstimatedTime(i int) *ProblemUpdate {
	pu.mutation.AddEstimatedTime(i)
	return pu
}

// SetLanguage sets the "language" field.
func (pu *ProblemUpdate) SetLanguage(s string) *ProblemUpdate {
	pu.mutation.SetLanguage(s)
	return pu
}

// SetRequestID sets the "request_id" field.
func (pu *ProblemUpdate) SetRequestID(s string) *ProblemUpdate {
	pu.mutation.SetRequestID(s)
	return pu
}

// Mutation returns the ProblemMutation object of the builder.
func (pu *ProblemUpdate) Mutation() *ProblemMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProblemUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, ProblemMutation](ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProblemUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProblemUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProblemUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *ProblemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(problem.Table, problem.Columns, sqlgraph.NewFieldSpec(problem.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.UUID(); ok {
		_spec.SetField(problem.FieldUUID, field.TypeString, value)
	}
	if value, ok := pu.mutation.Title(); ok {
		_spec.SetField(problem.FieldTitle, field.TypeString, value)
	}
	if value, ok := pu.mutation.Background(); ok {
		_spec.SetField(problem.FieldBackground, field.TypeString, value)
	}
	if value, ok := pu.mutation.Code(); ok {
		_spec.SetField(problem.FieldCode, field.TypeString, value)
	}
	if value, ok := pu.mutation.TestCode(); ok {
		_spec.SetField(problem.FieldTestCode, field.TypeString, value)
	}
	if pu.mutation.TestCodeCleared() {
		_spec.ClearField(problem.FieldTestCode, field.TypeString)
	}
	if value, ok := pu.mutation.EstimatedTime(); ok {
		_spec.SetField(problem.FieldEstimatedTime, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedEstimatedTime(); ok {
		_spec.AddField(problem.FieldEstimatedTime, field.TypeInt, value)
	}
	if value, ok := pu.mutation.Language(); ok {
		_spec.SetField(problem.FieldLanguage, field.TypeString, value)
	}
	if value, ok := pu.mutation.RequestID(); ok {
		_spec.SetField(problem.FieldRequestID, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{problem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// ProblemUpdateOne is the builder for updating a single Problem entity.
type ProblemUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProblemMutation
}

// SetUUID sets the "uuid" field.
func (puo *ProblemUpdateOne) SetUUID(s string) *ProblemUpdateOne {
	puo.mutation.SetUUID(s)
	return puo
}

// SetTitle sets the "title" field.
func (puo *ProblemUpdateOne) SetTitle(s string) *ProblemUpdateOne {
	puo.mutation.SetTitle(s)
	return puo
}

// SetBackground sets the "background" field.
func (puo *ProblemUpdateOne) SetBackground(s string) *ProblemUpdateOne {
	puo.mutation.SetBackground(s)
	return puo
}

// SetCode sets the "code" field.
func (puo *ProblemUpdateOne) SetCode(s string) *ProblemUpdateOne {
	puo.mutation.SetCode(s)
	return puo
}

// SetTestCode sets the "test_code" field.
func (puo *ProblemUpdateOne) SetTestCode(s string) *ProblemUpdateOne {
	puo.mutation.SetTestCode(s)
	return puo
}

// SetNillableTestCode sets the "test_code" field if the given value is not nil.
func (puo *ProblemUpdateOne) SetNillableTestCode(s *string) *ProblemUpdateOne {
	if s != nil {
		puo.SetTestCode(*s)
	}
	return puo
}

// ClearTestCode clears the value of the "test_code" field.
func (puo *ProblemUpdateOne) ClearTestCode() *ProblemUpdateOne {
	puo.mutation.ClearTestCode()
	return puo
}

// SetEstimatedTime sets the "estimated_time" field.
func (puo *ProblemUpdateOne) SetEstimatedTime(i int) *ProblemUpdateOne {
	puo.mutation.ResetEstimatedTime()
	puo.mutation.SetEstimatedTime(i)
	return puo
}

// SetNillableEstimatedTime sets the "estimated_time" field if the given value is not nil.
func (puo *ProblemUpdateOne) SetNillableEstimatedTime(i *int) *ProblemUpdateOne {
	if i != nil {
		puo.SetEstimatedTime(*i)
	}
	return puo
}

// AddEstimatedTime adds i to the "estimated_time" field.
func (puo *ProblemUpdateOne) AddEstimatedTime(i int) *ProblemUpdateOne {
	puo.mutation.AddEstimatedTime(i)
	return puo
}

// SetLanguage sets the "language" field.
func (puo *ProblemUpdateOne) SetLanguage(s string) *ProblemUpdateOne {
	puo.mutation.SetLanguage(s)
	return puo
}

// SetRequestID sets the "request_id" field.
func (puo *ProblemUpdateOne) SetRequestID(s string) *ProblemUpdateOne {
	puo.mutation.SetRequestID(s)
	return puo
}

// Mutation returns the ProblemMutation object of the builder.
func (puo *ProblemUpdateOne) Mutation() *ProblemMutation {
	return puo.mutation
}

// Where appends a list predicates to the ProblemUpdate builder.
func (puo *ProblemUpdateOne) Where(ps ...predicate.Problem) *ProblemUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ProblemUpdateOne) Select(field string, fields ...string) *ProblemUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Problem entity.
func (puo *ProblemUpdateOne) Save(ctx context.Context) (*Problem, error) {
	return withHooks[*Problem, ProblemMutation](ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProblemUpdateOne) SaveX(ctx context.Context) *Problem {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProblemUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProblemUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *ProblemUpdateOne) sqlSave(ctx context.Context) (_node *Problem, err error) {
	_spec := sqlgraph.NewUpdateSpec(problem.Table, problem.Columns, sqlgraph.NewFieldSpec(problem.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Problem.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, problem.FieldID)
		for _, f := range fields {
			if !problem.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != problem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.UUID(); ok {
		_spec.SetField(problem.FieldUUID, field.TypeString, value)
	}
	if value, ok := puo.mutation.Title(); ok {
		_spec.SetField(problem.FieldTitle, field.TypeString, value)
	}
	if value, ok := puo.mutation.Background(); ok {
		_spec.SetField(problem.FieldBackground, field.TypeString, value)
	}
	if value, ok := puo.mutation.Code(); ok {
		_spec.SetField(problem.FieldCode, field.TypeString, value)
	}
	if value, ok := puo.mutation.TestCode(); ok {
		_spec.SetField(problem.FieldTestCode, field.TypeString, value)
	}
	if puo.mutation.TestCodeCleared() {
		_spec.ClearField(problem.FieldTestCode, field.TypeString)
	}
	if value, ok := puo.mutation.EstimatedTime(); ok {
		_spec.SetField(problem.FieldEstimatedTime, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedEstimatedTime(); ok {
		_spec.AddField(problem.FieldEstimatedTime, field.TypeInt, value)
	}
	if value, ok := puo.mutation.Language(); ok {
		_spec.SetField(problem.FieldLanguage, field.TypeString, value)
	}
	if value, ok := puo.mutation.RequestID(); ok {
		_spec.SetField(problem.FieldRequestID, field.TypeString, value)
	}
	_node = &Problem{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{problem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}