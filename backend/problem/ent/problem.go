// Code generated by ent, DO NOT EDIT.

package ent

import (
	"code-connect/problem/ent/problem"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Problem is the model entity for the Problem schema.
type Problem struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UUID holds the value of the "uuid" field.
	UUID string `json:"uuid,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// RequestID holds the value of the "request_id" field.
	RequestID string `json:"request_id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Problem) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case problem.FieldID:
			values[i] = new(sql.NullInt64)
		case problem.FieldUUID, problem.FieldTitle, problem.FieldContent, problem.FieldRequestID:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Problem", columns[i])
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
		case problem.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				pr.Title = value.String
			}
		case problem.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				pr.Content = value.String
			}
		case problem.FieldRequestID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field request_id", values[i])
			} else if value.Valid {
				pr.RequestID = value.String
			}
		}
	}
	return nil
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
	builder.WriteString("title=")
	builder.WriteString(pr.Title)
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(pr.Content)
	builder.WriteString(", ")
	builder.WriteString("request_id=")
	builder.WriteString(pr.RequestID)
	builder.WriteByte(')')
	return builder.String()
}

// Problems is a parsable slice of Problem.
type Problems []*Problem
