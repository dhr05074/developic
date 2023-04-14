package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Problem holds the schema definition for the Problem entity.
type Problem struct {
	ent.Schema
}

// Fields of the Problem.
func (Problem) Fields() []ent.Field {
	return []ent.Field{
		field.String("uuid"),
		field.Int("difficulty"),
		field.String("language"),
		field.Text("statement").Optional(),
		field.Text("examples").Optional(),
		field.Strings("constraints").Optional(),
		field.Strings("evaluation_criteria").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
	}
}

// Edges of the Problem.
func (Problem) Edges() []ent.Edge {
	return nil
}
