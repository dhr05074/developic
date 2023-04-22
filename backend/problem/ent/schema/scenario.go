package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Scenario holds the schema definition for the Scenario entity.
type Scenario struct {
	ent.Schema
}

// Fields of the Scenario.
func (Scenario) Fields() []ent.Field {
	return []ent.Field{
		field.String("uuid"),
		field.String("title"),
		field.Text("content"),
		field.String("request_id"),
	}
}

// Edges of the Scenario.
func (Scenario) Edges() []ent.Edge {
	return nil
}
