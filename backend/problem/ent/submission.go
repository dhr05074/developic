// Code generated by ent, DO NOT EDIT.

package ent

import (
	"code-connect/problem/ent/problem"
	"code-connect/problem/ent/submission"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Submission is the model entity for the Submission schema.
type Submission struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UUID holds the value of the "uuid" field.
	UUID string `json:"uuid,omitempty"`
	// Code holds the value of the "code" field.
	Code string `json:"code,omitempty"`
	// SubmitterID holds the value of the "submitter_id" field.
	SubmitterID string `json:"submitter_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SubmissionQuery when eager-loading is set.
	Edges               SubmissionEdges `json:"edges"`
	problem_submissions *int
}

// SubmissionEdges holds the relations/edges for other nodes in the graph.
type SubmissionEdges struct {
	// Problem holds the value of the problem edge.
	Problem *Problem `json:"problem,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ProblemOrErr returns the Problem value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SubmissionEdges) ProblemOrErr() (*Problem, error) {
	if e.loadedTypes[0] {
		if e.Problem == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: problem.Label}
		}
		return e.Problem, nil
	}
	return nil, &NotLoadedError{edge: "problem"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Submission) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case submission.FieldID:
			values[i] = new(sql.NullInt64)
		case submission.FieldUUID, submission.FieldCode, submission.FieldSubmitterID:
			values[i] = new(sql.NullString)
		case submission.ForeignKeys[0]: // problem_submissions
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Submission", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Submission fields.
func (s *Submission) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case submission.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case submission.FieldUUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value.Valid {
				s.UUID = value.String
			}
		case submission.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				s.Code = value.String
			}
		case submission.FieldSubmitterID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field submitter_id", values[i])
			} else if value.Valid {
				s.SubmitterID = value.String
			}
		case submission.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field problem_submissions", value)
			} else if value.Valid {
				s.problem_submissions = new(int)
				*s.problem_submissions = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryProblem queries the "problem" edge of the Submission entity.
func (s *Submission) QueryProblem() *ProblemQuery {
	return NewSubmissionClient(s.config).QueryProblem(s)
}

// Update returns a builder for updating this Submission.
// Note that you need to call Submission.Unwrap() before calling this method if this Submission
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Submission) Update() *SubmissionUpdateOne {
	return NewSubmissionClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Submission entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Submission) Unwrap() *Submission {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Submission is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Submission) String() string {
	var builder strings.Builder
	builder.WriteString("Submission(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("uuid=")
	builder.WriteString(s.UUID)
	builder.WriteString(", ")
	builder.WriteString("code=")
	builder.WriteString(s.Code)
	builder.WriteString(", ")
	builder.WriteString("submitter_id=")
	builder.WriteString(s.SubmitterID)
	builder.WriteByte(')')
	return builder.String()
}

// Submissions is a parsable slice of Submission.
type Submissions []*Submission