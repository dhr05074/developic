package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"portfolio/schema"
)

// Report holds the schema definition for the Report entity.
type Report struct {
	ent.Schema
}

// Fields of the Report.
func (Report) Fields() []ent.Field {
	return []ent.Field{
		field.String("request_id"),
		field.String("status"),
		field.JSON("project_feedbacks", []schema.Feedback{}),
		field.JSON("tech_stack_feedbacks", []schema.Feedback{}),
		field.JSON("project_recommendations", []schema.Recommendation{}),
		field.JSON("tech_stack_recommendations", []schema.Recommendation{}),
	}
}

// Edges of the Report.
func (Report) Edges() []ent.Edge {
	return nil
}
