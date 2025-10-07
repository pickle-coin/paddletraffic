package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Location holds the schema definition for the Location entity.
type Location struct {
	ent.Schema
}

// Fields of the Location.
func (Location) Fields() []ent.Field {
	return []ent.Field{
		field.Text("addressLine").
			MaxLen(120),
		field.Text("city").
			MaxLen(80),
		field.Text("region").
			MaxLen(50),
		field.Text("postalCode").
			MaxLen(20),
		field.Text("countryCode").
			MaxLen(3),
		field.Text("timezone").
			MaxLen(40),
		field.Float("latitude").
			Max(90.0).
			Min(-90.0),
		field.Float("longitude").
			Max(180.0).
			Min(0.0),
		field.String("placeId").
			MaxLen(200).
			Optional(),
	}
}

// Edges of the Location.
func (Location) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("court", Court.Type).
			Ref("location").
			Unique().
			Required(),
	}
}
