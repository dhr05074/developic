// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ProblemsColumns holds the columns for the "problems" table.
	ProblemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uuid", Type: field.TypeString},
		{Name: "title", Type: field.TypeString},
		{Name: "background", Type: field.TypeString, Size: 2147483647},
		{Name: "code", Type: field.TypeString, Size: 2147483647},
		{Name: "test_code", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "estimated_time", Type: field.TypeInt, Default: 30},
		{Name: "language", Type: field.TypeString},
		{Name: "request_id", Type: field.TypeString},
	}
	// ProblemsTable holds the schema information for the "problems" table.
	ProblemsTable = &schema.Table{
		Name:       "problems",
		Columns:    ProblemsColumns,
		PrimaryKey: []*schema.Column{ProblemsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ProblemsTable,
	}
)

func init() {
}
