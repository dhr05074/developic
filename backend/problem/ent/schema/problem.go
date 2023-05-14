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
		field.String("title"),
		field.Text("background"),
		field.Text("code"),
		field.Text("test_code").Optional(),
		field.Int("estimated_time").Default(30),
		field.String("language"),
		field.String("request_id"),
	}
}

// Edges of the Problem.
func (Problem) Edges() []ent.Edge {
	return nil
}
