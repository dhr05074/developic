// Code generated by ent, DO NOT EDIT.

package problem

const (
	// Label holds the string label denoting the problem type in the database.
	Label = "problem"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldBackground holds the string denoting the background field in the database.
	FieldBackground = "background"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldTestCode holds the string denoting the test_code field in the database.
	FieldTestCode = "test_code"
	// FieldEstimatedTime holds the string denoting the estimated_time field in the database.
	FieldEstimatedTime = "estimated_time"
	// FieldLanguage holds the string denoting the language field in the database.
	FieldLanguage = "language"
	// FieldRequestID holds the string denoting the request_id field in the database.
	FieldRequestID = "request_id"
	// EdgeSubmissions holds the string denoting the submissions edge name in mutations.
	EdgeSubmissions = "submissions"
	// Table holds the table name of the problem in the database.
	Table = "problems"
	// SubmissionsTable is the table that holds the submissions relation/edge.
	SubmissionsTable = "submissions"
	// SubmissionsInverseTable is the table name for the Submission entity.
	// It exists in this package in order to avoid circular dependency with the "submission" package.
	SubmissionsInverseTable = "submissions"
	// SubmissionsColumn is the table column denoting the submissions relation/edge.
	SubmissionsColumn = "problem_submissions"
)

// Columns holds all SQL columns for problem fields.
var Columns = []string{
	FieldID,
	FieldUUID,
	FieldTitle,
	FieldBackground,
	FieldCode,
	FieldTestCode,
	FieldEstimatedTime,
	FieldLanguage,
	FieldRequestID,
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

var (
	// DefaultEstimatedTime holds the default value on creation for the "estimated_time" field.
	DefaultEstimatedTime int
)
