// Code generated by ent, DO NOT EDIT.

package ent

import (
	"code-connect/ent/predicate"
	"code-connect/ent/problem"
	"code-connect/ent/record"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RecordUpdate is the builder for updating Record entities.
type RecordUpdate struct {
	config
	hooks    []Hook
	mutation *RecordMutation
}

// Where appends a list predicates to the RecordUpdate builder.
func (ru *RecordUpdate) Where(ps ...predicate.Record) *RecordUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetUUID sets the "uuid" field.
func (ru *RecordUpdate) SetUUID(s string) *RecordUpdate {
	ru.mutation.SetUUID(s)
	return ru
}

// SetUserUUID sets the "user_uuid" field.
func (ru *RecordUpdate) SetUserUUID(s string) *RecordUpdate {
	ru.mutation.SetUserUUID(s)
	return ru
}

// SetNillableUserUUID sets the "user_uuid" field if the given value is not nil.
func (ru *RecordUpdate) SetNillableUserUUID(s *string) *RecordUpdate {
	if s != nil {
		ru.SetUserUUID(*s)
	}
	return ru
}

// SetCode sets the "code" field.
func (ru *RecordUpdate) SetCode(s string) *RecordUpdate {
	ru.mutation.SetCode(s)
	return ru
}

// SetReadability sets the "readability" field.
func (ru *RecordUpdate) SetReadability(i int) *RecordUpdate {
	ru.mutation.ResetReadability()
	ru.mutation.SetReadability(i)
	return ru
}

// SetNillableReadability sets the "readability" field if the given value is not nil.
func (ru *RecordUpdate) SetNillableReadability(i *int) *RecordUpdate {
	if i != nil {
		ru.SetReadability(*i)
	}
	return ru
}

// AddReadability adds i to the "readability" field.
func (ru *RecordUpdate) AddReadability(i int) *RecordUpdate {
	ru.mutation.AddReadability(i)
	return ru
}

// SetRobustness sets the "robustness" field.
func (ru *RecordUpdate) SetRobustness(i int) *RecordUpdate {
	ru.mutation.ResetRobustness()
	ru.mutation.SetRobustness(i)
	return ru
}

// SetNillableRobustness sets the "robustness" field if the given value is not nil.
func (ru *RecordUpdate) SetNillableRobustness(i *int) *RecordUpdate {
	if i != nil {
		ru.SetRobustness(*i)
	}
	return ru
}

// AddRobustness adds i to the "robustness" field.
func (ru *RecordUpdate) AddRobustness(i int) *RecordUpdate {
	ru.mutation.AddRobustness(i)
	return ru
}

// SetEfficiency sets the "efficiency" field.
func (ru *RecordUpdate) SetEfficiency(i int) *RecordUpdate {
	ru.mutation.ResetEfficiency()
	ru.mutation.SetEfficiency(i)
	return ru
}

// SetNillableEfficiency sets the "efficiency" field if the given value is not nil.
func (ru *RecordUpdate) SetNillableEfficiency(i *int) *RecordUpdate {
	if i != nil {
		ru.SetEfficiency(*i)
	}
	return ru
}

// AddEfficiency adds i to the "efficiency" field.
func (ru *RecordUpdate) AddEfficiency(i int) *RecordUpdate {
	ru.mutation.AddEfficiency(i)
	return ru
}

// SetProblemID sets the "problem" edge to the Problem entity by ID.
func (ru *RecordUpdate) SetProblemID(id int) *RecordUpdate {
	ru.mutation.SetProblemID(id)
	return ru
}

// SetProblem sets the "problem" edge to the Problem entity.
func (ru *RecordUpdate) SetProblem(p *Problem) *RecordUpdate {
	return ru.SetProblemID(p.ID)
}

// Mutation returns the RecordMutation object of the builder.
func (ru *RecordUpdate) Mutation() *RecordMutation {
	return ru.mutation
}

// ClearProblem clears the "problem" edge to the Problem entity.
func (ru *RecordUpdate) ClearProblem() *RecordUpdate {
	ru.mutation.ClearProblem()
	return ru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RecordUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RecordUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RecordUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RecordUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *RecordUpdate) check() error {
	if _, ok := ru.mutation.ProblemID(); ru.mutation.ProblemCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Record.problem"`)
	}
	return nil
}

func (ru *RecordUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(record.Table, record.Columns, sqlgraph.NewFieldSpec(record.FieldID, field.TypeInt))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.UUID(); ok {
		_spec.SetField(record.FieldUUID, field.TypeString, value)
	}
	if value, ok := ru.mutation.UserUUID(); ok {
		_spec.SetField(record.FieldUserUUID, field.TypeString, value)
	}
	if value, ok := ru.mutation.Code(); ok {
		_spec.SetField(record.FieldCode, field.TypeString, value)
	}
	if value, ok := ru.mutation.Readability(); ok {
		_spec.SetField(record.FieldReadability, field.TypeInt, value)
	}
	if value, ok := ru.mutation.AddedReadability(); ok {
		_spec.AddField(record.FieldReadability, field.TypeInt, value)
	}
	if value, ok := ru.mutation.Robustness(); ok {
		_spec.SetField(record.FieldRobustness, field.TypeInt, value)
	}
	if value, ok := ru.mutation.AddedRobustness(); ok {
		_spec.AddField(record.FieldRobustness, field.TypeInt, value)
	}
	if value, ok := ru.mutation.Efficiency(); ok {
		_spec.SetField(record.FieldEfficiency, field.TypeInt, value)
	}
	if value, ok := ru.mutation.AddedEfficiency(); ok {
		_spec.AddField(record.FieldEfficiency, field.TypeInt, value)
	}
	if ru.mutation.ProblemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   record.ProblemTable,
			Columns: []string{record.ProblemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(problem.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.ProblemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   record.ProblemTable,
			Columns: []string{record.ProblemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(problem.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{record.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// RecordUpdateOne is the builder for updating a single Record entity.
type RecordUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RecordMutation
}

// SetUUID sets the "uuid" field.
func (ruo *RecordUpdateOne) SetUUID(s string) *RecordUpdateOne {
	ruo.mutation.SetUUID(s)
	return ruo
}

// SetUserUUID sets the "user_uuid" field.
func (ruo *RecordUpdateOne) SetUserUUID(s string) *RecordUpdateOne {
	ruo.mutation.SetUserUUID(s)
	return ruo
}

// SetNillableUserUUID sets the "user_uuid" field if the given value is not nil.
func (ruo *RecordUpdateOne) SetNillableUserUUID(s *string) *RecordUpdateOne {
	if s != nil {
		ruo.SetUserUUID(*s)
	}
	return ruo
}

// SetCode sets the "code" field.
func (ruo *RecordUpdateOne) SetCode(s string) *RecordUpdateOne {
	ruo.mutation.SetCode(s)
	return ruo
}

// SetReadability sets the "readability" field.
func (ruo *RecordUpdateOne) SetReadability(i int) *RecordUpdateOne {
	ruo.mutation.ResetReadability()
	ruo.mutation.SetReadability(i)
	return ruo
}

// SetNillableReadability sets the "readability" field if the given value is not nil.
func (ruo *RecordUpdateOne) SetNillableReadability(i *int) *RecordUpdateOne {
	if i != nil {
		ruo.SetReadability(*i)
	}
	return ruo
}

// AddReadability adds i to the "readability" field.
func (ruo *RecordUpdateOne) AddReadability(i int) *RecordUpdateOne {
	ruo.mutation.AddReadability(i)
	return ruo
}

// SetRobustness sets the "robustness" field.
func (ruo *RecordUpdateOne) SetRobustness(i int) *RecordUpdateOne {
	ruo.mutation.ResetRobustness()
	ruo.mutation.SetRobustness(i)
	return ruo
}

// SetNillableRobustness sets the "robustness" field if the given value is not nil.
func (ruo *RecordUpdateOne) SetNillableRobustness(i *int) *RecordUpdateOne {
	if i != nil {
		ruo.SetRobustness(*i)
	}
	return ruo
}

// AddRobustness adds i to the "robustness" field.
func (ruo *RecordUpdateOne) AddRobustness(i int) *RecordUpdateOne {
	ruo.mutation.AddRobustness(i)
	return ruo
}

// SetEfficiency sets the "efficiency" field.
func (ruo *RecordUpdateOne) SetEfficiency(i int) *RecordUpdateOne {
	ruo.mutation.ResetEfficiency()
	ruo.mutation.SetEfficiency(i)
	return ruo
}

// SetNillableEfficiency sets the "efficiency" field if the given value is not nil.
func (ruo *RecordUpdateOne) SetNillableEfficiency(i *int) *RecordUpdateOne {
	if i != nil {
		ruo.SetEfficiency(*i)
	}
	return ruo
}

// AddEfficiency adds i to the "efficiency" field.
func (ruo *RecordUpdateOne) AddEfficiency(i int) *RecordUpdateOne {
	ruo.mutation.AddEfficiency(i)
	return ruo
}

// SetProblemID sets the "problem" edge to the Problem entity by ID.
func (ruo *RecordUpdateOne) SetProblemID(id int) *RecordUpdateOne {
	ruo.mutation.SetProblemID(id)
	return ruo
}

// SetProblem sets the "problem" edge to the Problem entity.
func (ruo *RecordUpdateOne) SetProblem(p *Problem) *RecordUpdateOne {
	return ruo.SetProblemID(p.ID)
}

// Mutation returns the RecordMutation object of the builder.
func (ruo *RecordUpdateOne) Mutation() *RecordMutation {
	return ruo.mutation
}

// ClearProblem clears the "problem" edge to the Problem entity.
func (ruo *RecordUpdateOne) ClearProblem() *RecordUpdateOne {
	ruo.mutation.ClearProblem()
	return ruo
}

// Where appends a list predicates to the RecordUpdate builder.
func (ruo *RecordUpdateOne) Where(ps ...predicate.Record) *RecordUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RecordUpdateOne) Select(field string, fields ...string) *RecordUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Record entity.
func (ruo *RecordUpdateOne) Save(ctx context.Context) (*Record, error) {
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RecordUpdateOne) SaveX(ctx context.Context) *Record {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RecordUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RecordUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *RecordUpdateOne) check() error {
	if _, ok := ruo.mutation.ProblemID(); ruo.mutation.ProblemCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Record.problem"`)
	}
	return nil
}

func (ruo *RecordUpdateOne) sqlSave(ctx context.Context) (_node *Record, err error) {
	if err := ruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(record.Table, record.Columns, sqlgraph.NewFieldSpec(record.FieldID, field.TypeInt))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Record.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, record.FieldID)
		for _, f := range fields {
			if !record.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != record.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.UUID(); ok {
		_spec.SetField(record.FieldUUID, field.TypeString, value)
	}
	if value, ok := ruo.mutation.UserUUID(); ok {
		_spec.SetField(record.FieldUserUUID, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Code(); ok {
		_spec.SetField(record.FieldCode, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Readability(); ok {
		_spec.SetField(record.FieldReadability, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.AddedReadability(); ok {
		_spec.AddField(record.FieldReadability, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.Robustness(); ok {
		_spec.SetField(record.FieldRobustness, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.AddedRobustness(); ok {
		_spec.AddField(record.FieldRobustness, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.Efficiency(); ok {
		_spec.SetField(record.FieldEfficiency, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.AddedEfficiency(); ok {
		_spec.AddField(record.FieldEfficiency, field.TypeInt, value)
	}
	if ruo.mutation.ProblemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   record.ProblemTable,
			Columns: []string{record.ProblemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(problem.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.ProblemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   record.ProblemTable,
			Columns: []string{record.ProblemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(problem.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Record{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{record.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}
