// Code generated by ent, DO NOT EDIT.

package record

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the record type in the database.
	Label = "record"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// FieldUserUUID holds the string denoting the user_uuid field in the database.
	FieldUserUUID = "user_uuid"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldReadability holds the string denoting the readability field in the database.
	FieldReadability = "readability"
	// FieldRobustness holds the string denoting the robustness field in the database.
	FieldRobustness = "robustness"
	// FieldEfficiency holds the string denoting the efficiency field in the database.
	FieldEfficiency = "efficiency"
	// EdgeProblem holds the string denoting the problem edge name in mutations.
	EdgeProblem = "problem"
	// Table holds the table name of the record in the database.
	Table = "records"
	// ProblemTable is the table that holds the problem relation/edge.
	ProblemTable = "records"
	// ProblemInverseTable is the table name for the Problem entity.
	// It exists in this package in order to avoid circular dependency with the "problem" package.
	ProblemInverseTable = "problems"
	// ProblemColumn is the table column denoting the problem relation/edge.
	ProblemColumn = "problem_records"
)

// Columns holds all SQL columns for record fields.
var Columns = []string{
	FieldID,
	FieldUUID,
	FieldUserUUID,
	FieldCode,
	FieldReadability,
	FieldRobustness,
	FieldEfficiency,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "records"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"problem_records",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultReadability holds the default value on creation for the "readability" field.
	DefaultReadability int
	// DefaultRobustness holds the default value on creation for the "robustness" field.
	DefaultRobustness int
	// DefaultEfficiency holds the default value on creation for the "efficiency" field.
	DefaultEfficiency int
)

// OrderOption defines the ordering options for the Record queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUUID orders the results by the uuid field.
func ByUUID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUUID, opts...).ToFunc()
}

// ByUserUUID orders the results by the user_uuid field.
func ByUserUUID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserUUID, opts...).ToFunc()
}

// ByCode orders the results by the code field.
func ByCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCode, opts...).ToFunc()
}

// ByReadability orders the results by the readability field.
func ByReadability(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReadability, opts...).ToFunc()
}

// ByRobustness orders the results by the robustness field.
func ByRobustness(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRobustness, opts...).ToFunc()
}

// ByEfficiency orders the results by the efficiency field.
func ByEfficiency(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEfficiency, opts...).ToFunc()
}

// ByProblemField orders the results by problem field.
func ByProblemField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProblemStep(), sql.OrderByField(field, opts...))
	}
}
func newProblemStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProblemInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ProblemTable, ProblemColumn),
	)
}
