package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Court holds the schema definition for the Court entity.
type Court struct {
	ent.Schema
}

// Fields of the Court.
func (Court) Fields() []ent.Field {
	return []ent.Field{
		field.Text("name").
			MaxLen(200),
		field.Int("courtCount"),
	}
}

// Edges of the Court.
func (Court) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("location", Location.Type).
			Unique(),
	}
}
