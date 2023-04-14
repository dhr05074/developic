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
	"entgo.io/ent/dialect/sql/sqljson"
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

// SetDifficulty sets the "difficulty" field.
func (pu *ProblemUpdate) SetDifficulty(i int) *ProblemUpdate {
	pu.mutation.ResetDifficulty()
	pu.mutation.SetDifficulty(i)
	return pu
}

// AddDifficulty adds i to the "difficulty" field.
func (pu *ProblemUpdate) AddDifficulty(i int) *ProblemUpdate {
	pu.mutation.AddDifficulty(i)
	return pu
}

// SetLanguage sets the "language" field.
func (pu *ProblemUpdate) SetLanguage(s string) *ProblemUpdate {
	pu.mutation.SetLanguage(s)
	return pu
}

// SetStatement sets the "statement" field.
func (pu *ProblemUpdate) SetStatement(s string) *ProblemUpdate {
	pu.mutation.SetStatement(s)
	return pu
}

// SetNillableStatement sets the "statement" field if the given value is not nil.
func (pu *ProblemUpdate) SetNillableStatement(s *string) *ProblemUpdate {
	if s != nil {
		pu.SetStatement(*s)
	}
	return pu
}

// ClearStatement clears the value of the "statement" field.
func (pu *ProblemUpdate) ClearStatement() *ProblemUpdate {
	pu.mutation.ClearStatement()
	return pu
}

// SetExamples sets the "examples" field.
func (pu *ProblemUpdate) SetExamples(s string) *ProblemUpdate {
	pu.mutation.SetExamples(s)
	return pu
}

// SetNillableExamples sets the "examples" field if the given value is not nil.
func (pu *ProblemUpdate) SetNillableExamples(s *string) *ProblemUpdate {
	if s != nil {
		pu.SetExamples(*s)
	}
	return pu
}

// ClearExamples clears the value of the "examples" field.
func (pu *ProblemUpdate) ClearExamples() *ProblemUpdate {
	pu.mutation.ClearExamples()
	return pu
}

// SetConstraints sets the "constraints" field.
func (pu *ProblemUpdate) SetConstraints(s []string) *ProblemUpdate {
	pu.mutation.SetConstraints(s)
	return pu
}

// AppendConstraints appends s to the "constraints" field.
func (pu *ProblemUpdate) AppendConstraints(s []string) *ProblemUpdate {
	pu.mutation.AppendConstraints(s)
	return pu
}

// ClearConstraints clears the value of the "constraints" field.
func (pu *ProblemUpdate) ClearConstraints() *ProblemUpdate {
	pu.mutation.ClearConstraints()
	return pu
}

// SetEvaluationCriteria sets the "evaluation_criteria" field.
func (pu *ProblemUpdate) SetEvaluationCriteria(s []string) *ProblemUpdate {
	pu.mutation.SetEvaluationCriteria(s)
	return pu
}

// AppendEvaluationCriteria appends s to the "evaluation_criteria" field.
func (pu *ProblemUpdate) AppendEvaluationCriteria(s []string) *ProblemUpdate {
	pu.mutation.AppendEvaluationCriteria(s)
	return pu
}

// ClearEvaluationCriteria clears the value of the "evaluation_criteria" field.
func (pu *ProblemUpdate) ClearEvaluationCriteria() *ProblemUpdate {
	pu.mutation.ClearEvaluationCriteria()
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
	if value, ok := pu.mutation.Difficulty(); ok {
		_spec.SetField(problem.FieldDifficulty, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedDifficulty(); ok {
		_spec.AddField(problem.FieldDifficulty, field.TypeInt, value)
	}
	if value, ok := pu.mutation.Language(); ok {
		_spec.SetField(problem.FieldLanguage, field.TypeString, value)
	}
	if value, ok := pu.mutation.Statement(); ok {
		_spec.SetField(problem.FieldStatement, field.TypeString, value)
	}
	if pu.mutation.StatementCleared() {
		_spec.ClearField(problem.FieldStatement, field.TypeString)
	}
	if value, ok := pu.mutation.Examples(); ok {
		_spec.SetField(problem.FieldExamples, field.TypeString, value)
	}
	if pu.mutation.ExamplesCleared() {
		_spec.ClearField(problem.FieldExamples, field.TypeString)
	}
	if value, ok := pu.mutation.Constraints(); ok {
		_spec.SetField(problem.FieldConstraints, field.TypeJSON, value)
	}
	if value, ok := pu.mutation.AppendedConstraints(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, problem.FieldConstraints, value)
		})
	}
	if pu.mutation.ConstraintsCleared() {
		_spec.ClearField(problem.FieldConstraints, field.TypeJSON)
	}
	if value, ok := pu.mutation.EvaluationCriteria(); ok {
		_spec.SetField(problem.FieldEvaluationCriteria, field.TypeJSON, value)
	}
	if value, ok := pu.mutation.AppendedEvaluationCriteria(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, problem.FieldEvaluationCriteria, value)
		})
	}
	if pu.mutation.EvaluationCriteriaCleared() {
		_spec.ClearField(problem.FieldEvaluationCriteria, field.TypeJSON)
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

// SetDifficulty sets the "difficulty" field.
func (puo *ProblemUpdateOne) SetDifficulty(i int) *ProblemUpdateOne {
	puo.mutation.ResetDifficulty()
	puo.mutation.SetDifficulty(i)
	return puo
}

// AddDifficulty adds i to the "difficulty" field.
func (puo *ProblemUpdateOne) AddDifficulty(i int) *ProblemUpdateOne {
	puo.mutation.AddDifficulty(i)
	return puo
}

// SetLanguage sets the "language" field.
func (puo *ProblemUpdateOne) SetLanguage(s string) *ProblemUpdateOne {
	puo.mutation.SetLanguage(s)
	return puo
}

// SetStatement sets the "statement" field.
func (puo *ProblemUpdateOne) SetStatement(s string) *ProblemUpdateOne {
	puo.mutation.SetStatement(s)
	return puo
}

// SetNillableStatement sets the "statement" field if the given value is not nil.
func (puo *ProblemUpdateOne) SetNillableStatement(s *string) *ProblemUpdateOne {
	if s != nil {
		puo.SetStatement(*s)
	}
	return puo
}

// ClearStatement clears the value of the "statement" field.
func (puo *ProblemUpdateOne) ClearStatement() *ProblemUpdateOne {
	puo.mutation.ClearStatement()
	return puo
}

// SetExamples sets the "examples" field.
func (puo *ProblemUpdateOne) SetExamples(s string) *ProblemUpdateOne {
	puo.mutation.SetExamples(s)
	return puo
}

// SetNillableExamples sets the "examples" field if the given value is not nil.
func (puo *ProblemUpdateOne) SetNillableExamples(s *string) *ProblemUpdateOne {
	if s != nil {
		puo.SetExamples(*s)
	}
	return puo
}

// ClearExamples clears the value of the "examples" field.
func (puo *ProblemUpdateOne) ClearExamples() *ProblemUpdateOne {
	puo.mutation.ClearExamples()
	return puo
}

// SetConstraints sets the "constraints" field.
func (puo *ProblemUpdateOne) SetConstraints(s []string) *ProblemUpdateOne {
	puo.mutation.SetConstraints(s)
	return puo
}

// AppendConstraints appends s to the "constraints" field.
func (puo *ProblemUpdateOne) AppendConstraints(s []string) *ProblemUpdateOne {
	puo.mutation.AppendConstraints(s)
	return puo
}

// ClearConstraints clears the value of the "constraints" field.
func (puo *ProblemUpdateOne) ClearConstraints() *ProblemUpdateOne {
	puo.mutation.ClearConstraints()
	return puo
}

// SetEvaluationCriteria sets the "evaluation_criteria" field.
func (puo *ProblemUpdateOne) SetEvaluationCriteria(s []string) *ProblemUpdateOne {
	puo.mutation.SetEvaluationCriteria(s)
	return puo
}

// AppendEvaluationCriteria appends s to the "evaluation_criteria" field.
func (puo *ProblemUpdateOne) AppendEvaluationCriteria(s []string) *ProblemUpdateOne {
	puo.mutation.AppendEvaluationCriteria(s)
	return puo
}

// ClearEvaluationCriteria clears the value of the "evaluation_criteria" field.
func (puo *ProblemUpdateOne) ClearEvaluationCriteria() *ProblemUpdateOne {
	puo.mutation.ClearEvaluationCriteria()
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
	if value, ok := puo.mutation.Difficulty(); ok {
		_spec.SetField(problem.FieldDifficulty, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedDifficulty(); ok {
		_spec.AddField(problem.FieldDifficulty, field.TypeInt, value)
	}
	if value, ok := puo.mutation.Language(); ok {
		_spec.SetField(problem.FieldLanguage, field.TypeString, value)
	}
	if value, ok := puo.mutation.Statement(); ok {
		_spec.SetField(problem.FieldStatement, field.TypeString, value)
	}
	if puo.mutation.StatementCleared() {
		_spec.ClearField(problem.FieldStatement, field.TypeString)
	}
	if value, ok := puo.mutation.Examples(); ok {
		_spec.SetField(problem.FieldExamples, field.TypeString, value)
	}
	if puo.mutation.ExamplesCleared() {
		_spec.ClearField(problem.FieldExamples, field.TypeString)
	}
	if value, ok := puo.mutation.Constraints(); ok {
		_spec.SetField(problem.FieldConstraints, field.TypeJSON, value)
	}
	if value, ok := puo.mutation.AppendedConstraints(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, problem.FieldConstraints, value)
		})
	}
	if puo.mutation.ConstraintsCleared() {
		_spec.ClearField(problem.FieldConstraints, field.TypeJSON)
	}
	if value, ok := puo.mutation.EvaluationCriteria(); ok {
		_spec.SetField(problem.FieldEvaluationCriteria, field.TypeJSON, value)
	}
	if value, ok := puo.mutation.AppendedEvaluationCriteria(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, problem.FieldEvaluationCriteria, value)
		})
	}
	if puo.mutation.EvaluationCriteriaCleared() {
		_spec.ClearField(problem.FieldEvaluationCriteria, field.TypeJSON)
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
