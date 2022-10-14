package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Users holds the schema definition for the Users entity.
type Advertisement struct {
	ent.Schema
}

// Fields of the Users.
func (Advertisement) Fields() []ent.Field {
	return []ent.Field{
		field.String("Date"),
		field.String("Contrat"),
		field.String("Entreprise"),
		field.String("Image"),
		field.String("Localisation"),
		field.String("Name"),
		field.String("Remuneration"),
		field.String("Sector"),
	}
}

// Edges of the Users.
func (Advertisement) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", Users.Type).Ref("advertisement").Unique(),
	}
}
