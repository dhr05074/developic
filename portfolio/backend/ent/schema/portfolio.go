package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"portfolio/schema"
)

// Portfolio holds the schema definition for the Portfolio entity.
type Portfolio struct {
	ent.Schema
}

// Fields of the Portfolio.
func (Portfolio) Fields() []ent.Field {
	return []ent.Field{
		field.String("request_id"),
		field.String("job"),
		field.Uint("career_years"),
		field.Strings("tech_stacks"),
		field.JSON("projects", []schema.Project{}),
		field.Text("preferred_company").Optional(),
	}
}

// Edges of the Portfolio.
func (Portfolio) Edges() []ent.Edge {
	return nil
}
