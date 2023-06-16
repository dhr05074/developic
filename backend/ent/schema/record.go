package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Record holds the schema definition for the Record entity.
type Record struct {
	ent.Schema
}

// Fields of the Record.
func (Record) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().StructTag("-"),
		field.String("uuid").Unique().StructTag("id"),
		field.String("user_uuid").Default("test"),
		field.Text("code"),
		field.Int("readability").Default(0),
		field.Int("robustness").Default(0),
		field.Int("efficiency").Default(0),
	}
}

// Edges of the Record.
func (Record) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("problem", Problem.Type).Ref("records").Unique().Required(),
	}
}
