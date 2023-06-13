// Code generated by ent, DO NOT EDIT.

package report

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the report type in the database.
	Label = "report"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRequestID holds the string denoting the request_id field in the database.
	FieldRequestID = "request_id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldProjectFeedbacks holds the string denoting the project_feedbacks field in the database.
	FieldProjectFeedbacks = "project_feedbacks"
	// FieldTechStackFeedbacks holds the string denoting the tech_stack_feedbacks field in the database.
	FieldTechStackFeedbacks = "tech_stack_feedbacks"
	// FieldProjectRecommendations holds the string denoting the project_recommendations field in the database.
	FieldProjectRecommendations = "project_recommendations"
	// FieldTechStackRecommendations holds the string denoting the tech_stack_recommendations field in the database.
	FieldTechStackRecommendations = "tech_stack_recommendations"
	// Table holds the table name of the report in the database.
	Table = "reports"
)

// Columns holds all SQL columns for report fields.
var Columns = []string{
	FieldID,
	FieldRequestID,
	FieldStatus,
	FieldProjectFeedbacks,
	FieldTechStackFeedbacks,
	FieldProjectRecommendations,
	FieldTechStackRecommendations,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Report queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByRequestID orders the results by the request_id field.
func ByRequestID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRequestID, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}
