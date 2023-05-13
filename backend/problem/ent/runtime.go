// Code generated by ent, DO NOT EDIT.

package ent

import (
	"code-connect/problem/ent/problem"
	"code-connect/problem/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	problemFields := schema.Problem{}.Fields()
	_ = problemFields
	// problemDescEstimatedTime is the schema descriptor for estimated_time field.
	problemDescEstimatedTime := problemFields[5].Descriptor()
	// problem.DefaultEstimatedTime holds the default value on creation for the estimated_time field.
	problem.DefaultEstimatedTime = problemDescEstimatedTime.Default.(int)
}
