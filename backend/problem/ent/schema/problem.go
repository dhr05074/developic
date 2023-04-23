package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Problem holds the schema definition for the Problem entity.
type Problem struct {
	ent.Schema
}

// Fields of the Problem.
func (Problem) Fields() []ent.Field {
	return []ent.Field{
		field.String("uuid"),
		field.String("title").Optional(),
		field.Text("content").Optional(),
		field.String("request_id"),
	}
}

// Edges of the Problem.
func (Problem) Edges() []ent.Edge {
	return nil
}
