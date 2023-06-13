// Code generated by ent, DO NOT EDIT.

package ent

import (
	"code-connect/ent/problem"
	"code-connect/ent/record"
	"code-connect/gateway"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProblemCreate is the builder for creating a Problem entity.
type ProblemCreate struct {
	config
	mutation *ProblemMutation
	hooks    []Hook
}

// SetUUID sets the "uuid" field.
func (pc *ProblemCreate) SetUUID(s string) *ProblemCreate {
	pc.mutation.SetUUID(s)
	return pc
}

// SetCode sets the "code" field.
func (pc *ProblemCreate) SetCode(s string) *ProblemCreate {
	pc.mutation.SetCode(s)
	return pc
}

// SetTitle sets the "title" field.
func (pc *ProblemCreate) SetTitle(s string) *ProblemCreate {
	pc.mutation.SetTitle(s)
	return pc
}

// SetLanguage sets the "language" field.
func (pc *ProblemCreate) SetLanguage(gl gateway.ProgrammingLanguage) *ProblemCreate {
	pc.mutation.SetLanguage(gl)
	return pc
}

// SetDifficulty sets the "difficulty" field.
func (pc *ProblemCreate) SetDifficulty(i int) *ProblemCreate {
	pc.mutation.SetDifficulty(i)
	return pc
}

// SetNillableDifficulty sets the "difficulty" field if the given value is not nil.
func (pc *ProblemCreate) SetNillableDifficulty(i *int) *ProblemCreate {
	if i != nil {
		pc.SetDifficulty(*i)
	}
	return pc
}

// SetReadability sets the "readability" field.
func (pc *ProblemCreate) SetReadability(i int) *ProblemCreate {
	pc.mutation.SetReadability(i)
	return pc
}

// SetNillableReadability sets the "readability" field if the given value is not nil.
func (pc *ProblemCreate) SetNillableReadability(i *int) *ProblemCreate {
	if i != nil {
		pc.SetReadability(*i)
	}
	return pc
}

// SetModularity sets the "modularity" field.
func (pc *ProblemCreate) SetModularity(i int) *ProblemCreate {
	pc.mutation.SetModularity(i)
	return pc
}

// SetNillableModularity sets the "modularity" field if the given value is not nil.
func (pc *ProblemCreate) SetNillableModularity(i *int) *ProblemCreate {
	if i != nil {
		pc.SetModularity(*i)
	}
	return pc
}

// SetEfficiency sets the "efficiency" field.
func (pc *ProblemCreate) SetEfficiency(i int) *ProblemCreate {
	pc.mutation.SetEfficiency(i)
	return pc
}

// SetNillableEfficiency sets the "efficiency" field if the given value is not nil.
func (pc *ProblemCreate) SetNillableEfficiency(i *int) *ProblemCreate {
	if i != nil {
		pc.SetEfficiency(*i)
	}
	return pc
}

// SetTestability sets the "testability" field.
func (pc *ProblemCreate) SetTestability(i int) *ProblemCreate {
	pc.mutation.SetTestability(i)
	return pc
}

// SetNillableTestability sets the "testability" field if the given value is not nil.
func (pc *ProblemCreate) SetNillableTestability(i *int) *ProblemCreate {
	if i != nil {
		pc.SetTestability(*i)
	}
	return pc
}

// SetMaintainablity sets the "maintainablity" field.
func (pc *ProblemCreate) SetMaintainablity(i int) *ProblemCreate {
	pc.mutation.SetMaintainablity(i)
	return pc
}

// SetNillableMaintainablity sets the "maintainablity" field if the given value is not nil.
func (pc *ProblemCreate) SetNillableMaintainablity(i *int) *ProblemCreate {
	if i != nil {
		pc.SetMaintainablity(*i)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *ProblemCreate) SetID(i int) *ProblemCreate {
	pc.mutation.SetID(i)
	return pc
}

// AddRecordIDs adds the "records" edge to the Record entity by IDs.
func (pc *ProblemCreate) AddRecordIDs(ids ...int) *ProblemCreate {
	pc.mutation.AddRecordIDs(ids...)
	return pc
}

// AddRecords adds the "records" edges to the Record entity.
func (pc *ProblemCreate) AddRecords(r ...*Record) *ProblemCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pc.AddRecordIDs(ids...)
}

// Mutation returns the ProblemMutation object of the builder.
func (pc *ProblemCreate) Mutation() *ProblemMutation {
	return pc.mutation
}

// Save creates the Problem in the database.
func (pc *ProblemCreate) Save(ctx context.Context) (*Problem, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProblemCreate) SaveX(ctx context.Context) *Problem {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ProblemCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ProblemCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *ProblemCreate) defaults() {
	if _, ok := pc.mutation.Difficulty(); !ok {
		v := problem.DefaultDifficulty
		pc.mutation.SetDifficulty(v)
	}
	if _, ok := pc.mutation.Readability(); !ok {
		v := problem.DefaultReadability
		pc.mutation.SetReadability(v)
	}
	if _, ok := pc.mutation.Modularity(); !ok {
		v := problem.DefaultModularity
		pc.mutation.SetModularity(v)
	}
	if _, ok := pc.mutation.Efficiency(); !ok {
		v := problem.DefaultEfficiency
		pc.mutation.SetEfficiency(v)
	}
	if _, ok := pc.mutation.Testability(); !ok {
		v := problem.DefaultTestability
		pc.mutation.SetTestability(v)
	}
	if _, ok := pc.mutation.Maintainablity(); !ok {
		v := problem.DefaultMaintainablity
		pc.mutation.SetMaintainablity(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProblemCreate) check() error {
	if _, ok := pc.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New(`ent: missing required field "Problem.uuid"`)}
	}
	if _, ok := pc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "Problem.code"`)}
	}
	if _, ok := pc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Problem.title"`)}
	}
	if _, ok := pc.mutation.Language(); !ok {
		return &ValidationError{Name: "language", err: errors.New(`ent: missing required field "Problem.language"`)}
	}
	if _, ok := pc.mutation.Difficulty(); !ok {
		return &ValidationError{Name: "difficulty", err: errors.New(`ent: missing required field "Problem.difficulty"`)}
	}
	if v, ok := pc.mutation.Difficulty(); ok {
		if err := problem.DifficultyValidator(v); err != nil {
			return &ValidationError{Name: "difficulty", err: fmt.Errorf(`ent: validator failed for field "Problem.difficulty": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Readability(); !ok {
		return &ValidationError{Name: "readability", err: errors.New(`ent: missing required field "Problem.readability"`)}
	}
	if _, ok := pc.mutation.Modularity(); !ok {
		return &ValidationError{Name: "modularity", err: errors.New(`ent: missing required field "Problem.modularity"`)}
	}
	if _, ok := pc.mutation.Efficiency(); !ok {
		return &ValidationError{Name: "efficiency", err: errors.New(`ent: missing required field "Problem.efficiency"`)}
	}
	if _, ok := pc.mutation.Testability(); !ok {
		return &ValidationError{Name: "testability", err: errors.New(`ent: missing required field "Problem.testability"`)}
	}
	if _, ok := pc.mutation.Maintainablity(); !ok {
		return &ValidationError{Name: "maintainablity", err: errors.New(`ent: missing required field "Problem.maintainablity"`)}
	}
	return nil
}

func (pc *ProblemCreate) sqlSave(ctx context.Context) (*Problem, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *ProblemCreate) createSpec() (*Problem, *sqlgraph.CreateSpec) {
	var (
		_node = &Problem{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(problem.Table, sqlgraph.NewFieldSpec(problem.FieldID, field.TypeInt))
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.UUID(); ok {
		_spec.SetField(problem.FieldUUID, field.TypeString, value)
		_node.UUID = value
	}
	if value, ok := pc.mutation.Code(); ok {
		_spec.SetField(problem.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := pc.mutation.Title(); ok {
		_spec.SetField(problem.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := pc.mutation.Language(); ok {
		_spec.SetField(problem.FieldLanguage, field.TypeString, value)
		_node.Language = value
	}
	if value, ok := pc.mutation.Difficulty(); ok {
		_spec.SetField(problem.FieldDifficulty, field.TypeInt, value)
		_node.Difficulty = value
	}
	if value, ok := pc.mutation.Readability(); ok {
		_spec.SetField(problem.FieldReadability, field.TypeInt, value)
		_node.Readability = value
	}
	if value, ok := pc.mutation.Modularity(); ok {
		_spec.SetField(problem.FieldModularity, field.TypeInt, value)
		_node.Modularity = value
	}
	if value, ok := pc.mutation.Efficiency(); ok {
		_spec.SetField(problem.FieldEfficiency, field.TypeInt, value)
		_node.Efficiency = value
	}
	if value, ok := pc.mutation.Testability(); ok {
		_spec.SetField(problem.FieldTestability, field.TypeInt, value)
		_node.Testability = value
	}
	if value, ok := pc.mutation.Maintainablity(); ok {
		_spec.SetField(problem.FieldMaintainablity, field.TypeInt, value)
		_node.Maintainablity = value
	}
	if nodes := pc.mutation.RecordsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   problem.RecordsTable,
			Columns: problem.RecordsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(record.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProblemCreateBulk is the builder for creating many Problem entities in bulk.
type ProblemCreateBulk struct {
	config
	builders []*ProblemCreate
}

// Save creates the Problem entities in the database.
func (pcb *ProblemCreateBulk) Save(ctx context.Context) ([]*Problem, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Problem, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProblemMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ProblemCreateBulk) SaveX(ctx context.Context) []*Problem {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ProblemCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ProblemCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}