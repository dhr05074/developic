// Code generated by ent, DO NOT EDIT.

package ent

import (
	"code-connect/ent/problem"
	"code-connect/gateway"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Problem is the model entity for the Problem schema.
type Problem struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty" -`
	// UUID holds the value of the "uuid" field.
	UUID string `json:"uuid,omitempty" id`
	// Code holds the value of the "code" field.
	Code string `json:"code,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Language holds the value of the "language" field.
	Language gateway.ProgrammingLanguage `json:"language,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Difficulty holds the value of the "difficulty" field.
	Difficulty int `json:"difficulty,omitempty"`
	// Readability holds the value of the "readability" field.
	Readability int `json:"readability,omitempty"`
	// Robustness holds the value of the "robustness" field.
	Robustness int `json:"robustness,omitempty"`
	// Efficiency holds the value of the "efficiency" field.
	Efficiency int `json:"efficiency,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProblemQuery when eager-loading is set.
	Edges        ProblemEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ProblemEdges holds the relations/edges for other nodes in the graph.
type ProblemEdges struct {
	// Records holds the value of the records edge.
	Records []*Record `json:"records,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// RecordsOrErr returns the Records value or an error if the edge
// was not loaded in eager-loading.
func (e ProblemEdges) RecordsOrErr() ([]*Record, error) {
	if e.loadedTypes[0] {
		return e.Records, nil
	}
	return nil, &NotLoadedError{edge: "records"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Problem) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case problem.FieldID, problem.FieldDifficulty, problem.FieldReadability, problem.FieldRobustness, problem.FieldEfficiency:
			values[i] = new(sql.NullInt64)
		case problem.FieldUUID, problem.FieldCode, problem.FieldTitle, problem.FieldLanguage, problem.FieldDescription:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Problem fields.
func (pr *Problem) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case problem.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = int(value.Int64)
		case problem.FieldUUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value.Valid {
				pr.UUID = value.String
			}
		case problem.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				pr.Code = value.String
			}
		case problem.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				pr.Title = value.String
			}
		case problem.FieldLanguage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field language", values[i])
			} else if value.Valid {
				pr.Language = gateway.ProgrammingLanguage(value.String)
			}
		case problem.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pr.Description = value.String
			}
		case problem.FieldDifficulty:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field difficulty", values[i])
			} else if value.Valid {
				pr.Difficulty = int(value.Int64)
			}
		case problem.FieldReadability:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field readability", values[i])
			} else if value.Valid {
				pr.Readability = int(value.Int64)
			}
		case problem.FieldRobustness:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field robustness", values[i])
			} else if value.Valid {
				pr.Robustness = int(value.Int64)
			}
		case problem.FieldEfficiency:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field efficiency", values[i])
			} else if value.Valid {
				pr.Efficiency = int(value.Int64)
			}
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Problem.
// This includes values selected through modifiers, order, etc.
func (pr *Problem) Value(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// QueryRecords queries the "records" edge of the Problem entity.
func (pr *Problem) QueryRecords() *RecordQuery {
	return NewProblemClient(pr.config).QueryRecords(pr)
}

// Update returns a builder for updating this Problem.
// Note that you need to call Problem.Unwrap() before calling this method if this Problem
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Problem) Update() *ProblemUpdateOne {
	return NewProblemClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the Problem entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Problem) Unwrap() *Problem {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Problem is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Problem) String() string {
	var builder strings.Builder
	builder.WriteString("Problem(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("uuid=")
	builder.WriteString(pr.UUID)
	builder.WriteString(", ")
	builder.WriteString("code=")
	builder.WriteString(pr.Code)
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(pr.Title)
	builder.WriteString(", ")
	builder.WriteString("language=")
	builder.WriteString(fmt.Sprintf("%v", pr.Language))
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(pr.Description)
	builder.WriteString(", ")
	builder.WriteString("difficulty=")
	builder.WriteString(fmt.Sprintf("%v", pr.Difficulty))
	builder.WriteString(", ")
	builder.WriteString("readability=")
	builder.WriteString(fmt.Sprintf("%v", pr.Readability))
	builder.WriteString(", ")
	builder.WriteString("robustness=")
	builder.WriteString(fmt.Sprintf("%v", pr.Robustness))
	builder.WriteString(", ")
	builder.WriteString("efficiency=")
	builder.WriteString(fmt.Sprintf("%v", pr.Efficiency))
	builder.WriteByte(')')
	return builder.String()
}

// Problems is a parsable slice of Problem.
type Problems []*Problem
