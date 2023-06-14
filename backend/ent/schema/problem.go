package schema

import (
	"code-connect/gateway"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Problem holds the schema definition for the Problem entity.
type Problem struct {
	ent.Schema
}

// Fields of the Problem.
func (Problem) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().StructTag("-"),
		field.String("uuid").Unique().StructTag("id"),
		field.Text("code"),
		field.String("title"),
		field.String("language").GoType(gateway.ProgrammingLanguage("")),
		field.Int("difficulty").Min(0).Max(3000).Default(1500),
		field.Int("readability").Default(0),
		field.Int("modularity").Default(0),
		field.Int("efficiency").Default(0),
		field.Int("testability").Default(0),
		field.Int("maintainability").Default(0),
	}
}

// Edges of the Problem.
func (Problem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("records", Record.Type),
	}
}
